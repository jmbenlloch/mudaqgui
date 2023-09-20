<script setup lang="ts">
import ChargeHistogram from '../components/ChargeHistogram.vue';
import { onMounted, ref, shallowRef, watch } from 'vue'
import type { Ref } from 'vue'
import { range } from 'lodash'

import { useEventStore } from '../stores/events'
import { useConfigStore } from '../stores/configuration'
import { storeToRefs } from 'pinia'

const eventStore = useEventStore()
const configStore = useConfigStore()
const { charges } = storeToRefs(eventStore)
const { selectedCard, disableForms } = storeToRefs(configStore)

const nChannels: Ref<number> = ref(32)
</script>

<template>
  <div class="flex flex-col">
    <div class="grid grid-cols-8">
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
      <p>Test</p>
    </div>
    <ChargeHistogram :data="range(0, 1024)" />
  </div>
</template>

