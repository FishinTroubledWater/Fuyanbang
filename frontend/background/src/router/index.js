import LoginVue from '@/views/Login.vue'
import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UserView from "@/views/UserView";
import RecipeView from "@/views/RecipeView";
import CommentView from "@/views/CommentView";
import FeedBackView from "@/views/FeedBackView";
import DashBoard from "@/components/DashBoard";
import NewsView from "@/views/NewsView";
import PostView from "@/views/PostView";
import ForgetPwd from "@/views/ForgetPwd";


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
                    path: '/post',
                    name: 'post',
                    component: PostView,
                },
                {
                    path: '/news',
                    name: 'news',
                    component: NewsView,
                },
                {
                    path: '/user',
                    name: 'user',
                    component: UserView,
                },
                {
                    path: '',
                    name: 'dashBoard',
                    component: DashBoard,
                },
                {
                    path: '/about',
                    name: 'about',
                    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
                },
            ]
        },
        {
            path: '/forget',
            name: 'forget',
            component: ForgetPwd,
        },
        {
            path: "/",
            redirect: "/login",
        },
    ]
})

// 挂载路由导航守卫
router.beforeEach((to,from,next)=>{
    if(to.path==='/login' || to.path==='/forget') return next();
    const tokenStr =window.sessionStorage.getItem('token')
    if(!tokenStr) return next('/login')
    next()
})

// const router = new VueRouter({
//     mode: 'history',
//     base: process.env.BASE_URL,
//     routes
// })

export default router
