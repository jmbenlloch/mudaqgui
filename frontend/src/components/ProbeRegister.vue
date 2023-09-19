<script setup lang="ts">
import { ref, watch } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const store = useConfigStore()
const { probe, selectedCard, disableForms } = storeToRefs(store)

const probeRegister = ref(-1)
const nChannels = ref(32)

watch(probeRegister, (value) => {
  if (value > -1) {
    const array: Array<number> = Array(nChannels.value).fill(0)
    array[probeRegister.value] = 1
    probe.value[selectedCard.value].peakSensingHG = array
  }
})
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Probe register</h2>
    <div class="border border-primary w-fit px-3">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">Select channel</span>
        </label>
        <select v-model="probeRegister" class="select select-bordered" :disabled="disableForms">
          <option :value="-1">None</option>
          <option v-for="n in nChannels" :value="n - 1">Channel {{ n }}</option>
        </select>
      </div>
    </div>
  </div>
</template>