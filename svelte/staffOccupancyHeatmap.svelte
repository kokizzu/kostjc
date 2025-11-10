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

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);

  function reColorBookings() {
    const total = bookingsPerMonth.length || 0;
    (bookingsPerMonth || []).forEach((booking, idx) => {
      const degree = Math.floor( (360 * idx) / (total + 1) ); // since 360 in hsl = 0
      const color = 'hsl(' + degree + ', 100%, 47%)';

      booking.color = color;
    })
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
   * @returns {BookingDetailPerMonth}
   */
  function getBookingByRoomName(roomName) {
    for (let i = 0; i < bookingsPerMonth.length; i++) {
      if (bookingsPerMonth[i].roomName === roomName) {
        return bookingsPerMonth[i];
      }
    }

    return {
      id: 0,
      roomId: 0,
      roomName: '',
      tenantId: 0,
      tenantName: '',
      dateStart: '',
      dateEnd: '',
      totalPaid: 0,
      totalPrice: 0,
      extraTenants: [],
      color: ''
    };
  }

  onMount(() => {
    refreshTotalDaysInMonth();
    reColorBookings();
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
        reColorBookings();
      }
    )
  }


  let tooltipData = null;
  let tooltipPosition = { x: 0, y: 0 };
  
  /**
   * @param {MouseEvent} event
   * @param {BookingDetailPerMonth} booking
   * @param {number} day
   * @param {string} room
   */
  function showTooltip(event, booking, day, room) {
    if (!booking) {
      return;
    }
    
    tooltipData = {
      tenant: tenants[booking.tenantId] || 'Unknown',
      room: room,
      dateStart: booking.dateStart,
      dateEnd: booking.dateEnd,
      day: day
    };
    
    tooltipPosition = {
      x: event.clientX,
      y: event.clientY
    };
  }

  function hideTooltip() {
    tooltipData = null;
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
                {@const isOccupied = isDateIncludeInDay(day + 1, bk.dateStart || '', (bk || {})['dateEnd'] || '')}
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div 
                  class="date-cell {isOccupied ? 'occupied' : 'not-occupied'}"
                  on:mouseenter={(e) => {
                    if (isOccupied && bk) {
                      showTooltip(e, bk, day + 1, room);
                    }
                  }}
                  on:mouseleave={hideTooltip}
                  aria-label="tool-tip"
                  style="background-color: {bk.color}"
                >
                </div>
              {/each}
          </div>
        {/each}
      </div>
    </div>
  </div>
</LayoutMain>

{#if tooltipData}
  <div 
    class="tooltip" 
    style="left: {tooltipPosition.x + 10}px; top: {tooltipPosition.y + 10}px"
  >
    <div class="tooltip-content">
      <strong>Penghuni:</strong> {tooltipData.tenant}<br/>
      <strong>Kamar:</strong> {tooltipData.room}<br/>
      <strong>Check-in:</strong> {tooltipData.dateStart}<br/>
      <strong>Check-out:</strong> {tooltipData.dateEnd}
    </div>
  </div>
{/if}

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

  .table-header,
  .table-row {
    display: grid;
    grid-template-columns: 100px repeat(auto-fit, minmax(25px, 1fr));
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
    padding: 8px;
    font-weight: 600;
    text-align: left;
    font-size: var(--font-md);
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

  .table-row:hover {
    background-color: var(--gray-001);
  }

  .table-row:hover .sticky-room {
    background-color: var(--gray-001);
  }

  .tooltip {
    position: fixed;
    z-index: 1000;
    pointer-events: none;
    animation: fadeIn 0.2s;
  }

  .tooltip-content {
    background-color: white;
    color: black;
    padding: 10px 12px;
    border-radius: 6px;
    font-size: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    white-space: nowrap;
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }
</style>