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
	const scale = computed(() => {
		try {
			return Math.max(...Object.values(props.films[props.year] || { 'default': 5 })) / 5
		}
		catch (e) {
			return 1
		}
	})

	const getDate = (year, day) => {
		var date = new Date(year, 0)
		return new Date(date.setDate(day))
	}

	const getWeekDay = (year, day) => {
		return getDate(year, day).getDay()
	}

	const getWeekNumber = (year, day) => {
		// TODO: figure out why this is
		if (getDate(year, 0).getDay() == 6) {
			return Math.ceil((getDate(year, 0).getDay() + day + 1) / 7) - 1
		}
		return Math.ceil((getDate(year, 0).getDay() + day + 1) / 7)
	}

	const filmsWatchedOn = (year, day) => {
		var date = getDate(year, day);
		var formattedDate = (date.getMonth() + 1) + '/' + date.getDate()
		
		try {
			return props.films[year][formattedDate] || 0
		}
		catch (e) {
			return 0
		}
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

	div::before, div::after
	{
		content: "";
		display: block;
		position: fixed;
		top: 0;
		right: 0;
		bottom: 0;

		width: var(--space);
		pointer-events: none;
		background: linear-gradient(to right,
			rgba(25, 30, 37, 0.000) 0%,
			rgba(25, 30, 37, 0.013) 9.6%,
			rgba(25, 30, 37, 0.049) 18.1%,
			rgba(25, 30, 37, 0.104) 25.6%,
			rgba(25, 30, 37, 0.175) 32.3%,
			rgba(25, 30, 37, 0.259) 38.4%,
			rgba(25, 30, 37, 0.352) 44.0%,
			rgba(25, 30, 37, 0.450) 49.3%,
			rgba(25, 30, 37, 0.550) 54.4%,
			rgba(25, 30, 37, 0.648) 59.6%,
			rgba(25, 30, 37, 0.741) 65.0%,
			rgba(25, 30, 37, 0.825) 70.6%,
			rgba(25, 30, 37, 0.896) 76.8%,
			rgba(25, 30, 37, 0.951) 83.7%,
			rgba(25, 30, 37, 0.987) 91.3%,
			rgba(25, 30, 37, 1.000) 100%);
	}

	div::before
	{
		left: 0;
		right: unset;
		background: linear-gradient(to left,
			rgba(25, 30, 37, 0.000) 0%,
			rgba(25, 30, 37, 0.013) 9.6%,
			rgba(25, 30, 37, 0.049) 18.1%,
			rgba(25, 30, 37, 0.104) 25.6%,
			rgba(25, 30, 37, 0.175) 32.3%,
			rgba(25, 30, 37, 0.259) 38.4%,
			rgba(25, 30, 37, 0.352) 44.0%,
			rgba(25, 30, 37, 0.450) 49.3%,
			rgba(25, 30, 37, 0.550) 54.4%,
			rgba(25, 30, 37, 0.648) 59.6%,
			rgba(25, 30, 37, 0.741) 65.0%,
			rgba(25, 30, 37, 0.825) 70.6%,
			rgba(25, 30, 37, 0.896) 76.8%,
			rgba(25, 30, 37, 0.951) 83.7%,
			rgba(25, 30, 37, 0.987) 91.3%,
			rgba(25, 30, 37, 1.000) 100%);
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