<script setup lang="ts">
import EnableTrigger from '../components/EnableTrigger.vue';
import EnableAmplifier from '../components/EnableAmplifier.vue'
import ProbeRegister from '../components/ProbeRegister.vue'
import GainBias from '../components/GainBias.vue';
import MenuBar from '../components/MenuBar.vue';
</script>

<template>
  <main>
    <h1 class="text-sm">NEXT Muon DAQ</h1>

    <div class="grid grid-flow-col">
      <div class="w-1/6">
        <MenuBar />
      </div>
      <div class="flex gap-3">
        <EnableTrigger/>
        <ProbeRegister />
        <GainBias />
      </div>
    </div>
  </main>
</template>
