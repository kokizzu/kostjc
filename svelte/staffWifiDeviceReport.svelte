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
   * @property {number} deletedAt
   */

  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { localeDateFromYYYYMMDD } from './_components/xFormatter';
  import { notifier } from './_components/xNotifier';
  import LayoutMain from './_layouts/main.svelte';
  import { StaffWifiDeviceReport } from './jsApi.GEN';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let wifiDeviceReports = /** @type {WifiDeviceReport[]} */ ([/* wifiDeviceReports */]);
  const rooms  = /** @type {Record<Number, string>} */ ({/* rooms */});
  const tenants = /** @type {Record<Number, string>} */ ({/* tenants */});

  function currentYearMonth() {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    return `${year}-${month}`;
  }

  let yearMonth = /** @type {string} */ (currentYearMonth());
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

  function isTodayGreater(dateStr) {
    const inputDate = new Date(dateStr);
    const today = new Date();
    return today > inputDate;
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
            <tr class="
              {isTodayGreater(data.endAt) ? 'reminding' : ''} 
              {data.deletedAt > 0 ? 'deleted' : ''}
            ">
              <td>{localeDateFromYYYYMMDD(data.startAt)}</td>
              <td>{localeDateFromYYYYMMDD(data.endAt)}</td>
              <td>{localeDateFromYYYYMMDD(data.paidAt)}</td>
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

  table tbody tr.reminding {
    background-color: var(--yellow-transparent);
    color: var(--yellow-006);
  }

  table tbody tr.deleted {
    background-color: var(--red-transparent) !important;
    color: var(--red-006) !important;
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