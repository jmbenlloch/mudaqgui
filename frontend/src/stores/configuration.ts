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

export const useConfigStore = defineStore('config', () => {
  const slowControl: Ref<slowControlConfig> = ref({
    channel_preamp_HG: [0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0],
    input_dac: [0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0],
    channel_preamp_disable: [0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0],
    discriminatorMask: [0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0],
    enable_or32: 0,
    dac1_code: 0,
    dac2_code: 0,
  })
  const probe: Ref<probeConfig> = ref({
    peakSensingHG: [0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0],
  })


  EventsOn("configSlowControl", (data) => {
    console.log("SC", data)
    slowControl.value = data
  })

  EventsOn("configProbe", (data) => {
    console.log("Probe", data)
    probe.value = data
  })

  return { slowControl, probe }
})
