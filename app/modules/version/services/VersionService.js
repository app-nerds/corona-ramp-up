/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

export class VersionService {
	constructor($http) {
		this.$http = $http;
		this.baseURL = "/api/version"
	}

	async GetVersion() {
		let response = await this.$http.get(this.baseURL);
		return response.body;
	}
}

export function VersionServiceInstaller(Vue) {
	Vue.prototype.versionService = new VersionService(Vue.http);
}