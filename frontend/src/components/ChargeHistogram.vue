<script setup lang="ts">
import { onMounted, ref, shallowRef, watch } from 'vue'
import type { Ref } from 'vue'
import type { ChartData, ChartOptions } from 'chart.js'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, PointElement, LineElement, CategoryScale, LinearScale } from 'chart.js'
import { range } from 'lodash'

import { useEventStore } from '../stores/events'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const eventStore = useEventStore()
const configStore = useConfigStore()
const { chargesRebin } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)


ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const props = defineProps({
  channel: { required: true, type: Number },
})

function createDataObject(): ChartData {
  let result: ChartData = {
    //    labels: range(0, 1024),
    labels: range(0, 128),
    datasets: [{
      data: chargesRebin.value[selectedCard.value].Charges[props.channel],
      backgroundColor: '#f87979',
    }]
  }
  return result
}

// @ts-ignore
const chartData: Ref<ChartData> = ref(createDataObject())

// @ts-ignore
const chartOptions: Ref<ChartOptions<'bar'>> = ref({
  // Turn off animations and data parsing for performance
  animation: false,
  //parsing: false,
  responsive: false,
  plugins: {
    legend: {
      display: false,
    }
  },
  scales: {
    x: {
      title : {
        display: true,
        text: "Charge (ADC counts)",
      },
    },
    y: {
      title : {
        display: true,
        text: "# Events",
      },
    },
  },
})

function updatePlot(){
  chartData.value = createDataObject()
}

watch(chargesRebin, (values) => {
  console.log("prop updated!")
  updatePlot()
})
</script>

<template>
  <div class="flex flex-col">
    <!-- @vue-ignore -->
    <Bar :options="chartOptions" :data="chartData" class="w-full h-36" />
  </div>
</template>

