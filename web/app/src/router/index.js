/*
 * File Created: Monday, 29th June 2020 11:52:52 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

import Vue from 'vue'
import VueRouter from 'vue-router'

import AuthLogin from "@/views/auth/Login.vue"
import AuthRegister from "@/views/auth/Register.vue"

const DashboardHome = () => import("@/views/dashboard/home/Home.vue")

import reportCourse from './reports/course' 
import setting from "./settings/setting"

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
    path: '/welcome',
    name: 'welcome',
    component: AuthRegister,
    meta : {
      requiredAuth : false,
      title : "Welcome"
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
  },
  ...reportCourse,
  ...setting
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

export default router
