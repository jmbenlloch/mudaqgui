<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';

const vcxo = ref(0)

function updateVXCO(value: number){
  console.log("Do something VCXO")
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
        <NumericInput :value="vcxo" @update-value="updateVXCO" class="mx-1"
          :min="0" :max="1024" />
      </div>
    </div>
  </div>
</template>