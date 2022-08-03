import { createApp } from 'vue'
import VueTippy from 'vue-tippy'
import { createPinia } from 'pinia'
import App from './App.vue'

const pinia = createPinia()
const app = createApp(App)
app.use(VueTippy, {
	theme: 'standard-dark',
	defaultProps: { 
		placement: 'right',
		delay: [0, 0],
		duration: 0,
		allowHTML: true,
		maxWidth: 150,
		theme: 'standard-dark'
	},
})
app.use(pinia)
app.mount('#app')