<script setup lang="ts">
import { onMounted, ref, shallowRef, watch } from 'vue'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, PointElement, LineElement, CategoryScale, LinearScale } from 'chart.js'
import { range } from 'lodash'

import { useEventStore } from '../stores/events'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const eventStore = useEventStore()
const configStore = useConfigStore()
const { events, t0, t1 } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)


ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const chartData = ref({
  labels: [1, 2, 3],
  datasets: [{
    data: [40, 20, 12],
    pointStyle: false as const,
  }]
})
const chartOptions = ref({
  responsive: true
})

watch(t1, (values) => {
  chartData.value = {
    labels: range(values[selectedCard.value].length),
    datasets: [{
      data: values[selectedCard.value],
      pointStyle: false as const,
    }]
  }
  console.log(t1.value)
})
</script>

<template>
  <Line :options="chartOptions" :data="chartData" width="700" />
</template>

