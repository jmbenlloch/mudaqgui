<script setup lang="ts">
import { onMounted, ref, shallowRef, watch } from 'vue'
import type { Ref } from 'vue'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, PointElement, LineElement, CategoryScale, LinearScale } from 'chart.js'
import { range } from 'lodash'

import { useEventStore } from '../stores/events'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const eventStore = useEventStore()
const configStore = useConfigStore()
const { charges } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const chartData = ref({
  labels: range(0, 1024),
  datasets: [{
    // @ts-ignore
    data: [],
    backgroundColor: '#f87979',
  }]
})


// @ts-ignore
const chartOptions: Ref<ChartOptions<"line">> = ref({
  // Turn off animations and data parsing for performance
  animation: false,
  //parsing: false,
  responsive: true
})

watch(charges, (values) => {
  chartData.value = {
    labels: range(0, 1024),
    datasets: [{
      data: charges.value[selectedCard.value].Charges[0],
      backgroundColor: '#f87979',
    }]
  }
})
</script>

<template>
  <Bar :options="chartOptions" :data="chartData" width="1200" />
</template>

