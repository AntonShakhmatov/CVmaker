<script setup lang="ts">
defineProps<{
  form: {
    first_name: string,
    last_name: string,
    competence: string,
    title: string
  }
}>()
</script>

<template v-if="ok">
  <section class="mx-auto w-full max-w-2xl">
    <h1 class="text-3xl font-bold tracking-tight">Header</h1>
      <div>
        <label class="block text-sm font-medium text-gray-700">First name</label>
        <input
          v-model="form.first_name"
          placeholder="Put your name here"
          class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Lastname</label>
        <input
          v-model="form.last_name"
          placeholder="Put your lastname here"
          class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Competence</label>
        <input
          v-model="form.competence"
          placeholder="Put your competence here"
          class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Title</label>
        <input
          v-model="form.title"
          placeholder="Put your title here"
          class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />
      </div>

      <div class="flex items-center gap-3 pt-2">
      </div>
  </section>
</template>
