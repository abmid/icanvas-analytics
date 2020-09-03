<script>
import { Pie } from "vue-chartjs";

export default {
  extends: Pie,
  props: {
    chartData: {
      type: Object,
      default: null
    },
    options: {
      type: Object,
      default: null
    }
  },
  watch: {
    chartData() {
      this.$data._chart.update();
    }
  },
  mounted() {
    this.renderChart(this.chartData, this.options);
  }
};
</script>
