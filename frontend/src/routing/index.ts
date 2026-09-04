import { createRouter, createWebHistory } from "vue-router"

import Dashboard from "../view/Dashboard/Dashboard.vue"
import SupabaseProject from "../Supabase/supabaseProject/SupabaseProject.vue"
import SupabaseMonitorView from "../view/Supabase/SupabaseMonitorView.vue"

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            component: Dashboard
        },
        {
            path: "/supabase/project/:id",
            component: SupabaseProject
        },
        {
            path: "/supabase/project/:id/status",
            component: SupabaseMonitorView
        },
        {
            path: "/projects/:id",
            redirect: to => `/supabase/project/${to.params.id}/status`
        }
    ]
})

export default router
