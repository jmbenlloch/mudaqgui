<script setup lang="ts">
import { ref, watch } from 'vue'
import NumericInput from './NumericInput.vue';
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { update } from 'lodash';

const nChannels = ref(32)
const value = ref(0)

const store = useConfigStore()
const { slowControl } = storeToRefs(store)

function updateValue(value: Event, index: number){
  console.log("update")
  console.log(value)
  console.log(index)
  slowControl.value.channel_preamp_HG[index] = value
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">HG preamp gain/bias</h2>
    <div class="border border-primary w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="(n, index) in nChannels">
        <label class="label cursor-pointer">
          <span class="label-text">Ch. {{ n }}</span> 
          <NumericInput :value="value" @update-value="updateValue($event, index)" class="mx-1" :min="0" :max="63"/>
          <!-- <NumericInput class="mx-1" :min="0" :max="254"/> -->
        </label>
      </div>
    </div>
  </div>
  {{  slowControl.channel_preamp_HG }}
</template>