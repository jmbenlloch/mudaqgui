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
  <div class="flex flex-col flex-wrap">
    <div class="flex gap-1">
      <EnableAmplifier />
      <EnableTrigger />
      <div class="flex flex-col gap-1">
        <ProbeRegister />
        <SetDAC />
        <SetVCXO />
      </div>
      <div class="flex flex-col gap-1">
        <SendCardConfig />
        <HVControl />
      </div>
    </div>
    <div>
      <GainBias />
    </div>
  </div>
</template>
