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
const { charges } = storeToRefs(eventStore)
const { selectedCard, disableForms } = storeToRefs(configStore)

const channel: Ref<number> = ref(0)
const nChannels: Ref<number> = ref(32)

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement)

const chartData: Ref<ChartData> = ref({
  labels: range(0, 1024),
  datasets: [{
    // @ts-ignore
    data: [],
    backgroundColor: '#f87979',
  }]
})


// @ts-ignore
const chartOptions: Ref<ChartOptions> = ref({
  // Turn off animations and data parsing for performance
  animation: false,
  //parsing: false,
  responsive: true,
  plugins: {
    legend: {
      display: false,
    }
  },
})

function updatePlot() {
  chartData.value = {
    labels: range(0, 1024),
    datasets: [{
      data: charges.value[selectedCard.value].Charges[channel.value],
      backgroundColor: '#f87979',
    }]
  }
}

watch(charges, (values) => {
  updatePlot()
})

watch(channel, (value) => {
  updatePlot()
})
</script>

<template>
  <div class="flex flex-col">
    <div class="form-control max-w-xs">
      <label class="label">
        <span class="label-text">Select channel</span>
      </label>
      <select v-model="channel" class="select select-bordered" :disabled="disableForms">
        <option :value="-1">None</option>
        <option v-for="n in nChannels" :value="n - 1">Channel {{ n }}</option>
      </select>
    </div>
    <!-- @vue-ignore -->
    <Bar :options="chartOptions" :data="chartData" width="1200" />
  </div>
</template>

