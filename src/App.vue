<template>
	<div>Test</div>
	<contrib-bar name="qjack" />
</template>

<script>
import ContribBar from "./components/ContribBar.vue";

export default {
	name: "App",
	components: {
		ContribBar
	}
};
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
		--accent-6: #01e054; /* for safety */
	}

	body
	{
		background: var(--dark);
		color: var(--light);

		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	#app
	{
		max-width: 1000px;
		margin: 60px auto;
	}
</style>
