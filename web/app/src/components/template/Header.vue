/*
 * File Created: Monday, 29th June 2020 5:22:22 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <!-- Navbar -->
  <nav class="main-header navbar navbar-expand navbar-white navbar-light">
    <!-- Left navbar links -->
    <ul class="navbar-nav">
      <li class="nav-item">
        <a class="nav-link" data-widget="pushmenu" href="#" role="button"><i class="fas fa-bars"></i></a>
      </li>
      <li class="nav-item d-none d-sm-inline-block">
        <a href="#" class="nav-link">GitHub</a>
      </li>
    </ul>

    <!-- Right navbar links -->
    <ul class="navbar-nav ml-auto">
      <!-- Notifications Dropdown Menu -->
      <li class="nav-item dropdown">
        <a class="nav-link" data-toggle="dropdown" href="#" aria-expanded="false">
          <i class="far fa-bell"></i>
          <span v-if="countNotif != 0" class="badge badge-warning navbar-badge">{{countNotif}}</span>
        </a>
        <div class="dropdown-menu dropdown-menu-lg dropdown-menu-right" style="left: inherit; right: 0px;">
          <span class="dropdown-item dropdown-header">{{countNotif}} Notifications</span>
          <span v-for="(data, index) in notification" :key="index">
            <div class="dropdown-divider"></div>
            <a href="#" class="dropdown-item custom-notif">
              <i class="fas fa-envelope mr-2"></i> {{ data.message }}
              <span class="float-right text-muted text-sm" style="font-size:13px !important">{{ data.created_at }}</span>
            </a>
          </span>
          <a href="#" class="dropdown-item dropdown-footer">See All Notifications</a>
        </div>
      </li>      
      <!-- Account  -->
      <li class="nav-item dropdown">
        <a class="nav-link" data-toggle="dropdown" href="#">
          <i class="fas fa-user"></i>&nbsp;<b>{{userName}}</b>
        </a>
        <div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
          <span class="dropdown-header">Menu</span>
          <div class="dropdown-divider"></div>
          <a href="#" class="dropdown-item">
            <i class="fas fa-key mr-2"></i> Change Password
            <span class="float-right text-muted text-sm">3 mins</span>
          </a>
          <div class="dropdown-divider"></div>
          <a href="#" @click="handleLogout()" class="dropdown-item dropdown-footer">Logout</a>
        </div>
      </li>
    </ul>
  </nav>
  <!-- /.navbar -->  
</template>

<script>
import iMixins from "@/helpers/mixins"
export default {
  mixins: [iMixins],  
  data(){
    return{
      ws : {
        connection : null
      },
    }
  },
  computed: {
    countNotif : function(){
      return this.$store.state.notification.notification.length
    },
    notification : function(){
      return this.$store.state.notification.notification
    },
    userName : function(){
      return this.$store.state.auth.currentUser.name
    }
  },
  created(){
    this.initWebSocket()
  },
  methods : {
    handleLogout : function(){
      this.$store.dispatch("auth/logout").then(res => {
        if (res.status == 200) {
          location.reload()
        }else{
          this.$$_TOAST_SHOW("danger", "Failed Logout", res.data.message)
        }
      }).catch(err => {
        console.log(err)
      })
    },
    initWebSocket : function(){
      const vm = this
      this.ws.connection = new WebSocket("ws://localhost:8000/v1/ws/server"),
          // container = $("#container")
      this.ws.connection.onopen = function() {
          // container.append("<p>Socket is open</p>");
      };
      this.ws.connection.onmessage = function (e) {
          // container.append("<p> Got some shit:" + e.data + "</p>");
          var parse = JSON.parse(e.data)
          vm.$store.dispatch("notification/addNotification",parse)
          vm.$$_TOAST_SHOW("success","Generate Data", parse.message)
      }
      this.ws.connection.onclose = function () {
          // container.append("<p>Socket closed</p>");
      }   
    }
  }
}
</script>

<style scoped>
.custom-notif{
  white-space: normal !important;
  overflow: hidden;
  width: 100%;
  height: 100%;
  font-size: 15px;
}
</style>