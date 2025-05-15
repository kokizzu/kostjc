<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Facility} Facility */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminFacility } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
    import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let facility    = /** @type {Facility} */ ({/* facility */});
  let facilities  = /** @type {any[][]} */([/* facilities */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitAddFacility = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminFacility( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddFacility = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        facilities = o.facilities;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      facility: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminFacility(i,
      /** @type {import('./jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        facilities = o.facilities;
        notifier.showSuccess(`Facility '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      facility: {
        id: row[0]
      },
      cmd: CmdDelete
    });
    await AdminFacility(i,
      /** @type {import('./jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        facilities = o.facilities;
        notifier.showSuccess(`Facility '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const facility = {
      id: id,
      facilityName: payloads[1],
      facilityType: payloads[2],
      extraChargeIDR: Number(payloads[3]),
      descriptionEN: payloads[4]
    }
    const i = /** @type {any}*/ ({
      pager,
      facility,
      cmd: CmdUpsert
    });
    await AdminFacility(i,
      /** @type {import('./jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        facilities = o.facilities;
        notifier.showSuccess(`Facility '${facility.facilityName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddFacility(/** @type any[] */ payloads) {
    isSubmitAddFacility = true;

    const facility = /** @type {any} */ ({
      facilityName: payloads[1],
      facilityType: payloads[2],
      extraChargeIDR: Number(payloads[3]),
      descriptionEN: payloads[4]
    });
    const i = /** @type {any} */ ({
      pager,
      facility,
      cmd: CmdUpsert
    });

    await AdminFacility(i,
      /** @type {import('../jsApi.GEN').AdminFacilityCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddFacility = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        facilities = o.facilities;
        notifier.showSuccess(`Facility '${facility.facilityName}' created !!`);
        
        OnRefresh(pager);
        popUpForms.Reset();
        popUpForms.Hide();
      }
    );
  }

  const FacilityTypes = [
    'Room',
    'Building',
    'Site'
  ]
</script>

{#if isPopUpFormReady}
  <PopUpForms
    bind:this={popUpForms}
    heading="Add Facility"
    FIELDS={fields}
    REFS={{
      'facilityType': FacilityTypes
    }}
    bind:isSubmitted={isSubmitAddFacility}
    OnSubmit={OnAddFacility}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-facility">
    <h2>Facility Management</h2>
    <MasterTable
      NAME="Facility"
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={facilities}
      REFS={{
        'facilityType': FacilityTypes
      }}

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
      title="add facility"
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
  .master-facility {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-facility h2 {
    margin: 0;
  }
</style>