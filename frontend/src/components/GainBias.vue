<script setup lang="ts">
import { ref } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const nChannels = ref(32)

const store = useConfigStore()
const { slowControl, disableForms } = storeToRefs(store)

function updateGain(value: number, index: number) {
  slowControl.value.channel_preamp_HG[index] = value
}

function updateBias(value: number, index: number) {
  slowControl.value.input_dac[index] = value
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">HG preamp gain/bias</h2>
    <div class="border border-primary w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="(n, index) in nChannels">
        <label class="label cursor-pointer">
          <span class="label-text">Ch. {{ n }}</span>
          <NumericInput :value="slowControl.channel_preamp_HG[index]" @update-value="updateGain($event, index)"
            class="mx-1" :min="0" :max="63" :disabled="disableForms"/>
          <NumericInput :value="slowControl.input_dac[index]" @update-value="updateBias($event, index)" class="mx-1"
            :min="0" :max="254" :disabled="disableForms"/>
        </label>
      </div>
    </div>
  </div>
</template>