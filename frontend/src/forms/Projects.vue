<script setup lang="ts">
defineProps<{
  form: {
    name: string,
    description: string,
  }
}>()
</script>

<template>
   <section class="mx-auto w-full max-w-2xl">
        <h1 class="text-3xl font-bold tracking-tight">Projects</h1>
                <p class="mt-2 text-sm text-gray-600">Name of the project</p>
                <input v-model="form.name" 
                placeholder="Input here name of your project"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                />

                <p class="mt-2 text-sm text-gray-600">Description of work have done</p>
                <textarea v-model="form.description" 
                placeholder="Input here description of your work"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                ></textarea><br><br>
    </section> 
</template>