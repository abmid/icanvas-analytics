/*
 * File Created: Monday, 6th July 2020 4:04:03 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <div class="row">
    <div class="col">
      <f-filter-account
          v-bind:value="accountValue"
          v-on:update:value="accountValue = $event"           
      />
    </div>
    <div class="col">
      <f-filter-date 
          v-bind:value="dateValue"
          v-on:update:value="dateValue = $event"       
      />
    </div>
  </div>
</template>

<script>
import FFilterAccount from "./FilterAccount"
import FFilterDate from "./FilterDate"
export default {
  components: {
    FFilterAccount,
    FFilterDate
  },
  props : {
    account: {
      type: Object
    },
    date: {
      type: String
    }
  },
  data() {
    return {
    };
  },
  computed : {
    accountValue : {
        get(){
            return this.account
        }, 
        set(newValue){
            return this.$emit('update:account', newValue)
        }
    },
    dateValue : {
        get(){
            return this.date
        }, 
        set(newValue){
            return this.$emit('update:date', newValue)
        }      
    }  
  },
  methods: {
  },
};
</script>

<style>
</style>