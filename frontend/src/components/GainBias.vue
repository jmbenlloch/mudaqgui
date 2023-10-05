<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import type { Ref } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const nChannels : Ref<Number> = ref(32)
const gains: Ref<Array<number>> = ref(Array(32).fill(0))
const biases: Ref<Array<number>> = ref(Array(32).fill(0))

const store = useConfigStore()
const { slowControl, disableForms, selectedCard } = storeToRefs(store)

function updateGain(value: number, index: number) {
  slowControl.value[selectedCard.value].channel_preamp_HG[index] = value
}

function updateBias(value: number, index: number) {
  slowControl.value[selectedCard.value].input_dac[index] = value
}

function loadConfiguration(){
  console.log("update values gain")
  gains.value = slowControl.value[selectedCard.value].channel_preamp_HG
  biases.value = slowControl.value[selectedCard.value].input_dac
  console.log(gains.value)
  console.log(biases.value)
}

watch(slowControl, (value) => {
  loadConfiguration()
})

watch(selectedCard, (value) => {
  loadConfiguration()
})

onMounted(() => {
  loadConfiguration()
})
</script>

<template>
  <div class="border p-2 m-2 w-fit">
    <h2 class="font-bold text-xl pl-2">HG preamp gain/bias</h2>
    <div class="w-fit grid grid-cols-4">
      <div class="form-control" v-for="(n, index) in nChannels" :key="`gain-${index}`">
        <label class="label cursor-pointer">
          <span class="label-text">Ch. {{ n }}</span>
          <NumericInput :value="gains[parseInt(index)]" @update-value="updateGain($event, parseInt(index))"
            @increment="updateGain($event, parseInt(index))"
            @decrement="updateGain($event, parseInt(index))"
            class="mx-1" :min="0" :max="63" :disabled="disableForms"/>
          <NumericInput :value="biases[parseInt(index)]" @update-value="updateBias($event, parseInt(index))" class="mx-1"
            @increment="updateBias($event, parseInt(index))"
            @decrement="updateBias($event, parseInt(index))"
            :min="0" :max="254" :disabled="disableForms"/>
        </label>
      </div>
    </div>
  </div>
</template>