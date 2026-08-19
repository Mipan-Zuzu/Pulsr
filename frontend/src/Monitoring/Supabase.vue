<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { onBeforeMount, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { supabaseProject } from '../store/supabase';
import { Redis } from '../store/redis';
import { RouterLink } from 'vue-router';
import { Skeleton } from 'primevue';
import type { Project } from '../type/typeSupabase';

const authkeys = ref<string>();
const authKeyInput = ref<string>('');
const isCheckingAuth = ref<boolean>(true);

const supabaseOrg = ref<Project[]>()
const loadingOrg = ref<boolean>(true)

const supabase = supabaseProject();
const redis = Redis();
const { authKey, errRedis } = storeToRefs(redis);

onBeforeMount(async () => {
    await redis.redisFind();

    if (errRedis.value !== '') {
        console.error(errRedis.value);
        isCheckingAuth.value = false;
        return;
    }

    await supabase.allproject(authKeyInput.value)
    const { projectDetail, loadingProject } = storeToRefs(supabase)

    loadingOrg.value = loadingProject.value
    supabaseOrg.value = projectDetail.value

    authkeys.value = authKey.value;
    isCheckingAuth.value = false;
     console.log(projectDetail.value, loadingOrg.value)
});

const submitAuthKey = async () => {
    // TODO: panggil action buat simpan authkey + set TTL
    // contoh: await redis.setAuthKey(authKeyInput.value)
    await supabase.allproject(authKeyInput.value)

    const { projectDetail, loadingProject } = storeToRefs(supabase)

    loadingOrg.value = loadingProject.value
    supabaseOrg.value = projectDetail.value

    console.log(projectDetail.value, loadingOrg.value)

    authkeys.value = authKeyInput.value;
};

const searchOrg = ref<string>('');
</script>

<template>
    <main class="flex flex-col gap-10 pt-22 justify-center items-center">
        <h1 class="text-3xl font-medium">Available Organizations</h1>
        <div v-if="!isCheckingAuth && authkeys === undefined"
            class="fixed inset-0 z-50 flex items-center justify-center">
            <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

            <section
                class="relative z-10 w-full max-w-sm mx-4 rounded-xl border border-gray-200 bg-white p-6 shadow-xl">
                <h2 class="text-xl font-semibold text-gray-900 mb-1">Insert Your AuthKey</h2>
                <p class="text-sm text-gray-500 mb-4">Enter your key to continue</p>

                <input v-model="authKeyInput" type="password" placeholder="Authkey" autocomplete="off"
                    class="w-full rounded-md border border-gray-300 px-3 py-2 text-sm outline-none transition-colors focus:border-gray-500" />

                <button type="button" @click="submitAuthKey"
                    class="mt-4 w-full rounded-md bg-gray-900 py-2 text-sm font-medium text-white transition-colors hover:bg-gray-800">
                    Continue
                </button>

                <div class="mt-4 flex gap-2 rounded-md bg-gray-50 p-3">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 shrink-0 mt-0.5 text-gray-400" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M9 12.75L11.25 15 15 9.75M21 12c0 4.556-3.03 8.25-6.75 9.75-3.72-1.5-6.75-5.194-6.75-9.75V6.75l6.75-3 6.75 3V12z" />
                    </svg>
                    <p class="text-xs text-gray-500 leading-relaxed">
                        Your AuthKey is stored securely and never shared with anyone. It has a limited lifetime and
                        expires automatically for your safety.
                    </p>
                </div>
            </section>
        </div>

        <div class="relative w-72">
            <Icon icon="boxicons:search" class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 pointer-events-none"
                aria-hidden="true" />
            <input id="search" v-model="searchOrg" type="text" autocomplete="off" placeholder="Search organizations"
                class="w-full rounded-md border bg-white pl-10 pr-3 py-2 text-sm text-gray-900 outline-none transition-all duration-150" />
        </div>

        <section class="flex justify-center items-center">
            <div class="grid grid-cols-3 gap-3">
                <!-- TODO: v-for loop dari organizations list -->
                <div v-if="loadingOrg">
                    <Skeleton size="2rem" />
                </div>
                <RouterLink to="" v-else 
                     
                    >
                    <article
                        class="border p-5 px-7 rounded-lg hover:scale-105 duration-100 cursor-pointer hover:shadow-lg bg-white">
                        <span class="flex gap-7 items-center">
                            <div class="p-2 rounded-full border" aria-hidden="true">
                                <Icon icon="fluent:organization-horizontal-16-regular" />
                            </div>
                            <div class="flex flex-col">
                                <h2 class="font-semibold text-xl">Mipan' Orgs</h2>
                                <div class="flex gap-2">
                                    <p>1 project</p>
                                    <p>Free plan</p>
                                </div>
                            </div>
                        </span>
                    </article>
                </RouterLink>
            </div>
        </section>
    </main>
</template>