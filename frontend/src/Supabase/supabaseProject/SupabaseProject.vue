<script setup>
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from 'vue';
import Sidebar from 'primevue/sidebar';
import SidebarBackdrop from 'primevue/sidebarbackdrop';
import SidebarAside from 'primevue/sidebaraside';
import SidebarContent from 'primevue/sidebarcontent';
import SidebarFooter from 'primevue/sidebarfooter';
import SidebarGroup from 'primevue/sidebargroup';
import SidebarGroupAction from 'primevue/sidebargroupaction';
import SidebarGroupContent from 'primevue/sidebargroupcontent';
import SidebarGroupLabel from 'primevue/sidebargrouplabel';
import SidebarHeader from 'primevue/sidebarheader';
import SidebarMain from 'primevue/sidebarmain';
import SidebarLayout from 'primevue/sidebarlayout';
import SidebarMenu from 'primevue/sidebarmenu';
import SidebarMenuAction from 'primevue/sidebarmenuaction';
import SidebarMenuBadge from 'primevue/sidebarmenubadge';
import SidebarMenuButton from 'primevue/sidebarmenubutton';
import SidebarMenuItem from 'primevue/sidebarmenuitem';
import SidebarMenuSub from 'primevue/sidebarmenusub';
import SidebarMenuSubButton from 'primevue/sidebarmenusubbutton';
import SidebarMenuSubItem from 'primevue/sidebarmenusubitem';
import SidebarPanel from 'primevue/sidebarpanel';
import SidebarRail from 'primevue/sidebarrail';
import SidebarSpacer from 'primevue/sidebarspacer';
import SidebarTrigger from 'primevue/sidebartrigger';
import { Icon } from '@iconify/vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { supabaseDetail } from '../../store/supabase';
import Skeleton from 'primevue/skeleton';

const route = useRoute()
const projectDetail = ref()
const supabasePorjectDetail = supabaseDetail()

const id = route.params.id

const detailProject = ref()
const {projectData, errProject, loadingProject} = storeToRefs(supabaseDetail())

onBeforeMount(() => {
    supabasePorjectDetail.allprojectData(id)
})

const isMobile = ref(false);
const open = ref(false);
let mql = null;
let onMqlChange = null;

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
        <SidebarLayout class="min-h-192! relative!">
            <SidebarBackdrop v-if="isMobile && open" class="absolute!" />
            <!-- Icon bar - always collapsed, opens on hover as overlay -->
            <Sidebar id="iconbar" side="left" :collapsible="isMobile ? 'offcanvas' : 'icon'" :overlay="true" :openOnHover="!isMobile" width="12rem" v-model:open="open">
                <SidebarSpacer />
                <SidebarAside>
                    <SidebarPanel>
                        <SidebarHeader>
                            <SidebarMenu>
                                <SidebarMenuItem>
                                    <SidebarMenuButton class="p-1!">
                                        <div class="flex size-6 shrink-0 items-center justify-center rounded-md text-white text-xs font-bold leading-none">
                                        <Icon class="text-stone-700" icon="tabler:brand-cake"  />
                                        </div>
                                        <span class="font-semibold text-sm">Pulsr Demo</span>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            </SidebarMenu>
                        </SidebarHeader>

                        <SidebarContent>
                            <SidebarGroup>
                                <SidebarGroupContent>
                                    <SidebarMenu>
                                        <SidebarMenuItem>
                                            <SidebarMenuButton :isActive="true">
                                                <Icon icon="lucide:home" />
                                                <span>Home</span>
                                            </SidebarMenuButton>
                                        </SidebarMenuItem>
                                        <SidebarMenuItem>
                                            <SidebarMenuButton>
                                                <Icon icon="lucide:database" />
                                                <span>Database</span>
                                            </SidebarMenuButton>
                                        </SidebarMenuItem>
                                        <SidebarMenuItem>
                                            <SidebarMenuButton>
                                                <Icon icon="lucide:key" />
                                                <span>Authentication</span>
                                            </SidebarMenuButton>
                                        </SidebarMenuItem>
                                        <SidebarMenuItem>
                                            <SidebarMenuButton>
                                                <Icon icon="lucide:logs" />
                                                <span>Edge Functions</span>
                                            </SidebarMenuButton>
                                        </SidebarMenuItem>
                                    </SidebarMenu>
                                </SidebarGroupContent>
                            </SidebarGroup>
                        </SidebarContent>
                    </SidebarPanel>
                </SidebarAside>
            </Sidebar>

            <SidebarMain v-if="loadingProject == true">
            <header class="flex h-12 items-center gap-2 border-b border-surface-200 dark:border-surface-700 px-4">
                    <SidebarTrigger v-if="isMobile" target="iconbar" severity="secondary" :text="true" size="small">
                        <Icon icon="lucide:panel-left" />
                    </SidebarTrigger>
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
            </SidebarMain>
            <SidebarMain v-else>
                <header class="flex h-12 items-center gap-2 border-b border-surface-200 dark:border-surface-700 px-4">
                    <SidebarTrigger v-if="isMobile" target="iconbar" severity="secondary" :text="true" size="small">
                        <Icon icon="lucide:panel-left" />
                    </SidebarTrigger>
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
            </SidebarMain>
        </SidebarLayout>
    </section>
</template>