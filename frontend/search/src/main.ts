import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
// import { init } from '@telegram-apps/sdk-vue'

const app = createApp(App)

// init();

app.use(createPinia())
app.use(router)

app.mount('#app')
