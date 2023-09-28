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
  const totalRate = computed(() => sum(Object.values(rates.value)))

  EventsOn("rate", (data : CardRates) => {
    for (let card in data){
      rates.value[card] = data[card]
    }
  })

  return { rates, totalRate }
})