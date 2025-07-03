<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
   /** @typedef {import('./_types/masters.js').ExtendedActionButton} ExtendedActionButton */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/cafe.js').Sale} Sale */
  /** @typedef {import('./_types/cafe.js').Menu} Menu */
  
  import { onMount } from 'svelte';

  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';
  import { notifier } from './_components/xNotifier';
  import PopUpAddSale from './_components/PopUpAddSale.svelte';
  
  import { AdminSale } from './jsApi.GEN';

  import { Icon } from './node_modules/svelte-icons-pack/dist';
  
  import { RiSystemAddBoxLine} from './node_modules/svelte-icons-pack/dist/ri';


  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let sale  = /** @type {Sale} */ ({/* sale */});
  let sales = /** @type {any[][]} */([/* sales */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});
  let tenants = /** @type {Record<Number, string>} */ ({/* tenants */});
  let menus     = /** @type {Record<Number, string>} */({/* menus */});

  const PaymentMethods = [
    'Cash',
    'QRIS',
    'Transfer',
    'Donation',
    'TopUp',
    'Debt'
  ];
  const PaymentStatuses = [
    'Paid',
    'Unpaid',
    'Refunded'
  ];

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddSale |any
  } */ (null);
  let isSubmitAddSale = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminSale( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddSale = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        sales = o.sales;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      sale: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminSale(i,
      /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        sales = o.sales;
        notifier.showSuccess(`Sales restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      sale: {
        id: row[0]
      },
      cmd: CmdDelete
    });
    await AdminSale(i,
      /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        sales = o.sales;
        notifier.showSuccess(`Sale deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  // @ts-ignore
  async function OnEdit(/** @type {any} */ id, /** @type {any[]} */ payloads) {
    const sale = {
    id: payloads[0],
    cashier: String(payloads[1]),
    tenantId: String(payloads[2]),
    buyerName: String(payloads[3]),
    menuIds: payloads[4],
    paymentMethod: String(payloads[5]),
    paymentStatus: String(payloads[6]),
    transferIDR: Number(payloads[7]),
    qrisIDR: Number(payloads[8]),
    cashIDR: Number(payloads[9]),
    debtIDR: Number(payloads[10]),
    topupIDR: Number(payloads[11]),
    donation: Number(payloads[12]),
    totalPriceIDR: Number(payloads[13]),
    salesDate: String(payloads[14]),
    paidAt: String(payloads[15]),
    note: String(payloads[16]),
  }
    const i = /** @type {any}*/ ({
      pager,
      sale,
      cmd: CmdUpsert
    });
    await AdminSale(i,
      /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        sales = o.sales;
        notifier.showSuccess(`Sale #${sale.id} updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddSale(/** @type {Sale} */ sale) {
    isSubmitAddSale = true;
    const i = /** @type {any} */ ({
      pager,
      sale,
      cmd: CmdUpsert
    });

    await AdminSale(i,
      /** @type {import('../jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddSale = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        sales = o.sales;
        notifier.showSuccess(`Sale created !!`);

        popUpForms.Reset();
        popUpForms.Hide();
        OnRefresh(pager);
      }
    );
  }

</script>

{#if isPopUpFormReady}
  <PopUpAddSale
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddSale}
    tenants={tenants}
    menus={menus}
    OnSubmit={OnAddSale}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-table">
    <h2>Sale Management</h2>
    <MasterTable
      NAME="Sales"
      ACCESS={segments}
      REFS={{
        'tenantId': tenants,
        'menuIds': menus,
        'paymentStatus': PaymentStatuses,
        'paymentMethod': PaymentMethods
      }}

      COL_WIDTHS={{
        'tenantId': 300,
        'menuIds': 300
      }}

      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={sales}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}

      SINGLE_SELECTED={true}
    >
    <button
      class="btn"
      on:click={() => popUpForms.Show()}
      title="add sale"
    >
      <Icon
        color="var(--gray-007)"
        size="18"
        src={RiSystemAddBoxLine}
      />
    </button>
    </MasterTable>
  </div>
</LayoutMain>

<style>
  .master-table {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-table h2 {
    margin: 0;
  }
</style>