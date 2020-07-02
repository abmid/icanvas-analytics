/*
 * File Created: Monday, 29th June 2020 11:52:52 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

import Vue from 'vue'
import VueRouter from 'vue-router'

import AuthLogin from "@/views/auth/Login.vue"

const DashboardHome = () => import("@/views/dashboard/home/Home.vue")

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'login',
    component: AuthLogin,
    meta : {
      requiredAuth : false,
      title : "Login"
    }
  },
  {
    path: '/home',
    name: 'dashboard.home',
    component: DashboardHome,
    meta : {
      requiredAuth : true,
      title : "Home"
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
