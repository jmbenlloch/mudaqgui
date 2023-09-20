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
const { chargesRebin } = storeToRefs(eventStore)
const { selectedCard } = storeToRefs(configStore)

const nChannels: Ref<number> = ref(32)
</script>

<template>
  <div class="flex flex-col">
    <div class="grid grid-cols-8">
      <div v-for="ch in 32">
        <ChargeHistogram :data="chargesRebin[selectedCard].Charges[ch-1]" />
      </div>
    </div>
  </div>
</template>

