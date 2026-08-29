<script setup lang="ts">
    import Sidebar from "../../components/Sidebar.vue"
    import {useNavSidebar} from "../../store/navside.ts"
    import { storeToRefs } from "pinia"

    import Vercel from '../../Monitoring/Vercel.vue'
    import Supabase from '../../Monitoring/Supabase.vue'

    const storeNav = useNavSidebar()

    const {myState} = storeToRefs(storeNav)
</script>

<template>
        <aside class="fixed z-50 ">
            <Sidebar />
        </aside>
        <div> 
            <section :class="myState === 'Vercel' ? 'block' : 'hidden'">
                <Vercel />
            </section>
             <section :class="myState === 'Supabase' ? 'block' : 'hidden'">
                <Supabase />
            </section>
        </div>
</template>
