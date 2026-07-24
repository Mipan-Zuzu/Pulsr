<script setup lang="ts">
import { Icon } from "@iconify/vue"

const monitoring = [
    { name: "Cloudflare", icon: "selfhst:cloudflare", status: false },
    { name: "Supabase", icon: "selfhst:supabase", status: true },
    { name: "Vercel", icon: "devicon:vercel", status: true },
    { name: "Redis", icon: "logos:upstash-icon", status: false },
    { name: "Http Status", icon: "ci:globe", status: false },
    { name: "Groq Api", icon: "bxl:groq-ai", status: false },
    { name: "Koyeb", icon: "simple-icons:koyeb", status: false },
    { name: "Mongo Db", icon: "devicon:mongodb", status: false },
    { name: "Domain SSL", icon: "material-icon-theme:url", status: false },
]

const mainMenu = [
    { name: "Notification", icon: "iconamoon:notification", path: "/notification" },
    { name: "Dashboard", icon: "ri:dashboard-fill", path: "/" },
    { name: "Setting", icon: "tdesign:setting", path: "/setting" },
]
</script>

<template>
    <aside class="border-r border-stone-200 h-screen w-52 flex flex-col">
        <!-- Header -->
        <div class="mt-4 flex items-center px-3 gap-3 shrink-0">
            <Icon class="text-stone-700" icon="tabler:brand-cake" width="42" height="42" />
            <section class="flex flex-col">
                <h1 class="text-lg font-semibold text-stone-700">Pulsr Demo</h1>
                <p class="text-[13px] text-stone-500 -mt-2">Free Plan</p>
            </section>
        </div>

        <!-- Main menu -->
        <nav class="flex flex-col mt-10 border-b border-stone-200 pb-3 px-2 gap-1 shrink-0">
            <router-link
                v-for="item in mainMenu"
                :key="item.name"
                :to="item.path"
                active-class="bg-stone-100 text-stone-800"
                class="group hover:bg-stone-50 duration-300 cursor-pointer rounded-lg p-2 flex items-center gap-3 w-full"
            >
                <Icon
                    class="text-stone-500 group-hover:text-stone-700 duration-100"
                    :icon="item.icon"
                    width="20"
                    height="20"
                />
                <h1 class="text-stone-500 group-hover:text-stone-700 duration-100 text-sm">
                    {{ item.name }}
                </h1>
            </router-link>
        </nav>

        <!-- Monitoring label -->
        <h1 class="ml-3 text-stone-500 text-sm mt-4 font-semibold shrink-0">
            Monitoring
        </h1>

        <!-- Monitoring list (scrollable) -->
        <nav class="flex flex-col gap-1 mt-3 px-2 flex-1 overflow-y-auto">
            <div
                v-for="item in monitoring"
                :key="item.name"
                class="group hover:bg-stone-100 duration-300 cursor-pointer rounded-lg p-2 flex items-center gap-3 w-full"
            >
                <!-- Status dot -->
                <span class="relative flex h-2 w-2 shrink-0">
                    <span
                        :class="item.status ? 'bg-green-400' : 'bg-yellow-400'"
                        class="animate-ping absolute inline-flex h-full w-full rounded-full opacity-75"
                    ></span>
                    <span
                        :class="item.status ? 'bg-green-500' : 'bg-yellow-500'"
                        class="relative inline-flex rounded-full h-2 w-2"
                    ></span>
                </span>

                <Icon
                    class="opacity-70 group-hover:opacity-100 duration-300 shrink-0"
                    :icon="item.icon"
                    width="16"
                    height="16"
                />
                <h1 class="text-stone-500 group-hover:text-stone-700 duration-300 text-sm truncate">
                    {{ item.name }}
                </h1>
            </div>
        </nav>

        <!-- Support (fix di bawah) -->
        <div class="shrink-0 border-t border-stone-200 p-3">
            <a href="#" class="flex items-center gap-3 text-stone-700 font-medium hover:text-stone-900 duration-300">
                <Icon icon="bx:support" width="20" height="20" />
                Support
            </a>
        </div>
    </aside>
</template>