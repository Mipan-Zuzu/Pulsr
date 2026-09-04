<script setup lang="ts">
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { supabaseProject, useSupabaseMonitorStore } from '../../store/supabase'
import { Redis } from '../../store/redis'
import Sidebar from '../../components/SupabaseSidebar/Sidebar.vue'
import SupabaseMenu from '../../components/SupabaseMenu/SupabaseMenu.vue'

const projectStore = supabaseProject()
const monitorStore = useSupabaseMonitorStore()
const redisStore = Redis()
const { projectDetail, loadingProject, projectDetailError } = storeToRefs(projectStore)

const quickStart = async (id: string) => {
  try {
    await monitorStore.startProject(id, redisStore.authKey)
    alert(`Project ${id} start/restore command sent!`)
    await projectStore.allproject(redisStore.authKey)
  } catch (err: any) {
    alert(err?.response?.data?.message || 'Failed to start project')
  }
}

const quickPause = async (id: string) => {
  if (!confirm(`Are you sure you want to PAUSE project ${id}?`)) return
  try {
    await monitorStore.pauseProject(id, redisStore.authKey)
    alert(`Project ${id} pause command sent!`)
    await projectStore.allproject(redisStore.authKey)
  } catch (err: any) {
    alert(err?.response?.data?.message || 'Failed to pause project')
  }
}

onMounted(async () => {
  await redisStore.redisFind()
  await projectStore.allproject(redisStore.authKey)
})
</script>

<template>
  <div class="flex h-screen flex-col bg-[#fafbfc]">
    <!-- Topbar: full width -->
    <SupabaseMenu class="w-full flex-shrink-0" />

    <!-- Baris bawah: sidebar overlay + area konten -->
    <div class="flex flex-1 overflow-hidden">
      <Sidebar class="z-50 fixed" />

      <main class="flex-1 overflow-y-auto p-6 md:p-10 md:ml-20 max-w-7xl">
        <div class="flex flex-wrap items-center justify-between gap-4 mb-8">
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold tracking-tight text-gray-950">Projects</h1>
            <p class="text-sm text-gray-500 mt-1">Select a project to monitor live health, metrics & status history</p>
          </div>
          <button
            @click="projectStore.allproject(redisStore.authKey)"
            class="rounded-lg border border-gray-200 bg-white px-3.5 py-1.5 text-xs font-medium text-gray-700 hover:bg-gray-50 transition-colors shadow-xs"
          >
            Refresh Projects
          </button>
        </div>

        <div v-if="loadingProject" class="py-16 text-center text-sm text-gray-500">
          Loading organization projects...
        </div>

        <div v-else-if="projectDetailError" class="rounded-xl border border-rose-200 bg-rose-50 p-4 text-sm text-rose-700">
          {{ projectDetailError }}
        </div>

        <section v-else class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <RouterLink
            v-for="project in projectDetail?.data || []"
            :key="project.id"
            :to="`/supabase/project/${project.id}/status`"
            class="group flex flex-col justify-between rounded-xl border border-gray-200/90 bg-white p-5 shadow-xs transition-all hover:border-gray-300 hover:shadow-lg hover:-translate-y-1"
          >
            <div>
              <div class="flex items-center justify-between">
                <h2 class="font-bold text-lg text-gray-900 group-hover:text-blue-600 transition-colors">
                  {{ project.name }}
                </h2>
                <span
                  class="inline-flex items-center gap-1.5 rounded-full px-2 py-0.5 text-xs font-medium"
                  :class="
                    project.status === 'ACTIVE_HEALTHY' || !project.status
                      ? 'bg-emerald-50 text-emerald-700'
                      : 'bg-amber-50 text-amber-700'
                  "
                >
                  <span
                    class="h-1.5 w-1.5 rounded-full"
                    :class="project.status === 'ACTIVE_HEALTHY' || !project.status ? 'bg-emerald-500' : 'bg-amber-500'"
                  />
                  {{ project.status || 'ACTIVE' }}
                </span>
              </div>

              <p class="text-xs text-gray-500 mt-2 flex items-center gap-2">
                <span>Region: <strong class="text-gray-700 font-mono">{{ project.region || 'ap-southeast' }}</strong></span>
                <span>&bull;</span>
                <span>Ref: <strong class="text-gray-700 font-mono">{{ project.ref || project.id }}</strong></span>
              </p>

              <div class="mt-4 rounded-lg bg-gray-50 p-2.5 text-xs text-gray-600 border border-gray-100 flex items-center justify-between">
                <span class="text-gray-500 font-mono">Engine: {{ project.database?.postgres_engine || 'PostgreSQL' }}</span>
                <span class="text-blue-600 font-semibold group-hover:underline flex items-center gap-1">
                  View Status &rarr;
                </span>
              </div>
            </div>

            <!-- Quick Action controls: Start & Pause directly from card -->
            <div class="mt-3 flex items-center gap-2 pt-2 border-t border-gray-100">
              <button
                @click.prevent="quickStart(project.id)"
                class="flex-1 rounded-md border border-emerald-200 bg-emerald-50 px-2 py-1 text-[11px] font-semibold text-emerald-700 hover:bg-emerald-100 transition-colors"
              >
                ▶ Start
              </button>
              <button
                @click.prevent="quickPause(project.id)"
                class="flex-1 rounded-md border border-amber-200 bg-amber-50 px-2 py-1 text-[11px] font-semibold text-amber-800 hover:bg-amber-100 transition-colors"
              >
                ⏸ Pause
              </button>
            </div>

            <div class="mt-2 flex items-center justify-between pt-1 text-[11px] text-gray-400">
              <span>90-Day Status</span>
              <span class="text-emerald-600 font-medium">100% Operational</span>
            </div>
          </RouterLink>
        </section>
      </main>
    </div>
  </div>
</template>
