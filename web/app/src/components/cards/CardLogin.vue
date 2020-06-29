<template>
  <div class="card">
    <div class="card-body login-card-body">
      <p class="login-box-msg">Sign in to start your session</p>
        <slot />
    </div>
    <!-- /.login-card-body -->
  </div>  
</template>

<script>
export default {

}
</script>

<style scoped>
.card, .card-body{
  border-radius: 8px !important;
}
.card-body{
  box-shadow: 6px 6px 7px 2px #0000001c;
}
</style>