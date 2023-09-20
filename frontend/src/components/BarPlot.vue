<script setup lang="ts">
import { onMounted, ref, shallowRef, watch } from 'vue'
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

function addData() {
  chartData.value = {
    labels: ['January', 'February', 'March', 'Test'],
    datasets: [{ data: [40, 20, 12, 24] }]
  }
  console.log(chartData.value.labels)
  console.log(chartData.value.datasets)
}

const chartData = ref({
  labels: ['January', 'February', 'March'],
  datasets: [{ data: [40, 20, 12] }]
})
const chartOptions = ref({
  responsive: true
})
</script>

<template>
  <Bar id="my-chart-id" :options="chartOptions" :data="chartData" />
  <button class="btn btn-primary" @click="addData">Add data 2</button>
</template>

