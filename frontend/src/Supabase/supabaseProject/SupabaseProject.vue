<script setup lang="ts">
import Sidebar from '../../components/SupabaseSidebar/Sidebar.vue'
import SupabaseMenu from '../../components/SupabaseMenu/SupabaseMenu.vue'

interface Project {
  id: string
  name: string
  provider: string
  region: string
  tier: string
}

const projects: Project[] = [
  {
    id: 'pulsr',
    name: 'Pulsr',
    provider: 'AWS',
    region: 'ap-southeast-2',
    tier: 'NANO',
  },
]
</script>

<template>
  <div class="flex h-screen flex-col">
    <!-- Topbar: full width, di luar konteks flex-row sidebar -->
    <SupabaseMenu class="w-full flex-shrink-0" />

    <!-- Baris bawah: sidebar overlay (fixed) + area konten -->
    <div class="flex flex-1 overflow-hidden">
      <Sidebar class="z-50 fixed" />

      <main class="flex-1 overflow-y-auto p-10 md:ml-20">
        <h1 class="text-2xl font-semibold mb-10">Projects</h1>
        <section class="grid grid-cols-1 gap-10 md:grid-cols-3">
          <RouterLink
            v-for="project in projects"
            :key="project.id"
            :to="`/projects/${project.id}`"
            class="flex flex-col gap-2 rounded-lg border p-2 px-10 transition-colors hover:bg-gray-50 hover:scale-105 hover:shadow-2xl"
          >
            <span class="flex flex-col gap-2">
              <h1 class="font-semibold">{{ project.name }}</h1>
              <p class="text-gray-500">
                <span>{{ project.provider }}</span> | <span>{{ project.region }}</span>
              </p>
            </span>
            <div class="w-fit rounded-2xl border px-2 py-0.5">
              <h1 class="text-sm">{{ project.tier }}</h1>
            </div>
          </RouterLink>
        </section>
      </main>
    </div>
  </div>
</template>