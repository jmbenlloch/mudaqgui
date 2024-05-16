<script setup lang="ts">
import {SelectCalibrationFile, LoadCalibrationFile } from "../../wailsjs/go/main/App";
import { onMounted, ref, shallowRef, watch } from 'vue'
import type { Ref } from 'vue'
import LogCalibration from "./LogCalibration.vue";

import { useEventStore } from '../stores/events'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const eventStore = useEventStore()
const configStore = useConfigStore()
const { chargesRebin } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)

const nChannels: Ref<number> = ref(32)
const configFile: Ref<string> = ref("")

function selectConfigurationFile() {
  SelectCalibrationFile().then((file: string) => {
    configFile.value = file
  });
}

function loadConfiguration() {
  LoadCalibrationFile(configFile.value).then(() => {
    console.log("saved")
  });
}
</script>

<template>
  <div class="flex flex-col">
    <div class="border p-2 m-2">
      <h2 class="text-xl font-bold pl-2">Select calibration file</h2>
      <div class="form-control max-w-xs">
        <label class="label">
          <span class="text-lg">Select file</span>
        </label>
        <div class="relative">
          <div class="absolute inset-y-0 left-0 flex items-center px-3 pointer-events-none
        bg-gray-800 border border-r-0 rounded-l-lg">
            <span class="text-gray-100">BROWSE</span>
          </div>
          <input @click="selectConfigurationFile" :value="`  ... ${configFile.slice(-20)}`" placeholder="Select file..."
            type="text" readonly class="input input-bordered block pl-24 w-full max-w-xs truncate ..." />
        </div>
        <div class="flex flex-row gap-2 m-2">
          <button @click="loadConfiguration" class="btn btn-primary mt-2 w-1/2">Load</button>
        </div>
      </div>
    </div>

    <LogCalibration />

  </div>
</template>