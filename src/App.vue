<template>
	<img alt="Vue logo" src="./assets/logo.png" />
	<contrib-bar :name="qjack" />
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
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	margin-top: 60px;
}
</style>
