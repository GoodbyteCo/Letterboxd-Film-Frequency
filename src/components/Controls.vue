<template>
	<form @submit.prevent>
		<fieldset>
			<legend>Graph settings</legend>

			<div class="username">
				<label for="username">Letterboxd username</label>
				<input type="text"
					id="username" 
					name="u"
					placeholder="ex: holopollock"
					:value="user.username"
					@change="user.username = ($event.target as HTMLInputElement).value"
					v-on:blur="user.username = ($event.target as HTMLInputElement).value" 
					v-on:keyup.enter="user.username = ($event.target as HTMLInputElement).value"
					required
				>
			</div>

			<div>
				<label for="year">Year</label>
				<select id="year" @change="selectedYear = +($event.target as HTMLSelectElement).value">
					<option v-for="year in range(currentYear, lowestYear)"
						:key="year"
						:value="year"
						:selected="year == currentYear"
					>
						{{ year }}
					</option>
				</select>
			</div>
		</fieldset>
	</form>
</template>

<script setup lang="ts">
	import { ref, watch } from 'vue'
import { useUserStore } from '../store/user';

	// Cant move to own file see: https://github.com/vuejs/vue-next/issues/4294
	type ControlsProps = {
		lowestYear?: number
	}
	const props = withDefaults(defineProps<ControlsProps>() ,{
		lowestYear: 2011
	})

	// Cant move to own file see: https://github.com/vuejs/vue-next/issues/4294
	type ControlEmit = {
		(event: 'changeYear', value: number): void
	}
	const emit = defineEmits<ControlEmit>()
	
	// Data
	const urlParams = new URLSearchParams(window.location.search)
	const currentYear = new Date().getFullYear()
	const user = useUserStore()
	user.username = urlParams.get("u") ?? ''  
	const selectedYear = ref(currentYear)

	// On Run

	// Watchers
	watch(selectedYear, (year, prevYear) => {
		if (year != prevYear) {
			emit("changeYear", year)
		}
	})

	// Helpers
	const range = (start: number, end: number) => {
		if (selectedYear.value < end) {
			selectedYear.value = currentYear
		}
		const targetLength = (start - end) + 1
		const arr = new Array<number | undefined>(targetLength)
		const empty = Array.apply(undefined, arr)
		const result = empty.map((_discard: unknown, n: number) => n + end)
		return result.reverse()
	}
</script>

<style scoped>
	form
	{
		margin: 0 var(--space) var(--space);
	}

	fieldset
	{
		display: flex;
		flex-wrap: nowrap;
		align-items: center;
		gap: 10px;

		border: none;
		padding: 0;
		margin: 0;
	}

	.username
	{
		flex-grow: 1;
	}

	label
	{
		display: block;
		margin: 10px;
		opacity: 0.4;

		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		user-select: none;
	}

	input, select
	{
		display: block;
		width: 100%;
		max-width: none;
		box-sizing: border-box;
		-webkit-appearance: none;

		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		font-size: 1.5rem;
		color: var(--light);

		padding: 0.5ch 1.2ch;
		border: 1px solid var(--accent-0);
		border-radius: 6px;
		background-color: hsla(0, 0%, 100%, 0.02);
		box-shadow:
			0 1px 2px rgba(0, 0, 0, 0.05),
			0 2px 4px rgba(0, 0, 0, 0.15);

		transition: all 0.15s cubic-bezier(0,.5,0,1);
	}

	input::placeholder
	{
		color: var(--light);
		opacity: 0.4;
	}

	select
	{
		background-image: url('data:image/svg+xml; utf8, <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 9.4"><path d="M0 1.4L1.4 0 8 6.6 14.6 0 16 1.4l-8 8z" fill="%23adbed0"/></svg>');
		background-position: right 1ch top 53%;
		background-size: 1ch;
		background-repeat: no-repeat;
		padding-right: 2.4ch;
		user-select: none;
	}

	input:hover, select:hover, input:focus, select:focus
	{
		box-shadow:
			0 2.5px 2.9px rgba(0, 0, 0, 0.25),
			0 4.1px 8.8px rgba(0, 0, 0, 0.543),
			0 10px 15px rgba(0, 0, 0, 0.66);
		background-color: hsla(0, 0%, 100%, 0.06);
		border-color: var(--accent-5);
		outline: none;
	}

	legend
	{
		opacity: 0;
		width: 0px;
		height: 0px;
		user-select: none;
	}

	@media screen and (max-width: 445px)
	{
		fieldset
		{
			flex-direction: column;
			margin-bottom: 60px;
		}

		fieldset *
		{
			width: 100%;
			max-width: none;
		}
	}
</style>