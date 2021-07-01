<template>
	<div>hello</div>
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
		console.log(props.name);
		const getData = async () => {
			const res = await fetch(`/api?user=${props.name}`);
			const json = await res.json();
			yearData.value = json.data;
			console.log(json.data);
		};

		onMounted(getData());

		return {
			yearData
		};
	}
};
</script>

<style>
</style>