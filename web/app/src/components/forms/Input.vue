/*
 * File Created: Monday, 29th June 2020 6:40:53 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
    <input 
        :type="type" 
        class="form-control" 
        :placeholder="placeholder" 
        v-bind:value="value" 
        v-on:input="$emit('input', $event.target.value)">
</template>

<script>
export default {
    props : {
        value: {
            type: String,
        },
        type: {
            type: String,
            required: true
        },
        placeholder: {
            type: String,
            required: true
        }
    },    
}
</script>

<style>

</style>