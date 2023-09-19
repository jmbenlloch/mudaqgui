<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { SetVCXO } from "../../wailsjs/go/main/App";
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'


const store = useConfigStore()
const { disableForms } = storeToRefs(store)

const vcxo = ref(0)

function updateVXCO(value: number) {
  console.log("Do something VCXO")
}

function setVCXO() {
  SetVCXO().then(() => {
    console.log("get rate")
  });
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">VCXO compensation</h2>
    <div class="border border-primary w-fit px-3">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">VCXO value</span>
        </label>
        <NumericInput :value="vcxo" @update-value="updateVXCO" class="mx-1" :min="0" :max="1024" :disabled="disableForms"/>
      </div>
      <button @click="setVCXO()" class="btn btn-primary">VCXO</button>
    </div>
  </div>
</template>