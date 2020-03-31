/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

export default {
	name: "Home",

	data() {
		return {
			regions: [],
			selectedRegions: [],
			startPoints: [],
			selectedStartPoint: "",
			seriesData: [],
			version: "",
			xAxis: {
				valueType: "Category",
				title: "Days",
			},
			yAxis: {
				title: "Confirmed Cases",
			},
			marker: {
				visible: true,
			},
			tooltip: {
				enable: true,
			},
		};
	},

	async created() {
		this.version = await this.versionService.GetVersion();
		this.getRegions();
		this.getStartPoints();
	},

	methods: {
		async getRegions() {
			try {
				let response = await this.$http.get(`/api/regions`);
				this.regions = response.body;
			} catch (e) {
				this.alertService.error(e);
			}
		},

		async getSeriesData(regions, startPoint) {
			try {
				let body = {
					regions: regions,
					startPoint: startPoint,
				};

				let response = await this.$http.post(`/api/lineseries`, body);
				this.seriesData = response.body;
			} catch (e) {
				this.alertService.error(e);
			}
		},

		async getStartPoints() {
			try {
				let response = await this.$http.get(`/api/startpoints`);
				this.startPoints = response.body;
			} catch (e) {
				this.alertService.error(e);
			}
		}
	},

	watch: {
		selectedRegions(newValue) {
			if (this.selectedStartPoint && newValue.length > 0) {
				this.getSeriesData(newValue, this.selectedStartPoint);
			} else {
				this.seriesData = [];
			}
		},

		selectedStartPoint(newValue) {
			if (newValue && this.selectedRegions.length > 0) {
				this.getSeriesData(this.selectedRegions, newValue);
			}
		},
	},

	template: `
	<div>
		<div class="row">
			<div class="col">
				<p>
					The chart allows you to select multiple countries and a start point. Once chosen the line graph
					shows the progression over a period of days of confirmed cases in each selected region. Click
					<router-link :to="{name: 'sources'}">here</router-link> to see the sources I used.
				</p>
				<p>
					There are two start points defined: <strong>First Major Step</strong>, and
					<strong>All Data</strong>. First Major Step is defined as the point when the selected region took
					a drastic measure to reduce the risk of spread, such as shutting down a city or prohibiting travel.
					When you choose this option the page displays the date and event I chose to use.

					For All Data this simply uses all the data from the start of recording in the John Hopkins database.
					This date is January 22nd, 2020.
				</p>
				<p>
					To begin, select the regions you wish to see, then choose a start point. The data should display
					in the chart below.
				</p>
			</div>
		</div>

		<div class="row mt-3">
			<div class="col">
				<div class="form-group">
					<label><strong>Regions:</strong></label>
					<ejs-multiselect
						:dataSource="regions"
						mode="CheckBox"
						:showSelectAll="true"
						v-model="selectedRegions">
					</ejs-multiselect>
				</div>

				<div class="form-group">
					<label><strong>Start Point:</strong></label>
					<div class="form-check">
						<div v-for="(s, index) in startPoints" :key="index">
							<input
								class="form-check-input"
								type="radio"
								name="starPoint"
								:id="'startRadio'+index"
								:value="s"
								v-model="selectedStartPoint" />
							<label class="form-check-label">{{s}}</label>
						</div>
					</div>
				</div>
			</div>

			<div class="col scrolling-info-panel">
				<p class="mt-3" v-if="selectedRegions.length && selectedStartPoint">
					Here is some information for the regions you selected:

					<ul>
						<li v-for="d in seriesData" :key="d.region">
							<strong>{{d.region}}:</strong>
							<ul>
								<li v-if="d.shutdownDescription">{{d.shutdownDescription}}</li>
								<li>{{dateTimeService.formatDate(d.effectiveStartDate)}}</li>
							</ul>
						</li>
					</ul>
				</p>
			</div>
		</div>

		<div class="row" v-if="seriesData.length">
			<div class="col">
				<ejs-chart title="Ramp Up" width="100%" height="550px" :primaryXAxis="xAxis" :primaryYAxis="yAxis" useGroupingSeparator="true" :tooltip="tooltip">
					<e-series-collection>
						<e-series
							v-for="d in seriesData"
							:key="d.region"
							:dataSource="d.seriesData"
							xName="dayNumber"
							yName="confirmedCases"
							:name="d.region"
							type="Line"
							width="3"
							:marker="marker">
						</e-series>
					</e-series-collection>
				</ejs-chart>
			</div>
		</div>
	</div>
	`
};

