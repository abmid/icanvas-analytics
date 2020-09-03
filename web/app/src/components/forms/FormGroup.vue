/*
 * File Created: Wednesday, 15th July 2020 3:00:41 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <div class="form-group">
    <label for="exampleInputEmail1">{{title}} <i v-if="loading" class="fas fa-sync-alt fa-spin"></i>      
    </label>
    <slot />
  </div>  
</template>

<script>
export default {
    props: {
        title: {
            type: String,
            required: true
        },
        loading: {
            type: Boolean,
            required: true
        }
    }
}
</script>

<style>

</style>