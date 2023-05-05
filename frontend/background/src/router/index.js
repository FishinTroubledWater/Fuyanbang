import LoginVue from '@/views/Login.vue'
import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UserView from "@/views/UserView";
import ContentView from "@/views/ContentView";
import RecipeView from "@/views/RecipeView";
import TrendsView from "@/views/TrendsView";
import CommentView from "@/views/CommentView";
import FeedBackView from "@/views/FeedBackView";


Vue.use(VueRouter)

const router = new VueRouter({
    routes: [
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
                    path: '/comment',
                    name: 'comment',
                    component: CommentView,
                },
                {
                    path: '/feedback',
                    name: 'feedback',
                    component: FeedBackView,
                },
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
})

//挂载路由导航收尾
// router.beforeEach((to,from,next)=>{
//     if(to.path==='/login') return next();
//     const tokenStr =window.sessionStorage.getItem('token')
//     if(!tokenStr) return next('/login')
//     next()
// })

// const router = new VueRouter({
//     mode: 'history',
//     base: process.env.BASE_URL,
//     routes
// })

export default router
