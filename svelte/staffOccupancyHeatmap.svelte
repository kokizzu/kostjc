<script>
    import MonthShifter from './_components/MonthShifter.svelte';
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Booking} Booking */
  /** @typedef {import('./_types/property.js').Facility} Facility */
  /** @typedef {import('./_types/property.js').Payment} Payment */

  import LayoutMain from './_layouts/main.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let roomNames = /** @type {string[]} */ ([/* roomNames */]);
  let tenants     = /** @type {Record<Number, string>} */({/* tenants */});
  let rooms       = /** @type {Record<Number, string>} */({/* rooms */});

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);

  async function RefreshData() {
    
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="occupancy-heatmap-container">
    <MonthShifter
      bind:yearMonth
      bind:isLoading
      OnChanges={RefreshData}
    />
    <div class="rooms-heatmap">
      {#each roomNames as room}
        <div class="heatmap-container">
          <h2>Room {room}</h2>
          <div class="heatmap-grid">
            {#each Array.from({ length: 31 }) as _, day}
              <span class="heatmap-cell">
                {day+1}
              </span>
            {/each}
          </div>
        </div>
      {/each}
    </div>
  </div>
</LayoutMain>

<style>
  .occupancy-heatmap-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  .rooms-heatmap {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .heatmap-container {
    display: grid;
    grid-template-columns: 100px 1fr;
    align-items: center;
    gap: 20px;
    width: 100%;
    height: fit-content;
  }

  .heatmap-container h2 {
    margin: 0;
  }

  .heatmap-container .heatmap-grid {
    display: grid;
    grid-template-columns: repeat(11, 1fr);
    gap: 5px;
    flex-grow: 1;
  }

  .heatmap-container .heatmap-grid .heatmap-cell {
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid var(--gray-002);
    background-color: var(--gray-001);
    padding: 10px;
    border-radius: 4px;
    width: 100%;
    font-size: var(--font-lg);
    font-weight: 600;
    cursor: pointer;
  }
</style>