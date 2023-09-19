<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const store = useConfigStore()
const { slowControl } = storeToRefs(store)

function updateDAC(value: number){
  slowControl.value.dac1_code = value
  slowControl.value.dac2_code = value
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">DAC threshold</h2>
    <div class="border border-primary w-fit px-3">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">DAC Threshold</span>
        </label>
        <NumericInput :value="slowControl.dac1_code" @update-value="updateDAC" class="mx-1"
          :min="0" :max="1024" />
      </div>
    </div>
  </div>
</template>