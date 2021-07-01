import { createApp } from 'vue'
import VueTippy from 'vue-tippy'
import App from './App.vue'

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

app.mount('#app')