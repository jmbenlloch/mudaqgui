<script setup lang="ts">
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { ScanDevices } from "../../wailsjs/go/main/App";

const store = useConfigStore()
const { cards, selectedCard } = storeToRefs(store)

function scanDevices() {
  ScanDevices().then(() => {
    console.log("scan")
  });
}
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Devices</h2>
    <div class="form-control max-w-xs">
      <label class="label">
        <span class="label-text">Select card</span>
      </label>
      <select v-model="selectedCard" class="select select-bordered">
        <option v-for="card in cards" :value="card">Card {{ card }}</option>
      </select>
    </div>

    <div class="flex flex-wrap gap-2 m-2">
      <h3 class="font-bold">Devices found:</h3>
      <div>
        <span v-for="card in cards">{{ card }}, </span>
        <span v-if="cards.length == 0">None</span>
      </div>
    </div>
    <button @click="scanDevices()" class="btn btn-primary w-full">Scan network</button>
  </div>
</template>