import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import { fill } from 'lodash'

type EventData = {
  [index: number]: [{
    T0: number,
    T1: number,
    LostBuffer: number,
    LostFPGA: number,
    Charges: Array<number>,
  }]
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
  const events: Ref<EventData> = ref({})

  const t0 = computed(() => {
    let t0s: { [index: string]: Array<Point> } = {}
    Object.entries(events.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t0s[card] = events.map((o, index) => {
        return { x: index, y: o.T0 }
      })
    });
    return t0s
  })

  const t1 = computed(() => {
    let t1s: { [index: string]: Array<Point> } = {}
    Object.entries(events.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t1s[card] = events.map((o, index) => {
        return { x: index, y: o.T1 }
      })
    });
    return t1s
  })

  EventsOn("events", (data) => {
    console.log("events", data)
    events.value = data
  })

  EventsOn("charges", (data) => {
    console.log("charges", data)
    charges.value = data
  })

  EventsOn("chargesRebin", (data) => {
    console.log("chargesRebin", data)
    chargesRebin.value = data
  })

  const charges: Ref<ChargeHistogram> = ref({
    1024: {
      Charges: Array(32).fill(Array(1024).fill(0)),
    },

    85: {
      Charges: Array(32).fill(Array(1024).fill(0)),
    },
    69: {
      Charges: Array(32).fill(Array(1024).fill(0)),
    },
  })

  const chargesRebin: Ref<ChargeHistogram> = ref({
    1024: {
      Charges: Array(32).fill(Array(128).fill(0)),
    },
    85: {
      Charges: Array(32).fill(Array(128).fill(0)),
    },
    69: {
      Charges: Array(32).fill(Array(128).fill(0)),
    },
  })

  return { events, t0, t1, charges, chargesRebin }
})