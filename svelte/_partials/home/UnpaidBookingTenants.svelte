<script>
  /** @typedef {import('../../_types/property.js').Payment} Payment */

  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { AiOutlineEye } from '../../node_modules/svelte-icons-pack/dist/ai';
  import { GetRelativeDayLabel } from "../../_components/xGenerator";
  import PopUpShowBookingPayments from '../../_components/PopUpShowBookingPayments.svelte';
  import { CmdForm } from '../../_components/xConstant';
  import { notifier } from '../../_components/xNotifier';
  import { AdminPayment } from '../../jsApi.GEN';

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
  const unpaidBookingTenants = /** @type {UnpaidBookingTenant[]} */ ([/* unpaidBookingTenants */]);

  const msPerDay = 1000 * 60 * 60 * 24;

  /**
   * @typedef {'green' | 'yellow' | 'red'} ProgressColor
   */

  /**
   * @typedef {Object} PaidProgress
   * @property {number} percentage
   * @property {ProgressColor} color
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
      color: color
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
</script>

<PopUpShowBookingPayments
  bind:bookingId={bookingIdToShowPayment}
  bind:payments={paymentsForBooking}
  bind:this={popUpShowPayments}
/>

<section class="empty-unpaidBookingTenants">
  <h1>Unpaid Bookings</h1>
  {#if unpaidBookingTenants && unpaidBookingTenants.length > 0}
    <div class="cards">
      {#each (unpaidBookingTenants || []) as ub}
        {@const prog = getOccupancyProgress(ub)}
        {@const paidPercent = getProgressPaidPercentage(ub)}
        <div class="card">
          <h3>{ub.tenantName}</h3>
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
              <label for="">Occupancy Progress</label>
              <div class="progress">
                <span class={prog.color} style="width: {prog.percentage}%;"></span>
              </div>
            </div>
            <div class="actions">
              <button class="btn" title="Show Payments" on:click={() => showPaymentsForBooking(ub.bookingId)}>
                <Icon
                  src={AiOutlineEye}
                  color="#FFF"
                  size="17"
                />
                <span>Show Payments</span>
              </button>
            </div>
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
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .empty-unpaidBookingTenants .cards .card hr {
    height: 1px;
    background-color: var(--gray-004) !important;
    border: none;
  }

  .empty-unpaidBookingTenants .cards .card .detail {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .empty-unpaidBookingTenants .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-unpaidBookingTenants .cards .card .desc {
    display: flex;
    flex-direction: column;
    gap: 5px;
    z-index: 20;
  }

  .empty-unpaidBookingTenants .cards .card .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
  }

  .empty-unpaidBookingTenants .cards .card .actions .btn {
    background-color: var(--blue-006);
    color: #FFF;
    padding: 5px 12px;
    border: none;
    border-radius: 9999px;
    display: flex;
    flex-direction: row;
    gap: 5px;
    align-items: center;
    cursor: pointer;
  }

  .empty-unpaidBookingTenants .cards .card .actions .btn:hover {
    background-color: var(--blue-005);
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
    gap: 5px;
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