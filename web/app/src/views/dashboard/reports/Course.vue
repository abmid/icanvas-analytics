<template>
  <container-dashboard>
    <content-header title="Course Reports" :breadcrumbs="breadcrumbs" />
    <content-main>
      <!-- Wrap Course Reports -->
      <i-card title="Course Reports" addClass :useHeader="true" header="Featured" :minimize="false">
        <template #card-text>
          <!-- Filter -->
          <i-card title="Filter" :useHeader="true" :minimize="true">
            <template #card-text>
              <!-- Filter by Subaccount -->
              <f-filter
                v-bind:account="filter.account"
                v-on:update:account="filter.account = $event"        
                v-bind:date="filter.date"
                v-on:update:date="filter.date = $event"                        
               />
            </template>
          </i-card>
          <!-- Info Box for Top 10 -->
          <f-top-ten 
            v-show="filter.account.id != ''" 
            :account="filter.account" 
            :date="filter.date"
            :teacherAnalytics="filter.teacher_analytics"/>
          <!-- Report Table -->
          <f-table 
            v-if="filter.account.id != ''" 
            :account="filter.account" 
            :date="filter.date"
            :teacherAnalytics="filter.teacher_analytics"  />
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
import FTopTen from "@/features/dashboard/reports/TopTen"
import FFilter from "@/features/dashboard/reports/Filter"
import FTable from "@/features/dashboard/reports/Table"
export default {
  name: "Home",
  components: {
    ContainerDashboard,
    ContentHeader,
    ContentMain,
    ICard,
    FFilter,
    FTopTen,
    FTable
  },
  data() {
    return {
      breadcrumbs: [
        { name: "Home", url: "/home" },
        { name: "Report Course", url: "#", isActive: true }
      ],
      filter: {
        account: {
          id: "",
          label: "Select sub account"
        },
        teacher_analytics: false,
        date: null
      },
    };
  },
};
</script>
