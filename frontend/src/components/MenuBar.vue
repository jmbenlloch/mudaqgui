<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { ScanDevices, UpdateGlobalConfig, WriteDataFile, HVOn, HVOff, StartRun, StopRun } from "../../wailsjs/go/main/App";
import Rate from './Rate.vue';
import NetworkInterface from './NetworkInterface.vue';
import LoadSaveConfig from './LoadSaveConfig.vue';


const store = useConfigStore()
const { slowControl, probe, cards, selectedCard, disableForms } = storeToRefs(store)

function scanDevices() {
  ScanDevices().then(() => {
    console.log("scan")
  });
}

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
  <div>
    <h2 class="font-bold text-xl">Menu</h2>

    <button @click="scanDevices()" class="btn btn-primary">Scan network</button>
    <button @click="startRun()" class="btn btn-primary">Start run</button>
    <button @click="stopRun()" class="btn btn-primary">Stop run</button>
    <button @click="hvOn()" class="btn btn-primary">HV on</button>
    <button @click="hvOff()" class="btn btn-primary">HV off</button>
    <button @click="updateGlobalConfig()" :disabled="disableForms" class="btn btn-primary">Send configuration to all cards</button>
    <button @click="writeData" class="btn btn-primary">Write data</button>

    <Rate />
    <NetworkInterface/>

    <div>
      <h3>Devices found</h3>
      <ul>
        <li v-for="card in cards">{{ card }}</li>
      </ul>

      <div class="form-control max-w-xs">
        <label class="label">
          <span class="label-text">Select card</span>
        </label>
        <select v-model="selectedCard" class="select select-bordered">
          <option v-for="card in cards" :value="card">Card {{ card }}</option>
        </select>
      </div>
    </div>

    <LoadSaveConfig />
  </div>
</template>