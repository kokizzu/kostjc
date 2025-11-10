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
    <div class="rooms-heatmap-wrapper">
      <div class="rooms-heatmap-table">
        <div class="table-header">
          <div class="room-label sticky-corner">Room</div>
          {#each Array.from({ length: totalDaysInSelectedMonth }) as _, day}
            <div class="date-header">{day + 1}</div>
          {/each}
        </div>
       
        {#each roomNames as room}
          {@const bk = getBookingByRoomName(room)}
          <div class="table-row">
            <div class="room-name sticky-room">Room {room}</div>
            {#each Array.from({ length: totalDaysInSelectedMonth }) as _, day}
              <div class="date-cell {isDateIncludeInDay(day + 1, (bk || {})['dateStart'] || '', (bk || {})['dateEnd'] || '') ? 'occupied' : 'not-occupied'}">
              </div>
            {/each}
          </div>
        {/each}
      </div>
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
    width: fit-content;
    min-width: 100%;
  }

  .table-header,
  .table-row {
    display: grid;
    grid-template-columns: 120px repeat(31, minmax(40px, 1fr));
    gap: 2px;
    align-items: center;
  }

  .table-header {
    position: sticky;
    top: 0;
    background-color: white;
    z-index: 20;
    font-weight: 700;
    border-bottom: 2px solid var(--gray-004);
    padding-bottom: 10px;
  }

  .room-label,
  .room-name {
    padding: 15px;
    font-weight: 600;
    text-align: left;
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
    font-size: var(--font-sm);
    color: var(--gray-008);
    padding: 5px;
    background-color: white;
  }

  .date-cell {
    height: 40px;
    border-radius: 4px;
    cursor: pointer;
  }

  .date-cell.occupied {
    border: 1px solid var(--blue-005);
    background-color: var(--blue-transparent);
  }

  .date-cell.not-occupied {
    border: 1px solid var(--gray-002);
    background-color: var(--gray-001);
  }

  .table-row:hover {
    background-color: var(--gray-001);
  }

  .table-row:hover .sticky-room {
    background-color: var(--gray-001);
  }
</style>