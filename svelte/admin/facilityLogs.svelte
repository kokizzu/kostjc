<script>
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/masters.js').Field} Field */
  /** @typedef {import('../_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/users.js').User} User */
  /** @typedef {import('../_types/property.js').Booking} Booking */
  /** @typedef {import('../_types/property.js').Facility} Facility */
  
  import LayoutMain from '../_layouts/main.svelte';
  import MasterTableActionlog from '../_components/MasterTableActionlog.svelte';
  import { AdminFacilityLogs } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier';
  import LogsSubMenu from '../_partials/LogsSubMenu.svelte';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let logs    = /** @type {any[][]} */([/* logs */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  async function refreshTableView(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn
    };
    await AdminFacilityLogs( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').AdminFacilityLogsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        logs = o.logs;
      }
    );
  }
</script>

<LayoutMain access={segments} user={user}>
  <div class="logs-container">
    <LogsSubMenu />
    <MasterTableActionlog arrayOfArray={false}
      bind:pager={pager}
      {fields}
      onRefreshTableView={refreshTableView}
      rows={logs}
    />
  </div>
</LayoutMain>

<style>
  .logs-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
</style>