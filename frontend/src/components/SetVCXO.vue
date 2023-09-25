<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { SetVCXO } from "../../wailsjs/go/main/App";
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'


const store = useConfigStore()
const { disableForms, selectedCard } = storeToRefs(store)

const vcxo = ref(0)

function updateVXCO(value: number) {
  vcxo.value = value
}

function setVCXO() {
  SetVCXO(selectedCard.value, vcxo.value).then(() => {
    console.log("get rate")
  });
}
</script>

<template>
  <div class="border p-2 m-1">
    <h2 class="font-bold text-xl pl-2">VCXO compensation</h2>
    <div class="w-fit px-3 flex flex-wrap gap-2 mt-2">
      <div class="form-control max-w-xs">
        <NumericInput :value="0" @update-value="updateVXCO" :min="0" :max="1024" :disabled="disableForms"/>
      </div>
      <button @click="setVCXO()" class="btn btn-primary btn-sm">VCXO</button>
    </div>
  </div>
</template>