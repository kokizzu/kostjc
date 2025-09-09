<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/cafe.js').Laundry} Laundry */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminLaundry } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let laundry     = /** @type {Laundry} */ ({/* laundry */});
  let laundries   = /** @type {any[][]} */([/* laundries */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitAddLaundry = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminLaundry( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminLaundryCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddLaundry = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        laundries = o.laundries;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      laundry: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminLaundry(i,
      /** @type {import('./jsApi.GEN').AdminLaundryCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        laundries = o.laundries;
        notifier.showSuccess(`Laundry '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      laundry: {
        id: row[0]
      },
      cmd: CmdDelete
    });
    await AdminLaundry(i,
      /** @type {import('./jsApi.GEN').AdminLaundryCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        laundries = o.laundries;
        notifier.showSuccess(`Laundry '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const laundry = {
      id: id,
      customer: payloads[1],
      items: payloads[2],
      status: payloads[3],
      note: payloads[4],
      laundryAt: payloads[5],
    }
    const i = /** @type {any}*/ ({
      pager,
      laundry,
      cmd: CmdUpsert
    });
    await AdminLaundry(i,
      /** @type {import('./jsApi.GEN').AdminLaundryCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        laundries = o.laundries;
        notifier.showSuccess(`Laundry #${id}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddlaundry(/** @type any[] */ payloads) {
    isSubmitAddLaundry = true;

    const laundry = /** @type {any} */ ({
      customer: payloads[1],
      items: payloads[2],
      status: payloads[3],
      note: payloads[4],
      laundryAt: payloads[5]
    });
    const i = /** @type {any} */ ({
      pager,
      laundry,
      cmd: CmdUpsert
    });

    await AdminLaundry(i,
      /** @type {import('../jsApi.GEN').AdminLaundryCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddLaundry = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        laundries = o.laundries;
        notifier.showSuccess(`Laundry '${laundry.items}' created !!`);
        
        OnRefresh(pager);
        popUpForms.Reset();
        popUpForms.Hide();
      }
    );
  }

  const laundryStatus = [
    'Pending',
    'In-Progress',
    'Completed'
  ]
</script>

{#if isPopUpFormReady}
  <PopUpForms
    bind:this={popUpForms}
    heading="Add Laundry"
    FIELDS={fields}
    REFS={{
      'status': laundryStatus
    }}
    bind:isSubmitted={isSubmitAddLaundry}
    OnSubmit={OnAddlaundry}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-laundry">
    <h2>Laundry Management</h2>
    <MasterTable
      NAME="Laundry"
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={laundries}
      REFS={{
        'status': laundryStatus
      }}
      FIELD_TO_SEARCH="items"

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
      title="add laundry"
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
  .master-laundry {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-laundry h2 {
    margin: 0;
  }
</style>