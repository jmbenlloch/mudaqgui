import { ref, computed, watch } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";
import { EventsOn } from '../../wailsjs/runtime/runtime'


dayjs.extend(utc)
dayjs.extend(timezone)

export type CalibrationConfig = {
    Duration: number,
    Events: number,
    Bias: number,
    Gain: number,
    Dac: number,
};

export type MessageType = {
    Timestamp: number,
    Configuration: CalibrationConfig,
};

export const useMessagesStore = defineStore('messages', () => {
    const messages: Ref<Array<MessageType>> = ref([])

    EventsOn("calibration", (data) => {
        console.log(data);
        messages.value.push(data)
        if (messages.value.length > 20) {
            messages.value.splice(0, messages.value.length - 20)
        }
    });

    return { messages }
})