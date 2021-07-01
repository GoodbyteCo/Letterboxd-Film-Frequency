<template>
	<div>
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 317 41" aria-labelledby="" class="graph">
			<rect v-for="day in daysInTheYear(year)" 
				:key="day" 
				:transform="'translate('+((getWeekNumber(year, day) - 1) * 6)+' '+(getWeekDay(year, day) * 6)+')'" 
				:fill="'var(--accent-'+Math.ceil(filmsWatchedOn(year, day) / scale)+')'"
				v-tippy="{ content: '<b>'+filmsWatchedOn(year, day)+' film(s)</b> watched on '+getDate(year, day).toLocaleDateString() }"
				width="5" height="5"
			/>
		</svg>
	</div>
	<p id="scroll-prompt">
		Scroll
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 36 24" class="icon">
			<path d="M28 4l-1.4 1.4 5.6 5.6H0v2h32.2l-5.6 5.6L28 20l8-8z"/>
		</svg>
	</p>
</template>

<script setup>
	import { defineProps, computed } from 'vue'
	import { directive } from 'vue-tippy'

	const props = defineProps({
		year: Number,
		films: Object
	})

	// scale increment = 1/5th the maximum watched in any one day
	// or if object is undefined, scale increment = 1
	const scale = computed(() => Math.max(...Object.values(props.films || { 'default': 5 })) / 5 )

	const getDate = (year, day) => {
		var date = new Date(year, 0)
		return new Date(date.setDate(day))
	}

	const getWeekDay = (year, day) => {
		return getDate(year, day).getDay()
	}

	const getWeekNumber = (year, day) => {
		return Math.ceil((getDate(year, 0).getDay() + day + 1) / 7)
	}

	const filmsWatchedOn = (year, day) => {
		var date = getDate(year, day);
		var formattedDate = (date.getMonth() + 1) + '/' + date.getDate()
		return props.films[formattedDate] || 0
	}

	const daysInTheYear = (year) => {
		var isLeapYear = (new Date(year, 1, 29).getDate() === 29)
		return isLeapYear ? 366 : 365
	}
</script>

<style scoped>
	svg.graph
	{
		width: 100%;
		min-width: 800px;
		box-sizing: border-box;
		padding: 0 var(--space) var(--space);
	}

	svg.graph rect
	{
		outline: none;
	}

	div
	{
		width: 100%;
		overflow: scroll;
	}

	#scroll-prompt
	{
		display: none;
		margin: 0 var(--space);
		opacity: 0.4;
		user-select: none;

		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		line-height: 1rem;
		text-align: right;
	}

	svg.icon
	{
		fill: var(--light);
		height: 1rem;
		margin-left: 1ch;
		transform: translateY(3px);
	}

	@media screen and (max-width: 795px)
	{
		#scroll-prompt
		{
			display: block;
		}
	}
</style>