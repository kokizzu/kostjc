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
  /** @typedef {'sortByRoom' | 'sortByPricePerDayPercent' | 'sortByPricePerRoomPercent'} SortBy */

  import LayoutMain from './_layouts/main.svelte';
  import MonthShifter from './_components/MonthShifter.svelte';
  import { StaffPricePerDayReport } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import Radio from './_components/Radio.svelte';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let pricePerDay = /** @type {PricePerDayReport[]} */ ([ /* pricePerDay */ ]);
  let sortedData  = /** @type {PricePerDayReport[]} */ ([]);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);
  let sortBy    = /** @type {SortBy} */ ('sortByRoom');

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


  /** @typedef {'sortByRoom' | 'sortByPricePerDayPercent' | 'sortByPricePerRoomPercent'} FilterValues */
  /** @typedef {import('./_types/masters.js').RadioOption<FilterValues>} RadioOption<FilterValues> */

  let selectedFilter = /** @type {FilterValues} */ ('sortByRoom');
  const filterName = /** @type {string} */ ('show-filter');
  const filterOptions = /** @type {RadioOption[]} */ ([
    {
      id: 'sortByRoom',
      name: filterName,
      value: 'sortByRoom',
      label: 'Sort By Room'
    },
    {
      id: 'sortByPricePerDayPercent',
      name: filterName,
      value: 'sortByPricePerDayPercent',
      label: 'Sort By Price Per Day Percent'
    },
    {
      id: 'sortByPricePerRoomPercent',
      name: filterName,
      value: 'sortByPricePerRoomPercent',
      label: 'Sort By Price Per Room Percent'
    },
  ]);

  function sortData() {
    sortedData = [...pricePerDay];

    switch(sortBy) {
      case 'sortByRoom':
        sortedData.sort((a, b) => compareRoomName(a.roomName, b.roomName));
        break;
      case 'sortByPricePerDayPercent':
        sortedData.sort((a, b) => (b.pricePerDayPercentage || 0) - (a.pricePerDayPercentage || 0));
        break;
      case 'sortByPricePerRoomPercent':
        sortedData.sort((a, b) => (b.pricePerRoomPercentage || 0) - (a.pricePerRoomPercentage || 0));
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

  $: if (selectedFilter) {
    handleSortChange(selectedFilter);
  }

  function formatNumber(num) {
    return Number(num).toFixed(1);
  }

  const durationPerDay = (1000 * 60 * 60 * 24)

  function calculateDurationDay(/** @type {string} */ dateStart, /** @type {string} */ dateEnd) {
    return Math.ceil((Date.parse(dateEnd) - Date.parse(dateStart)) / durationPerDay);
  }

  $: if (pricePerDay.length > 0) {
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
      <Radio
        className="filters"
        options={filterOptions}
        bind:selected={selectedFilter}
      />
    </div>

    <div class="table-wrapper">
      <div class="table-header">
        <div class="cell">Room</div>
        <div class="cell">Room Size</div>
        <div class="cell">Tenant</div>
        <div class="cell">Date Start</div>
        <div class="cell">Date End</div>
        <div class="cell">Duration</div>
        <div class="cell">Price Per Day</div>
        <div class="cell">Price Per Day Chart</div>
        <div class="cell">Price Per Day / Room</div>
        <div class="cell">Price Per Room Chart</div>
        <div class="cell">Total Paid</div>
        <div class="cell">Total Price</div>
      </div>

      <div class="table-body">
        {#each sortedData as data}
          <div class="table-row">
            <div class="cell">{data.roomName}</div>
            <div class="cell">{data.roomSize || '--'}</div>
            <div class="cell">{data.tenantName || '--'}</div>
            <div class="cell">{data.dateStart || '--'}</div>
            <div class="cell">{data.dateEnd || '--'}</div>
            <div class="cell">{calculateDurationDay(data.dateStart, data.dateEnd) || '0'} Days</div>
            <div class="cell">{formatNumber(data.pricePerDayValue || 0)}</div>
            <div class="cell">
              <div class="bar-container">
                <div 
                  class="bar bar-price-per-day" 
                  style="width: {data.pricePerDayPercentage || 0}%"
                  title="{formatNumber(data.pricePerDayValue)} ({formatNumber(data.pricePerDayPercentage)}%)"
                >
                  <span class="bar-label">{formatNumber(data.pricePerDayPercentage)}%</span>
                </div>
              </div>
            </div>
            <div class="cell">{formatNumber(data.pricePerRoomValue || 0)}</div>
            <div class="cell">
              <div class="bar-container">
                <div 
                  class="bar bar-price-per-room" 
                  style="width: {data.pricePerRoomPercentage || 0}%"
                  title="{formatNumber(data.pricePerRoomValue)} ({formatNumber(data.pricePerRoomPercentage)}%)"
                >
                  <span class="bar-label">{formatNumber(data.pricePerRoomPercentage)}%</span>
                </div>
              </div>
            </div>
            <div class="cell">{data.totalPaid || '0'}</div>
            <div class="cell">{data.totalPrice || '0'}</div>
          </div>
        {/each}
      </div>
    </div>
  </div>
</LayoutMain>

<style>
  :global(.filters) {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 10px;
    text-wrap: nowrap;
  }

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

  .table-wrapper {
    width: 100%;
    overflow-x: auto;
  }

  .table-header {
    display: grid;
    grid-template-columns: repeat(12, 1fr);
    background-color: var(--gray-002, #f8f9fa);
    position: sticky;
    top: 0;
    z-index: 10;
    border-bottom: 2px solid var(--gray-004, #e0e0e0);
  }

  .table-header .cell {
    padding: 12px;
    font-size: 13px;
    font-weight: 600;
    line-height: 1.4;
    text-align: left;
  }

  .table-body {
    display: flex;
    flex-direction: column;
  }

  .table-row {
    display: grid;
    grid-template-columns: repeat(12, 1fr);
    border-bottom: 1px solid var(--gray-004, #e0e0e0);
  }

  .table-row:hover {
    background-color: var(--gray-001, #fafafa);
  }

  .table-row .cell {
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

  @media only screen and (max-width: 1200px) {
    .report-container {
      padding: 12px;
    }

    .header-section {
      gap: 16px;
    }

    .table-wrapper {
      overflow-x: visible;
    }

    .table-header {
      grid-template-columns: repeat(6, 1fr);
      grid-template-rows: auto auto;
    }

    .table-row {
      grid-template-columns: repeat(6, 1fr);
      grid-template-rows: auto auto;
      gap: 8px 0;
      padding: 8px 0;
    }

    .table-header .cell,
    .table-row .cell {
      padding: 8px 6px;
      font-size: 11px;
    }

    .bar-container {
      width: 60px;
      height: 20px;
    }

    .bar-label {
      font-size: 9px;
    }
  }

  @media only screen and (max-width: 480px) {
    .table-header {
      grid-template-columns: repeat(4, 1fr);
      grid-template-rows: auto auto auto;
    }

    .table-row {
      grid-template-columns: repeat(4, 1fr);
      grid-template-rows: auto auto auto;
    }

    .table-header .cell,
    .table-row .cell {
      padding: 6px 4px;
      font-size: 10px;
    }
  }
</style>