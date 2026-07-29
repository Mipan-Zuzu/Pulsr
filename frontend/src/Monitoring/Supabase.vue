<script setup lang="ts">
import { storeToRefs } from "pinia"
import { supabaseProject } from '../store/supabase';
import { ref } from 'vue';
import { Icon } from "@iconify/vue"
import Password from 'primevue/password'

const authTokenSupabase = ref<string>("")
const storeprojectSupabase = supabaseProject()
const { projectDetail } = storeToRefs(storeprojectSupabase)

const handleauthToken = () => {
    storeprojectSupabase.allproject(authTokenSupabase.value)
}



const log = console.log
log(projectDetail.value.data)


// import component
import Card from 'primevue/card';
import Message from 'primevue/message';
import Button from 'primevue/button';

</script>

<template>
    <main class="flex flex-col ml-10 mr-10">
        <div class="flex justify-end">
            <Message severity="error">
                <template #icon>
                    <Receipt />
                </template>
                <span class=" flex items-center gap-3">
                    <Icon icon="material-symbols:error-outline" width="16" height="16" />
                    <p>Error, Need Authorization Token</p>
                </span>
            </Message>
        </div>
        <section class="grid grid-cols-1  mt-10 md:grid-cols-3 gap-5 justify-center">
            <Card class="max-w-sm">
                <template #title>
                    <span class="flex items-center gap-3">
                        <Icon icon="selfhst:supabase" width="16" height="16" />
                        <h1>Status Uptime</h1>
                    </span>
                </template>
                <template #content>
                    <p><span class="text-2xl">null</span> Supabase Uptime </p>
                </template>
                <template #footer>
                    <div class="flex flex-col gap-4">
                        <span class="flex items-center gap-3">
                            <p class="w-2 h-2 rounded-full bg-gray-500"></p>
                            <p class="w-2 h-2 rounded-full bg-gray-500 animate-ping absolute"></p>
                            <p>Offline</p>
                        </span>
                    </div>
                </template>
            </Card>
            <Card class="max-w-sm">
                <template #title>
                    <span class="flex items-center gap-3">
                        <Icon icon="selfhst:supabase" width="16" height="16" />
                        <h1>Total Request</h1>
                    </span>
                </template>
                <template #content>
                    <p><span class="text-2xl">null</span> Request Total </p>
                </template>
                <template #footer>
                    <div class="flex flex-col gap-4">
                        <span class="flex items-center gap-3">
                            <p class="w-2 h-2 rounded-full bg-gray-500"></p>
                            <p class="w-2 h-2 rounded-full bg-gray-500 animate-ping absolute"></p>
                            <p>Offline</p>
                        </span>
                    </div>
                </template>
            </Card>
            <Card class="max-w-sm">
                <template #title>
                    <span class="flex items-center gap-3">
                        <Icon icon="selfhst:supabase" width="16" height="16" />
                        <h1>Error history</h1>
                    </span>
                </template>
                <template #content>
                    <p class="text-2xl">Null</p>
                </template>
                <template #footer>
                    <div class="flex flex-col gap-4">
                        <span class="flex items-center gap-3">
                            <p class="w-2 h-2 rounded-full bg-gray-500"></p>
                            <p class="w-2 h-2 rounded-full bg-gray-500 animate-ping absolute"></p>
                            <p>Offline</p>
                        </span>
                    </div>
                </template>
            </Card>
        </section>
        <span>
            <div class="flex justify-start gap-2 mt-4">
                <FloatLabel variant="on">
                    <Password id="on_label" v-model="authTokenSupabase" :feedback="false" toggleMask />
                </FloatLabel>
                <Button @click="handleauthToken">Auth Key</Button>
            </div>
        </span>
        <section class="flex-col mt-10 justify-start">
            <h1 class="text-xl">Project</h1>
            <div class="flex justify-start mt-5" v-for="(item, index) in projectDetail"
               :key="index"
            >
                <button class="max-w-sm w-72 text-start">
                    <Card>
                        <template #title>
                            <span class="flex gap-2 items-center">
                                <Icon icon="selfhst:supabase" width="16" height="16" />
                                <h1>{{ item.Name }}</h1>
                            </span>
                        </template>
                        <template #subtitle>{{ item.Organization_slug }}</template>
                        <template #content>
                            <p class="m-0">Service Status : {{ item.Status }}</p>
                        </template>
                        <template #footer>
                            <span class="text-sm text-surface-500 dark:text-surface-400">
                                <div class="flex gap-2 items-center">
                                    <Icon icon="carbon:cics-region" width="16" height="16" />
                                    {{ item.region }}
                                </div>
                            </span>
                        </template>
                    </Card>
                </button>
            </div>
        </section>
    </main>
</template>