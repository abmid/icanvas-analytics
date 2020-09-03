<template>
  <container-dashboard>
    <content-header title="Settings" :breadcrumbs="breadcrumbs" />
    <content-main>
      <!-- Wrap Course Reports -->
      <i-card title="Setting" addClass :useHeader="true" header="Featured" :minimize="false">
        <template #card-text>
          <i-nav-tab :navs="nav" />
          <i-card />
        </template>
      </i-card>
    </content-main>
  </container-dashboard>
</template>

<script>
// @ is an alias to /src
import ContainerDashboard from "@/containers/Dashboard";
import ContentHeader from "@/components/contents/ContentHeader";
import ContentMain from "@/components/contents/ContentMain";
import ICard from "@/components/cards/Card";
import INavTab from "@/components/nav/NavTab"
import axios from "axios"
export default {
  name: "Home",
  components: {
    ContainerDashboard,
    ContentHeader,
    ContentMain,
    ICard,
    INavTab
  },
  data() {
    return {
      breadcrumbs: [
        { name: "Home", url: "/home" },
        { name: "Report Course", url: "#", isActive: true }
      ],
      form: [],
      nav: [
        { name: "Canvas Settings", link: "setting.canvas", isActive: true},
      ],
    };
  },
  methods : {
    submit(){
      var data = null
      data = [
        {"canvas_url" : "test"},
        {"canvas_token" : "test"}
      ]
      var setting = {
        "setting" : data
      }
      axios.post("/test", setting)
      .then()
      .catch()
    }
  },
  mounted(){
    this.submit()
    // setting[canvas_url] = abc, setting[canvas_token] = abc
  }
};
</script>
