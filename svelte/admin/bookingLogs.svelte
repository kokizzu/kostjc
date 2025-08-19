<script>
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/masters.js').Field} Field */
  /** @typedef {import('../_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/users.js').User} User */
  /** @typedef {import('../_types/masters.js').ExtendedActionButton} ExtendedActionButton */
  
  import LayoutMain from '../_layouts/main.svelte';
  import MasterTableActionlog from '../_components/MasterTableActionlog.svelte';
  import { AdminBookingLogs } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier';
  import LogsSubMenu from '../_partials/LogsSubMenu.svelte';
  import { RiSystemInformationLine, RiDesignShadowLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import PopUpCompareJson from '../_components/PopUpCompareJson.svelte';
  import PopUpDiffLogJson from '../_components/PopUpDiffLogJson.svelte';
  import { onMount } from 'svelte';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let logs    = /** @type {any[][]} */([/* logs */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});
  let users     = /** @type {Record<number, string>} */({/* users */});
  let tenants = /** @type {Record<number, string>} */({/* tenants */});

  async function refreshTableView(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn
    };
    console.log('req:', i)
    await AdminBookingLogs( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').AdminBookingLogsCallback} */
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
  
  let isPopUpReady = false;
  onMount(() => {
    isPopUpReady = true;
  })

  let popUpCompareJson;
  let popUpDiffLogJson;

  let beforeJson = '';
  let afterJson = '';

  const EXTENDED_BUTTONS = /** @type {ExtendedActionButton[]} */ ([
    {
      icon: RiSystemInformationLine,
      tooltip: 'Compare Data Before & After',
      action: (/** @type {Object} */ row) => {
        let dataBeforeJsonObj = row.beforeJson;
        let dataAfterJsonObj = row.afterJson;

        if (typeof row.beforeJson === 'string') {
          dataBeforeJsonObj = JSON.parse(row.beforeJson || '{}');
        }
        beforeJson = JSON.stringify(dataBeforeJsonObj, null, 2);

        if (typeof row.afterJson === 'string') {
          dataAfterJsonObj = JSON.parse(row.afterJson || '{}');
        }
        afterJson = JSON.stringify(dataAfterJsonObj, null, 2);

        popUpCompareJson.Show();
      }
    },
    {
      icon: RiDesignShadowLine,
      tooltip: 'Show Diff',
      action: (/** @type {Object} */ row) => {
        beforeJson = row.beforeJson;
        afterJson = row.afterJson;

        setTimeout(() => {
          popUpDiffLogJson.Show();
        }, 500);
      }
    }
  ]);
</script>

{#if isPopUpReady}
  <PopUpCompareJson
    bind:this={popUpCompareJson}
    bind:beforeJson
    bind:afterJson
  />

  <PopUpDiffLogJson
    bind:this={popUpDiffLogJson}
    bind:beforeJson
    bind:afterJson
    {users} {tenants}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="logs-container">
    <LogsSubMenu />
    <MasterTableActionlog arrayOfArray={false}
      bind:pager={pager}
      {fields}
      onRefreshTableView={refreshTableView}
      rows={logs}
      REFS={{
        'actorId': users
      }}
      COL_WIDTHS={{
        'actorId': 200
      }}
      {EXTENDED_BUTTONS}
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