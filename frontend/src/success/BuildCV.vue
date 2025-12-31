<script setup lang="ts">
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
    const res = 
    await fetch('http://localhost:5000/api/buildcv', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email }),
})

    if (!res.ok) throw new Error('Request failed')

    data.value = await res.json()
  } catch (e) {
    error.value = 'Failed to load resume'
  } finally {
    loading.value = false
  }
})
</script>

<template>
<h1>Your CV have been builded</h1>
</template>