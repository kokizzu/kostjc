<script>
  /**
   * @typedef {Object} UnpaidBookingTenant
   * @property {string} tenantName
   * @property {string} roomName
   * @property {number} totalPaid
   * @property {number} totalPrice
   * @property {string} dateStart
   */
   const unpaidBookingTenants = /** @type {UnpaidBookingTenant[]} */ ([/* unpaidBookingTenants */]);

  /**
   * @description Returns the relative day label based on the given date string (YYYY-MM-DD)
   * @param {string} dateStr - The date string in format YYYY-MM-DD
   * @returns {string} - The relative label like 'H', 'H-1', 'H+3', etc.
   */
  function getRelativeDayLabel(dateStr) {
    const inputDate = new Date(dateStr);
    const currentDate = new Date();
    
    inputDate.setHours(0, 0, 0, 0);
    currentDate.setHours(0, 0, 0, 0);

    const msInDay = 1000 * 60 * 60 * 24; // @ts-ignore
    const diffDays = Math.round((currentDate - inputDate) / msInDay);

    if (diffDays === 0) return 'H';
    if (diffDays > 0) return `H+${diffDays}`;
    
    return `H${diffDays}`;
  }
</script>

<section class="empty-unpaidBookingTenants">
  <h1>Unpaid Bookings</h1>
  {#if unpaidBookingTenants && unpaidBookingTenants.length > 0}
    <div class="cards">
      {#each (unpaidBookingTenants || []) as ub}
        <div class="card">
          <h3>{ub.tenantName}</h3>
          <div class="desc">
            <span>Room {ub.roomName}</span>
            <span>{ub.totalPaid}/{ub.totalPrice}</span>
            <span>{ub.dateStart} ({getRelativeDayLabel(ub.dateStart)})</span>
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
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .empty-unpaidBookingTenants .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-unpaidBookingTenants .cards .card .desc {
    padding-top: 10px;
    border-top: 1px solid var(--gray-004);
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
</style>