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

  let user              = /** @type {User} */ ({/* user */});
  let segments          = /** @type {Access} */ ({/* segments */});
  let bookingsPerMonth  = /** @type {any[]} */ ([/* bookingsPerMonth */]);
  let roomNames         = /** @type {string[]} */ ([/* roomNames */]);
  let tenants           = /** @type {Record<Number, string>} */({/* tenants */});
  let rooms             = /** @type {Record<Number, string>} */({/* rooms */});

  console.log('Bookings: ', bookingsPerMonth);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);

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

    const date = new Date(yearMonth + '-' + (day < 10 ? '0' + day : day));[]
    return date >= new Date(dateStart) && date <= new Date(dateEnd);
  }

  /**
   * @description Get booking by room name
   * @param {string} roomName
   * @returns {BookingDetailPerMonth}
   */
  function getBookingByRoomName(roomName) {
    for (let i = 0; i < bookingsPerMonth.length; i++) {
      if (bookingsPerMonth[i].roomName === roomName) {
        return bookingsPerMonth[i];
      }
    }

    return null;
  }

  onMount(() => {
    refreshTotalDaysInMonth();
  });

  async function RefreshData() {
    await StaffOccupancyHeatmap(// @ts-ignore
      { yearMonth },
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
      }
    )
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="occupancy-heatmap-container">
    <MonthShifter
      bind:yearMonth
      bind:isLoading
      OnChanges={RefreshData}
    />
    <div class="rooms-heatmap">
      {#each roomNames as room}
      {@const bk = getBookingByRoomName(room)}
        <div class="heatmap-container">
          <h2>Room {room}</h2>
          <div class="heatmap-grid">
            {#each Array.from({ length: totalDaysInSelectedMonth }) as _, day}
              <span class="heatmap-cell {isDateIncludeInDay(day, (bk || {})['dateStart'] || '', (bk || {})['dateEnd'] || '') ? 'occupied' : 'not-occupied'}">
                {day+1}
              </span>
            {/each}
          </div>
        </div>
      {/each}
    </div>
  </div>
</LayoutMain>

<style>
  .occupancy-heatmap-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  .rooms-heatmap {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .heatmap-container {
    display: grid;
    grid-template-columns: 100px 1fr;
    align-items: center;
    gap: 20px;
    width: 100%;
    height: fit-content;
  }

  .heatmap-container h2 {
    margin: 0;
  }

  .heatmap-container .heatmap-grid {
    display: grid;
    grid-template-columns: repeat(11, 1fr);
    gap: 5px;
    flex-grow: 1;
  }

  .heatmap-container .heatmap-grid .heatmap-cell {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px;
    border-radius: 4px;
    width: 100%;
    font-size: var(--font-lg);
    font-weight: 600;
    cursor: pointer;
  }

  .heatmap-container .heatmap-grid .heatmap-cell.occupied {
    border: 1px solid var(--blue-005);
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }

  .heatmap-container .heatmap-grid .heatmap-cell.not-occupied {
    border: 1px solid var(--gray-002);
    background-color: var(--gray-001);
    color: var(--gray-008);
  }
</style>