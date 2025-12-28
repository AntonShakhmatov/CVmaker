import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import Menu from './menu/Menu.vue'
import Home from './home/Home.vue'
import Contacts from './forms/Contacts.vue'
import Experience from './forms/Experience.vue'
import Header from './forms/Header.vue'
import Links from './forms/Links.vue'
import Projects from './forms/Projects.vue'
import Responsibilities from './forms/Responsibilities.vue'
import Technologies from './forms/Technologies.vue'
import Language from './forms/Language.vue'
import './assets/main.css'

const app = createApp(App)

app.use(router)
app.component('Menu', Menu)
app.component('Home', Home)
app.component('Language', Language)
app.component('Contacts', Contacts)
app.component('Experience', Experience)
app.component('Header', Header)
app.component('Links', Links)
app.component('Projects', Projects)
app.component('Responsibilities', Responsibilities)
app.component('Technologies', Technologies)

app.mount('#app')
