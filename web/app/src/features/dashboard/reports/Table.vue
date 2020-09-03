/*
 * File Created: Tuesday, 7th July 2020 5:31:06 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <div class="card card-warning card-outline">
    <div class="card-header ui-sortable-handle" style="cursor: move;">
      <h3 class="card-title">
        <i class="far fa-clipboard"></i>
        Reports in sub account
      </h3>

      <div class="card-tools">
        <div class="form-inline">
          <div class="form-group">
            <label>Search :&nbsp;</label>
            <input
              type="text"
              v-model="filterText"
              class="form-control"
              @keyup.enter="doFilter"
              placeholder="Search in here"
            />
          </div>
        </div>
      </div>
    </div>
    <div class="card-body" style="padding:0px !important;">
      <div>
        <!-- Table -->
        <vuetable
          ref="vuetable"
          :api-url="getUrl"
          :http-options="table.httpOptions"
          :fields="table.fields"
          :css="css.table"
          :queryParams="table.queryParams"
          :append-params="moreParams"
          pagination-path="pagination"
          data-path="data"
          @vuetable:pagination-data="onPaginationData"
        ></vuetable>
        <div class="vuetable-pagination ui basic segment grid icanvas-pagination">
          <vuetable-pagination-info :css="css.paginationInfo" ref="paginationInfo"></vuetable-pagination-info>

          <vuetable-pagination
            :css="css.pagination"
            ref="pagination"
            @vuetable-pagination:change-page="onChangePage"
          ></vuetable-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vuetable from "vuetable-2";
import VuetablePagination from "vuetable-2/src/components/VuetablePagination";
import VuetablePaginationInfo from "vuetable-2/src/components/VuetablePaginationInfo";
import CssConfig from "@/helpers/vuetable-bootstrap.js";
import _ from "lodash"
export default {
  props: {
    account: {
      type: Object,
      required: true
    },
    date: {
      type: String,
    },
    teacherAnalytics: {
      type: Boolean,
      required: true
    },
  },  
  components: {
    Vuetable,
    VuetablePagination,
    VuetablePaginationInfo
  },
  data() {
    return {
      table: {
        apiUrl : "http://localhost:8000/v1/analytics/courses?analytics_teacher=true",
        httpOptions: {
          withCredentials: true, // default
        },
        fields : [
          {
            name: "course_name",
            title: "Course Name"
          },
          {
            name: "teacher.name",
            title: "lecturer"
          },
          {
            name: "teacher.login_id",
            title: "email"
          },
          {
            name: "discussion_count",
            title: "Total Discussion"
          },
          {
            name: "assigment_count",
            title: "Total Assigment"
          },
          {
            name: "student_count",
            title: "Total Student"
          },
          {
            name: "average_grading",
            title: "Finish Grading Assigment"
          },
          {
            name: "final_score",
            title: "Final Score",
            callback: 'toFixed',
          }
        ],
        queryParams: {
          perPage: "limit",
          sort: "order_by",
          page: "page"
        }
      },
      filterText: "",
      css: CssConfig,
      moreParams: {}
    };
  },
  watch:{
    filterText(){
      this.doFilter()
    }
  },
  computed: {
    getUrl(){
      var query = ""
      if (this.account.id != "all" && this.account?.id != "undefined") {
        query = query + "&account_id=" + this.account.id
      }

      if (this.date != null) {
        query = query + "&date=" + this.date
      }
      
      return this.table.apiUrl + query
    }
  },
  methods: {
    //...
    toFixed(value){
      return value.toFixed(2) 
    },
    onPaginationData(paginationData) {
      this.$refs.pagination.setPaginationData(paginationData);
      this.$refs.paginationInfo.setPaginationData(paginationData); // <----
    },
    onChangePage(page) {
      this.$refs.vuetable.changePage(page);
    },
    doFilter : _.debounce(function(){
      console.log("doFilter:", this.filterText);
        this.moreParams = {
            'q': this.filterText
        }
        this.$nextTick( () => this.$refs.vuetable.refresh())      
      }, 500
    ),
    resetFilter() {
      this.filterText = "";
      console.log("resetFilter");
    }
  },
  mounted(){
    console.log(location.protocol + '//' + location.host)
  }
};
</script>

<style scoped>
.icanvas-pagination{
    padding: 0 19px;
    height: 46px;
    background: #fff !important;  
}
</style>