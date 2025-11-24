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
  /** @typedef {'room' | 'pricePerDayPercent' | 'pricePerRoomPercent' | 'pricePerDayValue'} SortBy */

  import LayoutMain from './_layouts/main.svelte';
  import MonthShifter from './_components/MonthShifter.svelte';
  import { StaffPricePerDayReport } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let pricePerDay = /** @type {PricePerDayReport[]} */ ([ /* pricePerDay */ ]);
  let sortedData  = /** @type {PricePerDayReport[]} */ ([]);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);
  let sortBy    = /** @type {SortBy} */ ('room');

  async function RefreshData() {
    isLoading = true;
    await StaffPricePerDayReport({ yearMonth },
    /** @type {import('./jsApi.GEN').StaffPricePerDayReportCallback} */
    /** @returns {Promise<void>} */
    function(/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        notifier.showError(o.error || 'something went wrong');
        isLoading = false;
        return
      }

      pricePerDay = o.pricePerDay || [];
      sortData();
      isLoading = false;
      return
    })
  }

  function sortData() {
    sortedData = [...pricePerDay];

    switch(sortBy) {
      case 'room':
        sortedData.sort((a, b) => compareRoomName(a.roomName, b.roomName));
        break;
      case 'pricePerDayPercent':
        sortedData.sort((a, b) => (b.pricePerDayPercentage || 0) - (a.pricePerDayPercentage || 0));
        break;
      case 'pricePerRoomPercent':
        sortedData.sort((a, b) => (b.pricePerRoomPercentage || 0) - (a.pricePerRoomPercentage || 0));
        break;
      case 'pricePerDayValue':
        sortedData.sort((a, b) => (b.pricePerDayValue || 0) - (a.pricePerDayValue || 0));
        break;
    }
  }

  function compareRoomName(room1, room2) {
    if (!room1 || !room2) return room1 < room2 ? -1 : 1;

    const { building: b1, floor: f1, number: n1 } = parseRoomName(room1);
    const { building: b2, floor: f2, number: n2 } = parseRoomName(room2);

    if (b1 !== b2) return b1 < b2 ? -1 : 1;
    if (f1 !== f2) return f1 - f2;
    return n1 - n2;
  }

  function parseRoomName(/** @type {string} */ roomName) {
    if (!roomName || roomName.length < 2) {
      return { building: '', floor: 0, number: 0 };
    }

    const building = roomName[0];
    const digits = roomName.slice(1);

    if (digits.length === 0) {
      return { building, floor: 0, number: 0 };
    }

    let floor = 0;
    let number = 0;

    if (digits.length === 2) {
      floor = parseInt(digits[0], 10);
      number = parseInt(digits[1], 10);
    } else if (digits.length === 3) {
      floor = parseInt(digits.substring(0, 2), 10);
      number = parseInt(digits[2], 10);
    } else {
      const floorStr = digits.substring(0, digits.length - 1);
      floor = parseInt(floorStr, 10) || 0;
      number = parseInt(digits[digits.length - 1], 10) || 0;
    }

    return { building, floor, number };
  }

  function handleSortChange(/** @type {SortBy} */ newSort) {
    sortBy = newSort;
    sortData();
  }

  function formatNumber(/** @type {number} */ num) {
    return Number(num).toFixed(1);
  }

  const durationPerDay = (1000 * 60 * 60 * 24)

  function calculateDurationDay(/** @type {string} */ dateStart, /** @type {string} */ dateEnd) {
    return Math.ceil((Date.parse(dateEnd) - Date.parse(dateStart)) / durationPerDay);
  }

  // Initial data sort on mount
  $: if ((pricePerDay || []).length > 0) {
    sortData();
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="header-section">
      <MonthShifter
        bind:yearMonth
        bind:isLoading
        OnChanges={RefreshData}
      />
      
      <div class="sort-controls">
        <label for="sort">Sort by:</label>
        <div class="sort-buttons" id="sort">
          <button 
            class="sort-btn" 
            class:active={sortBy === 'room'}
            on:click={() => handleSortChange('room')}
          >
            Room Name
          </button>
          <button 
            class="sort-btn" 
            class:active={sortBy === 'pricePerDayValue'}
            on:click={() => handleSortChange('pricePerDayValue')}
          >
            Price/Day
          </button>
          <button 
            class="sort-btn" 
            class:active={sortBy === 'pricePerDayPercent'}
            on:click={() => handleSortChange('pricePerDayPercent')}
          >
            Price/Day %
          </button>
          <button 
            class="sort-btn" 
            class:active={sortBy === 'pricePerRoomPercent'}
            on:click={() => handleSortChange('pricePerRoomPercent')}
          >
            Price/Room %
          </button>
        </div>
      </div>
    </div>

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
          {#each sortedData as data}
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
                  >
                    <span class="bar-label">{formatNumber(data.pricePerDayPercentage)}%</span>
                  </div>
                </div>
              </td>
              <td>{formatNumber(data.pricePerRoomValue || 0)}</td>
              <td>
                <div class="bar-container">
                  <div 
                    class="bar bar-price-per-room" 
                    style="width: {data.pricePerRoomPercentage || 0}%"
                    title="{formatNumber(data.pricePerRoomValue)} ({formatNumber(data.pricePerRoomPercentage)}%)"
                  >
                    <span class="bar-label">{formatNumber(data.pricePerRoomPercentage)}%</span>
                  </div>
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
    gap: 20px;
    padding: 20px;
  }

  .header-section {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .sort-controls {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
  }

  .sort-controls label {
    font-weight: 600;
    font-size: 14px;
    color: var(--gray-007, #333);
  }

  .sort-buttons {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
  }

  .sort-btn {
    padding: 8px 16px;
    border: 1px solid var(--gray-004, #ddd);
    background: white;
    border-radius: 6px;
    font-size: 14px;
    cursor: pointer;
    transition: all 0.2s ease;
    white-space: nowrap;
  }

  .sort-btn:hover {
    background: var(--gray-002, #f5f5f5);
    border-color: var(--gray-005, #ccc);
  }

  .sort-btn.active {
    background: var(--primary-color, #4CAF50);
    color: white;
    border-color: var(--primary-color, #4CAF50);
  }

  .table-container {
    overflow-x: auto;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    min-width: 1200px;
  }

  table thead {
    position: sticky;
    top: 0;
    background-color: var(--gray-002, #f8f9fa);
    z-index: 10;
    border: none;
  }

  table thead tr th {
    text-align: left;
    padding: 12px;
    white-space: normal;
    word-wrap: nowrap;
    max-width: 120px;
    font-size: 13px;
    font-weight: 600;
    line-height: 1.4;
  }

  table tbody tr {
    border-bottom: 1px solid var(--gray-004, #e0e0e0);
  }

  table tbody tr:hover {
    background-color: var(--gray-001, #fafafa);
  }

  table tbody tr td {
    padding: 12px;
    font-size: 14px;
  }

  .bar-container {
    width: 100px;
    height: 24px;
    background-color: var(--gray-003, #f0f0f0);
    border-radius: 4px;
    overflow: hidden;
    position: relative;
  }

  .bar {
    height: 100%;
    border-radius: 4px;
    transition: width 0.3s ease, opacity 0.2s ease;
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }

  .bar:hover {
    opacity: 0.85;
  }

  .bar-label {
    font-size: 11px;
    font-weight: 600;
    color: white;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.4);
    white-space: nowrap;
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
    .report-container {
      padding: 12px;
    }

    .header-section {
      gap: 16px;
    }

    .sort-controls {
      flex-direction: column;
      align-items: flex-start;
      gap: 8px;
    }

    .sort-buttons {
      width: 100%;
    }

    .sort-btn {
      flex: 1;
      min-width: 0;
      padding: 10px 12px;
      font-size: 13px;
    }

    .bar-container {
      width: 80px;
      height: 20px;
    }

    table {
      min-width: 1000px;
    }

    table thead tr th {
      padding: 10px 8px;
      font-size: 12px;
      max-width: 100px;
    }

    table tbody tr td {
      padding: 10px 8px;
      font-size: 13px;
    }
  }
</style>