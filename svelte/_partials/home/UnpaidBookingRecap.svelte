<script>
  import { formatPrice } from '../../_components/xFormatter';

  /**
   * @typedef {Object} UnpaidBookingTenant
   * @property {number} bookingId
   * @property {string} tenantName
   * @property {string} roomName
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} dateStart
   * @property {string} dateEnd
   */

  /**
   * @typedef {Object} UnpaidBookingRecapItem
   * @property {string} tenantName
   * @property {number} unpaidAmount
   * @property {string} earliestStartDate
   * @property {number} overdueDays
   * @property {string[]} roomNames
   */

  /** @typedef {'tenantName' | 'unpaidAmount' | 'earliestStartDate' | 'overdueDays' | 'roomNames'} SortKey */

  const unpaidBookingTenants = /** @type {UnpaidBookingTenant[]} */ ([/* unpaidBookingTenants */]);
  let sortBy = /** @type {SortKey} */ ('overdueDays');

  /**
   * @param {string} dateStr
   * @returns {Date}
   */
  function parseDate(dateStr) {
    const [year, month, day] = dateStr.split('-').map(Number);
    return new Date(year, month - 1, day);
  }

  /**
   * @param {string} dateStr
   * @returns {number}
   */
  function daysSince(dateStr) {
    const date = parseDate(dateStr);
    const today = new Date();
    date.setHours(0, 0, 0, 0);
    today.setHours(0, 0, 0, 0);
    const msPerDay = 1000 * 60 * 60 * 24;
    return Math.max(0, Math.floor((today.getTime() - date.getTime()) / msPerDay));
  }

  /**
   * @param {string} left
   * @param {string} right
   * @returns {string}
   */
  function earliestDate(left, right) {
    if (!left) return right;
    if (!right) return left;
    return parseDate(left) <= parseDate(right) ? left : right;
  }

  /**
   * @param {SortKey} key
   */
  function sortByColumn(key) {
    sortBy = key;
  }

  /**
   * @param {string} roomName
   * @returns {boolean}
   */
  function isPurpleRoom(roomName) {
    return /^[ABC]/i.test((roomName || '').trim());
  }

  /**
   * @param {string[]} roomNames
   * @returns {boolean}
   */
  function hasPurpleRoom(roomNames) {
    return (roomNames || []).some(isPurpleRoom);
  }

  const recapItems = /** @type {UnpaidBookingRecapItem[]} */ (
    Object.values(
      (unpaidBookingTenants || []).reduce(
        /**
         * @param {Record<string, UnpaidBookingRecapItem>} acc
         * @param {UnpaidBookingTenant} booking
         */
        (acc, booking) => {
          const unpaidAmount = Math.max(0, (booking.totalPrice || 0) - (booking.totalPaid || 0));
          if (unpaidAmount <= 0) return acc;

          const key = booking.tenantName || `booking-${booking.bookingId}`;
          if (!acc[key]) {
            acc[key] = {
              tenantName: booking.tenantName,
              unpaidAmount: 0,
              earliestStartDate: booking.dateStart,
              overdueDays: 0,
              roomNames: [],
            };
          }

          acc[key].unpaidAmount += unpaidAmount;
          acc[key].earliestStartDate = earliestDate(acc[key].earliestStartDate, booking.dateStart);
          if (!acc[key].roomNames.includes(booking.roomName)) {
            acc[key].roomNames = [...acc[key].roomNames, booking.roomName];
          }

          return acc;
        },
        {},
      ),
    )
      .map(item => ({
        ...item,
        overdueDays: daysSince(item.earliestStartDate),
      }))
  );

  /** @type {UnpaidBookingRecapItem[]} */
  let sortedRecapItems = [];

  $: sortedRecapItems = [...recapItems].sort((a, b) => {
    switch (sortBy) {
      case 'tenantName':
        return (b.tenantName || '').localeCompare(a.tenantName || '');
      case 'unpaidAmount':
        return b.unpaidAmount - a.unpaidAmount || b.overdueDays - a.overdueDays;
      case 'earliestStartDate':
        return parseDate(b.earliestStartDate).getTime() - parseDate(a.earliestStartDate).getTime() || b.overdueDays - a.overdueDays;
      case 'roomNames':
        return b.roomNames.join(', ').localeCompare(a.roomNames.join(', '));
      case 'overdueDays':
      default:
        return b.overdueDays - a.overdueDays || b.unpaidAmount - a.unpaidAmount;
    }
  });
</script>

<section class="unpaid-booking-recap">
  <div class="header">
    <h1>Unpaid Recap</h1>
    <p>Click a header to sort that column descending.</p>
  </div>

  {#if sortedRecapItems.length > 0}
    <div class="table-wrap">
      <table>
        <thead>
          <tr>
            <th>
              <button class="sort-header" class:active={sortBy === 'tenantName'} on:click={() => sortByColumn('tenantName')}>
                <span>Tenant</span>
                <span class="sort-icon" aria-hidden="true">▼</span>
              </button>
            </th>
            <th>
              <button class="sort-header" class:active={sortBy === 'unpaidAmount'} on:click={() => sortByColumn('unpaidAmount')}>
                <span>Total Unpaid</span>
                <span class="sort-icon" aria-hidden="true">▼</span>
              </button>
            </th>
            <th>
              <button class="sort-header" class:active={sortBy === 'earliestStartDate'} on:click={() => sortByColumn('earliestStartDate')}>
                <span>Earliest Start</span>
                <span class="sort-icon" aria-hidden="true">▼</span>
              </button>
            </th>
            <th>
              <button class="sort-header" class:active={sortBy === 'overdueDays'} on:click={() => sortByColumn('overdueDays')}>
                <span>Days</span>
                <span class="sort-icon" aria-hidden="true">▼</span>
              </button>
            </th>
            <th>
              <button class="sort-header" class:active={sortBy === 'roomNames'} on:click={() => sortByColumn('roomNames')}>
                <span>Rooms</span>
                <span class="sort-icon" aria-hidden="true">▼</span>
              </button>
            </th>
          </tr>
        </thead>
        <tbody>
          {#each sortedRecapItems as item (`${item.tenantName}-${item.earliestStartDate}`)}
            <tr class:room-purple={hasPurpleRoom(item.roomNames)} class:room-yellow={!hasPurpleRoom(item.roomNames)}>
              <td>{item.tenantName}</td>
              <td>{formatPrice(item.unpaidAmount, 'IDR')}</td>
              <td>{item.earliestStartDate}</td>
              <td class="days">{item.overdueDays}</td>
              <td>{item.roomNames.join(', ')}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {:else}
    <div class="empty-state">
      <p>No unpaid recap.</p>
    </div>
  {/if}
</section>

<style>
  .unpaid-booking-recap {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .header {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .header h1,
  .header p {
    margin: 0;
  }

  .header h1 {
    font-size: var(--font-xl);
  }

  .header p {
    color: var(--gray-007);
  }

  .table-wrap {
    overflow-x: auto;
    border: 1px solid var(--gray-003);
    border-radius: 12px;
    background-color: #fff;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    min-width: 680px;
  }

  thead {
    background:
      linear-gradient(135deg, var(--red-transparent), transparent 70%),
      var(--gray-001);
  }

  th,
  td {
    padding: 3px 14px;
    text-align: left;
    border-bottom: 1px solid var(--gray-003);
    vertical-align: top;
  }

  th {
    padding: 0;
  }

  td {
    font-weight: 500;
    transition: background-color 0.15s ease;
  }

  .sort-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 8px;
    width: 100%;
    padding: 3px 14px;
    border: none;
    background: transparent;
    text-align: left;
    font-size: 12px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.04em;
    color: var(--gray-007);
    cursor: pointer;
    transition: color 0.15s ease, background-color 0.15s ease;
  }

  .sort-icon {
    flex: none;
    font-size: 10px;
    opacity: 0.45;
    transition: opacity 0.15s ease, transform 0.15s ease;
  }

  .sort-header:hover,
  .sort-header.active {
    color: var(--red-006);
  }

  .sort-header:hover .sort-icon,
  .sort-header.active .sort-icon {
    opacity: 1;
    transform: translateY(1px);
  }

  .sort-header:focus-visible {
    outline: 2px solid var(--red-005);
    outline-offset: -2px;
  }

  tbody tr.room-purple td {
    background-color: #faf5ff;
  }

  tbody tr.room-yellow td {
    background-color: #fffdf2;
  }

  tbody tr:hover td {
    background-color: var(--red-transparent);
  }

  tbody tr:last-child td {
    border-bottom: none;
  }

  .days {
    font-weight: 700;
    color: var(--red-006);
  }

  .empty-state {
    padding: 18px;
    border-radius: 12px;
    background-color: var(--gray-001);
  }

  .empty-state p {
    margin: 0;
  }

  @media only screen and (max-width: 768px) {
    td {
      font-size: 13px;
      padding: 3px 12px;
    }

    th {
      padding: 0;
    }

    .sort-header {
      padding: 3px 12px;
      font-size: 13px;
    }
  }
</style>
