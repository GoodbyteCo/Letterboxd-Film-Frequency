function useDates() {
    const getDate = (year: number, day: number) => {
		const date = new Date(year, 0)
		return new Date(date.setDate(day))
	}

	const getWeekDay = (year: number, day: number) => {
		return getDate(year, day).getDay()
	}

	const getWeekNumber = (year:number, day:number) => {
		const firstDay = (getWeekDay(year, 0) + 1) % 7
		return Math.ceil((day + firstDay) / 7) - 1
	}

    const getDaysInTheYear = (year: number) => {
		const isLeapYear = (new Date(year, 1, 29).getDate() === 29)
		return isLeapYear ? 366 : 365
	}

    return {
        getDate,
        getWeekDay,
        getWeekNumber,
        getDaysInTheYear
    }
}

export default useDates