<template>
	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 317 41" aria-labelledby="">
		<rect v-for="day in daysInTheYear(year)" 
			:key="day" 
			:transform="'translate('+((getWeekNumber(year, day) - 1) * 6)+' '+(getWeekDay(year, day) * 6)+')'" 
			:fill="'var(--accent-'+Math.ceil(filmsWatchedOn(year, day) / scale)+')'"
			:content="filmsWatchedOn(year, day)+' films watched on '+getDate(year, day).toLocaleDateString()" v-tippy="{ trigger : 'mouseenter' }"
			width="5" height="5"
		/>
	</svg>
</template>

<script setup>
	import { defineProps, computed } from 'vue'

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

<style>
	svg
	{
		width: 100%;
	}
</style>