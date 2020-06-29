<template>
<div class="wrapper">
  <!-- Header -->
  <app-header />
  <!-- End : Header -->
  <!-- Sidebar -->
  <app-sidebar />
  <!-- End : Sidebar -->
  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
      <slot />
  </div>
  <!-- /.content-wrapper -->
  <!-- Sidebar Control -->
  <app-sidebar-control />
  <!-- End : Sidebar Control -->
  <!-- Footer -->
  <app-footer />
  <!-- End : Footer -->

</div>
</template>

<script>
// @ is an alias to /src
import AppHeader from "@/components/template/Header"
import AppSidebar from "@/components/template/Sidebar"
import AppFooter from "@/components/template/Footer"
import AppSidebarControl from "@/components/template/SidebarControl"
export default {
  name: 'Home',
  components: {
    AppHeader,
    AppSidebar,
    AppFooter,
    AppSidebarControl
  },
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
    this.toggleBodyClass('removeClass', 'login-page');
     this.toggleBodyClass('addClass', 'sidebar-mini');
  }   
}
</script>
