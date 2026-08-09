<script setup lang="ts">

// Third Party
import { storeToRefs } from "pinia"
import { Icon } from "@iconify/vue"
import Card from 'primevue/card';
import Password from 'primevue/password'
import Dialog from "../ui/Dialog.vue";
import { useRouter } from "vue-router";
// import { useRouter } from "vue-router";

// Internal
import { onBeforeMount, onMounted, ref, watch } from 'vue';
import { supabaseProject } from '../store/supabase';
import UptimeSupabase from "../Supabase/UptimeSupabase.vue";
import RequestTotal from "../Supabase/RequestTotal.vue";
import ErrorHistory from "../Supabase/ErrorHistory.vue";
import type { Project } from "../type/typeSupabase.ts";
import { Redis } from "../store/redis.ts";

// component
import Message from 'primevue/message';
import Button from 'primevue/button';


// Service
const router = useRouter()
const visible = ref<boolean>(false)
    const ProjectData = ref<Project[]>([])


const authTokenSupabase = ref<string>("")
const storeprojectSupabase = supabaseProject()
const redis = Redis()
const { projectDetail } = storeToRefs(storeprojectSupabase)
const {authKey} = storeToRefs(redis)

onBeforeMount(async() => {
    await redis.redisFind()
    storeprojectSupabase.allproject(authKey.value!)
})

const handleauthToken = () => {
    storeprojectSupabase.allproject(authTokenSupabase.value)
}

watch(projectDetail, (newValue) => {
    ProjectData.value = newValue.data
})

const handleDialog = () => {
    visible.value = true
}

console.log(ProjectData.value.length)

const projectNav = (value: string) => {
    if (value == null) {
        router.push(`/supabase/project/undefined`)
    }else {
        router.push(`/supabase/project/${value}`)
    }
}
</script>

<template>
    <section class="flex flex-col ml-10 mr-10">
        <div class="flex justify-end">
            <Message severity="error">
                <template #icon>
                    <Receipt />
                </template>
                <span class="flex items-center gap-3">
                    <Icon icon="material-symbols:error-outline" width="16" height="16" />
                    <p>Error, Need Authorization Token</p>
                </span>
            </Message>
        </div>
        <section class="grid grid-cols-1 mt-10 md:grid-cols-3 gap-5 justify-center">
            <UptimeSupabase />
            <RequestTotal />
            <ErrorHistory />
        </section>
        <span>
            <div class="flex justify-start gap-2 mt-4">
                <FloatLabel variant="on">
                    <Password id="on_label" v-model="authTokenSupabase" :feedback="false" toggleMask />
                </FloatLabel>
                <Button @click="handleauthToken">Key</Button>
            </div>
        </span>
        <section class="flex-col mt-10 justify-start">
            <h1 class="text-xl">Project</h1>
            <div v-if="ProjectData.length === 0" class="mt-10 flex items-center gap-2">
                <p class="text-gray-500">Empty Project data, Cannot Record analytic</p>
                <Dialog :visible="visible" />
                <button class="text-gray-500 underline cursor-pointer" @click="handleDialog">Inssert Auth Key</button>
            </div>
            <div v-else  class="flex justify-start mt-5" v-for="(item, index) in ProjectData"
                :key="index">
                <button 
                @click="projectNav(item.id)"
                class="max-w-sm w-72 text-start mt-10 hover:scale-105 duration-100 hover:opacity-70 cursor-pointer">
                    <Card>
                        <template #title>
                            <span class="flex gap-2 items-center">
                                <Icon icon="selfhst:supabase" width="16" height="16" />
                                <h1>{{ item.name }}</h1>
                            </span>
                        </template>
                        <template #content>
                            <p class="m-0">Status {{ item.status }}</p>
                        </template>
                        <template #footer>
                            <span class="text-sm text-surface-500 dark:text-surface-400">
                                <div class="flex gap-2 items-center">
                                    <Icon icon="twemoji:flag-australia" width="20" height="20" />
                                    <p class="text-lg "> AWS | {{ item.region }}</p>
                                </div>
                            </span>
                        </template>
                    </Card>
                </button>
            </div>
        </section>
    </section>
</template>