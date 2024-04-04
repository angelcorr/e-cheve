<script setup lang="ts">
import InputTerm from './InputTerm.vue';
import HeaderCheve from './HeaderCheve.vue';
import EmailItem from './EmailItem.vue';
import EmailDetails from './EmailDetails.vue';

import { store } from '../store';

</script>

<template>
  <div class="bg-white min-h-full h-screen">
    <section class="pb-5 mx-8">
      <HeaderCheve />
      <InputTerm />
    </section>
    <section class="w-full h-2/3 flex justify-center">
      <div class="flex flex-col overflow-y-auto" v-bind:class="[store.emailSelected ? 'w-1/2' : 'w-5/6']">
        <EmailItem
          v-for="email in store.emails"
          :email="email"
          :id="email.id"
          :key="email.id"
          @click="store.changeEmailSelected(email)"
        />
      </div>
      <div
        class="w-1/2 flex flex-col overflow-y-auto"
        v-if="Object.keys(store.emailSelected).length !== 0"
      >
        <EmailDetails :emailSelected="store.emailSelected" />
      </div>
    </section>
  </div>
</template>
