<template>
	<controls :username="username" :year="year"/>
	<status :message="statusMessage" :type="statusType"/>
	<graph :year="year" :films="films"/>
	<goodbyte-footer/>
</template>

<script setup>
	import { ref } from 'vue'
	import Controls from './components/Controls.vue'
	import Status from './components/Status.vue'
	import Graph from './components/Graph.vue'
	import GoodbyteFooter from './components/Footer.vue'

	const username = ref('')
	const year = ref(new Date().getFullYear())

	// TODO: replace with actual call to API
	const films = {'5/11': 1,'6/29': 3,'7/03': 1,'7/7': 1,'7/26': 1,'8/7': 1,'8/9': 3,'8/11': 1,'8/12': 2,'8/13': 1,'8/15': 5,'8/16': 3,'8/21': 2,'8/22': 2,'8/23': 1,'8/26': 4,'8/27': 1,'8/28': 1,'8/30': 7,'8/31': 1,'9/1': 1,'9/3': 1,'9/10': 1,'9/11': 1,'9/13': 1,'9/15': 1,'9/19': 3,'9/24': 1,'10/3': 1,'10/4': 1,'10/8': 1,'10/9': 6,'10/14': 1,'10/15': 2,'10/17': 3,'10/18': 2,'10/20': 2,'10/21': 2,'10/22': 1,'10/23': 2,'10/24': 2,'10/25': 3,'10/27': 2,'11/1': 2,'11/2': 2,'11/3': 1,'11/5': 2,'11/6': 2,'11/8': 2,'11/9': 6,'11/18': 1,'11/20': 3,'11/21': 1,'11/23': 1,'11/28': 3,'11/29': 1,'11/30': 1,'12/1': 2,'12/4': 2,'12/6': 1,'12/7': 2,'12/9': 1,'12/17': 1,'12/18': 1,'12/19': 3,'12/20': 1,'12/25': 3,'12/30': 1}
	const statusMessage = ref('Enter your Letterboxd username to get data.')
	const statusType = ref('') // 'error', 'info', or empty

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