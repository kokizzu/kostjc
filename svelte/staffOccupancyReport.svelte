<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Booking} Booking */
  /** @typedef {import('./_types/property.js').BookingDetail} BookingDetail */
  /** @typedef {import('./_types/property.js').Facility} Facility */
  /** @typedef {import('./_types/property.js').Payment} Payment */

  import LayoutMain from './_layouts/main.svelte';
  import { AdminBooking, AdminPayment, StaffOccupancyReport } from './jsApi.GEN';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import {
    RiArrowsArrowRightSLine, RiArrowsArrowLeftSLine,
    RiSystemExternalLinkLine, RiFinanceWallet3Line,
    RiSystemEyeLine, RiDesignBallPenLine
  } from './node_modules/svelte-icons-pack/dist/ri';
  import PopUpExtendBooking from './_components/PopUpExtendBooking.svelte';
  import Switcher from './_components/Switcher.svelte';
  import { notifier } from './_components/xNotifier';
  import { onMount } from 'svelte';
  import { dateISOFormatFromYYYYMMDD } from './_components/xFormatter';
  import { CmdForm, CmdUpsert } from './_components/xConstant';
  import PopUpShowBookingPayments from './_components/PopUpShowBookingPayments.svelte';
  import PopUpAddPayment from './_components/PopUpAddPayment.svelte';
  import PopUpEditBooking from './_components/PopUpEditBooking.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let bookingsPerQuartal = /** @type {BookingDetail[]} */ ([/* bookingsPerQuartal*/]);
  let roomNames = /** @type {string[]} */ ([/* roomNames */]);
  let tenants     = /** @type {Record<Number, string>} */({/* tenants */});
  let rooms       = /** @type {Record<Number, string>} */({/* rooms */});
  let facilities  = /** @type {Facility[]} */ ([/* facilities */]);

  async function refreshBookings() {
    await StaffOccupancyReport(// @ts-ignore
      { monthStart, monthEnd }, /** @type {import('./jsApi.GEN').StaffOccupancyReportCallback} */
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

  function formatYM(/** @type {Date} */ date) {
    const y = date.getFullYear();
    const m = String(date.getMonth() + 1).padStart(2, '0');
    return `${y}-${m}`;
  }

  function parseYM(/** @type {string} */ ym) {
    const [y, m] = ym.split('-').map(Number);
    return new Date(y, m - 1);
  }

  function getQuartalsFromEnd(/** @type {string} */ endDateStr) {
    const endDate = parseYM(endDateStr);
    const result = [];

    for (let i = 3; i >= 0; i--) {
      const d = new Date(endDate.getFullYear(), endDate.getMonth() - i);
      result.push(formatYM(d));
    }

    return result;
  }

  function getQuartalsFromStart(/** @type {string} */ startYM) {
    const [year, month] = startYM.split('-').map(Number);
    const result = [];
    for (let i = 0; i < 4; i++) {
      const d = new Date(year, month - 1 + i);
      result.push(formatYM(d));
    }
    return result;
  }

  const now = new Date();
  const initialStart = new Date(now.getFullYear(), now.getMonth()-2);
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

  function formatMonthYear(/** @type {string} */ ym) {
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
    return (yearMonth >= (booking.dateStart).slice(0, 7)) && (yearMonth <= (booking.dateEnd).slice(0, 7));
  }

  let showRefunded    = true;
  let showTenant      = true;
  let showDateStart   = true;
  let showDateEnd     = true;
  let showPaid        = true;
  let showPrice       = true;
  let showOnlyNotPaid = false;

  let isPopUpFormReady = false;
  onMount(() => isPopUpFormReady = true);

  let popupExtendBooking = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpExtendBooking | any} */ (null);
  let bookingToExtend = /** @type {BookingDetail} */ ({});
  let isSubmitExtendBooking = false;

  let dateStart = '';
  let dateEnd = '';

  async function SubmitExtendBooking(/** @type {Booking} */ booking, /** @type {number[]} */ facilities) {
    isSubmitExtendBooking = true;
    booking.tenantId = bookingToExtend.tenantId+'';
    booking.roomId = String(bookingToExtend.roomId);
    const i = /** @type {any} */ ({
      facilities: facilities,
      booking: booking,
      cmd: CmdUpsert
    });
    await AdminBooking(i,
      /** @type {import('../jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitExtendBooking = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        notifier.showSuccess(`Booking #${booking.id} extended !!`);

        popupExtendBooking.Reset();
        popupExtendBooking.Hide();
        refreshBookings();
      }
    );
  }

  function onExtendBooking(/** @type {BookingDetail} */ booking) {
    bookingToExtend = booking;
    dateStart = dateISOFormatFromYYYYMMDD(booking.dateEnd, 1);
    dateEnd = dateISOFormatFromYYYYMMDD(booking.dateEnd, 30);
    popupExtendBooking.Show();
  }

  let popUpShowPayments = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpShowBookingPayments | any} */ (null);
  let bookingIdToShowPayment = /** @type {number} */ (0);
  let paymentsForBooking = /** @type {Payment[]} */ ([]);

  async function showPaymentsForBooking(/** @type {number} */ bookingId) {
    bookingIdToShowPayment = bookingId;  // @ts-ignore
    await AdminPayment({
      cmd: CmdForm,
      bookingId: bookingId
    }, /** @type {import('../jsApi.GEN').AdminPaymentCallback} */
    /** @returns {Promise<void>} */
      function(/** @type any */ o) {
      if (o.error) {
        console.log(o);
        notifier.showError(o.error);
        return
      }

      paymentsForBooking = o.paymentsByBooking;
      console.log(paymentsForBooking);
    });

    popUpShowPayments.Show();
  }

  let popUpAddPayment = null;
  let bookingIdToPay = 0;
  let isSubmitAddPayment = false;

  function showInputPayment(/** @type {number} */ bookingId) {
    bookingIdToPay = bookingId;
    popUpAddPayment.Show();
  }

  async function submitAddPayment(/** @type {Payment} */ payment) {
    isSubmitAddPayment = true;
    payment.bookingId = bookingIdToPay+'';
    const i = /** @type {any} */ ({
      payment,
      cmd: 'upsert'
    });
    await AdminPayment(i,
      /** @type {import('../jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddPayment = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        notifier.showSuccess(`Payment created for booking #${bookingIdToPay} !!`);

        popUpAddPayment.Reset();
        popUpAddPayment.Hide();
        refreshBookings();
      }
    );
  }

  /**
   * @description Reorder bookings in a month, make sure the newest booking is in the end
   * @param {string} roomName
   * @param {string} quartal
   * @returns {BookingDetail[]}
   */
  function reOrderBookingPerQuartal(roomName, quartal) {
    let result = /** @type {BookingDetail[]} */ ([]);
    (bookingsPerQuartal || []).forEach((booking) => {
      if (booking.roomName === roomName && isBookingInThatMonth(booking, quartal)) {
        result.push(booking);
      }
    });

    result.sort((a, b) => {
      return new Date(a.dateStart).getTime() - new Date(b.dateStart).getTime();
    });

    return result;
  }

  let popUpEditBooking = null;
  let bookingIdToEdit = 0;
  let isSubmitEditBooking = false;

  /**
   * @description Show edit booking form
   * @param {BookingDetail} booking
   */
  async function showEditBooking(booking) {
    bookingIdToEdit = booking.id;
    popUpEditBooking.Show();
    popUpEditBooking.GetBookingById(bookingIdToEdit);
  }

  async function SubmitEditBooking(/** @type {Booking} */ booking) {
    const i = /** @type {any}*/ ({
      booking,
      cmd: CmdUpsert
    });
    await AdminBooking(i,
      /** @type {import('./jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        notifier.showSuccess(`Booking #${booking.id} updated !!`);
        refreshBookings();
        popUpEditBooking.Hide();
      }
    );
    isSubmitEditBooking = false;
  }
</script>

{#if isPopUpFormReady}
  <PopUpExtendBooking
    bind:this={popupExtendBooking}
    bind:isSubmitted={isSubmitExtendBooking}
    bind:bookingToExtend
    bind:dateStart
    bind:dateEnd
    facilities={facilities}
    tenants={tenants}
    OnSubmit={SubmitExtendBooking}
  />

  <PopUpShowBookingPayments
    bind:bookingId={bookingIdToShowPayment}
    bind:payments={paymentsForBooking}
    bind:this={popUpShowPayments}
    Refresh={refreshBookings}
  />

  <PopUpAddPayment
    isBookingReadOnly
    bind:bookingId={bookingIdToPay}
    bind:this={popUpAddPayment}
    bind:isSubmitted={isSubmitAddPayment}
    OnSubmit={submitAddPayment}
  />

  <PopUpEditBooking
    bind:this={popUpEditBooking}
    bind:bookingId={bookingIdToEdit}
    bind:isSubmitted={isSubmitEditBooking}
    tenants={tenants}
    rooms={rooms}
    OnSubmit={SubmitEditBooking}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <div class="quartal-shifter">
        <button on:click={() => shiftMonth(-1)} class="btn" aria-label="Previous">
          <Icon size="20" src={RiArrowsArrowLeftSLine} />
        </button>
        <span class="month-text">{formatMonthYear(monthStart)} - {formatMonthYear(monthEnd)}</span>
        <button on:click={() => shiftMonth(1)} class="btn" aria-label="Next">
          <Icon size="20" src={RiArrowsArrowRightSLine} />
        </button>
      </div>
      <div class="filters">
        <Switcher
          id="show-refunded"
          label="Show Refunded"
          bind:value={showRefunded}
        />
        <Switcher
          id="show-tenant"
          label="Show Tenant"
          bind:value={showTenant}
        />
        <Switcher
          id="show-date-start"
          label="Show Date Start"
          bind:value={showDateStart}
        />
        <Switcher
          id="show-date-end"
          label="Show Date End"
          bind:value={showDateEnd}
        />
        <Switcher
          id="show-paid"
          label="Show Paid"
          bind:value={showPaid}
        />
        <Switcher
          id="show-price"
          label="Show Price"
          bind:value={showPrice}
        />
        <Switcher
          id="show-only-not-paid"
          label="Show Only Not Paid"
          bind:value={showOnlyNotPaid}
        />
      </div>
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
                  {#each reOrderBookingPerQuartal(room.replace(' *', ''), quartal) as booking}
                    <div
                      class="cell
                      {booking.deletedAt > 0 ? 'refunded' : ''}
                      {showOnlyNotPaid && booking.amountPaid >= booking.totalPrice ? 'hidden' : ''}
                      {!showRefunded && booking.deletedAt > 0 ? 'hidden' : ''}
                      "
                    >
                      {#if booking.tenantName}
                        <span class="{showTenant ? '' : 'hidden'}">{booking.tenantName}</span> 
                        <span>
                          <span class="{showDateStart ? '' : 'hidden'}">{booking.dateStart}</span>
                          <span> s/d </span>
                          <span
                            class="
                            {!booking.isNearEnding && !booking.isExtended ? 'date-warning' : ''}
                            {!booking.isExtended && booking.isNearEnding ? 'date-alert' : ''}
                            {showDateEnd ? '' : 'hidden'}
                          ">
                            {booking.dateEnd}
                          </span>
                        </span>
                        <span class="{(booking.amountPaid == booking.totalPrice) ? "" : `${booking.deletedAt == 0 ? 'text-red' : ''}`}">
                          <span class="{showPaid ? '' : 'hidden'}">{booking.amountPaid}</span>
                          <span>/</span>
                          <span class="{showPrice ? '' : 'hidden'}">{booking.totalPrice}</span> 
                        </span>
                        <div class="actions">
                          <button class="btn" title="Extend Booking" on:click={() => onExtendBooking(booking)}>
                            <Icon
                              src={RiSystemExternalLinkLine}
                              size="17"
                            />
                          </button>
                          <button class="btn" title="Input Payment" on:click={() => showInputPayment(booking.id)}>
                            <Icon
                              src={RiFinanceWallet3Line}
                              size="17"
                            />
                          </button>
                          <button class="btn" title="Show Payments" on:click={() => showPaymentsForBooking(booking.id)}>
                            <Icon
                              src={RiSystemEyeLine}
                              size="17"
                            />
                          </button>
                          <button class="btn" title="Edit Booking" on:click={() => showEditBooking(booking)}>
                            <Icon
                              src={RiDesignBallPenLine}
                              size="17"
                            />
                          </button>
                        </div>
                      {/if}
                    </div>
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

  .report-container .actions {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 10px;
  }

  .report-container .actions .quartal-shifter {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
  }

  .report-container .actions .quartal-shifter .btn {
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

  .report-container .actions .quartal-shifter .btn:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }

  .report-container .actions .quartal-shifter .month-text {
    font-weight: 600;
    font-size: 18px;
    user-select: none;
  }

  .report-container .actions .filters {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 10px;
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
  }

  table tbody tr {
    border-bottom: 1px solid var(--gray-004);
  }

  table tbody tr td {
    padding: 8px 12px;
  }

  table tbody tr td .cells {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  table tbody tr td .cell {
    display: flex;
    flex-direction: column;
    gap: 1px;
    font-size: 12px;
    position: relative;
    padding: 2px 3px;
    border-radius: 4px;
    overflow: hidden;
    cursor: pointer;
  }

  table tbody tr td .cell .actions {
    display: none;
    flex-direction: row;
    justify-content: flex-end;
    align-items: flex-end;
    gap: 3px;
    position: absolute;
    right: 0;
    top: 0;
    padding: 0 5px 5px 0;
    background-color: var(--blue-transparent);
    width: 100%;
    height: 100%;
  }

  table tbody tr td .cell:hover .actions {
    display: flex;
  }

  table tbody tr td .cell .actions .btn {
    border: none;
    /* color: #fff; */
    background-color: transparent;
    width: fit-content;
    padding: 5px;
    border-radius: 8px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 3px;
    cursor: pointer;
  }

  table tbody tr td .cell .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

  :global(table tbody tr td .cell .actions .btn:hover svg) {
    fill: var(--blue-006);
  }

  table tbody tr td .date-warning {
    color: var(--yellow-006);
    padding: 2px 3px;
    border-radius: 4px;
    background-color: var(--yellow-transparent);
  }

  table tbody tr td .date-alert {
    color: var(--red-006);
    padding: 2px 3px;
    border-radius: 4px;
    background-color: var(--red-transparent);
  }

  table tbody tr td .cell.refunded {
    color: var(--violet-006) !important;
    padding: 4px 6px;
    border-radius: 4px;
    background-color: var(--violet-transparent);
  }

  .hidden {
    display: none !important;
  }
</style>