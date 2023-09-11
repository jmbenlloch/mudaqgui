<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps({
  min: { type: Number, required: true },
  max: { type: Number, required: true },
})

const value = ref(0)

function increment(){
  if (value.value < props.max){
    value.value++
  }
}

function decrement(){
  if (value.value > props.min){
    value.value--
  }
}

</script>

<template>
  <div class="custom-number-input h-10 w-16">
    <div class="flex flex-row h-10 w-full rounded-lg relative bg-transparent mt-1">
      <button @click="decrement"
        class=" bg-gray-300 text-gray-600 hover:text-gray-700 hover:bg-gray-400 h-full w-20 rounded-l cursor-pointer outline-none">
        <span class="m-auto text-2xl font-thin">−</span>
      </button>
      <input type="number" min="0" max="64"
        class="outline-none focus:outline-none text-center w-full bg-gray-300 hover:text-black focus:text-black  md:text-basecursor-default flex items-center text-gray-700"
        name="custom-input-number" v-model="value" />
      <button @click="increment"
        class="bg-gray-300 text-gray-600 hover:text-gray-700 hover:bg-gray-400 h-full w-20 rounded-r cursor-pointer">
        <span class="m-auto text-2xl font-thin">+</span>
      </button>
  </div>
</div></template>