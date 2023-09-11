<script setup lang="ts">
import { ref } from 'vue'

const probeRegister = ref([])
const nChannels = ref(32)

</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Probe register</h2>
    <div class="border border-primary w-fit px-3">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">Select channel</span>
        </label>
        <select class="select select-bordered">
          <option v-for="n in nChannels" :value="n">Channel {{ n }}</option>
        </select>
      </div>
    </div>
  </div>
</template>