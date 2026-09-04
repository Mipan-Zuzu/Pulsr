<script setup lang="ts">
import { ref } from 'vue'
import type { StatusHistoryDay } from '../../type/typeSupabase'

defineProps<{
  title: string
  subtitle?: string
  days: StatusHistoryDay[]
  overallUptime?: number
}>()

const hoveredDay = ref<StatusHistoryDay | null>(null)
const mousePos = ref<{ x: number; y: number }>({ x: 0, y: 0 })

const handleMouseEnter = (day: StatusHistoryDay, event: MouseEvent) => {
  hoveredDay.value = day
  const target = event.currentTarget as HTMLElement
  const rect = target.getBoundingClientRect()
  mousePos.value = {
    x: rect.left + rect.width / 2,
    y: rect.top,
  }
}

const handleMouseLeave = () => {
  hoveredDay.value = null
}
</script>

<template>
  <div class="rounded-xl border border-gray-200/80 bg-white p-6 shadow-sm transition-all duration-200">
    <div class="flex flex-wrap items-center justify-between gap-2 pb-4">
      <div>
        <div class="flex items-center gap-2">
          <span class="inline-block h-2 w-2 rounded-full bg-emerald-500 ring-4 ring-emerald-100 animate-pulse"></span>
          <h2 class="text-base font-semibold text-gray-900">{{ title }}</h2>
        </div>
        <p v-if="subtitle" class="text-xs text-gray-500 mt-0.5">{{ subtitle }}</p>
      </div>
      <div class="text-right">
        <span class="text-sm font-semibold text-emerald-600">
          {{ (overallUptime ?? 99.98).toFixed(2) }}% uptime
        </span>
        <p class="text-xs text-gray-400">Past 90 days</p>
      </div>
    </div>

    <!-- Status Strip like status.github.com -->
    <div class="relative mt-2">
      <div class="flex items-center justify-between gap-1 overflow-x-auto py-2">
        <div
          v-for="(day, idx) in days"
          :key="idx"
          @mouseenter="handleMouseEnter(day, $event)"
          @mouseleave="handleMouseLeave"
          :class="[
            'h-9 flex-1 min-w-[3px] rounded-xs cursor-pointer transition-all duration-150 hover:opacity-80 hover:scale-y-125',
            day.status === 'operational'
              ? 'bg-emerald-500'
              : day.status === 'degraded'
              ? 'bg-amber-400'
              : 'bg-rose-500'
          ]"
          :aria-label="`${day.formattedDate}: ${day.status} (${day.uptimePercentage}%)`"
        />
      </div>

      <div class="flex justify-between items-center text-[11px] text-gray-400 mt-1 font-mono">
        <span>90 days ago</span>
        <span class="inline-flex items-center gap-1.5">
          <span class="inline-block w-1.5 h-1.5 rounded-full bg-emerald-500"></span> Operational
        </span>
        <span>Today</span>
      </div>

      <!-- Floating Tooltip -->
      <teleport to="body">
        <div
          v-if="hoveredDay"
          class="fixed pointer-events-none z-50 -translate-x-1/2 -translate-y-full mb-2 rounded-lg bg-gray-900 px-3 py-2 text-xs text-white shadow-xl backdrop-blur-md"
          :style="{
            left: `${mousePos.x}px`,
            top: `${mousePos.y - 8}px`,
          }"
        >
          <div class="font-semibold text-gray-100">{{ hoveredDay.formattedDate }}</div>
          <div class="mt-0.5 flex items-center gap-1.5">
            <span
              class="h-1.5 w-1.5 rounded-full"
              :class="
                hoveredDay.status === 'operational'
                  ? 'bg-emerald-400'
                  : hoveredDay.status === 'degraded'
                  ? 'bg-amber-400'
                  : 'bg-rose-400'
              "
            />
            <span class="capitalize">{{ hoveredDay.status }}</span>
            <span class="text-gray-400">({{ hoveredDay.uptimePercentage }}% uptime)</span>
          </div>
          <div v-if="hoveredDay.incidents.length > 0" class="mt-1 border-t border-gray-800 pt-1 text-[11px] text-rose-300">
            {{ hoveredDay.incidents.join(', ') }}
          </div>
          <div v-else class="text-[10px] text-gray-400 mt-0.5">
            No incidents reported.
          </div>
        </div>
      </teleport>
    </div>
  </div>
</template>
