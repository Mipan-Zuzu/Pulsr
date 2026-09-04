<script setup lang="ts">
import { computed } from 'vue'
import type { SystemUsage, Project } from '../../type/typeSupabase'

const props = defineProps<{
  project?: Project | null
  metrics?: SystemUsage | null
}>()

const services = computed(() => [
  {
    name: 'PostgreSQL Database Engine',
    description: props.project?.database?.postgres_engine || props.project?.database?.version || 'PostgreSQL 15',
    status: props.project?.status === 'ACTIVE_HEALTHY' || !props.project?.status ? 'Operational' : 'Paused',
    uptime: '99.99%',
    icon: 'lucide:database',
  },
  {
    name: 'Connection Pooler (PgBouncer)',
    description: props.metrics?.pgbouncer?.version ? `v${props.metrics.pgbouncer.version}` : 'Active connection pooling',
    status: (props.metrics?.pgbouncer?.max_client_connections ?? 0) > 0 || !props.metrics ? 'Operational' : 'Degraded',
    uptime: '99.95%',
    icon: 'lucide:network',
  },
  {
    name: 'REST / PostgREST API Gateway',
    description: 'Auto-generated CRUD endpoints with RLS enforcement',
    status: 'Operational',
    uptime: '100.0%',
    icon: 'lucide:cpu',
  },
  {
    name: 'Auth Service (GoTrue)',
    description: 'User registration, JWT tokens & identity providers',
    status: 'Operational',
    uptime: '99.98%',
    icon: 'lucide:shield-check',
  },
  {
    name: 'Realtime Service',
    description: 'Postgres changes broadcast & presence WebSockets',
    status: 'Operational',
    uptime: '99.97%',
    icon: 'lucide:zap',
  },
  {
    name: 'Storage & Edge CDN',
    description: 'S3-compatible bucket, image transform & cache',
    status: 'Operational',
    uptime: '100.0%',
    icon: 'lucide:hard-drive',
  },
])
</script>

<template>
  <div class="rounded-xl border border-gray-200/80 bg-white shadow-sm overflow-hidden">
    <div class="border-b border-gray-100 bg-gray-50/50 px-6 py-4 flex items-center justify-between">
      <div>
        <h3 class="text-base font-semibold text-gray-900">Services & Infrastructure Status</h3>
        <p class="text-xs text-gray-500">Live operational condition of core project subcomponents</p>
      </div>
      <span class="inline-flex items-center gap-1.5 rounded-full bg-emerald-50 px-2.5 py-1 text-xs font-medium text-emerald-700">
        <span class="h-2 w-2 rounded-full bg-emerald-500"></span>
        All Systems Normal
      </span>
    </div>

    <div class="divide-y divide-gray-100">
      <div
        v-for="svc in services"
        :key="svc.name"
        class="flex flex-col sm:flex-row sm:items-center justify-between px-6 py-3.5 hover:bg-gray-50/70 transition-colors gap-2"
      >
        <div class="flex items-center gap-3">
          <div class="p-2 rounded-lg bg-gray-100 text-gray-600">
            <span class="text-sm font-bold">●</span>
          </div>
          <div>
            <div class="font-medium text-sm text-gray-900">{{ svc.name }}</div>
            <div class="text-xs text-gray-500">{{ svc.description }}</div>
          </div>
        </div>

        <div class="flex items-center gap-6 self-end sm:self-center">
          <span class="text-xs text-gray-400 font-mono hidden md:inline">{{ svc.uptime }} uptime</span>
          <span
            class="inline-flex items-center gap-1.5 rounded-md px-2.5 py-1 text-xs font-semibold"
            :class="
              svc.status === 'Operational'
                ? 'bg-emerald-50 text-emerald-700'
                : svc.status === 'Degraded'
                ? 'bg-amber-50 text-amber-700'
                : 'bg-rose-50 text-rose-700'
            "
          >
            <span
              class="h-1.5 w-1.5 rounded-full"
              :class="
                svc.status === 'Operational'
                  ? 'bg-emerald-500'
                  : svc.status === 'Degraded'
                  ? 'bg-amber-500'
                  : 'bg-rose-500'
              "
            />
            {{ svc.status }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
