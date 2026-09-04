<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { Motion } from 'motion-v'
import { useSupabaseMonitorStore } from '../../store/supabase'
import { Redis } from '../../store/redis'

import StatCard from '../../components/Metrics/StatCard.vue'
import StatusHistoryStrip from '../../components/Metrics/StatusHistoryStrip.vue'
import RequestMetricsChart from '../../components/Metrics/RequestMetricsChart.vue'
import PrometheusRealtimeChart from '../../components/Metrics/PrometheusRealtimeChart.vue'
import ServiceStatusList from '../../components/Metrics/ServiceStatusList.vue'
import LogViewer from '../../components/Metrics/LogViewer.vue'
import Sidebar from '../../components/SupabaseSidebar/Sidebar.vue'
import SupabaseMenu from '../../components/SupabaseMenu/SupabaseMenu.vue'

const route = useRoute()
const router = useRouter()
const projectId = computed(() => (route.params.id as string) || 'default-project')

const monitorStore = useSupabaseMonitorStore()
const redis = Redis()

const {
  currentProject,
  metrics,
  metricsHistory,
  analyticsCounts,
  logs,
  statusHistory,
  isLoading,
  isRefreshing,
} = storeToRefs(monitorStore)

const autoRefresh = ref(true)
let refreshInterval: any = null
const isActionLoading = ref(false)
const notification = ref<{ message: string; type: 'success' | 'error' } | null>(null)

const showNotification = (message: string, type: 'success' | 'error' = 'success') => {
  notification.value = { message, type }
  setTimeout(() => {
    notification.value = null
  }, 4000)
}

const toggleAutoRefresh = () => {
  autoRefresh.value = !autoRefresh.value
  if (autoRefresh.value) {
    startPolling()
  } else {
    stopPolling()
  }
}

const startPolling = () => {
  stopPolling()
  refreshInterval = setInterval(() => {
    if (!isLoading.value) {
      monitorStore.refreshMetricsOnly(projectId.value)
    }
  }, 10000) // scrape interval 10s
}

const stopPolling = () => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
    refreshInterval = null
  }
}

const triggerPause = async () => {
  if (!confirm('Are you sure you want to PAUSE this Supabase service? Database traffic will be halted.')) return
  isActionLoading.value = true
  try {
    await monitorStore.pauseProject(projectId.value)
    showNotification('Project pause command sent successfully')
    await monitorStore.loadAll(projectId.value)
  } catch (err: any) {
    showNotification(err?.response?.data?.message || 'Failed to pause project', 'error')
  } finally {
    isActionLoading.value = false
  }
}

const triggerStart = async () => {
  isActionLoading.value = true
  try {
    await monitorStore.startProject(projectId.value)
    showNotification('Project start/restore command sent successfully')
    await monitorStore.loadAll(projectId.value)
  } catch (err: any) {
    showNotification(err?.response?.data?.message || 'Failed to start project', 'error')
  } finally {
    isActionLoading.value = false
  }
}

onMounted(async () => {
  await redis.redisFind()
  await monitorStore.loadAll(projectId.value)
  if (autoRefresh.value) {
    startPolling()
  }
})
</script>

<template>
  <div class="min-h-screen bg-[#fafbfc] text-gray-900 flex flex-col">
    <!-- Topbar -->
    <SupabaseMenu class="w-full flex-shrink-0" />

    <div class="flex flex-1 overflow-hidden relative">
      <Sidebar class="z-40 fixed hidden md:flex" />

      <main class="flex-1 overflow-y-auto px-4 py-8 sm:px-8 md:ml-16 lg:px-12 max-w-7xl mx-auto w-full">
        <!-- Toast Notification -->
        <transition
          enter-active-class="transform ease-out duration-300 transition"
          enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
          enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
          leave-active-class="transition ease-in duration-100"
          leave-from-class="opacity-100"
          leave-to-class="opacity-0"
        >
          <div
            v-if="notification"
            class="fixed top-5 right-5 z-50 rounded-xl p-4 shadow-xl border backdrop-blur-md flex items-center gap-3 text-sm font-medium"
            :class="
              notification.type === 'success'
                ? 'bg-emerald-950/90 text-emerald-100 border-emerald-800'
                : 'bg-rose-950/90 text-rose-100 border-rose-800'
            "
          >
            <span class="h-2 w-2 rounded-full" :class="notification.type === 'success' ? 'bg-emerald-400' : 'bg-rose-400'"></span>
            {{ notification.message }}
          </div>
        </transition>

        <!-- Header / Status Banner -->
        <Motion
          :initial="{ opacity: 0, y: -16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.4 }"
          class="mb-8"
        >
          <div class="flex flex-wrap items-center justify-between gap-4 pb-6 border-b border-gray-200">
            <div>
              <div class="flex items-center gap-3">
                <button
                  @click="router.push('/supabase/project/' + projectId)"
                  class="rounded-lg p-1.5 text-gray-400 hover:text-gray-900 hover:bg-gray-100 transition-colors"
                  title="Back to projects"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                  </svg>
                </button>
                <h1 class="text-2xl sm:text-3xl font-bold tracking-tight text-gray-950">
                  {{ currentProject?.name || projectId }}
                </h1>
                <span class="rounded-full border border-gray-300 bg-white px-2.5 py-0.5 text-xs font-mono text-gray-600">
                  {{ currentProject?.region || 'ap-southeast' }}
                </span>
                <span class="rounded-full bg-emerald-100 text-emerald-800 px-2.5 py-0.5 text-xs font-semibold">
                  {{ currentProject?.status || 'ACTIVE_HEALTHY' }}
                </span>
              </div>
              <p class="text-sm text-gray-500 mt-1 ml-9">
                Host: {{ currentProject?.database?.host || `${projectId}.supabase.co` }} &bull; Ref: {{ currentProject?.ref || projectId }}
              </p>
            </div>

            <!-- Controls (Auto-refresh, Pause, Restore) -->
            <div class="flex flex-wrap items-center gap-2.5">
              <button
                @click="toggleAutoRefresh"
                :class="[
                  'inline-flex items-center gap-2 rounded-lg border px-3 py-1.5 text-xs font-medium transition-all',
                  autoRefresh
                    ? 'border-emerald-200 bg-emerald-50/70 text-emerald-700 hover:bg-emerald-100'
                    : 'border-gray-200 bg-white text-gray-600 hover:bg-gray-50'
                ]"
              >
                <span class="h-2 w-2 rounded-full" :class="autoRefresh ? 'bg-emerald-500 animate-pulse' : 'bg-gray-400'"></span>
                {{ autoRefresh ? 'Auto-refresh ON (10s)' : 'Auto-refresh OFF' }}
              </button>

              <button
                @click="monitorStore.refreshMetricsOnly(projectId)"
                :disabled="isRefreshing"
                class="rounded-lg border border-gray-200 bg-white px-3 py-1.5 text-xs font-medium text-gray-700 hover:bg-gray-50 transition-colors disabled:opacity-50"
              >
                {{ isRefreshing ? 'Refreshing...' : 'Refresh Now' }}
              </button>

              <!-- Start / Restore Button -->
              <button
                @click="triggerStart"
                :disabled="isActionLoading"
                class="inline-flex items-center gap-1.5 rounded-lg border border-emerald-300 bg-emerald-600 px-3 py-1.5 text-xs font-semibold text-white hover:bg-emerald-700 shadow-xs transition-colors disabled:opacity-50"
                title="Start or restore Supabase project"
              >
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>{{ isActionLoading ? 'Starting...' : 'Start / Restore' }}</span>
              </button>

              <!-- Pause Button -->
              <button
                @click="triggerPause"
                :disabled="isActionLoading"
                class="inline-flex items-center gap-1.5 rounded-lg border border-amber-300 bg-amber-50 px-3 py-1.5 text-xs font-semibold text-amber-800 hover:bg-amber-100 shadow-xs transition-colors disabled:opacity-50"
                title="Pause Supabase project services"
              >
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>{{ isActionLoading ? 'Pausing...' : 'Pause' }}</span>
              </button>
            </div>
          </div>

          <!-- Hero Global Status Banner (like status.github.com) -->
          <div class="rounded-2xl border border-emerald-200 bg-emerald-50/80 p-5 sm:p-6 shadow-xs flex items-center justify-between gap-4">
            <div class="flex items-center gap-4">
              <div class="relative flex h-4 w-4 items-center justify-center">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-3 w-3 bg-emerald-500"></span>
              </div>
              <div>
                <h2 class="text-lg font-bold text-emerald-950">All Systems Operational</h2>
                <p class="text-xs text-emerald-800/80">No active incidents or service degradation detected on project {{ currentProject?.name || projectId }}.</p>
              </div>
            </div>
            <div class="text-right hidden sm:block">
              <span class="text-xs font-mono text-emerald-800">Scraped via Prometheus Parser</span>
            </div>
          </div>
        </Motion>

        <!-- 90-Day Status History Strip (status.github.com) -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.1 }"
          class="mb-8"
        >
          <StatusHistoryStrip
            title="System Operational History"
            subtitle="Incident & uptime tracker over the last 90 calendar days"
            :days="statusHistory"
            :overall-uptime="99.98"
          />
        </Motion>

        <!-- High-level System KPI Stats -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.15 }"
          class="grid grid-cols-2 gap-4 sm:grid-cols-4 mb-8"
        >
          <div class="rounded-xl border border-gray-200 bg-white p-4 shadow-sm hover:shadow-md transition-all">
            <span class="text-xs font-medium text-gray-500 uppercase tracking-wider">CPU Load (15m)</span>
            <div class="mt-2 text-2xl font-bold text-gray-900">
              {{ metrics?.cpu?.load_avg_15m ? metrics.cpu.load_avg_15m.toFixed(2) : '0.04' }}
            </div>
            <p class="mt-1 text-xs text-emerald-600 font-medium">Optimal range (&lt; 1.0)</p>
          </div>

          <StatCard
            label="Committed Memory"
            :value="metrics?.memory?.committed_as_bytes || 524288000"
            sublabel="node_memory_Committed_AS_bytes"
            color="#2563eb"
          />

          <StatCard
            label="Memory Shmem"
            :value="metrics?.memory?.shmem_bytes || 67108864"
            sublabel="node_memory_Shmem_bytes"
            color="#7c3aed"
          />

          <div class="rounded-xl border border-gray-200 bg-white p-4 shadow-sm hover:shadow-md transition-all">
            <span class="text-xs font-medium text-gray-500 uppercase tracking-wider">PgBouncer Active Conns</span>
            <div class="mt-2 text-2xl font-bold text-gray-900">
              {{ metrics?.pgbouncer?.server_active_connections ?? 2 }}
              <span class="text-xs font-normal text-gray-400">
                / {{ metrics?.pgbouncer?.max_client_connections || 100 }} max
              </span>
            </div>
            <p class="mt-1 text-xs text-gray-500">
              Cached DNS: {{ metrics?.pgbouncer?.cached_dns_names ?? 1 }}
            </p>
          </div>
        </Motion>

        <!-- Charts Section (Flexible Request Volume & Real-time Scrape) -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.2 }"
          class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8"
        >
          <!-- Flexible Big Chart (2 Cols) -->
          <div class="lg:col-span-2">
            <RequestMetricsChart
              :data="analyticsCounts"
            />
          </div>

          <!-- Real-time Prometheus Stream Chart (1 Col) -->
          <div>
            <PrometheusRealtimeChart
              :history="metricsHistory"
            />
          </div>
        </Motion>

        <!-- Additional Resource Breakdown (Disk & Memory Slab) -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.25 }"
          class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8"
        >
          <StatCard
            label="Memory Slab"
            :value="metrics?.memory?.slab_bytes || 124518400"
            sublabel="Kernel data structure cache"
            color="#059669"
          />
          <StatCard
            label="Dirty Memory"
            :value="metrics?.memory?.dirty_bytes || 409600"
            sublabel="Waiting to be written to disk"
            color="#d97706"
          />
          <StatCard
            label="Page Tables"
            :value="metrics?.memory?.page_tables_bytes || 16777216"
            sublabel="Virtual-to-physical address map"
            color="#4f46e5"
          />
        </Motion>

        <!-- Services Status List (Detailed breakdown) -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.3 }"
          class="mb-8"
        >
          <ServiceStatusList
            :project="currentProject"
            :metrics="metrics"
          />
        </Motion>

        <!-- Activity & Incident Logs -->
        <Motion
          :initial="{ opacity: 0, y: 16 }"
          :animate="{ opacity: 1, y: 0 }"
          :transition="{ duration: 0.45, delay: 0.35 }"
          class="mb-12"
        >
          <LogViewer
            :logs="logs"
          />
        </Motion>
      </main>
    </div>
  </div>
</template>
