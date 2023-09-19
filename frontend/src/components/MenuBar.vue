<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { ScanDevices, UpdateGlobalConfig, SetDACThr, HVOn, HVOff, ReadData, GetRate, StartRun, StopRun } from "../../wailsjs/go/main/App";
import { EventsOn } from '../../wailsjs/runtime/runtime'
import Rate from './Rate.vue';


const store = useConfigStore()
const { slowControl, probe, cards, selectedCard } = storeToRefs(store)

function scanDevices() {
  ScanDevices().then(() => {
    console.log("scan")
  });
}

function getRate() {
  GetRate().then(() => {
    console.log("get rate")
  });
}

function setDACThr() {
  SetDACThr().then(() => {
    console.log("get rate")
  });
}

function readData() {
  ReadData().then(() => {
    console.log("read data")
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

function hvOn() {
  HVOn().then(() => {
    console.log("hv on")
  });
}

function hvOff() {
  HVOff().then(() => {
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
    <button @click="getRate()" class="btn btn-primary">Get Rate</button>
    <button @click="startRun()" class="btn btn-primary">Start run</button>
    <button @click="stopRun()" class="btn btn-primary">Stop run</button>
    <button @click="hvOn()" class="btn btn-primary">HV on</button>
    <button @click="hvOff()" class="btn btn-primary">HV off</button>
    <button @click="readData()" class="btn btn-primary">Read data</button>
    <button @click="setDACThr()" class="btn btn-primary">DAC</button>
    <button @click="updateGlobalConfig()" class="btn btn-primary">Send configuration to all cards</button>

    <Rate />

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

  </div>
</template>