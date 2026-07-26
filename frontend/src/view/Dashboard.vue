<script setup lang="ts">
    import Sidebar from '../components/Sidebar.vue'
    import {useNavSidebar} from "../store/navside.ts"
    import { storeToRefs } from "pinia"

    import Vercel from '../Monitoring/Vercel.vue'

    const storeNav = useNavSidebar()

    const {myState} = storeToRefs(storeNav)
</script>

<template>
    <body>
        <aside class="fixed">
            <Sidebar />
        </aside>
        <div>
            <section :class="myState === 'Vercel' ? 'block' : 'hidden'">
                <Vercel />
            </section>
            <!-- <section :class="">
                
            </section> -->
        </div>
    </body>
</template>
