<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { UpdateGlobalConfig, HVOn, HVOff, StartRun, StopRun } from "../../wailsjs/go/main/App";

const store = useConfigStore()
const { slowControl, probe, disableForms } = storeToRefs(store)


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

function updateGlobalConfig() {
  UpdateGlobalConfig(slowControl.value, probe.value).then(() => {
    console.log("update all configs")
  });
}
</script>

<template>
  <div class="border p-2 m-2">
    <h2 class="font-bold text-xl pl-2">Run control</h2>
    <div class="flex flex-wrap flex-col gap-2">
      <div class="flex flex-row gap-2 m-2">
        <button @click="startRun()" class="btn btn-success w-1/2">Start run</button>
        <button @click="stopRun()" class="btn btn-error w-1/2">Stop run</button>
      </div>
      <button @click="updateGlobalConfig()" :disabled="disableForms" class="mx-2 btn btn-primary">Send configuration to all cards</button>
    </div>
 </div>
</template>