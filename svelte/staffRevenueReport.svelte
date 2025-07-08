<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').ChartRevenueReport} ChartRevenueReport */
  /**
   * @typedef {Object} RevenueReport
   * @property {string} yearMonth
   * @property {number} bookingId
   * @property {number} revenueIDR
   * @property {number} donationIDR
   */

  import InputBox from './_components/InputBox.svelte';
    import MonthShifter from './_components/MonthShifter.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { formatYearMonth } from './_components/xFormatter';
  import { notifier } from './_components/xNotifier';
  import LayoutMain from './_layouts/main.svelte';
  import ChartRevenueDaily from './_partials/ChartRevenueMonthly.svelte';
  import { StaffRevenueReport } from './jsApi.GEN';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let revenueReports = /** @type {RevenueReport[]} */ ([/* revenueReports */]);
  let chartRevenueReports = /** @type {ChartRevenueReport[]} */ ([/* chartRevenueReports */]);
  let bookings = /** @type {Record<Number, string>} */ ({/* bookings */});

  let chartRevenueDaily = /** @type {ChartRevenueDaily|import('svelte').SvelteComponent} */ (null);

  let yearMonth = /** @type {string} */ (new Date().toISOString().slice(0, 7));
  let isFiltering = /** @type {boolean} */ (false);

  async function getRevenueReports() {
    isFiltering = true;
    await StaffRevenueReport(// @ts-ignore
      { yearMonth }, /** @type {import('./jsApi.GEN').StaffRevenueReportCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error || 'Error filtering revenue reports');
          return
        }
        
        revenueReports = o.revenueReports;
        chartRevenueReports = o.chartRevenueReports;

        chartRevenueDaily.updateData();
      }
    );
    isFiltering = false;
  }

  let sumRevenueIDR = 0;
  let sumDonationIDR = 0;

  $: {
    sumRevenueIDR = 0;
    sumDonationIDR = 0;
    for(let i=0; i< (revenueReports || []).length; i++) {
      sumRevenueIDR += revenueReports[i].revenueIDR|0
      sumDonationIDR += revenueReports[i].donationIDR|0
    }
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <div class="actions">
      <MonthShifter
        bind:yearMonth
        OnChanges={getRevenueReports}
      />
    </div>
    <ChartRevenueDaily
      bind:chartRevenueReports
      bind:this={chartRevenueDaily}
    />
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th style="min-width: 100px;">Month</th>
            <th style="min-width: 170px;">Booking</th>
            <th>Revenue (IDR)</th>
            <th>Donation (IDR)</th>
          </tr>
        </thead>
        <tbody>
          {#each (revenueReports || []) as data}
            <tr>
              <td>{formatYearMonth(data.yearMonth)}</td>
              <td>{bookings[data.bookingId]}</td>
              <td class="r">{data.revenueIDR}</td>
              <td class="r">{data.donationIDR}</td>
            </tr>
          {/each}

          {#if !revenueReports || !revenueReports.length}
            <tr>
              <td>No data</td>
            </tr>
          {/if}
        </tbody>
			<tfoot>
				<tr>
					<th></th>
					<th></th>
					<th class="r">{sumRevenueIDR}</th>
					<th class="r">{sumDonationIDR}</th>
				</tr>
			</tfoot>
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
  td.r, th.r {
	 text-align: right;
  }
</style>