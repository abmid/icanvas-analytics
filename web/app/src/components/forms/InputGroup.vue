<template>
<div class="input-group" v-bind:class="addClass">
    <input v-model="myValue" :type="type" class="form-control" :placeholder="placeholder">
    <div class="input-group-append">
    <div class="input-group-text">
        <span v-bind:class="icon"></span>
    </div>
    </div>
</div>  
</template>

<script>
export default {
    props:{
        value : {
            type : String
        },
        addClass : {
            type: String
        },
        type : {
            type: String,
            required : true
        },
        placeholder :{
            type: String,
            required : true
        },
        icon :{
            type: String,
            default : () => { "fas fa-envelope" }
        }
    },
    computed : {
        myValue : {
            get(){
                return this.value
            }, 
            set(newValue){
                return this.$emit('update:value', newValue)
            }
        }
    }
}
</script>

<style>

</style>