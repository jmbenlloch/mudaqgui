<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const triggerEnable: Ref<Array<number>> = ref([])
const nChannels = ref(32)

const store = useConfigStore()
const { slowControl, disableForms, selectedCard } = storeToRefs(store)

function loadConfiguration(){
  const array: Array<number> = []
  let mask = slowControl.value[selectedCard.value].discriminatorMask
  for (let i = 0; i < mask.length; i++) {
    if (mask[i] == 1){
      array.push(i)
    }
  }
  triggerEnable.value = array
}

watch(triggerEnable, (value) => {
  const array: Array<number> = Array(nChannels.value).fill(0)
  for (let ch of triggerEnable.value) {
    array[ch] = 1
  }
  slowControl.value[selectedCard.value].discriminatorMask = array
})

watch(slowControl, (value) => {
  loadConfiguration()
})

onMounted(() => {
  loadConfiguration()
})
</script>

<template>
  <div class="border p-2 m-1">
    <h2 class="font-bold text-xl pl-2">Enable Trigger</h2>
    <div class="w-fit grid grid-cols-4 gap-1">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <input v-model="triggerEnable" type="checkbox" class="checkbox checkbox-xs" :value="n - 1" :disabled="disableForms"/>
          <span class="label-text">Ch {{ n }}</span>
        </label>
      </div>
    </div>

    <div class="form-control w-fit">
      <label class="label cursor-pointer">
        <input v-model="slowControl[selectedCard].enable_or32" type="checkbox" class="checkbox checkbox-sm" :true-value="1" :false-value="0" :disabled="disableForms"/>
        <span class="label-text text-lg mx-3">Trigger OR 32</span>
      </label>
    </div>
  </div>
</template>