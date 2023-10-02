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
const { events, t0, t1 } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)


ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const chartData = ref({
  //labels: range(0, 2000),
  datasets: [{
    data: [],
    pointStyle: false as const,
    borderColor: '#FF6384',
  }]
})


// @ts-ignore
const chartOptions: Ref<ChartOptions<"line">> = ref({
  // Turn off animations and data parsing for performance
  animation: false,
  parsing: false,

  interaction: {
    mode: 'nearest',
    axis: 'x',
    intersect: false
  },
  plugins: {
    legend: {
      display: false,
    },
    decimation: {
      enabled: true,
      algorithm: 'lttb',
      samples: 200,
    }
  },
  scales: {
    x: {
      type: 'linear',
      title : {
        display: true,
        text: "Event",
      },
      suggestedMin: 0,
      suggestedMax: 100,
      ticks: {
        source: 'auto',
        // Disabled rotation for performance
        maxRotation: 0,
        autoSkip: true,
      }
    },
    y: {
      type: 'linear',
      title : {
        display: true,
        text: "t0",
      },
      suggestedMin: 0,
      suggestedMax: 100,
      ticks: {
        source: 'auto',
        // Disabled rotation for performance
        maxRotation: 0,
        autoSkip: true,
      }
    }
  }
})

watch(t1, (values) => {
  chartData.value = {
    //labels: range(values[selectedCard.value].length),
    datasets: [{
      data: values[selectedCard.value],
      pointStyle: false as const,
      borderColor: '#FF6384',
    }]
  }
  console.log(t1.value)
})
</script>

<template>
  <Line :options="chartOptions" :data="chartData" width="700" />
</template>