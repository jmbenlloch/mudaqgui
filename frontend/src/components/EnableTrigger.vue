<script setup lang="ts">
import { ref, watch } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const triggerEnable = ref([])
const nChannels = ref(32)

const store = useConfigStore()
const { slowControl, disableForms, selectedCard } = storeToRefs(store)

watch(triggerEnable, (value) => {
  const array: Array<number> = Array(nChannels.value).fill(0)
  for (let ch of triggerEnable.value) {
    array[ch] = 1
  }
  slowControl.value[selectedCard.value].discriminatorMask = array
})


</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Enable Trigger</h2>
    <div class="w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <input v-model="triggerEnable" type="checkbox" class="checkbox" :value="n - 1" :disabled="disableForms"/>
          <span class="label-text">Ch {{ n }}</span>
        </label>
      </div>
    </div>

    <div class="form-control">
      <label class="label cursor-pointer">
        <input v-model="slowControl[selectedCard].enable_or32" type="checkbox" class="checkbox" :true-value="1" :false-value="0" :disabled="disableForms"/>
        <span class="label-text text-lg">Trigger OR 32</span>
      </label>
    </div>
  </div>
</template>