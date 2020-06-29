import Vue from 'vue'
import VueRouter from 'vue-router'

import AuthLogin from "@/views/auth/Login.vue"

const DashboardHome = () => import("@/views/dashboard/home/Home.vue")

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'login',
    component: AuthLogin
  },
  {
    path: '/home',
    name: 'home',
    component: DashboardHome
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
