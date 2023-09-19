<script setup lang="ts">
import { UpdateCardConfig } from "../../wailsjs/go/main/App";
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const store = useConfigStore()
const { slowControl, probe, selectedCard, disableForms } = storeToRefs(store)

function updateConfigFromUI() {
  UpdateCardConfig(selectedCard.value, slowControl.value, probe.value).then(() => {
    console.log("test param")
  });
}
</script>

<template>
  <div>
    <button @click="updateConfigFromUI()" class="btn btn-primary" :disabled="disableForms">Send card configuration</button>
  </div>
</template>