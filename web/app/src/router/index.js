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
const DashboardReportCourse = () => import("@/views/dashboard/reports/Course.vue")
const DashboardSetting = () => import("@/views/dashboard/settings/Canvas")

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
  {
    path: '/report/course',
    name: 'dashboard.report.course',
    component: DashboardReportCourse,
    meta : {
      requiredAuth : true,
      title : "Course Reports"
    }
  },
  {
    path: '/setting',
    name: 'dashboard.setting',
    component: DashboardSetting,
    meta : {
      requiredAuth: true,
      title: "Setting"
    }
  }  
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
