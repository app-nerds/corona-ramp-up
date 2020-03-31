/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */
export default {
	template: `
		<header>
			<nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
				<router-link :to="{name: 'home'}" class="navbar-brand">COVID-19 Ramp Up Comparison</router-link>

				<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
			</nav>
		</header>
	`
}
