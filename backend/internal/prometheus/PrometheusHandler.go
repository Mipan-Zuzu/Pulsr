package prometheus

import (
	"strings"
	"math"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
	dto "github.com/prometheus/client_model/go"
	"pulsr/internal/struct"
	"time"
)

func sanitizeMetrics(raw string) string {
	lines := strings.Split(raw, "\n")
	seenHelp := map[string]bool{}
	seenType := map[string]bool{}

	var out []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "# HELP ") {
			parts := strings.SplitN(trimmed, " ", 4)
			if len(parts) >= 3 {
				name := parts[2]
				if seenHelp[name] {
					continue // skip duplikat
				}
				seenHelp[name] = true
			}
		}

		if strings.HasPrefix(trimmed, "# TYPE ") {
			parts := strings.SplitN(trimmed, " ", 4)
			if len(parts) >= 3 {
				name := parts[2]
				if seenType[name] {
					continue // skip duplikat
				}
				seenType[name] = true
			}
		}

		out = append(out, line)
	}

	return strings.Join(out, "\n")
}

func sanitizeValue(v float64) float64 {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0
	}
	return v
}


func FetchMetrics(raw string) (map[string]*dto.MetricFamily, error) {

	cleaned := sanitizeMetrics(raw)

	parser := expfmt.NewTextParser(model.LegacyValidation)

	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(cleaned))
	if err != nil {
		return nil, err
	}

	result := make(map[string]Struct.MetricsJSON)

	for name, mf := range metricFamilies {
		for _, m := range mf.GetMetric() {
			labels := map[string]string{}
			for _, l := range m.GetLabel() {
				labels[l.GetName()] = l.GetValue()
			}

			var value float64
			switch {
			case m.Counter != nil:
				value = m.Counter.GetValue()
			case m.Gauge != nil:
				value = m.Gauge.GetValue()
			case m.Untyped != nil:
				value = m.Untyped.GetValue()
			case m.Summary != nil:
				value = m.Summary.GetSampleSum()
			case m.Histogram != nil:
				value = m.Histogram.GetSampleSum()
			}

			result[name] = Struct.MetricsJSON{
				Name:   name,
				Help:   mf.GetHelp(),
				Type:   mf.GetType().String(),
				Labels: labels,
				Value:  sanitizeValue(value),
			}
		}
	}


	return parser.TextToMetricFamilies(strings.NewReader(cleaned))
}

func BuildSystemUsage(mfs map[string]*dto.MetricFamily, projectRef string) Struct.SystemUsage {
	usage := Struct.SystemUsage{
		ProjectRef: projectRef,
		Timestamp:  time.Now().Unix(),
	}

	getValue := func(m *dto.Metric) float64 {
		switch {
		case m.Gauge != nil:
			return m.Gauge.GetValue()
		case m.Counter != nil:
			return m.Counter.GetValue()
		case m.Untyped != nil:
			return m.Untyped.GetValue()
		}
		return 0
	}

	getLabel := func(m *dto.Metric, key string) string {
		for _, l := range m.GetLabel() {
			if l.GetName() == key {
				return l.GetValue()
			}
		}
		return ""
	}

	diskMap := map[string]*Struct.DiskUsage{}

	for name, mf := range mfs {
		for _, m := range mf.GetMetric() {
			v := getValue(m)

			switch name {
			case "node_load15":
				usage.CPU.LoadAvg15 = v
			case "node_memory_SwapTotal_bytes":
				usage.Memory.SwapTotalBytes = v
			case "node_memory_PageTables_bytes":
				usage.Memory.PageTablesBytes = v
			case "node_memory_Slab_bytes":
				usage.Memory.SlabBytes = v
			case "node_memory_Committed_AS_bytes":
				usage.Memory.CommittedASBytes = v
			case "node_memory_Dirty_bytes":
				usage.Memory.DirtyBytes = v
			case "node_memory_Shmem_bytes":
				usage.Memory.ShmemBytes = v

			case "node_disk_io_time_weighted_seconds_total":
				dev := getLabel(m, "device")
				if diskMap[dev] == nil {
					diskMap[dev] = &Struct.DiskUsage{Device: dev}
				}
				diskMap[dev].IOTimeWeightedSeconds = v

			case "node_disk_read_time_seconds_total":
				dev := getLabel(m, "device")
				if diskMap[dev] == nil {
					diskMap[dev] = &Struct.DiskUsage{Device: dev}
				}
				diskMap[dev].ReadTimeSeconds = v

			case "pgbouncer_version_info":
				usage.PgBouncer.Version = getLabel(m, "version")
			case "pgbouncer_config_max_client_connections":
				usage.PgBouncer.MaxClientConnections = v
			case "pgbouncer_pools_server_active_connections":
				usage.PgBouncer.ServerActiveConnections = v
			case "pgbouncer_pools_server_login_connections":
				usage.PgBouncer.ServerLoginConnections = v
			case "pgbouncer_cached_dns_names":
				usage.PgBouncer.CachedDNSNames = v
			}
		}
	}

	for _, d := range diskMap {
		usage.Disk = append(usage.Disk, *d)
	}

	return usage
}