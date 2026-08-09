<script>
  /** @typedef {import('../../_types/property.js').PaymentPromise} PaymentPromise */

  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiSystemHistoryFill } from '../../node_modules/svelte-icons-pack/dist/ri';
  import {
    compactDate,
    paymentPromisePayload,
    promiseDayLabel,
    promiseStatus,
    updatePaymentPromises,
  } from '../../_helper/paymentPromise';
  import PopUpAddPaymentPromise from '../../_components/PopUpAddPaymentPromise.svelte';
  import PopUpShowBookingMissingTenant from '../../_components/PopUpShowBookingMissingTenant.svelte';
  import { StaffMissingDataReport, StaffPaymentPromise } from '../../jsApi.GEN';
  import { CmdForm, CmdUpsert } from '../../_components/xConstant';
  import { notifier } from '../../_components/xNotifier';

  /**
   * @typedef {Object} AvailableRoom
   * @property {string} roomName
   * @property {string} availableAt
   * @property {boolean} isAvailableNow
   * @property {number} currentTenantId
   * @property {string} lastTenant
   * @property {number} basePriceIDR
   * @property {number} totalPriceIDR
   */
  const rooms = /** @type {AvailableRoom[]} */ ([/* availableRooms */]);
  /**
   * @typedef {Object} UnpaidBookingTenant
   * @property {number} bookingId
   * @property {number} tenantId
   * @property {string} tenantName
   * @property {string} roomName
   */
  const unpaidBookingTenants = /** @type {UnpaidBookingTenant[]} */ ([/* unpaidBookingTenants */]);
  let paymentPromises = /** @type {Record<number, PaymentPromise>} */ ({/* paymentPromises */});
  let sortedRooms = /** @type {AvailableRoom[]} */ ([]);
  let sortBy = /** @type {'default' | 'basePrice' | 'totalPrice'} */ ('default');
  let billingHistoryPopup = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpShowBookingMissingTenant | any} */ (null);
  let billingHistoryTenantName = '';
  let billingHistoryBookings = /** @type {import('../../_types/property').TenantBookingDetail[]} */ ([]);
  let billingHistoryPayments = /** @type {import('../../_types/property').PaymentOfBooking[]} */ ([]);
  let isBillingHistoryLoading = false;
  let popUpAddPaymentPromise = null;
  let paymentPromiseTenantId = 0;
  let paymentPromiseBookingId = 0;
  let paymentPromiseTenantName = '';
  let isSubmitPaymentPromise = false;

  function sortRooms() {
    sortedRooms = [...rooms];
    
    switch(sortBy) {
      case 'basePrice':
        sortedRooms.sort((a, b) => b.basePriceIDR - a.basePriceIDR);
        break;
      case 'totalPrice':
        sortedRooms.sort((a, b) => b.totalPriceIDR - a.totalPriceIDR);
        break;
      default:
        sortedRooms.sort((a, b) => {
          if (!a.availableAt) return 1;
          if (!b.availableAt) return -1;
          return new Date(a.availableAt).getTime() - new Date(b.availableAt).getTime();
        });
    }
  }

  $: if (rooms) {
    sortRooms();
  }

  $: if (sortBy) {
    sortRooms();
  }

  function formatDateLong(/** @type {string} */ dateStr, /** @type {number} */ dayTo = 0) {
    const dt = new Date(dateStr);
    dt.setDate(dt.getDate() + dayTo);
    return dt.toLocaleDateString('en-GB', {
      weekday: 'long',
      day: '2-digit',
      month: 'long',
      year: 'numeric'
    });
  }

  function getRelativeDayLabel(/** @type {string} */ dateStr, /** @type {number} */ dayTo = 0) {
    const inputDate = new Date(dateStr);
    inputDate.setDate(inputDate.getDate() + dayTo);
    const currentDate = new Date();
    
    inputDate.setHours(0, 0, 0, 0);
    currentDate.setHours(0, 0, 0, 0);

    const msInDay = 1000 * 60 * 60 * 24; // @ts-ignore
    const diffDays = Math.round((currentDate - inputDate) / msInDay);

    if (diffDays === 0) return 'H';
    if (diffDays > 0) return `H+${diffDays}`;
    
    return `H${diffDays}`;
  }

  /**
   * @param {string} availableAt
   * @returns {string}
   */ 
  function getAvailableDays(availableAt) {
    if (!availableAt) return '';

    const availableDate = new Date(availableAt);
    const now = new Date();

    // @ts-ignore
    const diffMs = now - availableDate;
    const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

    if (diffDays <= 0) return '(today)';

    return `(-<b>${diffDays}</b>d ago)`;
  }

  /**
   * Format number to IDR currency
   * @param {number} amount
   * @returns {string}
   */
  function formatIDR(amount) {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
    }).format(amount);
  }

  /**
   * @param {string} roomName
   * @returns {boolean}
   */
  function isPurpleRoom(roomName) {
    return /^[ABC]/i.test((roomName || '').trim());
  }

  /**
   * @param {AvailableRoom} room
   * @returns {boolean}
   */
  function findUnpaidBooking(room) {
    const tenantId = Number(room.currentTenantId || 0);
    if (tenantId > 0) {
      return (unpaidBookingTenants || []).find(item => Number(item.tenantId || 0) === tenantId) || null;
    }

    const tenantName = (room.lastTenant || '').trim().toLowerCase();
    if (!tenantName) return null;

    return (unpaidBookingTenants || []).find(item => {
      return Number(item.tenantId || 0) === 0 && (item.tenantName || '').trim().toLowerCase() === tenantName;
    }) || null;
  }

  /**
   * @param {AvailableRoom} room
   * @returns {boolean}
   */
  function hasUnpaidBooking(room) {
    return Boolean(findUnpaidBooking(room));
  }

  /**
   * @param {AvailableRoom} room
   * @returns {boolean}
   */
  function canShowBillingHistory(room) {
    return Number(room.currentTenantId || 0) > 0;
  }

  /**
   * @param {AvailableRoom} room
   */
  async function showBillingHistory(room) {
    const tenantId = Number(room.currentTenantId || 0);
    if (!tenantId) {
      notifier.showError('Cannot show billing history without tenant');
      return;
    }

    billingHistoryTenantName = room.lastTenant || `Tenant #${tenantId}`;
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

  /**
   * @param {number} tenantId
   * @returns {PaymentPromise | null}
   */
  function getPaymentPromise(tenantId) {
    return paymentPromises[tenantId] || paymentPromises[String(tenantId)] || null;
  }

  /**
   * @param {AvailableRoom} room
   */
  function showPaymentPromise(room) {
    const tenantId = Number(room.currentTenantId || 0);
    if (!tenantId) {
      notifier.showError('Cannot add payment promise without tenant');
      return;
    }

    paymentPromiseTenantId = tenantId;
    paymentPromiseTenantName = room.lastTenant || `Tenant #${tenantId}`;
    paymentPromiseBookingId = Number((findUnpaidBooking(room) || {}).bookingId || 0);
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
</script>

<PopUpShowBookingMissingTenant
  bind:this={billingHistoryPopup}
  tenantName={billingHistoryTenantName}
  bookings={billingHistoryBookings}
  payments={billingHistoryPayments}
  isLoading={isBillingHistoryLoading}
/>

<PopUpAddPaymentPromise
  bind:this={popUpAddPaymentPromise}
  bind:isSubmitted={isSubmitPaymentPromise}
  tenantName={paymentPromiseTenantName}
  OnSubmit={submitPaymentPromise}
/>

<section class="empty-rooms">
  <div class="header">
    <h1>Available Rooms</h1>
    <div class="sort-controls">
      <p>Sort by:</p>
      <div class="sort-buttons">
        <button 
          class="sort-btn" 
          class:active={sortBy === 'default'}
          on:click={() => sortBy = 'default'}
        >
          Available Date
        </button>
        <button 
          class="sort-btn" 
          class:active={sortBy === 'basePrice'}
          on:click={() => sortBy = 'basePrice'}
        >
          Base Price
        </button>
        <button 
          class="sort-btn" 
          class:active={sortBy === 'totalPrice'}
          on:click={() => sortBy = 'totalPrice'}
        >
          Total Price
        </button>
      </div>
    </div>
  </div>

  {#if sortedRooms && sortedRooms.length > 0}
    <div class="cards">
      {#each (sortedRooms || []) as r}
        {@const promise = getPaymentPromise(r.currentTenantId)}
        <div class="card" class:room-purple={isPurpleRoom(r.roomName)} class:room-yellow={!isPurpleRoom(r.roomName)}>
          <div class="room-actions">
            {#if canShowBillingHistory(r)}
              <button
                class="billing-history-button"
                title={`Show bills and payments for ${r.lastTenant || `Tenant #${r.currentTenantId}`}`}
                aria-label={`Show bills and payments for ${r.lastTenant || `Tenant #${r.currentTenantId}`}`}
                on:click={() => showBillingHistory(r)}
              >
                <Icon size="16" src={RiSystemHistoryFill} />
              </button>
            {/if}
            {#if hasUnpaidBooking(r)}
              <span class="warning-badge" title="still have outstanding payment" aria-label="still have outstanding payment">⚠</span>
            {/if}
          </div>
          <h3>Room {r.roomName}</h3>
          <div class="desc">
            <span>{@html r.isAvailableNow || r.availableAt == ''
              ? `Available Now ${getAvailableDays(r.availableAt)}`
              : 'Available on <b>' + formatDateLong(r.availableAt, 1)+' ('+ getRelativeDayLabel(r.availableAt, 1) +') </b>'
            }</span>
            <span>Last Tenant: <b>{r.lastTenant || '--'}</b></span>
            <div class="price-info">
              <span class="price-label">Base Price:</span>
              <span class="price-value">{formatIDR(r.basePriceIDR || 0)}</span>
            </div>
            <div class="price-info">
              <span class="price-label">Total Price:</span>
              <span class="price-value total">{formatIDR(r.totalPriceIDR || 0)}</span>
            </div>
          </div>
          <div class="promise-area">
            {#if canShowBillingHistory(r)}
              <button class="promise-add" on:click={() => showPaymentPromise(r)}>Promise</button>
            {/if}
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
      <p>No Available Rooms</p>
    </div>
  {/if}
</section>

<style>
  .empty-rooms {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    gap: 16px;
  }

  .empty-rooms h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .sort-controls {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
  }

  .sort-controls p {
    font-weight: 600;
    font-size: 14px;
    color: var(--gray-007, #333);
    margin: 0;
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

  .empty-rooms .cards {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    gap: 10px;
  }

  .empty-rooms .cards .card {
    background-color: var(--gray-001);
    padding: 10px;
    padding-top: 18px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    gap: 6px;
    position: relative;
    overflow: hidden;
  }

  .room-actions {
    position: absolute;
    top: 8px;
    right: 8px;
    z-index: 30;
    display: grid;
    grid-template-columns: repeat(2, 24px);
    grid-auto-rows: 24px;
    gap: 3px;
    justify-content: end;
  }

  .billing-history-button {
    width: 24px;
    height: 24px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 0;
    border: none;
    border-radius: 7px;
    background-color: #fff;
    color: var(--gray-007);
    cursor: pointer;
    box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.08);
  }

  .billing-history-button:hover,
  .billing-history-button:focus-visible {
    outline: none;
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }

  :global(.billing-history-button:hover svg),
  :global(.billing-history-button:focus-visible svg) {
    fill: var(--blue-006);
  }

  .warning-badge {
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 9999px;
    background: #fff7ed;
    color: #c2410c;
    font-size: 14px;
    line-height: 1;
    box-shadow: 0 0 0 1px rgba(194, 65, 12, 0.16);
  }

  .empty-rooms .cards .card.room-purple {
    background-color: #faf5ff;
  }

  .empty-rooms .cards .card.room-yellow {
    background-color: #fffdf2;
  }

  .empty-rooms .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-rooms .cards .card .desc {
    padding-top: 4px;
    border-top: 1px solid var(--gray-004);
    display: flex;
    flex-direction: column;
    flex: 1 1 auto;
    gap: 4px;
    z-index: 20;
  }

  .price-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1px 0;
  }

  .price-label {
    font-size: 13px;
    color: var(--gray-006, #666);
  }

  .price-value {
    font-weight: 600;
    font-size: 14px;
    color: var(--gray-008, #222);
  }

  .price-value.total {
    color: var(--primary-color, #4CAF50);
    font-size: 15px;
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

  .no-data {
    text-align: center;
    padding: 40px;
    color: var(--gray-006, #666);
  }

  @media only screen and (max-width: 1200px) {
    .empty-rooms .cards {
      grid-template-columns: repeat(3, 1fr);
    }
  }

  @media only screen and (max-width: 768px) {
    .header {
      flex-direction: column;
      align-items: flex-start;
    }

    .sort-controls {
      width: 100%;
      flex-direction: column;
      align-items: flex-start;
      gap: 8px;
    }

    .sort-buttons {
      width: 100%;
    }

    .sort-btn {
      flex: 1;
      text-align: center;
    }

    .empty-rooms .cards {
      grid-template-columns: 1fr;
    }
  }
</style>
