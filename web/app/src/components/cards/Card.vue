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
            <h5 class="card-title">{{ title }}</h5>
            <div class="card-tools">
                <slot name="generate-analytics"></slot>
                <button v-if="minimize" type="button" class="btn btn-tool" data-card-widget="collapse"><i class="fas fa-minus"></i>
                </button>
            </div>              
        </div>                      
        <div class="card-body">
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
        minimize: {
            type: Boolean
        }
    }
}
</script>

<style>

</style>