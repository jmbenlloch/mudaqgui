<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { Greet, ScanDevices, DevicesMacs } from "../../wailsjs/go/main/App";

function doGreeting() {
  Greet("testfn").then((result) => {
    console.log(result)
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

    <div>
      <h3>Devices found</h3>
      <ul>
        <li v-for="mac in macs">{{ mac }}</li>
      </ul>

    </div>

  </div>
</template>