<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { ScanDevices, SetVCXO, SetDACThr, HVOn, HVOff, ReadData, UpdateConfig, PrintT0, GetRate, StartRun, StopRun, DevicesMacs } from "../../wailsjs/go/main/App";

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

function setVCXO() {
  SetVCXO().then(() => {
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

const macs: Ref<Array<String>>  = ref([])

function showDevices() {
  DevicesMacs().then((result) => {
    macs.value = []
    for (let mac of result){
      macs.value.push(mac)
    }
  });
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Menu</h2>

    <button @click="scanDevices()" class="btn btn-primary">Scan network</button>
    <button @click="showDevices()" class="btn btn-primary">Show results</button>
    <button @click="getRate()" class="btn btn-primary">Get Rate</button>
    <button @click="startRun()" class="btn btn-primary">Start run</button>
    <button @click="stopRun()" class="btn btn-primary">Stop run</button>
    <button @click="hvOn()" class="btn btn-primary">HV on</button>
    <button @click="hvOff()" class="btn btn-primary">HV off</button>
    <button @click="readData()" class="btn btn-primary">Read data</button>
    <button @click="printT0()" class="btn btn-primary">T0</button>
    <button @click="updateConfig()" class="btn btn-primary">config</button>
    <button @click="setVCXO()" class="btn btn-primary">VCXO</button>
    <button @click="setDACThr()" class="btn btn-primary">DAC</button>

    <div>
      <h3>Devices found</h3>
      <ul>
        <li v-for="mac in macs">{{ mac }}</li>
      </ul>

    </div>

  </div>
</template>