<script>
    import { notifier } from './_components/xNotifier';
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').BookingDetail} BookingDetail */

  import LayoutMain from './_layouts/main.svelte';
  import { UserReport } from './jsApi.GEN';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiArrowsArrowRightSLine, RiArrowsArrowLeftSLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let bookingsPerQuartal = /** @type {BookingDetail[]} */ ([/* bookingsPerQuartal*/]);
  let roomNames = /** @type {string[]} */ ([/* roomNames */]);

  async function refreshBookings() {
    await UserReport(// @ts-ignore
      { monthStart, monthEnd }, /** @type {import('./jsApi.GEN').UserReportCallback} */
      /** @returns {Promise<void>} */
        function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        bookingsPerQuartal = o.bookingsPerQuartal
        roomNames = o.roomNames
      }
    );
  }

  function formatYM(date) {
    const y = date.getFullYear();
    const m = String(date.getMonth() + 1).padStart(2, '0');
    return `${y}-${m}`;
  }

  function parseYM(ym) {
    const [y, m] = ym.split('-').map(Number);
    return new Date(y, m - 1);
  }

  function getQuartalsFromEnd(endDateStr) {
    const endDate = parseYM(endDateStr);
    const result = [];

    for (let i = 3; i >= 0; i--) {
      const d = new Date(endDate.getFullYear(), endDate.getMonth() - i);
      result.push(formatYM(d));
    }

    return result;
  }

  function getQuartalsFromStart(startYM) {
    const [year, month] = startYM.split('-').map(Number);
    const result = [];
    for (let i = 0; i < 4; i++) {
      const d = new Date(year, month - 1 + i);
      result.push(formatYM(d));
    }
    return result;
  }

  const now = new Date();
  const initialStart = new Date(now.getFullYear(), now.getMonth());
  let monthStart = formatYM(initialStart);
  let quartals = getQuartalsFromStart(monthStart);
  let monthEnd = quartals[3];

  async function shiftMonth(direction) {
    const end = parseYM(monthEnd);
    end.setMonth(end.getMonth() + direction);
    monthEnd = formatYM(end);
    quartals = getQuartalsFromEnd(monthEnd);
    monthStart = quartals[0];

    await refreshBookings();
  }

  function formatMonthYear(ym) {
    const date = new Date(ym + "-01");
    return date.toLocaleString('default', { month: 'long', year: 'numeric' });
  }

  /**
   * @description Is Booking in that month
   * @param {BookingDetail} booking
   * @param {string} yearMonth
   * @returns {boolean}
   */
  function isBookingInThatMonth(booking, yearMonth) {
    return (booking.dateStart).includes(yearMonth) || (booking.dateEnd).includes(yearMonth);
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <header class="header" aria-label="Date range navigation">
        <button on:click={() => shiftMonth(-1)} class="btn" aria-label="Previous">
          <Icon size="20" src={RiArrowsArrowLeftSLine} />
        </button>
        <span class="header-text">{formatMonthYear(monthStart)} - {formatMonthYear(monthEnd)}</span>
        <button on:click={() => shiftMonth(1)} class="btn" aria-label="Next">
          <Icon size="20" src={RiArrowsArrowRightSLine} />
        </button>
      </header>
    </div>
    <table>
      <thead>
        <tr>
          <th></th>
          {#each (quartals || []) as quartal}
            <th>{formatMonthYear(quartal)}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each (roomNames || []) as room}
          <tr>
            <th>Room {room}</th>
            {#each (quartals || []) as quartal}
              <td>
                <div class="cells">
                  {#each (bookingsPerQuartal || []) as booking}
                    {#if booking.roomName == room && isBookingInThatMonth(booking, quartal)}
                      <div class="cell {booking.amountPaid >= booking.totalPrice ? '' :'unpaid'}">
                        {#if booking.tenantName}
                          <span>{booking.tenantName}</span> 
                          <span>{booking.dateStart} s/d {booking.dateEnd}</span> 
                          <span>{booking.amountPaid}/{booking.totalPrice}</span>
                        {/if}
                      </div>
                    {/if}
                  {/each}
                </div>
              </td>
            {/each}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</LayoutMain>

<style>
  .report-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  table thead tr th {
    text-align: left;
    padding: 8px 12px;
    border-bottom: 1px solid var(--gray-004)
  }

  table tbody tr {
    border-top: 1px solid var(--gray-002);
    border-bottom: 1px solid var(--gray-004);
  }

  table tbody tr td {
    padding: 8px 12px;
  }

  table tbody tr td .cells {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  table tbody tr td .cell {
    display: flex;
    flex-direction: column;
    gap: 2px;
    font-size: 12px;
  }

  table tbody tr td .cell.unpaid {
    padding: 5px 8px;
    background-color: var(--red-transparent);
    border-radius: 5px;
    color: var(--red-006);
  }

  .header {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
    margin-bottom: 32px;
  }
  .btn {
    background-color: var(--gray-002);
    border: none;
    border-radius: 9999px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--gray-008);
    padding: 8px;
  }
  .btn:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }
  .header-text {
    font-weight: 600;
    font-size: 18px;
    user-select: none;
  }
</style>