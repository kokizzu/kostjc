<script>
    import SubmitButton from "../../_components/SubmitButton.svelte";
    import { CmdUpsert } from "../../_components/xConstant";
    import { localeDateFromYYYYMMDD } from "../../_components/xFormatter";
    import { notifier } from "../../_components/xNotifier";
    import { AdminSettingFixInconsistencies } from "../../jsApi.GEN";

  export let tenants          = /** @type {Record<number, string>} */ ({/* tenants */});
  export let rooms            = /** @type {Record<number, string>} */ ({/* rooms */});

  /**
   * @typedef {Object} RoomBookingInconsistency
   * @property {number} roomId
   * @property {number} tenantId
   * @property {number} currentTenantId
   * @property {string} dateEnd
   * @property {boolean} isInconsistent
   */
  let roomIncosistencies = /** @type {RoomBookingInconsistency[]} */ ([/* roomIncosistencies */]);

  let isFixing = false;
  async function fixInconsistencies() {
    isFixing = true;
    await AdminSettingFixInconsistencies({// @ts-ignore
      cmd: CmdUpsert
      }, /** @type {import('./jsApi.GEN').AdminSettingFixInconsistenciesCallback} */
      /** @returns {Promise<void>} */
        function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error || 'failed to fix inconsistencies');
          return
        }
        
        roomIncosistencies = o.roomIncosistencies;
      }
    );
  }
</script>

<div class="container">
  <div class="heading">
    <h1>Rooms Inconsistencies</h1>
    <SubmitButton
      label="Fix Now"
      on:click={fixInconsistencies}
      isSubmitted={isFixing}
    />
  </div>
  <div class="table-container">
    <table>
      <thead>
        <tr>
          <th>Room</th>
          <th>Tenant</th>
          <th>Current Tenant</th>
          <th>Date End</th>
          <th>Consistent</th>
        </tr>
      </thead>
      <tbody>
        {#each (roomIncosistencies || []) as data}
          <tr class="">
            <td>{rooms[data.roomId]}</td>
            <td>{tenants[data.tenantId]}</td>
            <td>{tenants[data.currentTenantId]}</td>
            <td>{localeDateFromYYYYMMDD(data.dateEnd)}</td>
            <td>{data.isInconsistent ? '❌' : '✅'}</td>
          </tr>
        {/each}

        {#if !roomIncosistencies || (roomIncosistencies || []).length == 0}
          <tr>
            <td>No data</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .container .heading {
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: space-between;
  }

  .container .heading h1 {
    margin: 0;
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
    .container .heading {
      flex-direction: column;
    }

    :global(.container .heading .submit-btn) {
      width: 100% !important;
    }
  }
</style>