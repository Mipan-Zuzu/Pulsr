<script setup lang="ts">
import { ref, onMounted, watch, onBeforeUnmount } from 'vue'
import {
  Chart,
  LineController,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Tooltip,
  Legend,
  Filler,
} from 'chart.js'
import type { SystemUsage } from '../../type/typeSupabase'

Chart.register(
  LineController,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Tooltip,
  Legend,
  Filler
)

const props = defineProps<{
  history: SystemUsage[]
}>()

const canvasRef = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

const createOrUpdateChart = () => {
  if (!canvasRef.value || !props.history.length) return

  const labels = props.history.map((item, idx) => {
    if (item.timestamp) {
      const d = new Date(item.timestamp * 1000)
      return d.toLocaleTimeString([], { minute: '2-digit', second: '2-digit' })
    }
    return `T-${props.history.length - idx}`
  })

  const loadData = props.history.map((h) => h.cpu?.load_avg_15m || 0)
  const pgbouncerConns = props.history.map((h) => h.pgbouncer?.server_active_connections || 0)

  if (chartInstance) {
    chartInstance.destroy()
  }

  chartInstance = new Chart(canvasRef.value, {
    type: 'line',
    data: {
      labels,
      datasets: [
        {
          label: 'CPU Load (15m)',
          data: loadData,
          borderColor: '#0284c7', // Sky 600
          backgroundColor: 'rgba(2, 132, 199, 0.1)',
          fill: true,
          tension: 0.4,
          borderWidth: 2,
          pointRadius: 2,
          pointHoverRadius: 5,
        },
        {
          label: 'PgBouncer Active Conns',
          data: pgbouncerConns,
          borderColor: '#10b981', // Emerald 500
          backgroundColor: 'rgba(16, 185, 129, 0.05)',
          fill: true,
          tension: 0.4,
          borderWidth: 2,
          pointRadius: 2,
          pointHoverRadius: 5,
        },
      ],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      animation: { duration: 300 },
      plugins: {
        legend: {
          display: true,
          position: 'top',
          align: 'end',
          labels: {
            boxWidth: 8,
            boxHeight: 8,
            usePointStyle: true,
            font: { size: 11 },
          },
        },
        tooltip: {
          backgroundColor: '#111827',
          padding: 8,
          cornerRadius: 6,
        },
      },
      scales: {
        x: {
          grid: { display: false },
          ticks: { font: { size: 10 }, color: '#9ca3af', maxTicksLimit: 6 },
        },
        y: {
          beginAtZero: true,
          grid: { color: '#f3f4f6' },
          ticks: { font: { size: 10 }, color: '#9ca3af' },
        },
      },
    },
  })
}

onMounted(() => {
  createOrUpdateChart()
})

watch(() => props.history, () => {
  createOrUpdateChart()
}, { deep: true })

onBeforeUnmount(() => {
  if (chartInstance) {
    chartInstance.destroy()
  }
})
</script>

<template>
  <div class="rounded-xl border border-gray-200/80 bg-white p-5 shadow-sm">
    <div class="flex items-center justify-between border-b border-gray-100 pb-3">
      <div>
        <h3 class="text-sm font-semibold text-gray-900">CPU Load & Active Pool History</h3>
        <p class="text-xs text-gray-500">Real-time scrape from Prometheus exporter</p>
      </div>
      <span class="inline-flex items-center gap-1.5 rounded-full bg-emerald-50 px-2 py-0.5 text-xs font-medium text-emerald-700">
        <span class="h-1.5 w-1.5 rounded-full bg-emerald-500 animate-ping"></span> Live Scrape
      </span>
    </div>

    <div class="relative mt-4 h-56 w-full">
      <canvas ref="canvasRef"></canvas>
    </div>
  </div>
</template>
