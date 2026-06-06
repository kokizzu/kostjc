<script>
  /**
   * @typedef {import('../../_types/publicLoginSummary.js').PublicLoginSummary} PublicLoginSummary
   * @typedef {import('../../_types/publicLoginSummary.js').PublicUnpaidTenant} PublicUnpaidTenant
   * @typedef {import('../../_types/publicLoginSummary.js').PublicRoomNotice} PublicRoomNotice
   */

  export let publicLoginSummary = /** @type {PublicLoginSummary} */ ({/* publicLoginSummary */});

  /** @type {PublicUnpaidTenant[]} */
  const unpaidTenants = publicLoginSummary?.UnpaidTenants || [];
  /** @type {PublicRoomNotice[]} */
  const rooms = publicLoginSummary?.Rooms || [];

  /**
   * @param {number} days
   * @returns {string}
   */
  function formatDayLabel(days) {
    if (days === 0) return 'today';
    return String(days);
  }

  /**
   * @param {string} roomName
   * @returns {boolean}
   */
  function isPurpleRoom(roomName) {
    return /^[ABC]/i.test((roomName || '').trim());
  }
</script>

<section class="public-summary">
  <div class="summary-grid">
    <div class="summary-card">
      <h2>Still Not Paying (partial or fully)</h2>
      <p class="subtitle">&lt;30 = current month ongoing</p>
      {#if unpaidTenants.length}
        <div class="row-grid row-head">
          <span>Tenant</span>
          <span>Room</span>
          <span>Last date</span>
          <span>Late days</span>
        </div>
        <ul class="row-grid">
          {#each unpaidTenants as item}
            <li>
              <span class="tenant">{item.TenantName}</span>
              <span class:room-purple={isPurpleRoom(item.RoomName)} class:room-yellow={!isPurpleRoom(item.RoomName)} class="room">{item.RoomName}</span>
              <span class="date">{item.LastDate || '--'}</span>
              <span class="days">{formatDayLabel(item.DaysAgo)}</span>
            </li>
          {/each}
        </ul>
      {:else}
        <p class="empty">No unpaid tenant info available.</p>
      {/if}
    </div>

    <div class="summary-card">
      <h2>Not Yet Extended</h2>
      <p class="subtitle">negative = future</p>
      {#if rooms.length}
        <div class="row-grid row-head">
          <span>Tenant</span>
          <span>Room</span>
          <span>Last date</span>
          <span>Days ago</span>
        </div>
        <ul class="row-grid">
          {#each rooms as item}
            <li>
              <span class="tenant">{item.LastTenant || '--'}</span>
              <span class:room-purple={isPurpleRoom(item.RoomName)} class:room-yellow={!isPurpleRoom(item.RoomName)} class="room">{item.RoomName}</span>
              <span class="date">{item.LastDate || '--'}</span>
              <span class="days">{formatDayLabel(item.DaysAgo)}</span>
            </li>
          {/each}
        </ul>
      {:else}
        <p class="empty">No room extension info available.</p>
      {/if}
    </div>

    <div class="summary-card">
      <h2>Current Occupants</h2>
      {#if publicLoginSummary?.Occupants?.length}
        <div class="row-grid row-head occupant-head">
          <span>Tenant</span>
          <span>Rooms</span>
        </div>
        <ul class="occupant-list">
          {#each publicLoginSummary.Occupants as item}
            <li>
              <span class="tenant">{item.TenantName}</span>
              <span class:room-purple={isPurpleRoom(item.RoomNames?.[0] || '')} class:room-yellow={!isPurpleRoom(item.RoomNames?.[0] || '')} class="room occupant-rooms">{item.RoomNames.join(', ')}</span>
            </li>
          {/each}
        </ul>
      {:else}
        <p class="empty">No current occupants.</p>
      {/if}
    </div>
  </div>
</section>

<style>
  .public-summary {
    padding: 0 20px 20px;
    flex: 1 1 0;
    min-width: 0;
    width: 100%;
    box-sizing: border-box;
  }

  .summary-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 16px;
    align-items: start;
    width: 100%;
    min-width: 0;
  }

  .summary-card {
    width: 100%;
    min-width: 0;
    background: #fff;
    border: 1px solid var(--gray-003);
    border-radius: 16px;
    padding: 10px 12px;
    box-shadow: var(--shadow-md);
    align-self: start;
  }

  .summary-card h2 {
    margin: 0 0 12px;
    font-size: var(--font-lg);
  }

  .subtitle {
    margin: -6px 0 8px;
    color: var(--gray-006);
    font-size: var(--font-sm);
  }

  .summary-card ul {
    list-style: none;
    padding: 0;
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .row-head,
  .row-grid li {
    display: grid;
    grid-template-columns: minmax(0, 2.4fr) minmax(0, 1fr) 74px 56px;
    gap: 6px;
  }

  .occupant-head,
  .occupant-list li {
    display: grid;
    grid-template-columns: minmax(0, 2fr) minmax(0, 1.2fr);
    gap: 6px;
  }

  .row-grid li {
    padding: 5px 6px;
    border-radius: 10px;
    background: var(--gray-001);
    align-items: center;
    line-height: 1.15;
  }

  .row-head {
    margin-bottom: 8px;
    padding: 0 6px;
    color: var(--gray-006);
    font-size: var(--font-sm);
    font-weight: 700;
  }

  .row-head span {
    text-transform: uppercase;
    letter-spacing: 0.04em;
  }

  .tenant {
    font-weight: 700;
    display: block;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .room {
    color: var(--gray-009);
    font-weight: 700;
    display: block;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .occupant-rooms {
    font-weight: 600;
    display: block;
  }

  .room.room-purple {
    background-color: #faf5ff;
    color: #6b21a8;
  }

  .room.room-yellow {
    background-color: #fffdf2;
    color: #a16207;
  }

  .date,
  .days {
    color: var(--gray-006);
    font-size: var(--font-sm);
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    text-align: right;
  }

  .empty {
    margin: 0;
    color: var(--gray-006);
  }

  @media (max-width: 1080px) {
    .summary-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }

    .row-head,
    .row-grid li {
      grid-template-columns: minmax(0, 2.2fr) minmax(0, 1fr) 74px;
    }
  }

  @media (max-width: 820px) {
    .row-head,
    .row-grid li {
      grid-template-columns: minmax(0, 2fr) minmax(0, 1fr);
    }
  }

  @media (max-width: 620px) {
    .summary-grid {
      grid-template-columns: 1fr;
    }
  }

  @media (max-width: 720px) {
    .public-summary {
      padding-inline: 16px;
    }

    .row-head,
    .row-grid li {
      grid-template-columns: 1fr;
    }

    .occupant-head,
    .occupant-list li {
      grid-template-columns: 1fr;
    }
  }
</style>
