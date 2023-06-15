import { createRouter, createWebHistory } from "vue-router";
import Register from "../components/Register.vue";
import Login from "../components/Login.vue";
import Home from "../components/Home.vue";
import StuEdit from "../components/StuEdit.vue";
import Admin from "../components/admin.vue";
import Publish from "../components/publish.vue";
import ADregister from "../components/adminregister.vue";
const routes = [{
        path: "/register",
        component: Register,
    },
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/home",
        component: Home,
    },
    {
        path: "/stuedit",
        component: StuEdit,
    },
    {
        path: "/admin",
        component: Admin,
    },
    {
        path: "/pub",
        component: Publish,
    },
    {
        path: "/adr",
        component: ADregister,
    },
    // 其他路由规则
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;