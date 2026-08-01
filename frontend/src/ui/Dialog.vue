<script setup lang="ts">
import Dialog from 'primevue/dialog';
import { ref, watch } from 'vue';
const props = defineProps({
    visible: Boolean
})

const visibles = ref<boolean>(props.visible)

watch(() => props.visible, (val) => {
    visibles.value =  val
})

</script>

<template>
    <Dialog v-model:visible="visibles" maximizable modal header="Sync Supabase" :style="{ width: '50rem' }">
        <div class="flex flex-col gap-5">
            <h2 class="text-xl font-bold mt-0 mb-0">Connect your Supabase into Pulsr</h2>

            <div class="flex flex-col gap-1">
                <p class="font-semibold mb-0">1. Login ke Dashboard</p>
                <p class="mt-0 mb-0 ml-6 text-surface-600 dark:text-surface-400">
                    Buka
                    <a href="https://supabase.com/dashboard" target="_blank" class="text-primary font-medium underline">
                        supabase.com/dashboard
                    </a>
                </p>
            </div>

            <!-- Step 2 -->
            <div class="flex flex-col gap-1">
                <p class="font-semibold mb-0">2. Masuk ke Account Settings <span
                        class="font-normal text-surface-500">(bukan project settings!)</span></p>
                <ul class="list-disc ml-10 mt-1 mb-0 flex flex-col gap-1 text-surface-600 dark:text-surface-400">
                    <li>Klik avatar/foto profil kamu di pojok kanan atas</li>
                    <li>
                        Pilih <span class="font-medium">Account Preferences</span>, atau langsung buka
                        <a href="https://supabase.com/dashboard/account/tokens" target="_blank"
                            class="text-primary font-medium underline">
                            supabase.com/dashboard/account/tokens
                        </a>
                    </li>
                </ul>
            </div>
            <div class="flex flex-col gap-1">
                <p class="font-semibold mb-0">3. Buka Tab "Access Tokens"</p>
                <p class="mt-0 mb-0 ml-6 text-surface-600 dark:text-surface-400">
                    Di halaman account, cari dan klik tab <span class="font-medium">Access Tokens</span>.
                </p>
            </div>

            <div class="flex flex-col gap-1">
                <p class="font-semibold mb-0">4. Generate Token Baru</p>
                <ul class="list-disc ml-10 mt-1 mb-0 flex flex-col gap-1 text-surface-600 dark:text-surface-400">
                    <li>Klik tombol <span class="font-medium">Generate New Token</span></li>
                    <li>Kasih nama yang jelas, contoh: <code
                            class="bg-surface-100 dark:bg-surface-800 px-1 rounded">Pulsr Sync Token</code></li>
                    <li>Klik <span class="font-medium">Generate Token</span></li>
                </ul>
            </div>

            <div class="flex flex-col gap-1">
                <p class="font-semibold mb-0">5. Copy & Paste Token ke Pulsr</p>
                <p class="mt-0 mb-0 ml-6 text-surface-600 dark:text-surface-400">
                    Token hanya ditampilkan <span class="font-semibold">sekali</span>. Copy sekarang dan tempel ke kolom
                    di bawah ini.
                </p>
            </div>
        </div>

        <template #footer>
            <div class="flex items-center justify-between w-full text-sm text-surface-500 dark:text-surface-400">
                <span>Pulsr</span>
            </div>
        </template>
    </Dialog>
</template>