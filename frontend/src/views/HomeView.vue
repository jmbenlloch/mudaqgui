<script setup lang="ts">
import EnableTrigger from '../components/EnableTrigger.vue';
import EnableAmplifier from '../components/EnableAmplifier.vue'
import ProbeRegister from '../components/ProbeRegister.vue'
import GainBias from '../components/GainBias.vue';
import MenuBar from '../components/MenuBar.vue';
import SetDAC from '@/components/SetDAC.vue';
import SetVCXO from '@/components/SetVCXO.vue';
</script>

<template>
  <main>
    <h1 class="text-sm">NEXT Muon DAQ</h1>

    <div class="flex flex-row">
      <div class="w-1/6">
        <MenuBar />
      </div>
      <div class="flex gap-3">
        <EnableAmplifier/>
        <EnableTrigger/>
        <div class="flex flex-col gap-3">
          <ProbeRegister />
          <SetDAC />
          <SetVCXO />
        </div>
          <GainBias />
      </div>
    </div>
  </main>
</template>
