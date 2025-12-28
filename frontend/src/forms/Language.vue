<script setup lang="ts">
defineProps<{
  form: {
    language: string
  }
}>()
</script>

<template v-if="ok">
  <section class="mx-auto w-full max-w-2xl">
    <h1 class="text-3xl font-bold tracking-tight">Language</h1>
        <div class="flex justify-center">
            <select v-model="form.language" 
            class="mt-1 rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                        outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200">
            <option disabled value="">Choose language:</option>
            <option>CZ</option>
            <option>EN</option>
            <option>RU</option>
            </select>
        </div>
  </section>
</template>