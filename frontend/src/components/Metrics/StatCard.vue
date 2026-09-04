<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  value: number // in bytes
  label: string
  sublabel?: string
  color?: string
}>()

const formatBytes = (bytes: number) => {
  if (!bytes || bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`
}

const formatted = computed(() => formatBytes(props.value))
</script>

<template>
  <div class="rounded-xl border border-gray-200 bg-white p-4 shadow-sm transition-all hover:shadow-md">
    <div class="flex items-center justify-between">
      <span class="text-xs font-medium text-gray-500 uppercase tracking-wider">{{ label }}</span>
      <span
        v-if="color"
        class="h-2.5 w-2.5 rounded-full"
        :style="{ backgroundColor: color }"
      />
    </div>
    <div class="mt-2 text-2xl font-bold tracking-tight text-gray-900">
      {{ formatted }}
    </div>
    <div v-if="sublabel" class="mt-1 text-xs text-gray-500">
      {{ sublabel }}
    </div>
  </div>
</template>
