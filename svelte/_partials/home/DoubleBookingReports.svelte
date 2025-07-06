<script>
  /**
   * @typedef {Object} DoubleBookingReportData
   * @property {string|number} tenantId
   * @property {string} tenantName
   * @property {string} dateStart
   * @property {string} dateEnd
   */

  /**
   * @typedef {Object} DoubleBookingReport
   * @property {string|number} roomId
   * @property {string} roomName
   * @property {DoubleBookingReportData[]} tenants
   */

  const doubleBookingReports = /** @type {DoubleBookingReport[]} */ ([/* doubleBookingReports */]);
</script>

<section class="empty-doubleBookingReports">
  <h1>Double Bookings</h1>
  {#if doubleBookingReports && doubleBookingReports.length > 0}
    <div class="cards">
      {#each (doubleBookingReports || []) as db}
        <div class="card">
          <h3>Room {db.roomName}</h3>
          <div class="desc">
            {#each (db.tenants || []) as bk}
              <span>{bk.tenantName} ({bk.dateStart} s/d {bk.dateEnd})</span>
            {/each}
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
  .empty-doubleBookingReports {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .empty-doubleBookingReports h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .empty-doubleBookingReports .cards {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
  }

  .empty-doubleBookingReports .cards .card {
    background-color: var(--gray-001);
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .empty-doubleBookingReports .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-doubleBookingReports .cards .card .desc {
    padding-top: 10px;
    border-top: 1px solid var(--gray-004);
    display: flex;
    flex-direction: column;
    gap: 5px;
    z-index: 20;
    font-size: 12px;
  }

  @media only screen and (max-width : 768px) {
    .empty-doubleBookingReports .cards {
      grid-template-columns: 1fr;
    }
  }

  @media only screen and (max-width: 992px) {
    .empty-doubleBookingReports .cards {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>