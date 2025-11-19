<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/users.js').Tenant} Tenant */

  /**
   * @typedef {Object} PricePerDayReport
   * @property {string} roomName
   * @property {string} tenantName
   * @property {string} dateStart
   * @property {string} dateEnd
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} roomSize
   * @property {number} pricePerDayValue
   * @property {number} pricePerDayPercentage
   * @property {number} pricePerRoomValue
   * @property {number} pricePerRoomPercentage
   */

  import LayoutMain from './_layouts/main.svelte';
  import MonthShifter from './_components/MonthShifter.svelte';
  import { onMount } from 'svelte';
  import { StaffPricePerDayReport } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let pricePerDay = /** @type {PricePerDayReport[]} */ ([ /* pricePerDay */ ]);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);

  async function RefreshData() {
    await StaffPricePerDayReport({ yearMonth },
    /** @type {import('./jsApi.GEN').StaffPricePerDayReportCallback} */
    /** @returns {Promise<void>} */
    function(/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        notifier.showError(o.error || 'something went wrong');
        return
      }

      pricePerDay = o.pricePerDay;
      console.log('Price per day report:', pricePerDay);

      return
    })
  }

  function formatNumber(num) {
    return Number(num).toFixed(1);
  }

  /**
   * @description Calculate duration (days)
   * @param {string} dateStart
   * @param {string} dateEnd
   * @returns {number}
   */
  function calculateDurationDay(dateStart, dateEnd) {
    return Math.ceil((Date.parse(dateEnd) - Date.parse(dateStart)) / (1000 * 60 * 60 * 24));
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <MonthShifter
      bind:yearMonth
      bind:isLoading
      OnChanges={RefreshData}
    />
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Room</th>
            <th>Room Size</th>
            <th>Tenant</th>
            <th>Date Start</th>
            <th>Date End</th>
            <th>Duration</th>
            <th>Price Per Day</th>
            <th>Price Per Day Chart</th>
            <th>Price Per Day / Room</th>
            <th>Price Per Room Chart</th>
            <th>Total Paid</th>
            <th>Total Price</th>
          </tr>
        </thead>
        <tbody>
          {#each (pricePerDay || []) as data}
            <tr>
              <td>{data.roomName}</td>
              <td>{data.roomSize || '--'}</td>
              <td>{data.tenantName || '--'}</td>
              <td>{data.dateStart || '--'}</td>
              <td>{data.dateEnd || '--'}</td>
              <td>{calculateDurationDay(data.dateStart, data.dateEnd) || '0'} Days</td>
              <td>{formatNumber(data.pricePerDayValue || 0)}</td>
              <td>
                <div class="bar-container">
                  <div 
                    class="bar bar-price-per-day" 
                    style="width: {data.pricePerDayPercentage || 0}%"
                    title="{formatNumber(data.pricePerDayValue)} ({formatNumber(data.pricePerDayPercentage)}%)"
                  ><span class="bar-label">{formatNumber(data.pricePerDayPercentage)}%</span></div>
                </div>
              </td>
              <td>{formatNumber(data.pricePerRoomValue || 0)}</td>
              <td>
                <div class="bar-container">
                  <div 
                    class="bar bar-price-per-room" 
                    style="width: {data.pricePerRoomPercentage || 0}%"
                    title="{formatNumber(data.pricePerRoomValue)} ({formatNumber(data.pricePerRoomPercentage)}%)"
                  ><span class="bar-label">{formatNumber(data.pricePerRoomPercentage)}%</span></div>
                </div>
              </td>
              <td>{data.totalPaid || '0'}</td>
              <td>{data.totalPrice || '0'}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</LayoutMain>

<style>
  .report-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  .table-container {
    overflow-x: auto;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  table thead {
    position: sticky;
    top: 0;
    background-color: var(--gray-002);
    z-index: 10;
    border: none;
  }

  table thead tr th {
    text-align: left;
    padding: 8px 12px;
    text-wrap: nowrap;
  }

  table tbody tr {
    border-bottom: 1px solid var(--gray-004);
  }

  table tbody tr td {
    padding: 8px 12px;
  }

  /* Bar chart styles */
  .bar-container {
    width: 100px;
    height: 20px;
    background-color: var(--gray-003, #f0f0f0);
    border-radius: 4px;
    overflow: hidden;
    position: relative;
  }

  .bar {
    height: 100%;
    border-radius: 4px;
    transition: width 0.3s ease;
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .bar-label {
    font-size: 11px;
    font-weight: 600;
    color: white;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    white-space: nowrap;
    /* Hide jika bar terlalu kecil */
    display: none;
  }

  .bar:hover .bar-label {
    display: block;
  }

  .bar-price-per-day {
    background: linear-gradient(90deg, #4CAF50, #45a049);
  }

  .bar-price-per-room {
    background: linear-gradient(90deg, #2196F3, #1976D2);
  }

  @media only screen and (max-width : 768px) {
    .bar-container {
      width: 80px;
    }
  }
</style>