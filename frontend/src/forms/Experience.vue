<script setup lang="ts">
defineProps<{
  form: {
    company: string,
    location_field: string,
    position: string,
    date_from: string,
    date_to: string
  }
}>()
</script>

<template>
   <section class="mx-auto w-full max-w-2xl">
        <h1 class="text-3xl font-bold tracking-tight">Experience</h1>
                <p class="mt-2 text-sm text-gray-600">Company name</p>
                <input v-model="form.company" 
                placeholder="Put your company name here"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                />

                <p class="mt-2 text-sm text-gray-600">Location of the company</p>
                <input v-model="form.location_field" 
                placeholder="Put location of your company here"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                />

                <p class="mt-2 text-sm text-gray-600">What position you working in?</p>
                <input v-model="form.position" 
                placeholder="Put here your position in company"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                />

                <p class="mt-2 text-sm text-gray-600">What time you have began?</p>
                <input type="date" v-model="form.date_from" 
                placeholder="Put here date of the begining"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                />

                <p class="mt-2 text-sm text-gray-600">What time you end the work wit this company</p>
                <input type="date" v-model="form.date_to" 
                placeholder="Put here date when you have end"
                class="mt-1 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm
                 outline-none transition focus:border-blue-500 focus:ring-2 focus:ring-blue-200"
                /><br><br>
    </section>
</template>