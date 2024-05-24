<script setup lang="ts">
import { UpdateCardConfig } from "../../wailsjs/go/main/App";
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const store = useConfigStore()
const { slowControl, probe, selectedCard, disableForms } = storeToRefs(store)

function updateConfigFromUI() {
  UpdateCardConfig(selectedCard.value, slowControl.value[selectedCard.value], probe.value[selectedCard.value]).then(() => {
    console.log("test param")
  });
}
</script>

<template>
  <div class="border p-2 m-1">
    <h2 class="font-bold text-xl pl-2">Send card configuration</h2>
    <div class="w-fit px-3">
      <button @click="updateConfigFromUI()" class="btn btn-primary" :disabled="disableForms">Send card
        configuration</button>
    </div>
  </div>
</template>