import { createApp } from 'vue'
import VueTippy from 'vue-tippy'
import App from './App.vue'

const app = createApp(App)
app.use(VueTippy)
app.mount('#app')