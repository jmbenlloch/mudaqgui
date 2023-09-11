<script setup lang="ts">
import EnableTrigger from '@/components/EnableTrigger.vue';
import EnableAmplifier from '../components/EnableAmplifier.vue'
import ProbeRegister from '../components/ProbeRegister.vue'
import GainBias from '@/components/GainBias.vue';
</script>

<template>
  <main>
    <h1 class="text-sm">NEXT Muon DAQ</h1>

    <div class="flex gap-3">
      <EnableAmplifier/>
      <EnableTrigger/>
      <ProbeRegister />
      <GainBias />
    </div>
  </main>
</template>
