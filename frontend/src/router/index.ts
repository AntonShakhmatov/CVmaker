import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/home/Home.vue'
import Header from '@/forms/Header.vue'
import Contacts from '@/forms/Contacts.vue'
import Experience from '@/forms/Experience.vue'
import Projects from '@/forms/Projects.vue'
import Responsibilities from '@/forms/Responsibilities.vue'
import Technologies from '@/forms/Technologies.vue'
import Links from '@/forms/Links.vue'
import Language from '@/forms/Language.vue'
import Success from '@/success/Success.vue'

import { comment } from 'postcss'


const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/header',
    name: 'header',
    component: Header
  },
  {
    path: '/contacts',
    name: 'contacts',
    component: Contacts
  },
  {
    path: '/experience',
    name: 'experience',
    component: Experience
  },
  {
    path: '/projects',
    name: 'projects',
    component: Projects
  },
  {
    path: '/responsibilities',
    name: 'responsibilities',
    component: Responsibilities
  },
  {
    path: '/technologies',
    name: 'technologies',
    component: Technologies
  },
  {
    path: '/links',
    name: 'links',
    component: Links
  },
  {
    path: '/language',
    name: 'language',
    component: Language
  },
  {
    path: '/success',
    name: 'success',
    component: Success
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
