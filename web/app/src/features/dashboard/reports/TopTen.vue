/*
 * File Created: Monday, 6th July 2020 6:38:02 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

<template>
  <div>
    <i-info-box
      :subTitle="labelTitle"
      description="Complete results you can see in table at bellow"
      bg="bg-danger"
    />
    <chart-pie v-if="loaded" :chartData="chartData" :options="options" />
  </div>
</template>

<script>
import IInfoBox from "@/components/widgets/InfoBox";
import ChartPie from "@/components/charts/Pie";
import { courseAnalytics } from "@/api/reports/course";
export default {
  props: {
    account: {
      type: Object,
      required: true
    },
    date: {
      type: String
    },
    teacherAnalytics: {
      type: Boolean,
      required: true
    },
  },
  components: {
    IInfoBox,
    ChartPie
  },
  data() {
    return {
      loaded: false,
      chartData: {
        labels: ["Result"],
        datasets: [
          {
            label: "Top Ten",
            backgroundColor: ["#FFF"],
            data: [10]
          }
        ]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false
      }
    };
  },
  computed: {
    labelTitle() {
      var labelTop = this.teacherAnalytics == true ? "Lecturer" : "Course";
      return "Top 10 " + labelTop + " in Sub Account " + this.account.label;
    }
  },
  watch: {
    account: {
        deep: true,
        handler: function () {
            if (this.account.id != ""){
                this.getTopTen();
            }
        }
    },
    date: function(){
      if (this.date != null) {
        this.getTopTen()
      }
    }
  },
  methods: {
    getRandomColor() {
      var letters = "0123456789ABCDEF".split("");
      var color = "#";
      for (var i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    },
    clearDataChart() {
      this.chartData.labels = [];
      this.chartData.datasets[0].backgroundColor = [];
      this.chartData.datasets[0].data = [];
    },
    getTopTen: async function() {
      this.loaded = false;
      this.clearDataChart();
      var params = {
        account_id: this.account.id == "all" ? null : this.account.id,
        analytics_teacher: this.teacherAnalytics,
        date: this.date == null ? "" : this.date,
        page: "",
        limit: "",
        order_by: ""
      };
      let res = await courseAnalytics(params);
      
      if (res.status == 200 && res.data?.data != undefined) {
        this.loaded = true;
        let data = res.data.data;
        data.forEach(each => {
          this.chartData.labels.push(each.course_name);
          this.chartData.datasets[0].backgroundColor.push(
            this.getRandomColor()
          );
          this.chartData.datasets[0].data.push(each.final_score);
        });
      }
    }
  }
};
</script>

<style>
</style>