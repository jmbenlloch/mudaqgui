import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import { sum } from 'lodash'

type CardRates = {
  [index: number]: number,
}

export const useRateStore = defineStore('rate', () => {
  const rates: Ref<CardRates> = ref({})
  const lostFPGA: Ref<CardRates> = ref({})
  const lostBuffer: Ref<CardRates> = ref({})

  const totalRate = computed(() => sum(Object.values(rates.value)))
  const totalLostFPGA = computed(() => sum(Object.values(lostFPGA.value)))
  const totalLostBuffer = computed(() => sum(Object.values(lostBuffer.value)))

  EventsOn("rate", (data : CardRates) => {
    console.log("event rate", data)
    for (let card in data){
      rates.value[card] = data[card]
    }
  })

  EventsOn("lostFPGA", (data : CardRates) => {
    console.log("lostFPGA", data)
    for (let card in data){
      lostFPGA.value[card] = data[card]
    }
  })

  EventsOn("lostBuffer", (data : CardRates) => {
    console.log("lostBuffer", data)
    for (let card in data){
      lostBuffer.value[card] = data[card]
    }
  })

  return { rates, lostBuffer, lostFPGA, totalRate, totalLostBuffer, totalLostFPGA }
})