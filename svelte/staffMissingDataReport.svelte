<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */

  import LayoutMain from './_layouts/main.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let tenants     = /** @type {Record<Number, string>} */({/* tenants */});
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <table>
      <thead>
        <tr>
          <th>Room</th>
          <th>Tenant Name</th>
          <th>Telegram</th>
          <th>WhatsApp</th>
        </tr>
      </thead>
      <tbody>
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

  table thead {
    position: sticky;
    top: 0;
    background-color: var(--gray-002);
    z-index: 10;
    border: none;
  }

  table thead tr th {
    text-align: left;
    padding: 8px 12px;
  }

  table tbody tr {
    border-bottom: 1px solid var(--gray-004);
  }
</style>