/*
 * File Created: Wednesday, 15th July 2020 3:25:09 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <app-form-group :loading="false" title="Date">
    <app-date 
      v-model="myValue" 
      :only-date="true" 
      format="YYYY-MM-DD" 
      formatted="ll" 
      :auto-close="true"/>
  </app-form-group>
</template>

<script>
import AppDate from "vue-ctk-date-time-picker";
import AppFormGroup from "@/components/forms/FormGroup";
export default {
  components: {
    AppDate,
    AppFormGroup
  },
  props: {
    value: {
      type: String
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
  },  
};
</script>

<style>
</style>