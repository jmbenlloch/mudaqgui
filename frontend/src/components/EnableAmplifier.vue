<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Ref } from 'vue'
import { range } from 'lodash';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const ampEnable: Ref<Array<number>> = ref([])
const nChannels = ref(32)

function enableAll(){
  ampEnable.value = range(32)
}

function disableAll(){
  ampEnable.value = []
}

const store = useConfigStore()
const { slowControl, selectedCard, disableForms } = storeToRefs(store)

watch(ampEnable, (value) => {
  console.log(value)
  const array: Array<number> = Array(nChannels.value).fill(1)
  for (let ch of ampEnable.value){
    console.log(ch)
    array[ch] = 0
  }
  console.log(array)
  slowControl.value[selectedCard.value].channel_preamp_disable = array
})

watch(slowControl, (value) => {
  const array: Array<number> = []
  let channel_preamp_disable = slowControl.value[selectedCard.value].channel_preamp_disable
  console.log("config amp enable: ", channel_preamp_disable)
  for (let i = 0; i < channel_preamp_disable.length; i++) {
    if (channel_preamp_disable[i] == 0){
      array.push(i)
    }
  }
  ampEnable.value = array
  console.log("amp enable: ", ampEnable.value)
})
</script>

<template>
  <div class="border p-2 m-1">
    <h2 class="font-bold text-xl pl-2">Enable Amplifier</h2>
    <div class="w-fit grid grid-cols-4 gap-1">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <input v-model="ampEnable" type="checkbox" class="checkbox checkbox-xs" :value="n-1" :disabled="disableForms"/>
          <span class="label-text">Ch. {{ n-1 }}</span> 
        </label>
      </div>
    </div>
    <button @click="enableAll" :disabled="disableForms" class="btn btn-info m-2">Enable all</button>
    <button @click="disableAll" :disabled="disableForms" class="btn btn-warning m-2">Disable all</button>
  </div>
</template>