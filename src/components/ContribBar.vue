<template>
	<div></div>
</template>

<script>
import { ref, onMounted } from "vue";
export default {
	name: "ContribBar",
	props: {
		name: String
	},
	setup(props) {
		const yearData = ref({});
		const getData = async () => {
			const res = await fetch(`/api?user=${props.name}`);
			const json = await res.json();
			yearData.value = json.data;
			console.log(yearData);
		};

		onMounted(getData);

		return {
			yearData
		};
	}
};
</script>

<style>
</style>