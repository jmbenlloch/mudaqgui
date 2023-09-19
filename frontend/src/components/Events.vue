<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
// TODO: solve issue with plotly.js-dist types. Currently:
// mv node_modules/@types/plotly.js node_modules/@types/plotly.js-dist
import Plotly from 'plotly.js-dist'
import { EventsOn } from '../../wailsjs/runtime/runtime'

const test = ref<HTMLElement>()
onMounted(() => {
  EventsOn("events", (data) => {
    console.log("events", data)
  })

  Plotly.newPlot(test.value as HTMLElement, [{
    x: [1, 2, 3, 4, 5],
    y: [1, 2, 4, 8, 16]
  }], {
    margin: { t: 0 }
  });


})

</script>

<template>
  <div>
    <h1>Test</h1>
    <div ref="test" style="width:600px;height:250px;"></div>
  </div>
</template>