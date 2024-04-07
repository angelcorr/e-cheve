<script setup lang="ts">
import { store } from '../store'

const props = defineProps(['emailSelected'])
</script>

<template>
  <section class="h-full max-h-full m-4 p-4 overflow-y-auto border border-red-800 rounded-lg">
    <div>
      <header>
        <div class="flex justify-between">
          <p class="font-bold py-1">
            From: <span class="font-normal">{{ props.emailSelected.from }}</span>
          </p>
          <a class="text-lg cursor-pointer" @click="store.removeEmailSelected()">×</a>
        </div>
        <p class="font-bold py-1">
          To: <span class="font-normal">{{ props.emailSelected.to }}</span>
        </p>
        <p class="font-bold py-1">
          Date: <span class="font-normal">{{ props.emailSelected.date }}</span>
        </p>
        <p class="font-bold py-1">
          Subject: <span class="font-normal">{{ props.emailSelected.subject }}</span>
        </p>
      </header>
      <div>
        <p class="font-bold py-1">Message:</p>
        <p class="leading-7">
          {{ props.emailSelected.body }}
        </p>
      </div>
    </div>
  </section>
</template>
