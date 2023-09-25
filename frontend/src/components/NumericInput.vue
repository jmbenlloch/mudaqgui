<script setup lang="ts">
import { ref, watch } from 'vue'
import { debounce } from 'lodash';


const emit = defineEmits(['updateValue', 'increment', 'decrement'])
const props = defineProps({
  value: { type: Number, required: true },
  min: { type: Number, required: true },
  max: { type: Number, required: true },
  disabled: { type: Boolean, required: false },
})

function increment() {
  if (!props.disabled) {
    if (props.value < props.max) {
      emit("increment", props.value + 1)
    }
  }
}

function decrement() {
  if (!props.disabled) {
    if (props.value > props.min) {
      emit("decrement", props.value - 1)
    }
  }
}

function updateValue(event: Event) {
  if (!props.disabled) {
    // @ts-ignore
    let value = event.target.value
    if (value > props.max) {
      emit("updateValue", props.max)
    }
    else if (value < props.min) {
      emit("updateValue", props.min)
    }
    else {
      emit("updateValue", value)
    }
  }
}
</script>

<template>
  <div class="custom-number-input h-8 w-24">
    <div class="flex flex-row w-full rounded-lg relative bg-transparent">
      <button @click="decrement" class=" bg-gray-300 text-gray-600 w-20 rounded-l cursor-pointer outline-none"
        :class="[props.disabled ? '' : 'hover:text-gray-700 hover:bg-gray-400']">
        <span class="m-auto text-2xl font-thin">−</span>
      </button>
      <input type="number" :min="props.min" :max="props.max" :disabled="props.disabled"
        class="outline-none focus:outline-none text-center w-full bg-gray-300 hover:text-black focus:text-black  md:text-basecursor-default flex items-center text-gray-700"
        name="custom-input-number" :value="props.value" @input="updateValue" />
      <button @click="increment" class="bg-gray-300 text-gray-600 w-20 rounded-r cursor-pointer"
        :class="[props.disabled ? '' : 'hover:text-gray-700 hover:bg-gray-400']">
        <span class="m-auto text-2xl font-thin">+</span>
      </button>
    </div>
  </div>
</template>