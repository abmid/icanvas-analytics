<template>
  <container-dashboard>
    <content-header title="Dashboard Home" />
    <content-main>
      <app-card title="Testing card" addClass="card-warning card-outline" :useHeader="true" header="Featured">
          <template #card-text>
            <p>Now you see me</p>
          </template>
          <template #card-link>
            <a href="#" class="card-link">Card link</a>
            <a href="#" class="card-link">Another link</a>
          </template>
      </app-card>
    </content-main>
  </container-dashboard>
</template>

<script>
// @ is an alias to /src
import ContainerDashboard from "@/containers/Dashboard"
import ContentHeader from "@/components/contents/ContentHeader"
import ContentMain from "@/components/contents/ContentMain"
import AppCard from "@/components/cards/Card"
export default {
  name: 'Home',
  components: {
    ContainerDashboard,
    ContentHeader,
    ContentMain,
    AppCard
  }
}
</script>
