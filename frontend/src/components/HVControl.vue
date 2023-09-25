<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { HVOn, HVOff } from "../../wailsjs/go/main/App";


const store = useConfigStore()
const { slowControl, probe, cards, selectedCard, disableForms } = storeToRefs(store)

function hvOn() {
  HVOn(selectedCard.value).then(() => {
    console.log("hv on")
  });
}

function hvOff() {
  HVOff(selectedCard.value).then(() => {
    console.log("hv off")
  });
}
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">SiPM HV</h2>
    <div class="flex flex-wrap flex-col">
      <div class="flex flex-row gap-2 m-2">
        <button @click="hvOn()" class="btn btn-primary w-1/2">HV on</button>
        <button @click="hvOff()" class="btn btn-primary w-1/2">HV off</button>
      </div>
    </div>
  </div>
</template>