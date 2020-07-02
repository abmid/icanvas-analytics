/*
 * File Created: Monday, 29th June 2020 6:39:21 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <div class="card">
    <div class="card-body login-card-body">
      <p class="login-box-msg">Sign in to start your session</p>
        <slot />
    </div>
    <div v-if="loading" class="overlay"><i class="fas fa-2x fa-sync-alt fa-spin"></i></div>    
    <!-- /.login-card-body -->
  </div>  
</template>

<script>
export default {
  props:{
    loading : {
      type: Boolean,
      default : () => { false}
    }
  }
}
</script>

<style scoped>
.card, .card-body{
  border-radius: 8px !important;
}
.card-body{
  box-shadow: 6px 6px 7px 2px #0000001c;
}
</style>