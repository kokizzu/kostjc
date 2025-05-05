<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Location} Location */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminLocation } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let location  = /** @type {Location} */ ({/* location */});
  let locations = /** @type {any[][]} */([/* locations */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitAddLocation = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true;
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await AdminLocation( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminLocationCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddLocation = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        locations = o.locations;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      location: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await AdminLocation(i,
      /** @type {import('./jsApi.GEN').AdminLocationCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        locations = o.locations;
        notifier.showSuccess(`Location '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      location: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminLocation(i,
      /** @type {import('./jsApi.GEN').AdminLocationCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        locations = o.locations;
        notifier.showSuccess(`Location '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    console.log('Location ID to Edit: ' + String(id));
    const location = {
      id: payloads[0],
      name: payloads[1],
      address: payloads[2],
      gmapLocation: payloads[3]
    }
    const i = /** @type {any}*/ ({
      pager,
      location,
      cmd: 'upsert'
    });
    await AdminLocation(i,
      /** @type {import('./jsApi.GEN').AdminLocationCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        locations = o.locations;
        notifier.showSuccess(`Location '${location.name}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddLocation(/** @type any[] */ payloads) {
    isSubmitAddLocation = true;

    const location = /** @type {any} */ ({
      name: payloads[1],
      address: payloads[2],
      gmapLocation: payloads[3]
    });
    const i = /** @type {any} */ ({
      pager,
      location,
      cmd: 'upsert'
    });

    await AdminLocation(i,
      /** @type {import('./jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddLocation = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;
        notifier.showSuccess(`Location '${location.name}' created !!`);

        OnRefresh(pager);
        popUpForms.Reset();
        popUpForms.Hide();
      }
    );
  }
</script>

{#if isPopUpFormReady}
  <PopUpForms
    bind:this={popUpForms}
    heading="Add Location"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddLocation}
    OnSubmit={OnAddLocation}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-location">
    <h2>Location Management</h2>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={locations}

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
      title="add account"
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
  .master-location {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-location h2 {
    margin: 0;
  }
</style>