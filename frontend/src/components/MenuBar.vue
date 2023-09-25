<script setup lang="ts">
import RateGlobal from './RateGlobal.vue';
import NetworkInterface from './NetworkInterface.vue';
import LoadSaveConfig from './LoadSaveConfig.vue';
import SelectCard from './SelectCard.vue';
import RunControl from './RunControl.vue';
import HVControlGlobal from './HVControlGlobal.vue';
</script>

<template>
  <div>
    <RunControl />
    <HVControlGlobal />
    <RateGlobal />
    <NetworkInterface/>
    <SelectCard />
    <LoadSaveConfig />
  </div>
</template>