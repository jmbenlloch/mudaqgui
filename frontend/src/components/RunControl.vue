<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { UpdateGlobalConfig, WriteDataFile, HVOn, HVOff, StartRun, StopRun } from "../../wailsjs/go/main/App";

const store = useConfigStore()
const { slowControl, probe, disableForms } = storeToRefs(store)


function writeData() {
  WriteDataFile().then(() => {
    console.log("data written")
  });
}

function startRun() {
  StartRun().then(() => {
    console.log("start run")
  });
}

function stopRun() {
  StopRun().then(() => {
    console.log("stop run")
  });
}

const broadcastAddress : Ref<number> = ref(255)

function hvOn() {
  HVOn(broadcastAddress.value).then(() => {
    console.log("hv on")
  });
}

function hvOff() {
  HVOff(broadcastAddress.value).then(() => {
    console.log("hv off")
  });
}

function updateGlobalConfig() {
  UpdateGlobalConfig(slowControl.value, probe.value).then(() => {
    console.log("update all configs")
  });
}
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Run control</h2>
    <div class="flex flex-wrap flex-col">
      <div class="flex flex-row gap-2 m-2">
        <button @click="startRun()" class="btn btn-success w-1/2">Start run</button>
        <button @click="stopRun()" class="btn btn-error w-1/2">Stop run</button>
      </div>
      <button @click="updateGlobalConfig()" :disabled="disableForms" class="btn btn-primary">Send configuration to all cards</button>
      <button @click="writeData" class="btn btn-primary">Write data</button>
    </div>
 </div>
</template>