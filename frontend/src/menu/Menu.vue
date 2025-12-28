<template>
  <nav class="fixed inset-x-0 top-0 z-50 border-b border-gray-200 bg-white">
    <div class="mx-auto max-w-7xl px-4">
      <ul class="flex h-14 items-center gap-2 overflow-x-auto whitespace-nowrap">
        <li v-for="item in menuItems" :key="item.name" class="shrink-0">
          <router-link
            :to="item.path"
            class="inline-flex items-center rounded-md px-3 py-2 text-sm font-medium text-blue-600 transition-colors hover:underline"
            active-class="font-bold text-red-600 underline"
          >
            {{ item.name }}
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface MenuItem {
  name: string
  path: string
}

const menuItems = ref<MenuItem[]>([
  { name: 'Home', path: '/' },
  { name: 'Profile', path: '/success' },
])
</script>

<style scoped>

</style>
