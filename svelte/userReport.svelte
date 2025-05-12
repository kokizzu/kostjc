<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */

  import LayoutMain from './_layouts/main.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <header class="header" aria-label="Date range navigation">
        <button class="btn" aria-label="Previous">
          &#8249;
        </button>
        <span class="header-text">May 2025 - August 2025</span>
        <button class="btn" aria-label="Next">
          &#8250;
        </button>
      </header>
    </div>
    <table>
      <thead>
        <tr>
          <th></th>
          <th>May</th>
          <th>June</th>
          <th>July</th>
          <th>August</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th>Room A1</th>
          <td>
            <div class="cell">
              <span>Ahmad Habibi</span> 
              <span>10 Mei s/d 10 Juni</span> 
              <span>400/1000</span>
            </div>
          </td>
          <td>
            <div class="cell">
              <span>Ahmad Habibi</span> 
              <span>10 Mei s/d 10 Juni</span> 
              <span>400/1000</span>
            </div>
          </td>
          <td>
            <div class="cell unpaid">
              <span>Ahmad Habibi</span> 
              <span>10 Mei s/d 10 Juni</span> 
              <span>400/1000</span>
            </div>
          </td>
          <td>
            <div class="cell">
              <span>Ahmad Habibi</span> 
              <span>10 Mei s/d 10 Juni</span> 
              <span>400/1000</span>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</LayoutMain>

<style>
  .report-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  table thead tr th {
    text-align: left;
    padding: 8px 12px;
    border-bottom: 1px solid var(--gray-004)
  }

  table tbody tr td {
    padding: 8px 12px;
  }

  table tbody tr td .cell {
    display: flex;
    flex-direction: column;
    gap: 2px;
    font-size: 12px;
  }

  table tbody tr td .cell.unpaid {
    padding: 5px 8px;
    background-color: var(--red-transparent);
    border-radius: 5px;
    color: var(--red-006);
  }

  .header {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
    margin-bottom: 32px;
  }
  .btn {
    background-color: var(--gray-004);
    border: none;
    border-radius: 9999px;
    width: 36px;
    height: 36px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--gray-008);
    font-size: 16px;
    transition: background-color 0.2s ease;
  }
  .btn:hover {
    background-color: var(--gray-004);
  }
  .header-text {
    font-weight: 600;
    font-size: 18px;
    user-select: none;
  }
</style>