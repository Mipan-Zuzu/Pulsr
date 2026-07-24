import {createRouter, createWebHistory} from "vue-router"

import Dashboard from '../view/Dashboard.vue'

const router = createRouter({
    history: createWebHistory(),
    routes : [
        {
            path: "/",
            component: Dashboard
        }
    ]
})

export default router