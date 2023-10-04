<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useRateStore } from '../stores/rate'
import { storeToRefs } from 'pinia'


const store = useRateStore()
const { rates, lostBuffer, lostFPGA, totalRate, totalLostBuffer, totalLostFPGA } = storeToRefs(store)
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Rates</h2>
    <div>
      <table class="table">
        <thead>
          <tr>
            <th>Card</th>
            <th>Rate (Hz)</th>
            <th>Events lost (buffer)</th>
            <th>Events lost (FPGA)</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(value, key) in rates">
            <th>{{ key }}</th>
            <td>{{ value }}</td>
            <td>{{ lostBuffer[key] }}</td>
            <td>{{ lostFPGA[key] }}</td>
          </tr>

          <tr>
            <th>Total</th>
            <td>{{ totalRate }}</td>
            <td>{{ totalLostBuffer }}</td>
            <td>{{ totalLostFPGA }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>