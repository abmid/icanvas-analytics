<template>
  <button v-bind:type="type" class="btn" v-bind:class="addClass">Sign In</button>
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