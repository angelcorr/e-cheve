<script setup lang="ts">
import { store } from '../store'
import { getEmails as vGetEmails } from '../api/index'

const submit = async () => {
  await vGetEmails(false)

  if (store.term.length < 1) {
    // This is to clean the results stored in the store
    // and the email selected to preview
    store.emails = []
    store.emailSelected = {}
  }
}
</script>

<template>
  <section class="input-wrapper w-full flex justify-center">
    <input
      v-model="store.term"
      v-on:input="submit"
      type="text"
      name="email-search"
      placeholder="Search for emails"
      class="w-1/2 p-4 text-lg border border-solid border-red-800 rounded-md placeholder:text-red-950 placeholder:opacity-25 outline-black"
    />
  </section>
</template>
