<script setup lang="ts">
import { onMounted, ref } from "vue";
import type { Ref } from "vue";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";
import { useMessagesStore } from "@/stores/messages";
import { storeToRefs } from 'pinia'

dayjs.extend(utc)
dayjs.extend(timezone)

function parseTimestamp(timestamp: number): string {
  let result = dayjs.unix(timestamp).format("DD/MM/YY - HH:mm:ss");
  return result;
}

const store = useMessagesStore()
const { messages } = storeToRefs(store)
</script>

<template>
  <div class="border mx-auto p-2">
    <h2 class="text-xl font-bold m-4">Messages</h2>
    <div class="bg-gray-200 m-4 overflow-scroll">
      <p v-for="message in messages" class="mx-2 my-1">
        <span v-if="!message.Configuration.Finished"> {{ parseTimestamp(message.Timestamp) }} - Taking data during {{
        message.Configuration.Duration }} seconds. Bias {{ message.Configuration.Bias }}, gain {{
        message.Configuration.Gain }}, DAC {{ message.Configuration.Dac }} </span>

        <span v-if="message.Configuration.Finished"> {{ parseTimestamp(message.Timestamp) }} - Calibration finished</span>
      </p>
    </div>
  </div>
</template>
