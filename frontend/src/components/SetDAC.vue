<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const store = useConfigStore()
const { slowControl, disableForms, selectedCard } = storeToRefs(store)

function updateDAC(value: number) {
  slowControl.value[selectedCard.value].dac1_code = value
  slowControl.value[selectedCard.value].dac2_code = value
}
</script>

<template>
  <div class="border p-2 m-1">
    <h2 class="font-bold text-xl pl-2">DAC threshold</h2>
    <div class="w-fit px-3 mt-2">
      <div class="form-control max-w-xs">
        <NumericInput :value="slowControl[selectedCard].dac1_code" class="mx-1" @update-value="updateDAC"
          @increment="updateDAC" @decrement="updateDAC" :disabled="disableForms" :min="0" :max="1023" />
      </div>
    </div>
  </div>
</template>