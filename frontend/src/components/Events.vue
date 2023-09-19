<script setup lang="ts">
import { onMounted, onUnmounted, ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'

onMounted(() => {
  EventsOn("events", (data) => {
    console.log("events", data)
  })
})

</script>

<template>
  <div>
    <h1>Test</h1>
  </div>
</template>