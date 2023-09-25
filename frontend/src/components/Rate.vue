<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useRateStore } from '../stores/rate'
import { storeToRefs } from 'pinia'


const store = useRateStore()
const { rates, totalRate } = storeToRefs(store)
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Rates</h2>
    <div>
      <table class="table">
        <!-- head -->
        <thead>
          <tr>
            <th>Card</th>
            <th>Rate (Hz)</th>
          </tr>
        </thead>
        <tbody>
          <!-- row 1 -->
          <tr v-for="(value, key) in rates">
            <th>{{ key }}</th>
            <td>{{ value }}</td>
          </tr>

          <tr>
            <th>Total</th>
            <td>{{ totalRate }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>