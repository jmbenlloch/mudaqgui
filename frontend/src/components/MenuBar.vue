<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'
import { ScanDevices, SetVCXO, SetDACThr, HVOn, HVOff, ReadData, UpdateConfig, PrintT0, GetRate, StartRun, StopRun } from "../../wailsjs/go/main/App";
import { EventsOn } from '../../wailsjs/runtime/runtime'


const store = useConfigStore()
const { cards, selectedCard } = storeToRefs(store)

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

function updateConfig() {
  UpdateConfig().then(() => {
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

function printT0() {
  PrintT0().then(() => {
    console.log("start run")
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

const rate = ref(0)
const card = ref(0)

onMounted(() => {
  EventsOn("rate", (data) => {
    console.log("event rate", data)
    rate.value = data.rate
    card.value = data.card
  })
})
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
    <button @click="printT0()" class="btn btn-primary">T0</button>
    <button @click="updateConfig()" class="btn btn-primary">config</button>
    <button @click="setDACThr()" class="btn btn-primary">DAC</button>

    <h1>Rate {{ card }}: {{ rate }}</h1>

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