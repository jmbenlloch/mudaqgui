<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { range } from 'lodash';

const ampEnable: Ref<Array<Number>> = ref([])
const nChannels = ref(32)

function enableAll(){
  ampEnable.value = range(32)
}

function disableAll(){
  ampEnable.value = []
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Enable Amplifier</h2>
    <div class="border border-primary w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <input v-model="ampEnable" type="checkbox" class="checkbox" :value="n-1"/>
          <span class="label-text">Ch. {{ n-1 }}</span> 
        </label>
      </div>
    </div>
    <button @click="enableAll" class="btn btn-info m-2">Enable all</button>
    <button @click="disableAll" class="btn btn-warning m- m-22">Disable all</button>
  </div>
</template>