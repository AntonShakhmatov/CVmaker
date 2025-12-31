<script setup lang="ts">
import type MenuVue from '@/menu/Menu.vue'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useRouter } from 'vue-router'

const router = useRouter()

const route = useRoute()
const data = ref(null)
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  const email = route.query.email as string

  if (!email) {
    error.value = 'Email not provided'
    loading.value = false
    return
  }

  try {
    const res = await fetch(
      `http://localhost:5000/api/success?email=${encodeURIComponent(email)}`
    )

    if (!res.ok) throw new Error('Request failed')

    data.value = await res.json()
  } catch (e) {
    error.value = 'Failed to load resume'
  } finally {
    loading.value = false
  }
})

const handleSubmit = async () => {
  await router.push({
    path: '/buildcv',
    query: {
      email: route.query.email,
    },
  })
}
</script>

<template>
  <section class="min-h-screen bg-gray-100 py-10">
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>

    <div v-else-if="data">
    <div class="mx-auto max-w-4xl rounded-lg bg-white p-8 shadow">

      <!-- HEADER -->
      <header class="border-b pb-6 mb-6">
        <h1 class="text-3xl font-bold text-gray-900">
          {{ data.cv.header.first_name }} {{ data.cv.header.last_name }}
        </h1>
        <p class="mt-2 text-lg text-gray-600">
          {{ data.cv.header.title }}
        </p>
        <p class="mt-1 text-sm text-gray-500">
          {{ data.cv.header.competence }}
        </p>
      </header>

      <!-- CONTACTS -->
      <section class="mb-8">
        <h2 class="section-title">Contacts</h2>
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 text-sm text-gray-700">
          <div>📞 {{ data.cv.contacts.phone }}</div>
          <div>✉️ {{ data.cv.contacts.email }}</div>
          <div>📍 {{ data.cv.contacts.address_field }}</div>
        </div>
      </section>

      <!-- EXPERIENCE -->
      <section class="mb-8">
        <h2 class="section-title">Experience</h2>

        <div class="mt-4">
          <h3 class="font-semibold text-lg">
            {{ data.cv.experience.position }}
          </h3>
          <p class="text-gray-600">
            {{ data.cv.experience.company }} — {{ data.cv.experience.location_field }}
          </p>
          <p class="text-sm text-gray-500">
            {{ data.cv.experience.date_from }} – {{ data.cv.experience.date_to }}
          </p>
        </div>
      </section>

      <!-- PROJECT -->
      <section class="mb-8">
        <h2 class="section-title">Project</h2>

        <h3 class="font-semibold">
          {{ data.cv.projects.name }}
        </h3>
        <p class="mt-2 text-gray-700">
          {{ data.cv.projects.description }}
        </p>
      </section>

      <!-- TECHNOLOGIES -->
      <section class="mb-8">
        <h2 class="section-title">Technologies</h2>
        <div class="mt-3 flex flex-wrap gap-2">
          <span
            v-for="tech in data.cv.technologies.name?.split(',')"
            :key="tech"
            class="rounded bg-blue-100 px-3 py-1 text-sm text-blue-700"
          >
            {{ tech }}
          </span>
        </div>
      </section>

      <!-- LINKS -->
      <section>
        <h2 class="section-title">Links</h2>
        <a
          :href="data.cv.links.label"
          target="_blank"
          class="text-blue-600 hover:underline"
        >
          {{ data.cv.links.label }}
        </a>
      </section>
    <form @submit.prevent="handleSubmit">
      <button 
      type="submit"
      class="inline-flex items-center justify-center rounded-md bg-blue-600 px-4 py-2 text-sm
      font-semibold text-white transition hover:bg-blue-700
      focus:outline-none focus:ring-2 focus:ring-blue-300"
      >Submit</button>
    </form>
    </div>
    </div>
  </section>
</template>
