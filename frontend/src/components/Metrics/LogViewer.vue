<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LogItem } from '../../type/typeSupabase'

const props = defineProps<{
  logs: LogItem[]
}>()

const search = ref('')
const selectedStatus = ref<string>('all')

const filteredLogs = computed(() => {
  return props.logs.filter((log) => {
    const req = log.metadata?.[0]?.request?.[0]
    const res = log.metadata?.[0]?.response?.[0]
    const method = req?.method || ''
    const path = req?.path || log.event_message || ''
    const statusCode = res?.status_code?.toString() || ''

    const matchesSearch =
      path.toLowerCase().includes(search.value.toLowerCase()) ||
      method.toLowerCase().includes(search.value.toLowerCase()) ||
      statusCode.includes(search.value)

    if (!matchesSearch) return false

    if (selectedStatus.value === 'error') {
      return (res?.status_code ?? 0) >= 400
    }
    if (selectedStatus.value === 'success') {
      return (res?.status_code ?? 0) >= 200 && (res?.status_code ?? 0) < 400
    }
    return true
  })
})

const getStatusBadge = (code?: number) => {
  if (!code) return 'bg-gray-100 text-gray-700'
  if (code >= 500) return 'bg-rose-100 text-rose-800'
  if (code >= 400) return 'bg-amber-100 text-amber-800'
  if (code >= 300) return 'bg-blue-100 text-blue-800'
  return 'bg-emerald-100 text-emerald-800'
}

const formatTime = (timestamp: number) => {
  if (!timestamp) return '-'
  const d = new Date(timestamp / 1000)
  return d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' })
}
</script>

<template>
  <div class="rounded-xl border border-gray-200/80 bg-white shadow-sm overflow-hidden">
    <div class="flex flex-wrap items-center justify-between gap-3 border-b border-gray-100 p-4">
      <div>
        <h3 class="text-base font-semibold text-gray-900">Incident & Request Activity Logs</h3>
        <p class="text-xs text-gray-500">Live inspection from /analytics/endpoints/logs.all</p>
      </div>

      <div class="flex items-center gap-2">
        <input
          v-model="search"
          type="text"
          placeholder="Filter logs or status..."
          class="rounded-lg border border-gray-200 bg-gray-50 px-3 py-1.5 text-xs text-gray-900 outline-none focus:border-blue-500 focus:bg-white"
        />
        <select
          v-model="selectedStatus"
          class="rounded-lg border border-gray-200 bg-gray-50 px-2.5 py-1.5 text-xs text-gray-700 outline-none"
        >
          <option value="all">All Status</option>
          <option value="success">Success (2xx/3xx)</option>
          <option value="error">Errors (4xx/5xx)</option>
        </select>
      </div>
    </div>

    <div class="overflow-x-auto max-h-96 divide-y divide-gray-100">
      <div
        v-if="filteredLogs.length === 0"
        class="py-12 text-center text-sm text-gray-500"
      >
        No activity logs matching the criteria found.
      </div>
      <div
        v-for="log in filteredLogs"
        :key="log.id"
        class="flex items-center justify-between p-3.5 hover:bg-gray-50/80 transition-colors text-xs font-mono"
      >
        <div class="flex items-center gap-3 min-w-0 pr-4">
          <span
            class="rounded px-2 py-0.5 font-bold uppercase shrink-0"
            :class="getStatusBadge(log.metadata?.[0]?.response?.[0]?.status_code)"
          >
            {{ log.metadata?.[0]?.response?.[0]?.status_code || 200 }}
          </span>
          <span class="font-semibold text-gray-700 shrink-0">
            {{ log.metadata?.[0]?.request?.[0]?.method || 'GET' }}
          </span>
          <span class="truncate text-gray-800" :title="log.metadata?.[0]?.request?.[0]?.path || log.event_message">
            {{ log.metadata?.[0]?.request?.[0]?.path || log.event_message }}
          </span>
        </div>

        <div class="flex items-center gap-4 shrink-0 text-gray-400">
          <span v-if="log.metadata?.[0]?.response?.[0]?.origin_time" class="text-gray-500">
            {{ log.metadata?.[0]?.response?.[0]?.origin_time }}ms
          </span>
          <span>{{ formatTime(log.timestamp) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
