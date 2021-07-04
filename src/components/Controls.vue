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
					:value="username"
					@change="username = $event.target.value"
					v-on:blur="username = $event.target.value" 
					v-on:keyup.enter="username = $event.target.value"
					required
				>
			</div>

			<div>
				<label for="year">Year</label>
				<select id="year" @change="$emit(changeYear, $event.target.value); selectedYear = $event.target.value">
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

<script setup>
	import { defineProps, defineEmit, ref, watch } from 'vue'
	
	const currentYear = new Date().getFullYear()
	const selectedYear = ref(currentYear)

	const urlParams = new URLSearchParams(window.location.search)
	const username = ref(urlParams.get("u"))
	const emit = defineEmit(["changeUsername", "changeYear"])

	if (username.value != null) {
		emit("changeUsername", username.value)
	}

	const props = defineProps({
		lowestYear: {
			type: Number,
			default: 2011
		}
	})

	const range = (start, end) => {
		if (selectedYear.value < end) {
			emit("changeYear", currentYear)
		}
		const targetLength = (start - end) + 1
		const arr = new Array(targetLength)
		const b = Array.apply(null, arr)
		const result = b.map((discard, n) => n + end)
		return result.reverse()
	}

	watch(username, (user, prevUser) => {
		console.log(user)
		if (username.value != prevUser) {
			emit("changeUsername", user)
		}
	})
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