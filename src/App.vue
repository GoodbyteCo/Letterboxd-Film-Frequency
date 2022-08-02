<template>
	<controls 
		v-on:changeYear="year = +$event"
		:lowest-year="lowestYear"
	/>
	<status :message="statusMessage" :type="statusType"/>
	<graph :year="year" :films="films"/>
	<goodbyte-footer/>
</template>

<script setup lang="ts">
	import { ref, watch } from 'vue'
	import Controls from './components/Controls.vue'
	import Status from './components/Status.vue'
	import Graph from './components/Graph.vue'
	import GoodbyteFooter from './components/Footer.vue'
import { useUserStore } from './store/user'

	// Data
	const user = useUserStore()
	const year = ref(new Date().getFullYear())
	const lowestYear = ref(2011) // default to Letterboxd create date
	const films = ref<Record<string, Record<string, number>>>({})
	const statusMessage = ref('Enter your Letterboxd username to get data.')
	const statusType = ref('') // 'error', 'info', or empty

	watch(user, (newUser) => {
		updateGraph(newUser.username)
	})

	const updateGraph = (username: string) => {

		if (username.trim().length <= 0) {
			statusMessage.value = 'Enter your Letterboxd username to get data.'
			statusType.value = ''
			window.history.replaceState(null, "", '/') // clear url params
			return
		}

		statusMessage.value = 'Loading...'
		statusType.value = 'info'
		window.history.replaceState(null, "", '?u=' + username);

		fetch(`/api?user=${username}`)
			.then(function (res) {

				if (res.status == 404) {
					statusMessage.value = 'Error: username not found'
					statusType.value = 'error'
					return ""
				}
				else if (res.status == 500) {
					statusMessage.value = 'Internal server error'
					statusType.value = 'error'
					return ""
				}
				else {
					statusMessage.value = ''
					statusType.value = ''
				}

				return res.json()
			})
			.then(function(json) {
				console.log(json.data)
				films.value = json.data
				const years = Object.keys(json.data).map<number>(value => +value)
				lowestYear.value = Math.min(...years)
			})
	}

</script>

<style>
	:root
	{
		--dark: #191e25;
		--light: #adbed0;

		--accent-0: #2b343d;
		--accent-1: #34664c;
		--accent-2: #27844e;
		--accent-3: #1aa350;
		--accent-4: #0dc152;
		--accent-5: #01e054;

		--blue: #00a7ff;
		--red: #c14141;

		--space: 30px;
	}

	.tippy-box[data-theme~='standard-dark']
	{
		background-color: #1d2229;
		border: 1px solid rgba(173, 190, 208, 0.4);
		color: rgba(173, 190, 208, 0.6);
		padding: 5px 10px;
		border-radius: 6px;
		box-shadow:
			0 1px 2px rgba(0, 0, 0, 0.05),
			0 2px 4px rgba(0, 0, 0, 0.15);
	}

	.tippy-box[data-theme~='standard-dark'] b
	{
		color: var(--light);
		font-weight: normal;
	}

	body
	{
		margin: 0;
		background: var(--dark);
		color: var(--light);

		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	::selection
	{
		background: var(--accent-5);
		color: var(--dark);
		text-shadow: none;
	}

	#app
	{
		max-width: 1000px;
		margin: calc(2 * var(--space)) auto;
	}
</style>