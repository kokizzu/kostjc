<script>
  import { onMount } from "svelte";
    import { GetRelativeDayLabel } from "../../_components/xGenerator";

  /**
   * @typedef {Object} UnpaidBookingTenant
   * @property {string} tenantName
   * @property {string} roomName
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} dateStart
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
   * @description Get progress paid percentage, by total price, total paid, and days occupied
   * @param {UnpaidBookingTenant} data
   * @returns {PaidProgress}
   */
  function getProgressOccupiedPercentage(data) {
    const now = new Date();
    const startDate = new Date(data.dateStart);
    // @ts-ignore
    const daysOccupied = Math.floor((now - startDate) / msPerDay);

    let percentPaid = (data.totalPaid / data.totalPrice) * 100;
    percentPaid = Math.min(percentPaid, 100);

    if (daysOccupied > 30 && data.totalPaid < data.totalPrice) {
      percentPaid = 100;
    }

    const dayFactor = Math.min((daysOccupied / 30) * 100, 100);

    // 70% weight to payment
    // 30% weight to days occupied
    let percentage = (percentPaid * 0.7) + (dayFactor * 0.3);
    percentage = Math.min(percentage, 100);

    let color = /** @type {ProgressColor} */ ('green');

    if (percentage < 30) color = 'green';
    else if (percentage < 70) color = 'yellow';
    else color = 'red';

    if (percentage == 0) color = 'green';
    if (percentage < 5) percentage = 5;

    return {
      percentage: Math.round(percentage),
      color: color
    };
  }

  /**
   * @description Get progress paid percentage, by total price, total paid
   * @param {UnpaidBookingTenant} data
   * @returns {number} percentage
   */
  function getProgressPaidPercentage(data) {
    if (data.totalPaid >= data.totalPrice) return 100;
    
    const percentage = (data.totalPaid / data.totalPrice) * 100;
    return Math.min(percentage, 100);
  }
</script>

<section class="empty-unpaidBookingTenants">
  <h1>Unpaid Bookings</h1>
  {#if unpaidBookingTenants && unpaidBookingTenants.length > 0}
    <div class="cards">
      {#each (unpaidBookingTenants || []) as ub}
        {@const prog = getProgressOccupiedPercentage(ub)}
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
              <label for="">Paid progress</label>
              <div class="progress">
                <span class="blue" style="width: {paidPercent}%;"></span>
              </div>
            </div>
            <div class="progress-container">
              <label for="">Occupancy progress</label>
              <div class="progress">
                <span class={prog.color} style="width: {prog.percentage}%;"></span>
              </div>
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