<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import BarPlot from './BarPlot.vue'
// TODO: solve issue with plotly.js-dist types. Currently:
// mv node_modules/@types/plotly.js node_modules/@types/plotly.js-dist

// Maybe change to c3.js?

import Plotly from 'plotly.js-dist'
import { EventsOn } from '../../wailsjs/runtime/runtime'

function addData(){
  xs.value.push(45)
  ys.value.push(32)
  console.log(plot.value)
  console.log(xs.value)
  console.log(ys.value)
  updatePlot()
}

const xs = ref([1,2,3,4,5])
const ys = ref([1,2,4,8,16])
const plot = ref()

function updatePlot(){
  plot.value = Plotly.react(test.value as HTMLElement, [{
    x: xs.value,
    y: ys.value,
  }], {
    margin: { t: 0 }
  });
}

const test = ref<HTMLElement>()
onMounted(() => {
  EventsOn("events", (data) => {
    console.log("events", data)
  })

  updatePlot()
})

</script>

<template>
  <div>
    <h1>Test</h1>
    <BarPlot />
    <div ref="test" style="width:600px;height:250px;"></div>

    <button class="btn btn-primary" @click="addData">Add data</button>
  </div>
</template>