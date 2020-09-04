<template>
  <container-dashboard>
    <content-header title="Settings" :breadcrumbs="breadcrumbs" />
    <content-main>
      <!-- Wrap Course Reports -->
      <i-card title="Setting" addClass :useHeader="true" header="Featured" :minimize="false">
        <template #card-text>
          <i-nav-tab :navs="nav" />
          <!-- Form Canvas Setting -->
          <i-card title="">
            <template #card-link>
              <form @submit.prevent="handleSubmit()">
                <app-form-group :loading="status.loading" title="Canvas URL">
                    <app-input placeholder="Canvas URL" type="text" v-model="form.canvas_url" />
                </app-form-group>
                <app-form-group :loading="status.loading" title="Canvas Access Token">
                    <app-input placeholder="Canvas Access Token" type="text" v-model="form.canvas_token" />
                </app-form-group>     
                <app-btn type="submit" addClass="btn-success" title="Submit" />
              </form>
            </template>
          </i-card>
          <!-- End Form Canvas Setting -->
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
import AppInput from "@/components/forms/Input"
import AppFormGroup from "@/components/forms/FormGroup"
import AppBtn from "@/components/buttons/Button"
import {storeOrUpdateCanvasConfig, settings} from "@/api/settings/setting"
import iMixins from "@/helpers/mixins"

export default {
  mixins: [iMixins],    
  name: "Home",
  components: {
    ContainerDashboard,
    ContentHeader,
    ContentMain,
    ICard,
    INavTab,
    AppInput,
    AppFormGroup,
    AppBtn
  },
  data() {
    return {
      status: {
        loading: false
      },
      breadcrumbs: [
        { name: "Home", url: "/home" },
        { name: "Report Course", url: "#", isActive: true }
      ],
      form:{
        canvas_url: null,
        canvas_token: null
      },
      nav: [
        { name: "Canvas Settings", link: "setting.canvas", isActive: true},
      ],
    };
  },
  methods : {
    handleSubmit(){
      storeOrUpdateCanvasConfig(this.form)
      .then(res => {
        if(res.status == 200){
          this.$$_TOAST_SHOW("success","info", "Success save configuration canvas")
          let payload = {
            url : this.form.canvas_url,
            token : this.form.canvas_token
          }
          this.$store.commit("setting/CANVAS_CONFIG_EXISTS", payload)          
        }else{
          this.$$_TOAST_SHOW("danger","info", "Failed save configuration canvas")
        }
      })
      .catch(err => {
        console.log(err)
      })
    },
    getConfig(){
      settings()
      .then(res => {
        if (res.status >= 200 && res.data.length > 0) {
          let dataCanvasUrl = res.data.find(setting => setting.category == "canvas" && setting.name == "url")
          let dataCanvasToken = res.data.find(setting => setting.category == "canvas" && setting.name == "token")
          this.form.canvas_url = dataCanvasUrl.value
          this.form.canvas_token = dataCanvasToken.value
        }
      })
      .catch(err => {
        console.log(err)
      })
    }
  },
  mounted(){
    this.getConfig()
  }
};
</script>
