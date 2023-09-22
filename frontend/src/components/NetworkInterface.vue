<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import type { Ref } from 'vue'
import { GetNetworkInterfaces, StartConnection } from "../../wailsjs/go/main/App";

const interfaces: Ref<Array<string>> = ref([])
const iface : Ref<string> = ref("")

onMounted(() =>{
  GetNetworkInterfaces().then((result: Array<string>) => {
    interfaces.value = result
  });
})
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Network interface</h2>
    <div class="border border-primary w-fit px-3">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">Select interface</span>
        </label>
        <select v-model="iface" class="select select-bordered">
          <option>None</option>
          <option v-for="item in interfaces" :value="item">{{ item }}</option>
        </select>
      </div>
    </div>
    <button @click="StartConnection(iface)" class="btn btn-primary">Start connection</button>
  </div>
</template>