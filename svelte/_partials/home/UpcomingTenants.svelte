<script>
  /**
   * @typedef {Object} UpcomingTenant
   * @property {string} prevTenant
   * @property {string} nextTenant
   * @property {string} roomName
   * @property {string} dateStart
   * @property {string} dateEnd
   */
   const upcomingTenants = /** @type {UpcomingTenant[]} */ ([/* upcomingTenants */]);
</script>

<section class="empty-upcomingTenants">
  <h1>Upcoming Tenants</h1>
  {#if upcomingTenants && upcomingTenants.length > 0}
    <div class="cards">
      {#each (upcomingTenants || []) as ub}
        <div class="card">
          <div class="tenant">
            <span>Previous Tenant:</span>
            <h3>{ub.prevTenant}</h3>
          </div>
          <div class="tenant">
            <span>Next Tenant:</span>
            <h3>{ub.nextTenant}</h3>
          </div>
          <div class="desc">
            <span>Room {ub.roomName}</span>
            <span>{ub.dateStart} s/d {ub.dateEnd} </span>
          </div>
        </div>
      {/each}
    </div>
  {:else}
    <div class="no-data">
      <p>No Upcoming Tenants</p>
    </div>
  {/if}
</section>

<style>
  .empty-upcomingTenants {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .empty-upcomingTenants h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .empty-upcomingTenants .cards {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    gap: 10px;
  }

  .empty-upcomingTenants .cards .card {
    background-color: var(--gray-001);
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .empty-upcomingTenants .cards .card .tenant {
    display: flex;
    flex-direction: column;
  }

  .empty-upcomingTenants .cards .card .tenant span {
    font-size: 12px;
  }

  .empty-upcomingTenants .cards .card .tenant h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-upcomingTenants .cards .card .desc {
    padding-top: 10px;
    border-top: 1px solid var(--gray-004);
    display: flex;
    flex-direction: column;
    gap: 5px;
    z-index: 20;
  }

  @media only screen and (max-width : 768px) {
    .empty-upcomingTenants .cards {
      grid-template-columns: 1fr;
    }
  }

  @media only screen and (max-width: 992px) {
    .empty-upcomingTenants .cards {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>