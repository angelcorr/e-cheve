<script lang="ts">
export default {
  props: ['email']
}
</script>

<template>
  <div class="h-48 m-4 p-4 bg-linen border border-red-800 rounded-lg cursor-pointer">
    <header class="flex w-full justify-between">
      <h3 class="font-bold">
        From: <span class="font-normal">{{ email.from }}</span>
      </h3>
      <h3 class="font-bold">
        To: <span class="font-normal">{{ email.to }}</span>
      </h3>
      <h3 class="font-bold">
        Date: <span class="font-normal">{{ email.date }}</span>
      </h3>
    </header>
    <section>
      <p class="font-bold py-1">
        Subject: <span class="font-normal">{{ email.subject }}</span>
      </p>
      <p class="h-6 truncate">{{ email.body }}</p>
    </section>
  </div>
</template>
