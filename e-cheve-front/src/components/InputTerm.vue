<script setup lang="ts">
import { store, type Email } from '../store'

interface ApiHit {
  _id: string
  _source: {
    From: string
    To: string
    Date: string
    Subject: string
    Body: string
  }
}

interface ApiResponse {
  hits: {
    hits: ApiHit[]
  }
}

const submit = async () => {
  if (store.term.length < 1) return;

  const endpoint = `${import.meta.env.VITE_API_BASE_URL}/emails?term=${store.term}`;
  const response = await fetch(endpoint);
  if (response.status !== 200) return 'There was a problem'; //To improve

  const apiResponse: ApiResponse = await response.json();

  const emails: Email[] = apiResponse.hits.hits.map((hit) => ({
    id: hit._id,
    from: hit._source.From,
    to: hit._source.To,
    date: hit._source.Date,
    subject: hit._source.Subject,
    body: hit._source.Body
  }));

  store.emails = emails;
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
