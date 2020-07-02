/*
 * File Created: Monday, 29th June 2020 5:48:20 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
    <div class="card" v-bind:class="addClass">
        <!-- props use header el -->
        <div v-if="useHeader" class="card-header">
            <h5 class="m-0">{{ header }}</h5>
        </div>        
        <div class="card-body">
        <h5 class="card-title">Card title</h5>
            <p class="card-text">
                <!-- Slot for content Card -->
                <slot name="card-text"></slot>
            </p>
            <!-- Additional content link or footer -->
            <slot name="card-link"></slot>
        </div>
    </div>  
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
        useHeader : {
            type : Boolean,
            default : () => { false }
        },
        header : {
            type : String
        }
    }
}
</script>

<style>

</style>