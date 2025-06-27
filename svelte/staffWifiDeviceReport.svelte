<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /**
   * @typedef {Object} WifiDeviceReport
   * @property {number} tenantId
   * @property {number} roomId
   * @property {string} startAt
   * @property {string} endAt
   * @property {string} paidAt
   */

  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { notifier } from './_components/xNotifier';
  import LayoutMain from './_layouts/main.svelte';
  import { StaffWifiDeviceReport } from './jsApi.GEN';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let wifiDeviceReports = /** @type {WifiDeviceReport[]} */ ([/* wifiDeviceReports */]);
  const rooms  = /** @type {Record<Number, string>} */ ({/* rooms */});
  const tenants = /** @type {Record<Number, string>} */ ({/* tenants */});

  console.log('wifiDeviceReports:', wifiDeviceReports);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isFiltering = /** @type {boolean} */ (false);

  async function getWifiDeviceReports() {
    isFiltering = true;
    await StaffWifiDeviceReport(// @ts-ignore
      { yearMonth }, /** @type {import('./jsApi.GEN').StaffWifiDeviceReportCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error || 'Error filtering wifi device reports');
          return
        }
        
        wifiDeviceReports = o.wifiDeviceReports;
      }
    );
    isFiltering = false;
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <InputBox
        className="year-month"
        id="yearMonth"
        type="month"
        placeholder="Select Month"
        label=""
        bind:value={yearMonth}
      />
      <SubmitButton
        label="Filter"
        on:click={getWifiDeviceReports}
        isSubmitted={isFiltering}
      />
    </div>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Start At</th>
            <th>End At</th>
            <th>Paid At</th>
            <th>Tenant</th>
            <th>Room</th>
          </tr>
        </thead>
        <tbody>
          {#each (wifiDeviceReports || []) as data}
            <tr>
              <td>{data.startAt || '--'}</td>
              <td>{data.endAt || '--'}</td>
              <td>{data.paidAt || '--'}</td>
              <td>{tenants[data.tenantId] || '--'}</td>
              <td>{rooms[data.roomId] || '--'}</td>
            </tr>
          {/each}

          {#if !wifiDeviceReports || (wifiDeviceReports || []).length == 0}
            <tr>
              <td>No data</td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</LayoutMain>

<style>
  .report-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  .report-container .actions {
    display: flex;
    flex-direction: row;
    gap: 10px;
    justify-content: flex-end;
  }

  :global(.report-container .actions .year-month) {
    width: 300px;
  }

  .table-container {
    overflow-x: auto;
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
    text-wrap: nowrap;
  }

  table tbody tr {
    border-bottom: 1px solid var(--gray-004);
  }

  table tr td, table tr th {
    padding: 4px 12px;
  }

  @media only screen and (max-width : 768px) {
    .report-container .actions {
      flex-direction: column;
    }

    :global(.report-container .actions .year-month) {
      width: 100%;
    }

    :global(.report-container .actions .submit-btn) {
      width: 100% !important;
    }
  }
</style>