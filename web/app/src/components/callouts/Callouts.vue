/*
 * File Created: Thursday, 2nd July 2020 12:49:45 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
    <div class="callout" v-bind:class="addClass">
        <h5>{{ title }}</h5>
        <!-- <p> Lorem Ipsum </p> -->
        <slot />
    </div>  
</template>

<script>
export default {
    props : {
        addClass : {
            type: String
        },
        title: {
            type: String,
            required: true
        }
    }
}
</script>

<style>

</style>