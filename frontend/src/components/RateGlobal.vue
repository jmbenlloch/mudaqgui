<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useRateStore } from '../stores/rate'
import { useEventStore } from '../stores/events'
import { storeToRefs } from 'pinia'


const store = useRateStore()
const eventStore = useEventStore()
const { rates, totalRate } = storeToRefs(store)
const { nEvents } = storeToRefs(eventStore)
</script>

<template>
  <div class="border p-2 m-2 flex flex-wrap">
    <div>
      <h2 class="font-bold text-xl pl-2">Rate</h2>
    </div>
    <div class="ml-8 p-1">
      <span>{{ totalRate }} Hz</span>
    </div>
    <div>
      <h2 class="font-bold text-xl pl-2">Events</h2>
    </div>
    <div class="ml-8 p-1">
      <span>{{ nEvents }}</span>
    </div>
  </div>
</template>