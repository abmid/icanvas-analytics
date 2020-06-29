<template>
<container-auth>
  <i-card-login>
      <form @submit.prevent="handleSubmit()" method="post">
      <!-- Input Email -->
        <i-input-group 
          v-bind:value="form.email"
          v-on:update:value="form.email = $event"
          addClass="mb-3" 
          type="email" 
          placeholder="Email" 
          icon="fas fa-envelope"/>
        <!-- Input Password -->
        <i-input-group 
          v-bind:value="form.password"
          v-on:update:value="form.password = $event"
          addClass="mb-3" 
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
export default {
  components : {
    ContainerAuth,
    ICardLogin,
    IInputGroup,
    IButton
  },
  data() {
      return {
          form: {
              email: '',
              password: ''
          },             
      };
  },
  methods : {
    handleSubmit(){
      alert("OK")
    }
  }
}
</script>

<style>

</style>