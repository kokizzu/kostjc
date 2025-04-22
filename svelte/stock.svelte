<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Stock} Stock */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminStock } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let stock  = /** @type {Stock} */ ({/* stock */});
  let stocks = /** @type {any[][]} */([/* stocks */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitAddFacility = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await AdminStock( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminStockCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddFacility = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        stocks = o.stocks;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      stock: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await AdminStock(i,
      /** @type {import('./jsApi.GEN').AdminStockCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        stocks = o.stocks;
        notifier.showSuccess(`Stock '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      stock: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminStock(i,
      /** @type {import('./jsApi.GEN').AdminStockCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        stocks = o.stocks;
        notifier.showSuccess(`Stock '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const stock = {
      id: payloads[0],
      stockName: payloads[1],
      stockAddedAt: payloads[2],
      quantity: Number(payloads[3]),
      priceIDR: Number(payloads[4])
    }
    const i = /** @type {any}*/ ({
      pager,
      stock,
      cmd: 'upsert'
    });
    await AdminStock(i,
      /** @type {import('./jsApi.GEN').AdminStockCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        stocks = o.stocks;
        notifier.showSuccess(`Stock '${stock.stockName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddFacility(/** @type any[] */ payloads) {
    isSubmitAddFacility = true;

    const stock = /** @type {any} */ ({
      stockName: payloads[1],
      stockAddedAt: payloads[2],
      quantity: Number(payloads[3]),
      priceIDR: Number(payloads[4])
    });
    const i = /** @type {any} */ ({
      pager,
      stock,
      cmd: 'upsert'
    });

    await AdminStock(i,
      /** @type {import('../jsApi.GEN').AdminStockCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddFacility = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        stocks = o.stocks;
        notifier.showSuccess(`Stock '${stock.facilityName}' created !!`);

        popUpForms.Reset();

        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormReady}
  <PopUpForms
    bind:this={popUpForms}
    heading="Add Stock"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddFacility}
    OnSubmit={OnAddFacility}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-stock">
    <h2>Master Stock</h2>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={stocks}

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
    <button
      class="btn"
      on:click={() => popUpForms.Show()}
      title="add stock"
    >
      <Icon
        color="var(--gray-007)"
        size="16"
        src={RiSystemAddBoxLine}
      />
    </button>
    </MasterTable>
  </div>
</LayoutMain>

<style>
  .master-stock {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-stock h2 {
    margin: 0;
  }
</style>