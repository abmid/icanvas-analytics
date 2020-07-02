<template>
<container-auth>
  <i-card-login :loading="status.loading">
      <!-- Notifications -->
      <i-callout v-if="errors.message" title="Login Failed" addClass="callout-danger">
        <p>{{errors.message}}</p>
      </i-callout>
      <form @submit.prevent="handleSubmit()" method="post">
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
        <div class="row">
          <div class="col-8">
            <div class="icheck-primary">
              <input type="checkbox" id="remember">
              <label for="remember">
                Remember Me
              </label>
            </div>
          </div>
          <!-- /.col -->
          <div class="col-4">
            <i-button type="submit" addClass="btn-primary btn-block" title="Sign In" />
          </div>
          <!-- /.col -->
        </div>
      </form>

      <div class="social-auth-links text-center mb-3">
        <p>- OR -</p>
        <router-link :to="{ name: 'dashboard.home'}">User</router-link>
        <a href="#" class="btn btn-block btn-danger">
          <i></i> Forgot my password
        </a>
      </div>
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
export default {
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
            email: '',
            password: ''
        },             
      };
  },
  methods : {
    handleSubmit(){
      this.errors.message = null
      this.status.loading = true
      this.$store.dispatch("auth/login", this.form).then(
        res => {
          if (res.status == 200) {
            this.$router.push({ name: 'dashboard.home', params: { userId: 123 }})
          }
          this.status.loading = false
          this.errors.message = res.data.message
        }
      ).catch(err => {
        this.status.loading = false
        this.errors.message = err.response.data.message
      })
    }
  }
}
</script>

<style>

</style>