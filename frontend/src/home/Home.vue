

<script setup lang="ts">
import Language from '@/forms/Language.vue';
import { reactive } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()

const form = reactive({
  language: '',
  first_name: '',
  last_name: '',
  competence: '',
  title: '',
  phone: '',
  email: '',
  address_field: '',
  company: '',
  location_field: '',
  position: '',
  date_from: '',
  date_to: '',
  name: '',
  description: '',
  responsibilities_description: '',
  technologies_name: '',
  label: ''
})

  const resetForm = () => {
    form.language = ''
    form.first_name = ''
    form.last_name = ''
    form.competence = ''
    form.title = ''
    form.phone = ''
    form.email = ''
    form.address_field = ''
    form.company = ''
    form.location_field = ''
    form.position = ''
    form.date_from = ''
    form.date_to = ''
    form.name = ''
    form.description = ''
    form.responsibilities_description = ''
    form.technologies_name = ''
    form.label = ''
  }

 const handleSubmit = async () => {
  const payload = {
    header: {
      first_name: form.first_name,
      last_name: form.last_name,
      title: form.title,
      competence: form.competence,
    },
    contacts: {
      phone: form.phone,
      email: form.email,
      address_field: form.address_field,
    },
    language: {
      language: form.language,
    },
    experience: {
      company: form.company,
      location_field: form.location_field,
      position: form.position,
      date_from: form.date_from,
      date_to: form.date_to,
    },
    responsibilities: {
      description: form.responsibilities_description,
    },
    technologies: {
      name: form.technologies_name,
    },
    projects: {
      name: form.name,
      description: form.description,
    },
    links: {
      label: form.label
    },
  }

  const res = await fetch("http://localhost:5000/api/build", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  })

  await router.push({
    path: '/success',
    query: {
      email: form.email,
    },
  })
}


  </script>

  <template>
    <section class="mx-auto w-full max-w-2xl">
      <form @submit.prevent="handleSubmit" class="mt-6 space-y-5">
        <Language :form="form" />
        <Header :form="form"></Header>
        <Contacts :form="form" />
        <Experience :form="form" />
        <Responsibilities :form="form" />
        <Projects :form="form" />
        <Technologies :form="form" />
        <Links :form="form" />

          <button 
          type="submit"
          class="inline-flex items-center justify-center rounded-md bg-blue-600 px-4 py-2 text-sm
          font-semibold text-white transition hover:bg-blue-700
          focus:outline-none focus:ring-2 focus:ring-blue-300"
          >Submit</button>

          <button
          type="button"
          class="rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-semibold
                  text-gray-700 transition hover:bg-gray-50"
          @click="resetForm"
          >
          Reset
          </button>
      </form>
    </section>
  </template>