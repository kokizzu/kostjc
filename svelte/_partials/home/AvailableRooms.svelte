<script>
  /**
   * @typedef {Object} AvailableRoom
   * @property {string} roomName
   * @property {string} availableAt
   * @property {boolean} isAvailableNow
   */
  const rooms = /** @type {AvailableRoom[]} */ ([/* availableRooms */]);

  function formatDateLong(/** @type {string} */ dateStr) {
    const dt = new Date(dateStr);
    return dt.toLocaleDateString('en-GB', {
      weekday: 'long',
      day: '2-digit',
      month: 'long',
      year: 'numeric'
    });
  }
</script>

<section class="empty-rooms">
  <h1>Available Rooms</h1>
  {#if rooms && rooms.length > 0}
    <div class="cards">
      {#each (rooms || []) as r}
        <div class="card">
          <h3>Room {r.roomName}</h3>
          <div class="desc">
            <span>{@html r.isAvailableNow
              ? 'Available Now'
              : 'Available on <b>' + formatDateLong(r.availableAt)+'</b>'
            }</span>
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

  .empty-rooms h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .empty-rooms .cards {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    gap: 10px;
  }

  .empty-rooms .cards .card {
    background-color: var(--gray-001);
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .empty-rooms .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .empty-rooms .cards .card .desc {
    padding-top: 10px;
    border-top: 1px solid var(--gray-004);
    display: flex;
    flex-direction: column;
    gap: 5px;
    z-index: 20;
  }
</style>