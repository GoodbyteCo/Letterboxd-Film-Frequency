<template>
	<footer>
		<p>
			<a href="./">Film Frequency</a> is a <a href="http://goodbyte.ca">Goodbyte</a> project. Read
			<a href="https://github.com/GoodbyteCo/Letterboxd-Contribution-Graph">more about it on Github</a>.
		</p>
	</footer>
</template>

<script setup lang="ts"></script>

<style scoped>
	footer
	{
		padding: calc(2 * var(--space)) var(--space);
		opacity: 0.4;
		position: fixed;
		bottom: 0;
	}

	p
	{
		margin: 0;
	}

	a
	{
		color: var(--accent-5);
		text-decoration: none;
		background-image: linear-gradient(var(--accent-5), var(--accent-5));
		background-size: 100% 1px;
		background-position: left bottom;
		background-repeat: no-repeat;
		text-shadow: 1px 1px var(--dark), -1px 1px var(--dark);
	}

	@media screen and (max-height: 430px)
	{
		footer
		{
			position: relative;
			padding-bottom: 0;
		}
	}
</style>

