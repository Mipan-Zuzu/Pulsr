<template>
    <MegaMenu :model="items">
        <template #item="{ item, label }">
            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                <a v-ripple class="flex items-center" :href="href" @click="navigate">
                    <span class="ml-2">{{ label }}</span>
                </a>
            </router-link>
            <a v-else v-ripple :href="item.url" :target="item.target">
                <span class="ml-2">{{ label }}</span>
            </a>
        </template>
    </MegaMenu>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import MegaMenu from 'primevue/megamenu';
const items = ref([
    {
        label: 'Dashboard',
        icon: 'prime:palette',
        items: [
            [
                {
                    label: 'Router Link',
                    items: [
                        { label: 'Theming', route: '/theming' },
                        { label: 'UI Kit', route: '/uikit' }
                    ]
                }
            ]
        ]
    },
    {
        label: 'Project',
    },
    {
        label: 'Free',
    }
]);
</script>