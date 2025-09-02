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
   * @property {number} pricePerDay
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} roomSize
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

  /**
   * @description Calculate duration (days)
   * @param {string} dateStart
   * @param {string} dateEnd
   * @returns {number}
   */
  function calculateDurationDay(dateStart, dateEnd) {
    return Math.ceil((Date.parse(dateEnd) - Date.parse(dateStart)) / (1000 * 60 * 60 * 24));
  }

  /**
   * @description Calculate price per day
   * @param {number} totalPrice
   * @param {string} dateStart
   * @param {string} dateEnd
   * @returns {number}
   */
  function calculatePricePerDay(totalPrice, dateStart, dateEnd) {
    return Number((totalPrice / calculateDurationDay(dateStart, dateEnd)).toFixed(1));
  }

  /**
   * @description Calculate price per day per room
   * @param {number} totalPrice
   * @param {string} dateStart
   * @param {string} dateEnd
   * @param {string} roomSize
   * @returns {number}
   */
  function calculatePricePerDayPerRoom(totalPrice, dateStart, dateEnd, roomSize) {
    const [ roomSize1, roomSize2 ] = roomSize.split('x');
    const duration = calculateDurationDay(dateStart, dateEnd);

    return Number((totalPrice / duration / (Number(roomSize1) * Number(roomSize2))).toFixed(1));
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
            <th>Price Per Day / Room</th>
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
              <td>{calculatePricePerDay(data.totalPrice, data.dateStart, data.dateEnd) || '0'}</td>
              <td>{calculatePricePerDayPerRoom(data.totalPrice, data.dateStart, data.dateEnd, data.roomSize) || '0'}</td>
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

  @media only screen and (max-width : 768px) {

  }
</style>