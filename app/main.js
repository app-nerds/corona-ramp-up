/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

import router from "/app/router.js";
import Header from "/app/modules/ui/components/header.js";
import { AlertServiceInstaller } from "/app/modules/ui/services/AlertService.js";
import { DateTimeServiceInstaller } from "/app/modules/datetime/services/DateTimeService.js";
import { InstallAppHttpInterceptors, InstallGlobalHttpInterceptors } from "/app/HttpInterceptors.js";
import { VersionServiceInstaller } from "/app/modules/version/services/VersionService.js";

/*
 * Core plugins
 */
Vue.use(VueRouter);
Vue.use(VueResource);
Vue.use(VueLoading);

/*
 * Syncfusion plugins
 */
Vue.use(ejs.buttons.ButtonPlugin);
Vue.use(ejs.grids.GridPlugin);
Vue.use(ejs.schedule.SchedulePlugin);
Vue.use(ejs.popups.DialogPlugin);
Vue.use(ejs.notifications.ToastPlugin);
Vue.use(ejs.charts.ChartPlugin);

/*
 * Services
 */
Vue.use(AlertServiceInstaller);
Vue.use(DateTimeServiceInstaller);
Vue.use(VersionServiceInstaller);

Vue.component("loading", VueLoading);

InstallGlobalHttpInterceptors(Vue);

new Vue({
	el: "#app",
	router,

	components: {
		Header,
	},

	async beforeCreate() {
		InstallAppHttpInterceptors(Vue, this);
	},

	template: `
	<div style="width: 100%">
		<Header></Header>

		<main role="main" class="flex-shrink-0 main-body">
			<div class="container-fluid">
				<router-view></router-view>
			</div>
		</main>

		<ejs-toast ref="toast" id="toast" width="100%"></ejs-toast>
	</div>
	`
});


