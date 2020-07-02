/*
 * File Created: Monday, 29th June 2020 7:23:00 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <button v-bind:type="type" class="btn" v-bind:class="addClass">{{title}}</button>
</template>

<script>
export default {
    props : {
        title : {
            type : String,
            required : true
        },
        addClass : {
            type : String
        },
        type : {
            type : String,
            default : () => { "button" }
        }
    }
}
</script>

<style>

</style>