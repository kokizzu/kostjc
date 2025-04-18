<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Building} Building */
  /** @typedef {import('./_types/property.js').Location} Location */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { UserBuilding } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpAddBuilding from './_components/PopUpAddBuilding.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let building  = /** @type {Building} */ ({/* building */});
  let buildings = /** @type {any[][]} */([/* buildings */]);
  let locations = /** @type {Record<Number, string>} */({/* locations */});
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddBuilding |any
  } */ (null);
  let isSubmitAddBuilding = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true
    console.log('Locations: ', locations);
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await UserBuilding( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').UserBuildingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBuilding = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        buildings = o.buildings;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      building: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await UserBuilding(i,
      /** @type {import('./jsApi.GEN').UserBuildingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        buildings = o.buildings;
        notifier.showSuccess(`Building '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      building: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await UserBuilding(i,
      /** @type {import('./jsApi.GEN').UserBuildingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        buildings = o.buildings;
        notifier.showSuccess(`Building '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    console.log('Building ID to Edit: ' + String(id));
    const building = {
      id: payloads[0],
      buildingName: String(payloads[1]),
      locationId: String(payloads[2]),
      facilitiesObj: String(payloads[3]),
    }
    const i = /** @type {any}*/ ({
      pager,
      building,
      cmd: 'upsert'
    });
    await UserBuilding(i,
      /** @type {import('./jsApi.GEN').UserBuildingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        buildings = o.buildings;
        notifier.showSuccess(`Building '${building.buildingName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddBuilding(/** @type {Building} */ building) {
    isSubmitAddBuilding = true;
    const i = /** @type {any} */ ({
      pager,
      building,
      cmd: 'upsert'
    });

    await UserBuilding(i,
      /** @type {import('../jsApi.GEN').TenantAdminProductsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBuilding = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        buildings = o.buildings;
        notifier.showSuccess(`Building '${building.buildingName}' created !!`);

        popUpForms.Reset();

        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddBuilding
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddBuilding}
    locations={locations}
    OnSubmit={OnAddBuilding}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-building">
    <h2>Master Building</h2>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'locationId': locations
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={buildings}

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
      title="add building"
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
  .master-building {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-building h2 {
    margin: 0;
  }
</style>