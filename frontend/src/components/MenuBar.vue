<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { Greet, HVOn, HVOff, ReadData, PrintT0, GetRate, StartAcquisition, DevicesMacs } from "../../wailsjs/go/main/App";

function doGreeting() {
  Greet("testfn").then((result) => {
    console.log(result)
  });
}

function getRate() {
  GetRate().then(() => {
    console.log("get rate")
  });
}

function readData() {
  ReadData().then(() => {
    console.log("read data")
  });
}

function startRun() {
  StartAcquisition().then(() => {
    console.log("start run")
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

function scanDevices() {
  DevicesMacs().then((result) => {
    for (let mac of result){
      macs.value.push(mac)
    }
  });
}
</script>

<template>
  <div>
    <h2 class="font-bold text-xl">Menu</h2>

    <button @click="doGreeting()" class="btn btn-primary">Scan network</button>
    <button @click="scanDevices()" class="btn btn-primary">Show results</button>
    <button @click="getRate()" class="btn btn-primary">Get Rate</button>
    <button @click="startRun()" class="btn btn-primary">Start run</button>
    <button @click="hvOn()" class="btn btn-primary">HV on</button>
    <button @click="hvOff()" class="btn btn-primary">HV off</button>
    <button @click="readData()" class="btn btn-primary">Read data</button>
    <button @click="printT0()" class="btn btn-primary">T0</button>

    <div>
      <h3>Devices found</h3>
      <ul>
        <li v-for="mac in macs">{{ mac }}</li>
      </ul>

    </div>

  </div>
</template>