<script setup lang="ts">
defineProps<{
  form: {
    label: string,
  }
}>()
</script>

<template>
    <section class="mx-auto w-full max-w-2xl">
        <h1 class="text-3xl font-bold tracking-tight">Links</h1>
            <p class="mt-2 text-sm text-gray-600">Write here your github link</p>
            <input v-model="form.label" 
            placeholder="Put your github here"
            class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
            outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
            /><br><br>
    </section>
</template>