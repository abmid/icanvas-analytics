<template>
<div class="login-box">
  <div class="login-logo">
    iCanvas Analytics
  </div>
  <!-- /.login-logo -->
    <slot />
</div>  
</template>

<script>
export default {
  methods: {
    toggleBodyClass(addRemoveClass, className) {
      const el = document.body;

      if (addRemoveClass === 'addClass') {
        el.classList.add(className);
      } else {
        el.classList.remove(className);
      }
    },
  }, 
  mounted(){
    this.toggleBodyClass('removeClass', 'sidebar-mini');
     this.toggleBodyClass('addClass', 'login-page');
  } 
}
</script>

<style>
.login-page, .register-page{
  background-image: linear-gradient( 135deg, #FFE985 10%, #FA742B 100%) !important;
}
.login-logo{
  color: #fff !important;
  font-weight: bold;
}
</style>