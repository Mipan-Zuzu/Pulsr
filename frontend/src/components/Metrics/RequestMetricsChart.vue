<script setup lang="ts">
import { ref, onMounted, watch, onBeforeUnmount } from 'vue'
import {
  Chart,
  LineController,
  LineElement,
  PointElement,
  LinearScale,
  Title,
  CategoryScale,
  Tooltip,
  Legend,
  Filler,
  BarController,
  BarElement
} from 'chart.js'
import type { AnalyticsCount } from '../../type/typeSupabase'

Chart.register(
  LineController,
  LineElement,
  PointElement,
  LinearScale,
  Title,
  CategoryScale,
  Tooltip,
  Legend,
  Filler,
  BarController,
  BarElement
)

const props = withDefaults(
  defineProps<{
    data: AnalyticsCount[]
    timeframe?: string
  }>(),
  {
    timeframe: '24h'
  }
)

const canvasRef = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

const chartType = ref<'line' | 'bar'>('line')
const activeMetric = ref<'all' | 'rest' | 'auth' | 'storage' | 'realtime'>('all')

const createOrUpdateChart = () => {
  if (!canvasRef.value) return

  // Build labels and datasets
  const labels = props.data.map((item) => {
    try {
      const d = new Date(item.timestamp)
      return d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    } catch {
      return item.timestamp
    }
  })

  // Dataviz calibrated distinct accessible palette
  const datasets: any[] = []

  if (activeMetric.value === 'all' || activeMetric.value === 'rest') {
    datasets.push({
      label: 'REST API',
      data: props.data.map((d) => d.total_rest_requests || 0),
      borderColor: '#2563eb', // Blue
      backgroundColor: chartType.value === 'line' ? 'rgba(37, 99, 235, 0.08)' : 'rgba(37, 99, 235, 0.85)',
      fill: chartType.value === 'line',
      tension: 0.35,
      borderWidth: 2,
      pointRadius: props.data.length > 30 ? 0 : 3,
      pointHoverRadius: 6,
    })
  }

  if (activeMetric.value === 'all' || activeMetric.value === 'auth') {
    datasets.push({
      label: 'Auth Requests',
      data: props.data.map((d) => d.total_auth_requests || 0),
      borderColor: '#059669', // Emerald
      backgroundColor: chartType.value === 'line' ? 'rgba(5, 150, 105, 0.08)' : 'rgba(5, 150, 105, 0.85)',
      fill: chartType.value === 'line',
      tension: 0.35,
      borderWidth: 2,
      pointRadius: props.data.length > 30 ? 0 : 3,
      pointHoverRadius: 6,
    })
  }

  if (activeMetric.value === 'all' || activeMetric.value === 'storage') {
    datasets.push({
      label: 'Storage API',
      data: props.data.map((d) => d.total_storage_requests || 0),
      borderColor: '#d97706', // Amber
      backgroundColor: chartType.value === 'line' ? 'rgba(217, 119, 6, 0.08)' : 'rgba(217, 119, 6, 0.85)',
      fill: chartType.value === 'line',
      tension: 0.35,
      borderWidth: 2,
      pointRadius: props.data.length > 30 ? 0 : 3,
      pointHoverRadius: 6,
    })
  }

  if (activeMetric.value === 'all' || activeMetric.value === 'realtime') {
    datasets.push({
      label: 'Realtime WebSocket',
      data: props.data.map((d) => d.total_realtime_requests || 0),
      borderColor: '#7c3aed', // Purple
      backgroundColor: chartType.value === 'line' ? 'rgba(124, 58, 237, 0.08)' : 'rgba(124, 58, 237, 0.85)',
      fill: chartType.value === 'line',
      tension: 0.35,
      borderWidth: 2,
      pointRadius: props.data.length > 30 ? 0 : 3,
      pointHoverRadius: 6,
    })
  }

  if (chartInstance) {
    chartInstance.destroy()
  }

  chartInstance = new Chart(canvasRef.value, {
    type: chartType.value,
    data: {
      labels,
      datasets,
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      interaction: {
        mode: 'index',
        intersect: false,
      },
      plugins: {
        legend: {
          display: true,
          position: 'top',
          align: 'end',
          labels: {
            usePointStyle: true,
            boxWidth: 8,
            boxHeight: 8,
            font: {
              size: 11,
              family: 'system-ui, sans-serif',
            },
            color: '#4b5563',
          },
        },
        tooltip: {
          backgroundColor: '#111827',
          titleColor: '#f9fafb',
          bodyColor: '#e5e7eb',
          borderColor: '#374151',
          borderWidth: 1,
          padding: 10,
          cornerRadius: 8,
          usePointStyle: true,
        },
      },
      scales: {
        x: {
          grid: {
            display: false,
          },
          ticks: {
            font: { size: 10 },
            color: '#9ca3af',
            maxRotation: 0,
            autoSkip: true,
            maxTicksLimit: 8,
          },
        },
        y: {
          beginAtZero: true,
          grid: {
            color: '#f3f4f6',
          },
          ticks: {
            font: { size: 10 },
            color: '#9ca3af',
          },
        },
      },
    },
  })
}

onMounted(() => {
  createOrUpdateChart()
})

watch([() => props.data, chartType, activeMetric], () => {
  createOrUpdateChart()
}, { deep: true })

onBeforeUnmount(() => {
  if (chartInstance) {
    chartInstance.destroy()
  }
})
</script>

<template>
  <div class="rounded-xl border border-gray-200/80 bg-white p-6 shadow-sm">
    <div class="flex flex-wrap items-center justify-between gap-4 border-b border-gray-100 pb-4">
      <div>
        <h3 class="text-base font-semibold text-gray-900">Endpoint Request Volume & Activity</h3>
        <p class="text-xs text-gray-500">Live analytics counts across REST, Auth, Realtime, and Storage</p>
      </div>

      <div class="flex flex-wrap items-center gap-2">
        <!-- Metric filter pills -->
        <div class="inline-flex rounded-lg border border-gray-200 bg-gray-50/70 p-0.5 text-xs">
          <button
            @click="activeMetric = 'all'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              activeMetric === 'all' ? 'bg-white text-gray-900 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            All
          </button>
          <button
            @click="activeMetric = 'rest'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              activeMetric === 'rest' ? 'bg-white text-blue-600 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            REST
          </button>
          <button
            @click="activeMetric = 'auth'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              activeMetric === 'auth' ? 'bg-white text-emerald-600 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            Auth
          </button>
          <button
            @click="activeMetric = 'storage'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              activeMetric === 'storage' ? 'bg-white text-amber-600 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            Storage
          </button>
          <button
            @click="activeMetric = 'realtime'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              activeMetric === 'realtime' ? 'bg-white text-purple-600 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            Realtime
          </button>
        </div>

        <!-- Chart Type Switcher -->
        <div class="inline-flex rounded-lg border border-gray-200 bg-gray-50/70 p-0.5 text-xs">
          <button
            @click="chartType = 'line'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              chartType === 'line' ? 'bg-white text-gray-900 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            Line
          </button>
          <button
            @click="chartType = 'bar'"
            :class="[
              'rounded-md px-2.5 py-1 font-medium transition-colors',
              chartType === 'bar' ? 'bg-white text-gray-900 shadow-xs' : 'text-gray-500 hover:text-gray-900'
            ]"
          >
            Bar
          </button>
        </div>
      </div>
    </div>

    <div class="relative mt-4 h-72 w-full">
      <canvas ref="canvasRef"></canvas>
    </div>
  </div>
</template>
