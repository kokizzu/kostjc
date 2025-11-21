<script>
  /**
   * @typedef {Object} AvailableRoom
   * @property {string} roomName
   * @property {string} availableAt
   * @property {boolean} isAvailableNow
   * @property {string} lastTenant
   * @property {number} basePriceIDR
   * @property {number} totalPriceIDR
   */
  const rooms = /** @type {AvailableRoom[]} */ ([/* availableRooms */]);
  let sortedRooms = /** @type {AvailableRoom[]} */ ([]);
  let sortBy = /** @type {'default' | 'basePrice' | 'totalPrice'} */ ('default');

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
    const availableDate = new Date(availableAt);
    const now = new Date();

    // @ts-ignore
    const diffMs = now - availableDate;
    const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

    if (diffDays <= 0) return '(today)';

    return `(since ${diffDays} day${diffDays > 1 ? 's' : ''} ago)`;
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
</script>

<section class="empty-rooms">
  <div class="header">
    <h1>Available Rooms</h1>
    <div class="sort-controls">
      <label>Sort by:</label>
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
        <div class="card">
          <h3>Room {r.roomName}</h3>
          <div class="desc">
            <span>{@html r.isAvailableNow || r.availableAt == ''
              ? `Available Now <b>${getAvailableDays(r.availableAt)}</b>`
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

  .sort-controls label {
    font-weight: 600;
    font-size: 14px;
    color: var(--gray-007, #333);
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
    gap: 8px;
    z-index: 20;
  }

  .price-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 6px 0;
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