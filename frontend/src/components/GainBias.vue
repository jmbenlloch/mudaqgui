<script setup lang="ts">
import { ref } from 'vue'
import NumericInput from './NumericInput.vue';

const probeRegister = ref([])
const nChannels = ref(32)

</script>

<template>
  <div>
    <h2 class="font-bold text-xl">HG preamp gain/bias</h2>
    <div class="border border-primary w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <span class="label-text">Ch. {{ n }}</span> 
          <NumericInput class="mx-1" :min="0" :max="63"/>
          <NumericInput class="mx-1" :min="0" :max="254"/>
        </label>
      </div>
    </div>
  </div>
</template>