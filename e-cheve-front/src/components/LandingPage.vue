<script setup lang="ts">
import InputTerm from './InputTerm.vue'
import HeaderCheve from './HeaderCheve.vue'
import EmailItem from './EmailItem.vue'
import EmailDetails from './EmailDetails.vue'
import { getEmails as vGetEmails } from '../api/index'

import { store } from '../store'

</script>

<template>
  <div class="bg-white min-h-full h-screen">
    <section class="pb-5 mx-8">
      <HeaderCheve />
      <InputTerm />
    </section>
    <section class="w-full h-2/3 flex justify-center">
      <div
        class="flex flex-col overflow-y-auto"
        v-bind:class="[store.emailSelected ? 'w-1/2' : 'w-5/6']"
      >
        <EmailItem
          v-for="email in store.emails"
          :email="email"
          :id="email.id"
          :key="email.id"
          @click="store.changeEmailSelected(email)"
        />
        <div class="w-full flex flex-col items-center mb-5">
          <figure class="w-10 my-6" v-if="store.isLoading">
            <img src="../assets/spinner.png" alt="Spinner" class="animate-spin">
          </figure>
          <a
            @click="vGetEmails(true)"
            class="py-3 px-4 cursor-pointer border border-red-800 rounded-lg bg-linen hover:bg-white hover:scale-110 ease-in-out duration-300"
            v-if="store.emails.length !== 0"
          >
            Load more emails
          </a>
        </div>
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
