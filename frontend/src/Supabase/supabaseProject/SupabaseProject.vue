<script setup lang="ts">
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from 'vue';
import Sidebar from 'primevue/sidebar';
import { Icon } from '@iconify/vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { supabaseDetail } from '../../store/supabase';
import Skeleton from 'primevue/skeleton';

const route = useRoute()
const supabasePorjectDetail = supabaseDetail()

const id = route.params.id as string

const {loadingProject} = storeToRefs(supabaseDetail())

onBeforeMount(() => {
    supabasePorjectDetail.allprojectData(id)
})

const isMobile = ref(false);
const open = ref(false);
let mql: MediaQueryList | null = null;
let onMqlChange: ((e : MediaQueryListEvent) => void) | null = null;

const stats = [
    { label: 'Database Size', value: '1.2 GB', trend: 4, icon: 'lucide:database' },
    { label: 'API Requests', value: '84.3k', trend: 12, icon: 'lucide:activity' },
    { label: 'Active Connections', value: '23', trend: -6, icon: 'lucide:plug' },
    { label: 'Storage Used', value: '340 MB', trend: 2, icon: 'lucide:hard-drive' },
];

onMounted(() => {
    if (typeof window === 'undefined') return;

    mql = window.matchMedia('(max-width: 1023px)');
    isMobile.value = mql.matches;
    onMqlChange = (event) => {
        isMobile.value = event.matches;
    };
    mql.addEventListener('change', onMqlChange);
});

onBeforeUnmount(() => {
    if (mql && onMqlChange) {
        mql.removeEventListener('change', onMqlChange);
    }
})


</script>

<template>
    <section class="border border-surface-200 dark:border-surface-700 rounded-lg overflow-hidden">
        <div class="min-h-192 relative">
            <div
                v-if="isMobile && open"
                class="absolute inset-0 bg-black/20 z-40"
                @click="open = false"
            />

            <Sidebar
                v-model:visible="open"
                :position="isMobile ? 'left' : 'left'"
                :modal="isMobile"
                :show-close-icon="false"
                class="w-48"
            >
                <div class="flex h-full flex-col bg-surface-0 dark:bg-surface-900">
                    <div class="flex items-center gap-2 border-b border-surface-200 dark:border-surface-700 px-4 py-3">
                        <div class="flex size-6 shrink-0 items-center justify-center rounded-md bg-stone-100 text-stone-700 text-xs font-bold leading-none">
                            <Icon icon="tabler:brand-cake" />
                        </div>
                        <span class="font-semibold text-sm">Pulsr Demo</span>
                    </div>

                    <nav class="flex flex-col gap-1 p-3">
                        <button class="flex items-center gap-3 rounded-md bg-surface-100 dark:bg-surface-800 px-3 py-2 text-left text-sm font-medium text-surface-900 dark:text-surface-0">
                            <Icon icon="lucide:home" />
                            <span>Home</span>
                        </button>
                        <button class="flex items-center gap-3 rounded-md px-3 py-2 text-left text-sm text-surface-600 dark:text-surface-300 hover:bg-surface-100 dark:hover:bg-surface-800">
                            <Icon icon="lucide:database" />
                            <span>Database</span>
                        </button>
                        <button class="flex items-center gap-3 rounded-md px-3 py-2 text-left text-sm text-surface-600 dark:text-surface-300 hover:bg-surface-100 dark:hover:bg-surface-800">
                            <Icon icon="lucide:key" />
                            <span>Authentication</span>
                        </button>
                        <button class="flex items-center gap-3 rounded-md px-3 py-2 text-left text-sm text-surface-600 dark:text-surface-300 hover:bg-surface-100 dark:hover:bg-surface-800">
                            <Icon icon="lucide:logs" />
                            <span>Edge Functions</span>
                        </button>
                    </nav>
                </div>
            </Sidebar>

            <main v-if="loadingProject == true" class="min-h-192">
                <header class="flex h-12 items-center gap-2 border-b border-surface-200 dark:border-surface-700 px-4">
                    <button v-if="isMobile" class="inline-flex items-center justify-center rounded-md border border-surface-200 bg-surface-0 px-2 py-1 text-sm dark:border-surface-700 dark:bg-surface-900" @click="open = true">
                        <Icon icon="lucide:panel-left" />
                    </button>
                    <span class="text-sm font-medium">Suapabse </span>
                    <span class="text-xs text-muted-color">| Monitoring</span>
                </header>
                <div class="flex-1 p-4 flex flex-col gap-4 overflow-y-auto">
                    <div class="grid grid-cols-2 lg:grid-cols-4 gap-3">
                        <Skeleton />
                        <Skeleton />
                        <Skeleton />
                        <Skeleton />
                    </div>
                </div>
            </main>

            <main v-else class="min-h-192">
                <header class="flex h-12 items-center gap-2 border-b border-surface-200 dark:border-surface-700 px-4">
                    <button v-if="isMobile" class="inline-flex items-center justify-center rounded-md border border-surface-200 bg-surface-0 px-2 py-1 text-sm dark:border-surface-700 dark:bg-surface-900" @click="open = true">
                        <Icon icon="lucide:panel-left" />
                    </button>
                    <span class="text-sm font-medium">Suapabse </span>
                    <span class="text-xs text-muted-color">| Monitoring</span>
                </header>
                <div class="flex-1 p-4 flex flex-col gap-4 overflow-y-auto">
                    <div class="grid grid-cols-2 lg:grid-cols-4 gap-3">
                        <div v-for="stat in stats" :key="stat.label"
                            class="rounded-lg border border-surface-200 dark:border-surface-700 bg-surface-0 dark:bg-surface-900 p-4 flex flex-col gap-2">
                            <div class="flex items-center justify-between">
                                <span class="text-xs text-muted-color">{{ stat.label }}</span>
                                <Icon :icon="stat.icon" class="text-muted-color" width="16" height="16" />
                            </div>
                            <div class="flex items-end justify-between">
                                <span class="text-2xl font-semibold">{{ stat.value }}</span>
                                <span :class="['text-xs font-medium', stat.trend >= 0 ? 'text-emerald-600 dark:text-emerald-400' : 'text-red-500']">
                                    {{ stat.trend >= 0 ? '+' : '' }}{{ stat.trend }}%
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    </section>
</template>