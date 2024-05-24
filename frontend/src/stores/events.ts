import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import { fill } from 'lodash'

type TimeData = {
  [index: number]: [number]
}

type Point = {
  x: number,
  y: number,
}

type ChargeHistogram = {
  [index: number]: {
    Charges: Array<Array<number>>,
  }
}

export const useEventStore = defineStore('events', () => {
  const t0s_backend: Ref<TimeData> = ref({})
  const t1s_backend: Ref<TimeData> = ref({})

  const dataTaking: Ref<boolean> = ref(false)
  const nEvents: Ref<number> = ref(0)

  const t0 = computed(() => {
    let t0s: { [index: string]: Array<Point> } = {}
    Object.entries(t0s_backend.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t0s[card] = events.map((o, index) => {
        return { x: index, y: o }
      })
    });
    return t0s
  })

  const t1 = computed(() => {
    let t1s: { [index: string]: Array<Point> } = {}
    Object.entries(t1s_backend.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t1s[card] = events.map((o, index) => {
        return { x: index, y: o }
      })
    });
    console.log("t1s: ", t1)
    return t1s
  })

  EventsOn("t0s", (data) => {
    console.log("t0s backend", data)
    t0s_backend.value = data
  })

  EventsOn("t1s", (data) => {
    console.log("t1s backend", data)
    t1s_backend.value = data
  })

  EventsOn("charges", (data) => {
    console.log("charges", data)
    charges.value = data
  })

  EventsOn("chargesRebin", (data) => {
    console.log("chargesRebin", data)
    chargesRebin.value = data
  })

  EventsOn("dataTaking", (data) => {
    console.log("dataTaking", data)
    dataTaking.value = data
  })

  EventsOn("nEvents", (data) => {
    //console.log("nEvents", data)
    nEvents.value = data
  })

  const charges: Ref<ChargeHistogram> = ref({
    1024: {
      Charges: Array(32).fill(Array(4096).fill(0)),
    },

    85: {
      Charges: Array(32).fill(Array(4096).fill(0)),
    },
    69: {
      Charges: Array(32).fill(Array(4096).fill(0)),
    },
  })

  const chargesRebin: Ref<ChargeHistogram> = ref({
    1024: {
      Charges: Array(32).fill(Array(32).fill(0)),
    },
    85: {
      Charges: Array(32).fill(Array(32).fill(0)),
    },
    69: {
      Charges: Array(32).fill(Array(32).fill(0)),
    },
  })

  return { t0, t1, charges, chargesRebin, dataTaking, nEvents }
})