import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import { sum } from 'lodash'

type EventData = {
  [index: number]: [{
    T0: number,
    T1: number,
    LostBuffer: number,
    LostFPGA: number,
    Charges: Array<number>,
  }]
}

export const useEventStore = defineStore('events', () => {
  const events: Ref<EventData> = ref({})

  const t0 = computed(() => {
    let t0s: {[index: string] : Array<number>} = {}
    Object.entries(events.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t0s[card] = events.map(o => o.T0)
    });
    return t0s
  })

  const t1 = computed(() => {
    let t1s: {[index: string] : Array<number>} = {}
    Object.entries(events.value).forEach(entry => {
      const [card, events] = entry;
      console.log("data: ", card, events);
      t1s[card] = events.map(o => o.T1)
    });
    return t1s
  })

  EventsOn("events", (data) => {
    console.log("events", data)
    events.value = data
  })

  return { events, t0, t1 }
})