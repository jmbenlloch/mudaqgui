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
  <div>
    <button @click="hvOn()" class="btn btn-primary">HV on</button>
    <button @click="hvOff()" class="btn btn-primary">HV off</button>
  </div>
</template>