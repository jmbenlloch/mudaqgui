<script setup lang="ts">
import { onMounted, ref, shallowRef, watch } from 'vue'
import type { Ref } from 'vue'
import type { ChartData, ChartOptions } from 'chart.js'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, PointElement, LineElement, CategoryScale, LinearScale } from 'chart.js'
import { range } from 'lodash'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const props = defineProps({
  data: { required: true, type: Array<number> },
})

function createDataObject(data: Array<number>): ChartData {
  let result: ChartData = {
    labels: range(0, 1024),
    datasets: [{
      data: props.data,
      backgroundColor: '#f87979',
    }]
  }
  return result
}

// @ts-ignore
const chartData: Ref<ChartData> = ref(createDataObject(props.data))

// @ts-ignore
const chartOptions: Ref<ChartOptions<'bar'>> = ref({
  // Turn off animations and data parsing for performance
  animation: false,
  //parsing: false,
  responsive: true
})

function updatePlot(data: Array<number>) {
  chartData.value = createDataObject(data)
}

watch(props.data, (values) => {
  updatePlot(props.data)
})
</script>

<template>
  <div class="flex flex-col">
    <!-- @vue-ignore -->
    <Bar :options="chartOptions" :data="chartData" width="1200" />
  </div>
</template>

