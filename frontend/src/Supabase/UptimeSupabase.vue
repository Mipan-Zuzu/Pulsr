<script setup lang="ts">
import Card from 'primevue/card';
import { Icon } from '@iconify/vue';
</script>

<template>
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
</template>