<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').ChartRevenueReport} ChartRevenueReport */
  /**
   * @typedef {Object} RevenueReport
   * @property {string} dateStart
   * @property {number} bookingId
   * @property {number} cashIDR
   * @property {number} transferIDR
   * @property {number} revenueIDR
   * @property {number} donationIDR
   */

  import MonthShifter from './_components/MonthShifter.svelte';
  import { notifier } from './_components/xNotifier';
  import LayoutMain from './_layouts/main.svelte';
  import ChartRevenueMonthly from './_partials/ChartRevenueMonthly.svelte';
  import { AdminRevenueReport } from './jsApi.GEN';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let revenueReports = /** @type {RevenueReport[]} */ ([/* revenueReports */]);
  let chartRevenueReports = /** @type {ChartRevenueReport[]} */ ([/* chartRevenueReports */]);
  let bookings = /** @type {Record<Number, string>} */ ({/* bookings */});

  let chartRevenueMonthly = /** @type {ChartRevenueMonthly|import('svelte').SvelteComponent} */ (null);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isLoading = /** @type {boolean} */ (false);

  async function getRevenueReports() {
    isLoading = true;
    await AdminRevenueReport(// @ts-ignore
      { yearMonth }, /** @type {import('./jsApi.GEN').AdminRevenueReportCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error || 'Error filtering revenue reports');
          return
        }
        
        revenueReports = o.revenueReports;
        chartRevenueReports = o.chartRevenueReports;

        chartRevenueMonthly.updateData(chartRevenueReports);
      }
    );
    isLoading = false;
  }

  let sumRevenueIDR = 0;
  let sumCashIDR = 0;
  let sumTransferIDR = 0;
  let sumDonationIDR = 0;

  $: {
    sumCashIDR = 0;
    sumTransferIDR = 0;
    sumRevenueIDR = 0;
    sumDonationIDR = 0;
    for(let i=0; i< (revenueReports || []).length; i++) {
      sumCashIDR += revenueReports[i].cashIDR|0
      sumTransferIDR += revenueReports[i].transferIDR|0
      sumRevenueIDR += revenueReports[i].revenueIDR|0
      sumDonationIDR += revenueReports[i].donationIDR|0
    }
  }

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

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <MonthShifter
        bind:yearMonth
        bind:isLoading
        OnChanges={getRevenueReports}
      />
    </div>
    <ChartRevenueMonthly
      bind:chartRevenueReports
      bind:this={chartRevenueMonthly}
    />
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th style="min-width: 100px;">Date Start</th>
            <th style="min-width: 170px;">Booking</th>
            <th>Cash</th>
            <th>Transfer</th>
            <th>Total</th>
            <th>Donation</th>
          </tr>
        </thead>
        <tbody>
          <tr class="summary-row">
            <td><strong>Total</strong></td>
            <td></td>
            <td class="r"><strong>{sumCashIDR}</strong></td>
            <td class="r"><strong>{sumTransferIDR}</strong></td>
            <td class="r"><strong>{sumRevenueIDR}</strong></td>
            <td class="r"><strong>{sumDonationIDR}</strong></td>
          </tr>
          {#each (revenueReports || []) as data}
            <tr>
              <td>{formatDateLong(data.dateStart)}</td>
              <td>{bookings[data.bookingId]}</td>
              <td class="r">{data.cashIDR}</td>
              <td class="r">{data.transferIDR}</td>
              <td class="r">{data.revenueIDR}</td>
              <td class="r">{data.donationIDR}</td>
            </tr>
          {/each}

          {#if !revenueReports || !revenueReports.length}
            <tr>
              <td colspan="6">No data</td>
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
    justify-content: center;
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

  /* align right */
  td.r {
    text-align: right;
  }

  .summary-row {
    position: sticky;
    top: 0;
    background: var(--gray-002);
    z-index: 9;
  }

  .summary-row td {
    border-bottom: 2px solid var(--gray-004);
  }
</style>
