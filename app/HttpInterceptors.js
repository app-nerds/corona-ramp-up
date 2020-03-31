/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

export function InstallGlobalHttpInterceptors(Vue) {
	/*
	 * Installs a loading spinner for all HTTP AJAX calls
	 */
	Vue.http.interceptors.push(function () {
		let loading = true;
		let loader = null;

		setTimeout(() => {
			if (loading) Vue.$loading.show();
		}, 1000);

		return function () {
			loading = false;
			if (loader) loader.hide();
		};
	});

	/*
	This is an example of using an interceptor to ensure that a token is sent to API
	requests.

	Vue.http.interceptors.push(function(request) {
		let s = Vue.prototype.$session;
		let session = {
			encryptedToken: "",
		};

		if (s.exists() && s.get("session")) {
			session = new SessionService(null, s).GetSession();
		}

		request.headers.set("Authorization", session.encryptedToken);
	});
	 */
}

export function InstallAppHttpInterceptors(Vue, self) {
	Vue.http.interceptors.push(function () {
		return function (response) {
			if (
				response.status === 403 &&
				response.body === "Session expired"
			) {
				/*
				 * Here is where you could handle a session expiration.
				 * May want to destroy a local storage session
				 * and redirect to a login page or something...
				 */
			}
		};
	});
}
