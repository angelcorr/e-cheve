<script setup lang="ts">
const props = defineProps(['email'])
</script>

<template>
  <div class="h-48 max-h-48 m-4 p-4 bg-linen border border-red-800 rounded-lg cursor-pointer">
    <header class="flex w-full justify-between">
      <div>
        <h3 class="font-bold">
          From: <span class="font-normal">{{ props.email.from }}</span>
        </h3>
        <h3 class="font-bold h-6 text-ellipsis overflow-hidden">
          To: <span class="font-normal">{{ props.email.to }}</span>
        </h3>
      </div>
      <h3 class="font-bold">
        Date: <span class="font-normal">{{ new Date(props.email.date).toLocaleDateString() }}</span>
      </h3>
    </header>
    <section>
      <p class="font-bold py-1">
        Subject: <span class="font-normal">{{ props.email.subject }}</span>
      </p>
      <p class="h-6 truncate">{{ props.email.body }}</p>
    </section>
  </div>
</template>
