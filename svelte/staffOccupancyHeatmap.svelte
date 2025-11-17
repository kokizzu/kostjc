<script>
	/** @typedef {import('./_types/masters.js').Access} Access */
	/** @typedef {import('./_types/users.js').User} User */
	/** @typedef {import('./_types/property.js').Booking} Booking */
	/** @typedef {import('./_types/property.js').BookingDetailPerMonth} BookingDetailPerMonth */
	/** @typedef {import('./_types/property.js').Facility} Facility */
	/** @typedef {import('./_types/property.js').Payment} Payment */

	import LayoutMain from './_layouts/main.svelte';
	import MonthShifter from './_components/MonthShifter.svelte';
	import { StaffOccupancyHeatmap } from './jsApi.GEN';
	import { notifier } from './_components/xNotifier';
	import { onMount } from 'svelte';
	import { localeDateFromYYYYMMDD } from './_components/xFormatter';
	import Tooltip from './_components/Tooltip.svelte';

	let user = /** @type {User} */ ({/* user */});
	let segments = /** @type {Access} */ ({/* segments */});
	let bookingsPerMonth = /** @type {any[]} */ ([/* bookingsPerMonth */]);
	let roomNames = /** @type {string[]} */ ([/* roomNames */]);
	let tenants = /** @type {Record<Number, string>} */({/* tenants */});

	let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
	let isLoading = /** @type {boolean} */ (false);

	function reColorBookings() {
		const tenantColors = {};

		// make list of tenantIds
		(() => {
			const uniqTenantIds = {}
			bookingsPerMonth.forEach((booking) => {
				uniqTenantIds[booking.tenantId] = '';
			})
			const tenantIds = []
			for(const tenantId in uniqTenantIds) {
				tenantIds.push(tenantId)
			}
			tenantIds.sort((a, b) => a - b);
			for(const [index, tenantId] of tenantIds.entries()) {
				const degree = Math.floor((360 * index) / (tenantIds.length + 1));
				tenantColors[tenantId] = `hsl(${degree}, 100%, 47%)`;
			}
		})()


		bookingsPerMonth.forEach((booking) => {
			booking.color = tenantColors[booking.tenantId];
		});
	}

	let totalDaysInSelectedMonth = /** @type {number} */ (31);

	function refreshTotalDaysInMonth() {
		const date = new Date(yearMonth + '-01');
		totalDaysInSelectedMonth = new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();
	}

	/**
	 * @description Check if date is include in n-day
	 * @param {number} day
	 * @param {string} dateStart
	 * @param {string} dateEnd
	 * @returns {boolean}
	 */
	function isDateIncludeInDay(day, dateStart, dateEnd) {
		if (!dateStart || !dateEnd) {
			return false
		}

		const date = new Date(yearMonth + '-' + (day < 10 ? '0' + day : day));
		return date >= new Date(dateStart) && date <= new Date(dateEnd);
	}

	/**
	 * @description Get booking by room name
	 * @param {string} roomName
	 * @returns {BookingDetailPerMonth[]}
	 */
	function getBookingsByRoomName(roomName) {
		let bookings = /** @type {BookingDetailPerMonth[]} */ ([]);

		for(let i = 0; i < bookingsPerMonth.length; i++) {
			if (bookingsPerMonth[i].roomName === roomName) {
				bookings.push(bookingsPerMonth[i]);
			}
		}

		return bookings;
	}

	onMount(async() => {
		refreshTotalDaysInMonth();
		await RefreshData();
		reColorBookings();
	});

	async function RefreshData() {
		try {
			isLoading = true;
			await StaffOccupancyHeatmap(// @ts-ignore
				{yearMonth},
				/** @type {import('./jsApi.GEN').StaffOccupancyHeatmapCallback} */
				/** @returns {Promise<void>} */
				function(/** @type any */ o) {
					if (o.error) {
						console.log(o);
						notifier.showError(o.error);
						return
					}

					bookingsPerMonth = o.bookingsPerMonth;
					roomNames = o.roomNames;

					refreshTotalDaysInMonth();
					reColorBookings();
				}
			)
		} catch(error) {
			console.error('Failed to load data:', error);
			notifier.showError('Failed to load data');
		} finally {
			isLoading = false;
		}
	}

	let tooltipComp = /** @type {Tooltip} */ (null);
	let tooltipData = /** @type {{tenant: string; room: string; dateStart: string; dateEnd: string}} */ ({
		tenant: '',
		dateStart: '',
		dateEnd: '',
		room: ''
	});

	/**
	 * @Description Show tooltip
	 * @param {MouseEvent & { currentTarget: EventTarget & HTMLDivElement }} event
	 * @param {BookingDetailPerMonth} bk
	 * @param {string} roomName
	 */
	function showTooltip(event, bk, roomName) {
		tooltipData = {
			tenant: tenants[bk.tenantId],
			room: roomName,
			dateStart: localeDateFromYYYYMMDD(bk.dateStart),
			dateEnd: localeDateFromYYYYMMDD(bk.dateEnd)
		}
		const rect = event.currentTarget.getBoundingClientRect();

		tooltipComp.Show(rect);
	}

	/**
	 * @param {BookingDetailPerMonth[]} bks
	 * @param {number} day
	 * @returns {[BookingDetailPerMonth|null, boolean]}
	 */
	function getBookingIsOccupied(bks, day) {
		const bk = bks.find(b => {
			return isDateIncludeInDay(day, b.dateStart, b.dateEnd);
		});

		return [bk, !!bk];
	}
</script>

<LayoutMain access={segments} user={user}>
	<div class="occupancy-heatmap-container">
		<MonthShifter
			bind:yearMonth
			bind:isLoading
			OnChanges={RefreshData}
		/>
		<div class="rooms-heatmap-wrapper">
			<div class="rooms-heatmap-table">
				<div class="table-header">
					<div class="room-label sticky-corner">Room / Day</div>
					<div class="days">
						{#each Array.from({length: totalDaysInSelectedMonth}) as _, day}
							<div class="date-header">{day + 1}</div>
						{/each}
					</div>
				</div>

				{#each roomNames as room}
					{@const bookings = getBookingsByRoomName(room)}
					<div class="table-row">
						<div class="room-name sticky-room">{room}</div>
						<div class="blocks">
							{#each Array.from({length: totalDaysInSelectedMonth}) as _, day}
								{@const [bk, avail] = getBookingIsOccupied(bookings, day + 1)}
								<!-- svelte-ignore a11y-no-static-element-interactions -->
								{#if avail}
									<div
										class="date-cell occupied"
										on:mouseenter={(e) => { if (bk) showTooltip(e, bk, room) }}
										on:mouseleave={() => tooltipComp.Hide()}
										aria-label="tool-tip"
										style="background-color: {bk.color}"
									></div>
								{:else}
									<div class="date-cell not-occupied"></div>
								{/if}
							{/each}
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
</LayoutMain>

<Tooltip bind:this={tooltipComp}>
	<div class="tooltip-content">
		<strong>Tenant:</strong> {tooltipData.tenant}<br/>
		<strong>Room:</strong> {tooltipData.room}<br/>
		<strong>Check-in:</strong> {tooltipData.dateStart}<br/>
		<strong>Check-out:</strong> {tooltipData.dateEnd}
	</div>
</Tooltip>

<style>
	.occupancy-heatmap-container {
		display: flex;
		flex-direction: column;
		gap: 30px;
		padding: 20px;
	}

	.rooms-heatmap-wrapper {
		position: relative;
		width: 100%;
		overflow: auto;
		max-height: calc(100vh - 200px);
	}

	.rooms-heatmap-table {
		display: flex;
		flex-direction: column;
		gap: 2px;
		width: 100%;
	}

	.table-header {
		display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center;
		position: sticky;
		top: 0;
		background-color: white;
		z-index: 20;
		font-weight: 700;
		border-bottom: 2px solid var(--gray-004);
		padding-bottom: 10px;
	}

	.table-header .days {
		width: 100%;
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(25px, 1fr));
		gap: 2px;
	}

	.room-label {
		padding: 8px;
		font-weight: 600;
		text-align: left;
		font-size: var(--font-md);
		width: 70px;
		text-overflow: ellipsis;
	}

	.sticky-corner {
		position: sticky;
		left: 0;
		z-index: 30;
		background-color: white;
	}

	.sticky-room {
		position: sticky;
		left: 0;
		z-index: 10;
		background-color: white;
	}

	.date-header {
		text-align: center;
		font-size: 10px;
		color: var(--gray-008);
		padding: 2px;
		background-color: white;
	}

	.date-cell {
		min-width: 28px;
		width: 28px;
		height: 28px;
		border-radius: 3px;
		cursor: pointer;
		transition: all 0.2s ease;
		position: relative;
	}

	.date-cell.occupied:hover {
		transform: scale(1.15);
		box-shadow: 0 2px 8px rgba(0, 100, 255, 0.3);
		z-index: 5;
		border-color: var(--blue-007);
		background-color: var(--blue-004);
	}

	.date-cell.not-occupied {
		border: 1px solid var(--gray-002) !important;
		background-color: var(--gray-001) !important;
		cursor: default;

	}

	.date-cell.not-occupied:hover {
		background-color: var(--gray-002) !important;
	}

	.table-row {
		display: flex;
		flex-direction: row;
		gap: 10px;
	}

	.table-row .room-name {
		padding: 8px;
		font-weight: 600;
		text-align: left;
		font-size: var(--font-md);
		width: 70px;
	}

	.table-row:hover {
		background-color: var(--gray-001);
	}

	.table-row:hover .sticky-room {
		background-color: var(--gray-001);
	}

	.table-row .blocks {
		width: 100%;
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(25px, 1fr));
		gap: 2px;
	}

	.tooltip-content {
		color: var(--gray-008);
		padding: 10px 12px;
		border-radius: 6px;
		font-size: 12px;
		white-space: nowrap;
	}

	@media (max-width: 768px) {
	.date-cell {
		min-width: 24px;
		width: 24px;
		height: 24px;
		border-radius: 2px;
	}
	
	.table-row .blocks {
		gap: 1px;
	}
	
	.table-header .days {
		gap: 1px;
	}
	
	
	.room-label,
	.table-row .room-name {
		width: 50px;
		font-size: 12px;
		padding: 6px;
	}
	}
	@keyframes fadeIn {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
</style>