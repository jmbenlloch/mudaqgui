<script setup lang="ts">
import EnableTrigger from '../components/EnableTrigger.vue';
import EnableAmplifier from '../components/EnableAmplifier.vue'
import ProbeRegister from '../components/ProbeRegister.vue'
import GainBias from '../components/GainBias.vue';
import SetDAC from '@/components/SetDAC.vue';
import SetVCXO from '@/components/SetVCXO.vue';
import SendCardConfig from '@/components/SendCardConfig.vue';
import HVControl from '@/components/HVControl.vue';
</script>

<template>
  <div class="flex gap-3">
    <EnableAmplifier />
    <EnableTrigger />
    <div class="flex flex-col gap-3">
      <ProbeRegister />
      <SetDAC />
      <SetVCXO />
    </div>
    <GainBias />
    <SendCardConfig />
    <HVControl />
  </div>
</template>
