/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

const router = new VueRouter({
	routes: [
		{
			path: "/",
			name: "home",
			component: () => import("/app/pages/home.js"),
			meta: {
				title: "Home",
			},
		},
		{
			path: "/sources",
			name: "sources",
			component: () => import("/app/pages/sources.js"),
			meta: {
				title: "Sources",
			},
		},
	],

	scrollBehavior: () => {
		return {
			x: 0,
			y: 0,
		};
	},
});

router.beforeEach(async (to, from, next) => {
	/*
	 * Here is where you can do things like check to see if the user has a valid
	 * session, setup global Vue objects on successful session validation, etc...
	 */

	document.title = `${to.meta.title} | COVID-19 Ramp Up Comparison`;
	return next();
});

export default router;
