<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useRateStore } from '../stores/rate'
import { storeToRefs } from 'pinia'


const store = useRateStore()
const { rates, totalRate } = storeToRefs(store)
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Rates</h2>
    <div>
      <ul>
        <li v-for="(value, key) in rates">{{ key }} : {{ value }} Hz</li>
        <li>Total rate: {{ totalRate }} Hz</li>
      </ul>
    </div>
  </div>
</template>