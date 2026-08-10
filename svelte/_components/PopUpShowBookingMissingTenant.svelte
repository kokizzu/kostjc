<script>
  /** @typedef {import('../_types/property').PaymentOfBooking} PaymentOfBooking */
  /** @typedef {import('../_types/property').TenantBookingDetail} TenantBookingDetail */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import {
    compactIDRFromThousands,
    compactPaymentIDRFromThousands,
    paymentTiming,
  } from '../_helper/billingHistory';

  export let tenantName = '';
  export let bookings = /** @type {TenantBookingDetail[]} */ ([]);
  export let payments = /** @type {PaymentOfBooking[]} */ ([]);
  export let isLoading = false;

  let isShow = false;

  export const Show = () => (isShow = true);
  export const Hide = () => (isShow = false);

  /**
   * @param {number} bookingId
   * @returns {PaymentOfBooking[]}
   */
  function getPaymentsByBookingId(bookingId) {
    return payments
      .filter(payment => payment.bookingId == bookingId)
      .sort((left, right) => {
        const dateOrder = (left.paymentAt || '').localeCompare(right.paymentAt || '');
        return dateOrder || left.paymentId - right.paymentId;
      });
  }

  /**
   * @param {TenantBookingDetail} booking
   * @returns {number}
   */
  function debtAmount(booking) {
    return Math.max(0, (booking.totalPriceIDR || 0) - (booking.totalPaidIDR || 0));
  }

  /**
   * @param {string} paymentAt
   * @param {string} bookingStart
   * @returns {string}
   */
  function paymentTimingClass(paymentAt, bookingStart) {
    const timing = paymentTiming(paymentAt, bookingStart);
    if (timing.includes('late')) return 'late';
    if (timing.includes('early')) return 'early';
    return 'ontime';
  }

  /** @param {KeyboardEvent} event */
  function onKeydown(event) {
    if (isShow && event.key === 'Escape') Hide();
  }
</script>

<svelte:window on:keydown={onKeydown} />

<div class={`popup-container ${isShow ? 'show' : ''}`} role="presentation" on:click|self={Hide}>
  <div class="popup" role="dialog" aria-modal="true" aria-label={`Billing and payment history for ${tenantName}`} tabindex="-1">
    <header class="header">
      <div class="title">
        <h2>Billing &amp; Payment History</h2>
      </div>
      <div class="header-right">
        <p>{tenantName}</p>
        <button class="close-button" title="Close" aria-label="Close billing and payment history" on:click={Hide}>
          <Icon size="22" color="var(--red-005)" src={IoClose} />
        </button>
      </div>
    </header>

    <div class="bookings">
      {#if isLoading}
        <div class="message" aria-live="polite">Loading billing history…</div>
      {:else if bookings.length === 0}
        <div class="message">No bookings found for this tenant.</div>
      {:else}
        {#each bookings as booking (booking.bookingId)}
          <article class="booking">
            <section class="booking-summary" aria-label={`Booking ${booking.bookingId}`}>
              <h3>Booking #{booking.bookingId} <span>Room <strong>{booking.roomName || '—'}</strong></span></h3>
              <p>{booking.dateStart} to {booking.dateEnd}</p>
              <p class="amounts">
                <strong>Bill</strong> {compactIDRFromThousands(booking.totalPriceIDR)}
                <strong>Paid</strong> {compactIDRFromThousands(booking.totalPaidIDR)}
                {#if debtAmount(booking) > 0}
                  <span class="debt"><strong>Debt</strong> {compactIDRFromThousands(debtAmount(booking))}</span>
                {/if}
              </p>
            </section>

            <section class="payments" aria-label={`Payments for booking ${booking.bookingId}`}>
              {#each getPaymentsByBookingId(booking.bookingId) as payment (payment.paymentId)}
                <div class="payment">
                  <p class="payment-main">
                    <strong>Payment #{payment.paymentId}</strong>
                    {#if payment.paymentMethod}
                      <span>{payment.paymentMethod}</span>
                    {/if}
                    <strong class="payment-amount">{compactPaymentIDRFromThousands(payment.paidIDR)}</strong>
                  </p>
                  <div class="payment-meta">
                    <span>{payment.paymentAt}</span>
                    <strong class={`timing ${paymentTimingClass(payment.paymentAt, booking.dateStart)}`}>
                      {paymentTiming(payment.paymentAt, booking.dateStart)}
                    </strong>
                  </div>
                  {#if payment.note}
                    <p class="payment-note">{payment.note}</p>
                  {/if}
                </div>
              {:else}
                <p class="no-payments">No payments associated with this booking.</p>
              {/each}
            </section>
          </article>
        {/each}
      {/if}
    </div>
  </div>
</div>

<style>
  .popup-container {
    display: none;
    position: fixed;
    inset: 0;
    z-index: 2000;
    overflow: auto;
    align-items: flex-start;
    justify-content: center;
    padding: 2px;
    background-color: rgba(0 0 0 / 40%);
    backdrop-filter: blur(1px);
  }

  .popup-container.show {
    display: flex;
  }

  .popup {
    width: fit-content;
    max-width: calc(100vw - 4px);
    height: fit-content;
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    background-color: #fff;
    box-shadow: 0 18px 50px rgba(0 0 0 / 20%);
  }

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 4px;
    padding: 4px;
    border-bottom: 1px solid var(--gray-004);
  }

  .title,
  .header-right {
    display: flex;
    align-items: center;
    gap: 4px;
    min-width: 0;
  }

  .header-right {
    justify-content: flex-end;
  }

  .header h2,
  .header p {
    margin: 0;
  }

  .header h2 {
    font-size: var(--font-lg);
  }

  .header p {
    color: var(--gray-007);
    font-weight: 600;
    white-space: nowrap;
  }

  .close-button {
    display: flex;
    align-items: center;
    justify-content: center;
    flex: 0 0 auto;
    padding: 2px;
    border: none;
    border-radius: 50%;
    background-color: transparent;
    cursor: pointer;
  }

  .close-button:hover {
    background-color: var(--red-transparent);
  }

  .bookings {
    min-height: 120px;
    padding: 2px 4px;
    width: fit-content;
    max-width: 100%;
    display: flex;
    flex-direction: column;
    gap: 0;
  }

  .message {
    padding: 4px;
    border-radius: 8px;
    background-color: var(--gray-001);
    color: var(--gray-007);
    text-align: center;
  }

  .booking {
    display: grid;
    grid-template-columns: minmax(260px, max-content) minmax(360px, max-content);
    gap: 4px;
    padding: 3px 0;
    border-bottom: 1px solid var(--gray-003);
  }

  .booking:last-child {
    border-bottom: none;
  }

  .booking-summary {
    min-width: 0;
  }

  .booking-summary h3,
  .booking-summary p,
  .payment-main,
  .payment-note,
  .no-payments {
    margin: 0;
  }

  .booking-summary h3 {
    display: flex;
    flex-wrap: nowrap;
    gap: 4px;
    align-items: baseline;
    font-size: 14px;
    white-space: nowrap;
  }

  .booking-summary h3 span {
    color: var(--gray-007);
    font-size: 12px;
    font-weight: 400;
  }

  .booking-summary p,
  .payment-meta,
  .payment-note,
  .no-payments {
    color: var(--gray-007);
    font-size: 12px;
    line-height: 1.35;
  }

  .amounts {
    display: flex;
    flex-wrap: wrap;
    gap: 3px 6px;
  }

  .debt {
    color: var(--red-006);
    font-weight: 700;
  }

  .payment {
    display: grid;
    grid-template-columns: minmax(220px, 1fr) max-content;
    column-gap: 8px;
    padding: 0 0 3px;
    border-bottom: 1px solid var(--gray-002);
  }

  .payment:last-child {
    padding-bottom: 0;
    border-bottom: none;
  }

  .payment-main {
    display: flex;
    flex-wrap: wrap;
    gap: 3px 6px;
    align-items: baseline;
  }

  .payment-main span {
    color: var(--gray-008);
  }

  .payment-meta {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-end;
    gap: 3px 6px;
    text-align: right;
    white-space: nowrap;
  }

  .payment-note {
    grid-column: 1 / -1;
  }

  .payment-amount,
  .timing {
    font-weight: 700;
  }

  .timing {
    padding: 1px 3px;
    border-radius: 4px;
  }

  .timing.late {
    color: var(--red-006);
    background-color: var(--red-transparent);
  }

  .timing.early {
    color: var(--blue-006);
    background-color: var(--blue-transparent);
  }

  .timing.ontime {
    color: var(--green-007);
    background-color: var(--green-transparent);
  }

  .no-payments {
    padding: 0;
  }

  @media only screen and (max-width: 768px) {
    .popup-container {
      padding: 2px;
    }

    .booking {
      grid-template-columns: 1fr;
    }

    .payment {
      grid-template-columns: 1fr;
    }

    .payment-meta {
      justify-content: flex-start;
      text-align: left;
    }
  }
</style>
