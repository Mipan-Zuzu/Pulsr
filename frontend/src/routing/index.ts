import {createRouter, createWebHistory} from "vue-router"

import Dashboard from '../view/Dashboard.vue'
import SupabaseProject from "../Supabase/supabaseProject/SupabaseProject.vue"

const router = createRouter({
    history: createWebHistory(),
    routes : [
        {
            path: "/",
            component: Dashboard
        },
        {
            path: "/supabase/project/:id",
            component: SupabaseProject
        }
    ]
})

export default router