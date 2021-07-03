<template>
	<div v-show="message.trim().length > 0" aria-live="polite" role="alert">
		<p :class="type">{{ message }}</p>
	</div>
</template>

<script setup>
	import { defineProps } from 'vue'

	const props = defineProps({
		type: String,
		message: String
	})
</script>

<style scoped>
	div
	{
		display: flex;
		align-items: center;
		position: absolute;
		left: 0;
		right: 0;
		z-index: 99;
		pointer-events: none;

		padding: 0 var(--space);
		height: 12vw;
		min-height: 95px;
		max-height: 120px;
	}

	p
	{
		margin: 0 auto;
		width: fit-content;
		max-width: 330px;
		text-align: center;
		user-select: none;

		background: var(--dark);
		padding: 1ch 2ch;
		border: 1px solid rgba(173, 190, 208, 0.4);
		border-radius: 6px;
		box-shadow:
			0 1px 2px rgba(0, 0, 0, 0.05),
			0 2px 4px rgba(0, 0, 0, 0.15);
	}

	p.info
	{
		border-color: var(--blue);
	}

	p.error
	{
		border-color: var(--red);
	}
</style>