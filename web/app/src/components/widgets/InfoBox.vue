/*
 * File Created: Saturday, 4th July 2020 1:18:00 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
<div class="info-box" v-bind:class="bg">
    <span class="info-box-icon"><i v-bind:class="icon"></i></span>
    <div class="info-box-content">
        <span class="info-box-text">{{title}}</span>
        <span class="info-box-number">{{subTitle}}</span>

        <div class="progress">
            <div class="progress-bar" style="width: 70%"></div>
        </div>
        <span class="progress-description">
            {{description}}
        </span>
    </div>
    <!-- /.info-box-content -->
</div>  
</template>

<script>
export default {
    props : {
        bg: {
            type: String
        },
        icon: {
            type: String,
            default : () => { return "far fa-bookmark" }
        },
        title: {
            type: String
        },
        subTitle: {
            type: String
        },
        description: {
            type: String
        }
    }
}
</script>

<style>

</style>