/*
 * File Created: Monday, 29th June 2020 7:26:47 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <container-auth>
    <i-card-login :loading="status.loading" title="Welcome to iCanvas Analytics Dashboard, please create first account to start your sessions">
        <!-- Notifications -->
        <i-callout v-if="errors.message" title="Something wrong" addClass="callout-danger">
            <p>{{errors.message}}</p>
        </i-callout>
        <form @submit.prevent="handleSubmit()" method="post">
            <!-- Input Email -->
            <i-input-group 
            v-bind:value="form.name"
            v-on:update:value="form.name = $event"
            addClass="mb-3" 
            :required="true"
            type="text" 
            placeholder="Full name" 
            icon="fas fa-user"/>            
            <!-- Input Email -->
            <i-input-group 
            v-bind:value="form.email"
            v-on:update:value="form.email = $event"
            addClass="mb-3" 
            :required="true"
            type="email" 
            placeholder="Email" 
            icon="fas fa-envelope"/>
            <!-- Input Password -->
            <i-input-group 
            v-bind:value="form.password"
            v-on:update:value="form.password = $event"
            addClass="mb-3" 
            :required="true"
            type="password" 
            placeholder="Password" 
            icon="fas fa-lock"/>  
           <!-- Input Repeat Password -->
            <i-input-group 
            v-bind:value="form.passwordRepeat"
            v-on:update:value="form.passwordRepeat = $event"
            addClass="mb-3" 
            :required="true"
            type="password" 
            placeholder="Repeat password" 
            icon="fas fa-lock"/>                       
            <!-- Button Action -->
            <div class="row">
                <div class="col-6">
                </div>
                <!-- /.col -->
                <div class="col-6">
                    <i-button type="submit" addClass="btn-success btn-block" title="Create Account" />
                </div>
                <!-- /.col -->
            </div>
        </form>
        <!-- /.social-auth-links -->    
    </i-card-login>      
  </container-auth>
</template>

<script>
import ContainerAuth from "@/containers/Auth"
import ICardLogin from "@/components/cards/CardLogin"
import IInputGroup from "@/components/forms/InputGroup"
import IButton from "@/components/buttons/Button"
import ICallout from "@/components/callouts/Callouts"
import iMixins from "@/helpers/mixins"
export default {
    mixins: [iMixins],    
    components : {
        ContainerAuth,
        ICardLogin,
        IInputGroup,
        IButton,
        ICallout
    },
    data() {
        return {
            errors: {
                message : null
            },
            status : {
                loading : false
            },
            form: {
                name: '',
                email: '',
                password: '',
                passwordRepeat: ''
            },             
        };
    }, 
    watch : {
        "form.passwordRepeat" : function(){
            if (this.form.passwordRepeat != "" && this.form.passwordRepeat != null) {
                this.checkPasswordMatch()
            }else{
                this.errors.message = null
            }
        },
        "form.password" : function(){
            if (this.form.passwordRepeat != "" && this.form.passwordRepeat != null) {
                this.checkPasswordMatch()
            }
        }
    },
    methods : {
        checkPasswordMatch(){
            if (this.form.password != null && this.form.password != "") {
                if (this.form.password != this.form.passwordRepeat) {
                    this.errors.message = "Password and repeat password not same, please fix this"
                }else{
                    this.errors.message = null
                }
            }
        },
        checkForm(e){
            this.errors.message = null;
            if (this.form.password != this.form.passwordRepeat) {
                this.errors.message = "Password and Repeat Password not same !"
            }

            if (!this.errors.message) {
                return true;
            }

            e.preventDefault();            
        },
        handleSubmit(e){
            if (this.checkForm(e)) {
                this.status.loading = true
                this.errors.message = null
                this.$store.dispatch("auth/register", this.form).
                then(res => {
                    if (res.status == 201) {
                        this.$$_TOAST_SHOW("default","Info", "Your account administrator successfully created, please login to start")
                        this.$router.push({name : "login"})
                    }
                    this.status.loading = false
                }).
                catch(err => {
                    this.errors.message = err.response.data.message
                    this.status.loading = false
                })
            }
        }
    },
    beforeMount(){
        this.$store.dispatch("auth/registerCheck")
        .then(res => {
            if (res.status == 200) {
                if (res.data.status == false) {
                    this.$router.push({name: "login"})
                }
            }
        })
        .catch(err => {
            console.log(err)
        })
    }  
}
</script>

<style>

</style>