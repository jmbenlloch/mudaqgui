import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'

type slowControlConfig = {
  channel_preamp_HG: Array<number>,
  input_dac: Array<number>,
  channel_preamp_disable: Array<number>,
  discriminatorMask: Array<number>,
  enable_or32: number,
  dac1_code: number,
  dac2_code: number,
}

type probeConfig = {
  peakSensingHG: Array<number>,
}

type slowControlConfigs = {
  [index: number]: slowControlConfig,
}

type probeConfigs = {
  [index: number]: probeConfig,
}

function slowControlDefaultValue(): slowControlConfig {
  const configuration: slowControlConfig = {
    channel_preamp_HG: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    input_dac: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    channel_preamp_disable: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    discriminatorMask: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    enable_or32: 0,
    dac1_code: 0,
    dac2_code: 0,
  }
  return configuration
}

function probeDefaultValue(): probeConfig {
  const configuration: probeConfig = {
    peakSensingHG: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
  }
  return configuration
}

export const useConfigStore = defineStore('config', () => {
  const nonValidCardID = 1024
  const slowControl: Ref<slowControlConfigs> = ref({
    [nonValidCardID]: slowControlDefaultValue(),
  })
  const probe: Ref<probeConfigs> = ref({
    [nonValidCardID]: probeDefaultValue(),
  })

  const cards: Ref<Array<Number>> = ref([])
  // 1024 is a special value for initialization. Max card id is 255.
  const selectedCard: Ref<number> = ref(nonValidCardID)
  const disableForms = computed(() => selectedCard.value == nonValidCardID)

  EventsOn("configSlowControl", (data) => {
    console.log("SC", data)
    slowControl.value = data
    if (selectedCard.value == nonValidCardID) {
      selectedCard.value = parseInt(Object.keys(slowControl.value)[0])
    }
  })

  EventsOn("configProbe", (data) => {
    console.log("Probe", data)
    probe.value = data
    if (selectedCard.value == nonValidCardID) {
      selectedCard.value = parseInt(Object.keys(probe.value)[0])
    }
  })

  EventsOn("cards", (data: Array<Number>) => {
    console.log("Cards", data)
    cards.value = data
  })

  return { slowControl, probe, cards, selectedCard, disableForms }
})
