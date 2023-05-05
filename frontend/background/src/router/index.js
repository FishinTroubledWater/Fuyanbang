import LoginVue from '@/views/Login.vue'
import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UserView from "@/views/UserView";
import ContentView from "@/views/ContentView";
import RecipeView from "@/views/RecipeView";
import TrendsView from "@/views/TrendsView";


Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        component: LoginVue,
    },
    {
        path: '/home',
        name: 'home',
        component: HomeView,
        children:[

            {
                path: '/recipe',
                name: 'recipe',
                component: RecipeView,
            },
            {
                path: '/content',
                name: 'content',
                component: ContentView,
            },
            {
                path: '/trends',
                name: 'trends',
                component: TrendsView,
            },
            {
                path: '/user',
                name: 'user',
                component: UserView,
            },
        ]
    },
    {
        path: '/about',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
    {
        path: "/",
        redirect: "/login",
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
