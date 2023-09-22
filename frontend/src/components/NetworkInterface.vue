<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import type { Ref } from 'vue'
import { GetNetworkInterfaces, StartConnection } from "../../wailsjs/go/main/App";
import { CheckCircleIcon, XCircleIcon } from '@heroicons/vue/24/outline'

const interfaces: Ref<Array<string>> = ref([])
const iface: Ref<string> = ref("")
const connected: Ref<boolean> = ref(false)

function startConnection() {
  StartConnection(iface.value).then((result: boolean) => {
    console.log(result)
    connected.value = result
  })
}

onMounted(() => {
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
    <div v-if="!connected" class="flex flex-wrap gap-1">
      <XCircleIcon class="text-sm text-red-600 h-6 w-6" />
      <span class="text-red-600">Not connected</span>
    </div>
    <div v-if="connected" class="flex flex-wrap gap-1">
      <CheckCircleIcon class="text-sm text-green-600 h-6 w-6" />
      <span class="text-green-600">Connected!</span>
    </div>
    <button @click="startConnection" class="btn btn-primary">Start connection</button>
  </div>
</template>