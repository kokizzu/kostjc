<script>
  /** @typedef {import('../../_types/property.js').Payment} Payment */
  /** @typedef {import('../../_types/property.js').PaymentPromise} PaymentPromise */
  /** @typedef {import('../../_types/property.js').BookingDetail} BookingDetail */
  /** @typedef {import('../../_types/property.js').Booking} Booking */

  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    RiFinanceWallet3Line, RiSystemEyeLine, RiDesignBallPenLine, RiSystemHistoryFill
  } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { GetRelativeDayLabel } from "../../_components/xGenerator";
  import {
    compactDate,
    paymentPromisePayload,
    promiseDayLabel,
    promiseStatus,
    updatePaymentPromises,
  } from '../../_helper/paymentPromise';
  import PopUpShowBookingPayments from '../../_components/PopUpShowBookingPayments.svelte';
  import { CmdForm, CmdUpsert } from '../../_components/xConstant';
  import { notifier } from '../../_components/xNotifier';
  import { AdminBooking, AdminPayment, StaffMissingDataReport, StaffPaymentPromise } from '../../jsApi.GEN';
  import { onMount } from 'svelte';
  import PopUpAddPayment from '../../_components/PopUpAddPayment.svelte';
  import PopUpAddPaymentPromise from '../../_components/PopUpAddPaymentPromise.svelte';
  import PopUpEditBooking from '../../_components/PopUpEditBooking.svelte';
  import PopUpShowBookingMissingTenant from '../../_components/PopUpShowBookingMissingTenant.svelte';

  const tenants = /** @type {Record<number, string>} */ ({/* tenants */});
  const rooms = /** @type {Record<number, string>} */ ({/* rooms */});
  let paymentPromises = /** @type {Record<number, PaymentPromise>} */ ({/* paymentPromises */});

  function refreshPage() {
    window.location.href = '/';
  }

  let isPopUpFormReady = false;
  onMount(() => {
    isPopUpFormReady = true;
  })

  /**
   * @typedef {Object} UnpaidBookingTenant
   * @property {number} bookingId
   * @property {number} tenantId
   * @property {string} tenantName
   * @property {string} roomName
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} dateStart
   * @property {string} dateEnd
   */
  const unpaidBookingTenants = /** @type {UnpaidBookingTenant[]} */ ([/* unpaidBookingTenants */]);

  const msPerDay = 1000 * 60 * 60 * 24;

  /**
   * @typedef {'green' | 'yellow' | 'red'} ProgressColor
   */

  /**
   * @typedef {Object} PaidProgress
   * @property {number} percentage
   * @property {ProgressColor} color
   * @property {number} daysLeft
   */

  /**
   * @param {UnpaidBookingTenant} data
   * @returns {PaidProgress}
   */
  function getOccupancyProgress(data) {
    const now = new Date();
    const dateStart = new Date(data.dateStart);
    const dateEnd = new Date(data.dateEnd);

     // @ts-ignore
    const daysOccupied = Math.floor((now - dateStart) / msPerDay);
    // @ts-ignore
    const totalDays = (dateEnd - dateStart) / msPerDay;

    let percentage = (daysOccupied / totalDays) * 100;
    if (daysOccupied <= 0) {
      percentage = 0;
    }

    const daysPaid = (data.totalPaid / data.totalPrice) * totalDays;
    const daysLeft = daysPaid - Math.max(daysOccupied, 0);

    let color = /** @type {ProgressColor} */ ('green');

    if (daysOccupied < daysPaid) {
      color = 'green'
    }
    
    if (daysOccupied >= (daysPaid - 3)) {
      color = 'yellow';
    }
    
    if (daysOccupied > daysPaid) {
      color = 'red';
    }

    return {
      percentage: Math.round(percentage),
      color: color,
      daysLeft: parseFloat(daysLeft.toFixed(1))
    };
  }

  /**
   * @description Get progress paid percentage, by total price, total paid
   * @param {UnpaidBookingTenant} data
   * @returns {{ percentage: number; daysPaid: number; }}
   */
  function getProgressPaidPercentage(data) {
    const dateStart = new Date(data.dateStart);
    const dateEnd = new Date(data.dateEnd);

    if (data.totalPaid >= data.totalPrice) {
      const now = new Date();
      // @ts-ignore
      const daysOccupied = Math.floor((now - dateStart) / msPerDay);

      return {
        percentage: 100,
        daysPaid: daysOccupied
      };
    }

    // @ts-ignore
    const totalDays = (dateEnd - dateStart) / msPerDay;
    
    const percentage = (data.totalPaid / data.totalPrice) * 100;
    const daysPaid = (data.totalPaid / data.totalPrice) * totalDays;

    return {
      percentage: Math.min(percentage, 100),
      daysPaid: parseFloat(daysPaid.toFixed(1))
    }
  }

  /**
   * @param {number} daysLeft
   * @returns {'left' | 'over'}
   */
  function getDaysLeftStatus(daysLeft) {
    return daysLeft < 0 ? 'over' : 'left';
  }

  let popUpShowPayments = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpShowBookingPayments | any} */ (null);
  let bookingIdToShowPayment = /** @type {number} */ (0);
  let paymentsForBooking = /** @type {Payment[]} */ ([]);
  let billingHistoryPopup = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpShowBookingMissingTenant | any} */ (null);
  let billingHistoryTenantName = '';
  let billingHistoryBookings = /** @type {import('../../_types/property').TenantBookingDetail[]} */ ([]);
  let billingHistoryPayments = /** @type {import('../../_types/property').PaymentOfBooking[]} */ ([]);
  let isBillingHistoryLoading = false;

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
        notifier.showError(o.error || 'failed to fetch payments');
        return
      }
      paymentsForBooking = o.paymentsByBooking;
      popUpShowPayments.Show();
    });
  }

  /**
   * @param {number} tenantId
   * @param {string} tenantName
   */
  async function showBillingHistory(tenantId, tenantName) {
    if (!tenantId) {
      notifier.showError('Cannot show billing history without tenant');
      return;
    }

    billingHistoryTenantName = tenantName;
    billingHistoryBookings = [];
    billingHistoryPayments = [];
    isBillingHistoryLoading = true;
    billingHistoryPopup.Show();

    await StaffMissingDataReport(
      { tenantId, cmd: CmdForm, yearMonth: '' },
      /** @type {import('../../jsApi.GEN').StaffMissingDataReportCallback} */
      /** @returns {Promise<void>} */
      function(/** @type {any} */ output) {
        isBillingHistoryLoading = false;
        if (output.error) {
          console.log(output);
          notifier.showError(output.error || 'Could not load billing history');
          return;
        }

        billingHistoryBookings = output.bookings || [];
        billingHistoryPayments = output.payments || [];
      },
    );
  }

  let popUpAddPayment = null;
  let bookingIdToPay = 0;
  let isSubmitAddPayment = false;
  let popUpAddPaymentPromise = null;
  let paymentPromiseTenantId = 0;
  let paymentPromiseBookingId = 0;
  let paymentPromiseTenantName = '';
  let isSubmitPaymentPromise = false;

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
          notifier.showError(o.error || 'failed to create payment');
          return
        }
        popUpAddPayment.Hide();
        refreshPage();
      }
    );
  }

  let popUpEditBooking = null;
  let bookingIdToEdit = 0;
  let isSubmitEditBooking = false;

  async function showEditBooking(bookingId) {
    bookingIdToEdit = bookingId;
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
        popUpEditBooking.Hide();
        refreshPage();
      }
    );
    isSubmitEditBooking = false;
  }

  function showInputPayment(/** @type {number} */ bookingId) {
    bookingIdToPay = bookingId;
    popUpAddPayment.Show();
  }

  /**
   * @param {number} tenantId
   * @returns {PaymentPromise | null}
   */
  function getPaymentPromise(tenantId) {
    return paymentPromises[tenantId] || paymentPromises[String(tenantId)] || null;
  }

  /**
   * @param {number} tenantId
   * @param {string} tenantName
   * @param {number} bookingId
   */
  function showPaymentPromise(tenantId, tenantName, bookingId) {
    if (!tenantId) {
      notifier.showError('Cannot add payment promise without tenant');
      return;
    }

    paymentPromiseTenantId = tenantId;
    paymentPromiseTenantName = tenantName;
    paymentPromiseBookingId = bookingId;
    popUpAddPaymentPromise.Show();
  }

  async function submitPaymentPromise(/** @type {PaymentPromise} */ promise) {
    isSubmitPaymentPromise = true;
    const payload = paymentPromisePayload(promise, paymentPromiseTenantId, paymentPromiseBookingId);

    await StaffPaymentPromise({
      cmd: CmdUpsert,
      promise: payload,
    }, /** @type {import('../../jsApi.GEN').StaffPaymentPromiseCallback} */
    /** @returns {Promise<void>} */
      function(/** @type any */ o) {
      isSubmitPaymentPromise = false;
      if (o.error) {
        console.log(o);
        notifier.showError(o.error || 'failed to save payment promise');
        return
      }

      paymentPromises = updatePaymentPromises(paymentPromises, o.promise);
      popUpAddPaymentPromise.Hide();
    });
  }

  /**
   * @param {string} roomName
   * @returns {boolean}
   */
  function isPurpleRoom(roomName) {
    return /^[ABC]/i.test((roomName || '').trim());
  }
</script>

{#if isPopUpFormReady}
  <PopUpShowBookingPayments
    bind:bookingId={bookingIdToShowPayment}
    bind:payments={paymentsForBooking}
    bind:this={popUpShowPayments}
  />

  <PopUpShowBookingMissingTenant
    bind:this={billingHistoryPopup}
    tenantName={billingHistoryTenantName}
    bookings={billingHistoryBookings}
    payments={billingHistoryPayments}
    isLoading={isBillingHistoryLoading}
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

  <PopUpAddPaymentPromise
    bind:this={popUpAddPaymentPromise}
    bind:isSubmitted={isSubmitPaymentPromise}
    tenantName={paymentPromiseTenantName}
    OnSubmit={submitPaymentPromise}
  />
{/if}

<section class="empty-unpaidBookingTenants">
  <h1>Unpaid Bookings</h1>
  {#if unpaidBookingTenants && unpaidBookingTenants.length > 0}
    <div class="cards">
      {#each (unpaidBookingTenants || []) as ub}
        {@const prog = getOccupancyProgress(ub)}
        {@const paidPercent = getProgressPaidPercentage(ub)}
        {@const promise = getPaymentPromise(ub.tenantId)}
        <div class="card" class:room-purple={isPurpleRoom(ub.roomName)} class:room-yellow={!isPurpleRoom(ub.roomName)}>
          <div class="heading">
            <h3>{ub.tenantName}</h3>
            <div class="actions">
              <button class="btn" title="Billing History" on:click={() => showBillingHistory(ub.tenantId, ub.tenantName)}>
                <Icon
                  src={RiSystemHistoryFill}
                  size="17"
                />
              </button>
              <button class="btn" title="Input Payment" on:click={() => showInputPayment(ub.bookingId)}>
                <Icon
                  src={RiFinanceWallet3Line}
                  size="17"
                />
              </button>
              <button class="btn" title="Show Payments" on:click={() => showPaymentsForBooking(ub.bookingId)}>
                <Icon
                  src={RiSystemEyeLine}
                  size="17"
                />
              </button>
              <button class="btn" title="Edit Booking" on:click={() => showEditBooking(ub.bookingId)}>
                <Icon
                  src={RiDesignBallPenLine}
                  size="17"
                />
              </button>
            </div>
          </div>
          <div class="detail">
            <hr />
            <div class="desc">
              <span>Room {ub.roomName}</span>
              <span>{ub.totalPaid}/{ub.totalPrice}</span>
              <span>Start at {ub.dateStart} ({GetRelativeDayLabel(ub.dateStart)})</span>
            </div>
            <hr />
            <div class="progress-container">
              <label for="">Paid <b>{paidPercent.daysPaid}</b> days</label>
              <div class="progress">
                <span class="blue" style="width: {paidPercent.percentage}%;"></span>
              </div>
            </div>
            <div class="progress-container">
              <label for="">Occupancy Progress <b>{prog.daysLeft}</b>d {getDaysLeftStatus(prog.daysLeft)}</label>
              <div class="progress">
                <span class={prog.color} style="width: {prog.percentage}%;"></span>
              </div>
            </div>
          </div>
          <div class="promise-area">
            <button class="promise-add" on:click={() => showPaymentPromise(ub.tenantId, ub.tenantName, ub.bookingId)}>Promise</button>
            {#if promise}
              <div class={`promise-chip ${promiseStatus(promise.promisePayAt)}`}>
                <span>Promised <b>{compactDate(promise.promisedAt)}</b> pay <b>{compactDate(promise.promisePayAt)}</b></span>
                <strong>{promiseDayLabel(promise.promisePayAt)}</strong>
              </div>
              {#if promise.note}
                <div class="promise-note">{promise.note}</div>
              {/if}
            {/if}
          </div>
        </div>
      {/each}
    </div>
  {:else}
    <div class="no-data">
      <p>No Unpaid Bookings</p>
    </div>
  {/if}
</section>

<style>
  .empty-unpaidBookingTenants {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .empty-unpaidBookingTenants h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .empty-unpaidBookingTenants .cards {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    gap: 10px;
  }

  .empty-unpaidBookingTenants .cards .card {
    background-color: var(--gray-001);
    padding: 10px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 6px;
    position: relative;
    overflow: hidden;
  }

  .empty-unpaidBookingTenants .cards .card.room-purple {
    background-color: #faf5ff;
  }

  .empty-unpaidBookingTenants .cards .card.room-yellow {
    background-color: #fffdf2;
  }

  .empty-unpaidBookingTenants .cards .card .heading {
    display: block;
    min-height: 56px;
  }

  .empty-unpaidBookingTenants .cards .card hr {
    height: 1px;
    background-color: var(--gray-004) !important;
    border: none;
    margin: 0;
  }

  .empty-unpaidBookingTenants .cards .card .detail {
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .empty-unpaidBookingTenants .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
    max-width: calc(100% - 64px);
  }

  .empty-unpaidBookingTenants .cards .card .desc {
    display: flex;
    flex-direction: column;
    gap: 2px;
    z-index: 20;
  }

  .empty-unpaidBookingTenants .cards .card .actions {
    display: grid;
    grid-template-columns: repeat(2, 27px);
    grid-template-rows: repeat(2, 27px);
    justify-content: center;
    align-items: center;
    gap: 3px;
    position: absolute;
    top: 8px;
    right: 8px;
    width: 57px;
    height: 57px;
    z-index: 30;
  }

  .empty-unpaidBookingTenants .cards .card .actions .btn {
    border: none;
    /* color: #fff; */
    background-color: transparent;
    width: 27px;
    height: 27px;
    padding: 2px;
    border-radius: 8px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 3px;
    cursor: pointer;
  }

  .empty-unpaidBookingTenants .cards .card .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

  :global(.empty-unpaidBookingTenants .cards .card .actions .btn:hover svg) {
    fill: var(--blue-006);
  }

  .promise-area {
    display: flex;
    flex-direction: column;
    gap: 3px;
    margin-top: 2px;
    z-index: 20;
  }

  .promise-add {
    border: 1px solid #fca5a5;
    background-color: #fee2e2;
    color: #991b1b;
    border-radius: 6px;
    cursor: pointer;
    font-size: 12px;
    font-weight: 800;
    line-height: 1;
    padding: 4px;
    text-align: center;
  }

  .promise-add:hover,
  .promise-add:focus-visible {
    outline: none;
    background-color: #fecaca;
  }

  .promise-chip {
    border-radius: 6px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 4px;
    font-size: 12px;
    font-weight: 700;
    line-height: 1.2;
    padding: 3px 4px;
  }

  .promise-chip.left {
    background-color: #dbeafe;
    color: #1e40af;
  }

  .promise-chip.over {
    background-color: #fee2e2;
    color: #991b1b;
  }

  .promise-chip strong {
    flex: 0 0 auto;
    font-weight: 900;
  }

  .promise-note {
    color: var(--gray-007);
    font-size: 12px;
    font-weight: 600;
    line-height: 1.2;
    overflow-wrap: anywhere;
  }

  @media only screen and (max-width : 768px) {
    .empty-unpaidBookingTenants .cards {
      grid-template-columns: 1fr;
    }
  }

  @media only screen and (max-width: 992px) {
    .empty-unpaidBookingTenants .cards {
      grid-template-columns: 1fr 1fr;
    }
  }

  .progress-container {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .progress-container label {
    font-size: 12px;
    margin-left: 5px;
  }

  .progress {
    position: relative;
    width: 100%;
    height: 14px;
    background-color: var(--gray-003);
    border-radius: 9999px;
    overflow: hidden;
  }

  .progress span {
    height: 100%;
    display: block;
  }

  :global(.progress span.green) {
    background-image: linear-gradient(var(--green-004), var(--green-006));
  }

  :global(.progress span.yellow) {
    background-image: linear-gradient(var(--yellow-004), var(--yellow-006));
  }

  :global(.progress span.red) {
    background-image: linear-gradient(var(--red-004), var(--red-006));
  }

  :global(.progress span.blue) {
    background-image: linear-gradient(var(--blue-004), var(--blue-006));
  }
</style>
