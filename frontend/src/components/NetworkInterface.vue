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
  <div class="border p-2 m-2">
    <div class="flex justify-between">
      <h2 class="font-bold text-xl pl-2">Network interface</h2>
      <div class="mx-4 mt-1 object-center pr-4">
        <div v-if="!connected" class="flex flex-wrap gap-1">
          <XCircleIcon class="text-red-600 w-6" />
          <span class="text-red-600">Not connected</span>
        </div>
        <div v-if="connected" class="flex flex-wrap gap-1">
          <CheckCircleIcon class="text-green-600 w-6" />
          <span class="text-green-600">Connected!</span>
        </div>
      </div>
    </div>

    <div class="w-fit px-3 flex items-end">
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">Select interface</span>
        </label>
        <select v-model="iface" class="select select-bordered">
          <option>None</option>
          <option v-for="item in interfaces" :value="item">{{ item }}</option>
        </select>
      </div>
      <button @click="startConnection" class="btn btn-primary mx-2">Start connection</button>
    </div>
 </div>
</template>