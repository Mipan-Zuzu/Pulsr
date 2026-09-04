<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { onBeforeMount, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { supabaseProject, supabaseDetailOrg } from '../store/supabase';
import { Redis } from '../store/redis';
import { RouterLink } from 'vue-router';
import { Skeleton } from 'primevue';
import Loading from '../components/Loading/Loading.vue';

import type { Project, OrgDetail } from '../type/typeSupabase';

const authkeys = ref<string>();
const authKeyInput = ref<string>('');
const isCheckingAuth = ref<boolean>(true);
const isSubmittingAuth = ref<boolean>(false);

const id = ref<string>("")

const supabaseOrg = ref<Project[]>()
const supabaseDetailOrgs = ref<OrgDetail>()
const loadingOrg = ref<boolean>(true)

const supabase = supabaseProject();
const redis = Redis();
const { authKey, errRedis } = storeToRefs(redis);
const { projectDetail, loadingProject } = storeToRefs(supabase);
const orgDetail = supabaseDetailOrg();

onBeforeMount(async () => {
    await redis.redisFind();

    if (errRedis.value !== '') {
        console.error(errRedis.value);
        isCheckingAuth.value = false;
        return;
    }

    await supabase.allproject(authKey.value as string)

    loadingOrg.value = loadingProject.value
    const projects = projectDetail.value?.data as Project[] | undefined
    supabaseOrg.value = projects
    id.value = projects?.[0]?.organization_id ?? ""
    authkeys.value = authKey.value
    isCheckingAuth.value = false

    await orgDetail.detailOrg(id.value, authKey.value as string)
    const { DetailOrg, Errors } = storeToRefs(orgDetail)

    if (Errors.value != null) {
        console.log(Errors.value)
        return
    }

    supabaseDetailOrgs.value = DetailOrg.value
    loadingOrg.value = false
});

const submitAuthKey = async () => {
    if (isSubmittingAuth.value || authKeyInput.value.trim() === '') {
        return
    }
    
    isSubmittingAuth.value = true
    loadingOrg.value = true
    
    try {
        await supabase.allproject(authKeyInput.value)
        
        authkeys.value = authKeyInput.value
        const projects = projectDetail.value?.data as Project[] | undefined
        if (!projects?.length) {
            return
        }
        console.log("pepek teli")

        supabaseOrg.value = projects
        id.value = projects[0].organization_id

        await orgDetail.detailOrg(id.value, authKeyInput.value)

        const { DetailOrg, Errors } = storeToRefs(orgDetail)
        if (Errors.value != null) {
            console.log(Errors.value)
            return
        }

        supabaseDetailOrgs.value = DetailOrg.value
        console.log(authKey.value)
        console.log(authKeyInput.value)
    } finally {
        loadingOrg.value = false
        isSubmittingAuth.value = false
    }
};


const skeletonStyle = [
    { height: "5rem", width: "100%" },
    { height: "5rem", width: "100%" },
    { height: "5rem", width: "100%" },
]
const searchOrg = ref<string>('');
</script>

<template>
    <main class="flex flex-col gap-6 sm:gap-10 pt-16 sm:pt-22 px-4 sm:px-6 lg:px-8 justify-center items-center">
        <header class="text-center">
            <h1 class="text-2xl sm:text-3xl font-medium">Available Organizations</h1>
        </header>

        <div
            v-if="!isCheckingAuth && authkeys === undefined "
            class="fixed inset-0 z-50 flex items-center justify-center px-4"
            role="dialog"
            aria-modal="true"
            aria-labelledby="authkey-title"
        >
            <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

            <section
                class="relative z-10 w-full max-w-sm sm:max-w-md mx-auto rounded-xl border border-gray-200 bg-white p-5 sm:p-6 shadow-xl"
            >
                <h2 id="authkey-title" class="text-lg sm:text-xl font-semibold text-gray-900 mb-1">
                    Insert Your AuthKey
                </h2>
                <p class="text-sm text-gray-500 mb-4">Enter your key to continue</p>

                <form @submit.prevent="submitAuthKey">
                    <label for="authkey-input" class="sr-only">AuthKey</label>
                    <input
                        id="authkey-input"
                        v-model="authKeyInput"
                        type="password"
                        placeholder="Authkey"
                        autocomplete="off"
                        class="w-full rounded-md border border-gray-300 px-3 py-2 text-sm outline-none transition-colors focus:border-gray-500"
                    />

                    <button
                        type="submit"
                        :disabled="isSubmittingAuth"
                        class="mt-4 w-full rounded-md bg-gray-900 py-2 text-sm font-medium text-white transition-colors hover:bg-gray-800 disabled:cursor-not-allowed disabled:opacity-50"
                    >
                        {{ isSubmittingAuth ? 'Loading...' : 'Continue' }}
                    </button>
                </form>

                <aside class="mt-4 flex gap-2 rounded-md bg-gray-50 p-3">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-4 h-4 shrink-0 mt-0.5 text-gray-400"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        stroke-width="2"
                        aria-hidden="true"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            d="M9 12.75L11.25 15 15 9.75M21 12c0 4.556-3.03 8.25-6.75 9.75-3.72-1.5-6.75-5.194-6.75-9.75V6.75l6.75-3 6.75 3V12z"
                        />
                    </svg>
                    <p class="text-xs text-gray-500 leading-relaxed">
                        Your AuthKey is stored securely and never shared with anyone. It has a limited lifetime and
                        expires automatically for your safety.
                    </p>
                </aside>
            </section>
        </div>

        <form role="search" class="relative w-full max-w-md sm:max-w-lg" @submit.prevent>
            <label for="search" class="sr-only">Search organizations</label>
            <Icon
                icon="boxicons:search"
                class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 pointer-events-none text-gray-400"
                aria-hidden="true"
            />
            <input
                id="search"
                v-model="searchOrg"
                type="search"
                autocomplete="off"
                placeholder="Search organizations"
                class="w-full rounded-md border bg-white pl-10 pr-3 py-2 text-sm text-gray-900 outline-none transition-all duration-150"
            />
        </form>

        <section class="w-full" aria-label="Organization list">
            <ul class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3 list-none p-0 m-0">
                <template v-if="loadingOrg">
                    <li v-for="(item, index) in skeletonStyle" :key="`skeleton-${index}`">
                        <Skeleton :width="item.width" :height="item.height" />
                    </li>
                </template>

                <template v-else>
                    <li v-if="supabaseDetailOrgs"
                        v-for="(item, index) in supabaseOrg"
                        :key="index"
                    >
                        <RouterLink :to="`/supabase/project/${item.id}`" class="block">
                            <article
                                class="border p-4 sm:p-5 sm:px-7 rounded-lg hover:scale-105 duration-100 cursor-pointer hover:shadow-lg bg-white h-full"
                            >
                                <div class="flex gap-4 sm:gap-7 items-center">
                                    <div class="p-2 rounded-full border shrink-0" aria-hidden="true">
                                        <Icon icon="fluent:organization-horizontal-16-regular" />
                                    </div>
                                    <div class="flex flex-col min-w-0">
                                        <h2 class="font-semibold text-lg sm:text-xl truncate">
                                            {{ supabaseDetailOrgs?.name }}
                                        </h2>
                                        <div class="flex gap-2 text-sm text-gray-500">
                                            <span>{{ supabaseOrg?.length }} project(s)</span>
                                            <span>&middot;</span>
                                            <span>{{ supabaseDetailOrgs?.plan }}</span>
                                        </div>
                                    </div>
                                </div>
                            </article>
                        </RouterLink>
                    </li>

                    <li v-else class="col-span-full text-center text-sm text-gray-500 py-10 flex items-center gap-3">
                        <Loading />
                        <h1> This will take a fiew time <span class="font-semibold">tips! check your credential</span></h1>
                    </li>
                </template>
            </ul>
        </section>
    </main>
</template>