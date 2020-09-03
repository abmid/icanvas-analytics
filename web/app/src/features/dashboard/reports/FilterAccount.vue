/*
 * File Created: Wednesday, 15th July 2020 3:25:00 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <app-form-group :loading="status.loading" title="Sub Account">
    <v-select v-model="myValue" :options="filter.options" placeholder="Select Sub Account" />
  </app-form-group>
</template>

<script>
import {listAccount} from "@/api/canvas/accounts.js"
import AppFormGroup from "@/components/forms/FormGroup";
export default {
  components: {
    AppFormGroup
  },
  props : {
    value: {
      type: Object
    },
  },
  data() {
    return {
      status: {
        loading: true
      },
      filter: {
        options: []
      }
    };
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
  methods: {
      getListAccount : async function() {
          let res = await listAccount(1)
          if (res.status == 200) {
              this.filter.options = [
                { label: "Select Sub Account", id: ""},
                { label: "All", id: "all"}                
              ]
              var data = res.data
              data.forEach(each => {
                this.filter.options.push({
                  label: each.name,
                  id: each.id
                })
              });
          }
          this.status.loading = false
      }
  },
  mounted(){
      this.getListAccount()
  }
};
</script>

<style>
.vs__dropdown-toggle{
    height: 42px !important;
    min-height: 42px !important;   
}
</style>