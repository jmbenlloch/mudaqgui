<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import * as d3 from "d3";

const width = ref(640)
const height = ref(400)
const marginTop = ref(20)
const marginRight = ref(20)
const marginBottom = ref(20)
const marginLeft = ref(2)

const data = ref([1, 2, 3, 4, 5, 6, 7])

const x = d3.scaleLinear([0, data.value.length - 1], [marginLeft.value, width.value - marginRight.value]);
const y = d3.scaleLinear(d3.extent(data.value), [height.value - marginBottom.value, marginTop.value]);
const line = d3.line((d, i) => x(i), y);

const yAxis = d3.scaleLinear()
  .domain([0, 100])
  .range([height.value - marginBottom.value, marginTop.value]);



onMounted(() => {
  EventsOn("events", (data) => {
    console.log("events", data)
  })
})

</script>

<template>
  <div>
    <h1>Test</h1>
    <svg :width="width" :height="height">
      <path fill="none" stroke="currentColor" stroke-width="1.5" :d="line(data)" />
      <g fill="white" stroke="currentColor" stroke-width="1.5">
        {data.map((d, i) => (
        <circle key={i} cx={x(i)} cy={y(d)} r="2.5" />))}
      </g>
    </svg>
  </div>
</template>