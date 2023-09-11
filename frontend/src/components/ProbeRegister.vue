<script setup lang="ts">
import { ref } from 'vue'

const probeRegister = ref([])
const nChannels = ref(32)

</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Probe register</h2>
    <div class="border border-primary w-fit grid grid-cols-4 gap-2">
      <div class="form-control" v-for="n in nChannels">
        <label class="label cursor-pointer">
          <input type="radio" class="radio" />
          <span class="label-text">Ch. {{ n }}</span> 
        </label>
      </div>
    </div>
  </div>
</template>