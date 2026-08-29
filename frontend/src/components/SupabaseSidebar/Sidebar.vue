<script setup lang="ts">
/**
 * Reusable Sidebar Component
 * - Vue 3 + Tailwind CSS
 * - Icons: Iconify (@iconify/vue), pakai icon set "lucide"
 * - Collapsed by default (icon only), expands on hover
 */
import { Icon } from '@iconify/vue'

interface SidebarItem {
  label: string
  href: string
  icon: string
  active?: boolean
}

interface Props {
  items?: SidebarItem[]
  brandLabel?: string
  brandIcon?: string
}

withDefaults(defineProps<Props>(), {
  items: () => [
    { label: 'Projects', href: '#projects', icon: 'lucide:box', active: true },
    { label: 'Team', href: '#team', icon: 'lucide:users' },
    { label: 'Integrations', href: '#integrations', icon: 'lucide:layout-grid' },
    { label: 'Usage', href: '#usage', icon: 'lucide:line-chart' },
    { label: 'Billing', href: '#billing', icon: 'lucide:receipt' },
    { label: 'Organization Settings', href: '#settings', icon: 'lucide:settings' },
  ],
})
</script>

<template>
  <nav
    aria-label="Navigasi utama"
    class="group/sidebar flex h-screen w-16 shrink-0 flex-col overflow-hidden border-r border-gray-200 bg-white transition-[width] duration-200 ease-in-out hover:w-56"
  >

    <div class="flex-1 py-2">
      <ul class="flex flex-col gap-0.5">
        <li v-for="item in items" :key="item.label">
          <a
            :href="item.href"
            :aria-current="item.active ? 'page' : undefined"
            class="mx-2.5 flex h-10 items-center gap-3 rounded-lg px-2.5 text-gray-500 transition-colors duration-200 hover:bg-gray-50 hover:text-gray-900"
            :class="item.active && 'bg-gray-100 font-semibold text-gray-900'"
          >
            <Icon :icon="item.icon" class="h-5 w-5 shrink-0" />
            <span
              class="whitespace-nowrap text-sm opacity-0 transition-opacity duration-200 group-hover/sidebar:opacity-100"
            >
              {{ item.label }}
            </span>
          </a>
        </li>
      </ul>
    </div>

    <div class="shrink-0 py-4">
      <a
        href="#collapse"
        class="mx-2.5 flex h-10 items-center gap-3 rounded-lg px-2.5 text-gray-500 transition-colors duration-200 hover:bg-gray-50 hover:text-gray-900"
      >
        <Icon icon="lucide:panel-left" class="h-5 w-5 shrink-0" />
        <span
          class="whitespace-nowrap text-sm opacity-0 transition-opacity duration-200 group-hover/sidebar:opacity-100"
        >
          Collapse
        </span>
      </a>
    </div>
  </nav>
</template>