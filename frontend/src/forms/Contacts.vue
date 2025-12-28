<script setup lang="ts">
defineProps<{
  form: {
    phone: string,
    email: string,
    address_field: string
  }
}>()
</script>

<template v-if="ok">
    <section class="mx-auto w-full max-w-2xl">
        <h1 class="text-3xl font-bold tracking-tight">Contacts</h1>
        <p class="mt-2 text-sm text-gray-600">Input phone</p>
        <input v-model="form.phone" 
        placeholder="Put your phone here" 
        class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />

        <p class="mt-2 text-sm text-gray-600">Input email</p>
        <input v-model="form.email" 
        placeholder="Put your email here" 
        class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        />

        <p class="mt-2 text-sm text-gray-600">Input address</p>
        <input v-model="form.address_field" 
        placeholder="Put your address here" 
        class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
        /><br><br>
    </section>
</template>