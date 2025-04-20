<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/users.js').Tenant} Tenant */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { UserTenants } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let tenant  = /** @type {Tenant} */ ({/* tenant */});
  let tenants = /** @type {any[][]} */([/* tenants */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitTenant = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await UserTenants( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').UserTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitTenant = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        tenants = o.tenants;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      tenant: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await UserTenants(i,
      /** @type {import('./jsApi.GEN').UserTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        tenants = o.tenants;
        notifier.showSuccess(`Tenant '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      tenant: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await UserTenants(i,
      /** @type {import('./jsApi.GEN').UserTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        tenants = o.tenants;
        notifier.showSuccess(`Tenant '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    console.log('Tenant ID to Edit: ' + String(id));
    const tenant = {
      id: payloads[0],
      tenantName: payloads[1],
      ktpRegion: payloads[2],
      ktpNumber: payloads[3],
      ktpName: payloads[4],
      ktpPlaceBirth: payloads[5],
      ktpDateBirth: payloads[6],
      ktpGender: payloads[7],
      ktpAddress: payloads[8],
      ktpRtRw: payloads[9],
      ktpKelurahanDesa: payloads[10],
      ktpKecamatan: payloads[11],
      ktpReligion: payloads[12],
      ktpMaritalStatus: payloads[13],
      ktpCitizenship: payloads[14],
      telegramUsername: payloads[15],
      whatsappNumber: payloads[16]
    }
    const i = /** @type {any}*/ ({
      pager,
      tenant,
      cmd: 'upsert'
    });
    await UserTenants(i,
      /** @type {import('./jsApi.GEN').UserTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        tenants = o.tenants;
        notifier.showSuccess(`Tenant '${tenant.facilityName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddTenant(/** @type any[] */ payloads) {
    isSubmitTenant = true;

    const tenant = /** @type {any} */ ({
      tenantName: payloads[1],
      ktpRegion: payloads[2],
      ktpNumber: payloads[3],
      ktpName: payloads[4],
      ktpPlaceBirth: payloads[5],
      ktpDateBirth: payloads[6],
      ktpGender: payloads[7],
      ktpAddress: payloads[8],
      ktpRtRw: payloads[9],
      ktpKelurahanDesa: payloads[10],
      ktpKecamatan: payloads[11],
      ktpReligion: payloads[12],
      ktpMaritalStatus: payloads[13],
      ktpCitizenship: payloads[14],
      telegramUsername: payloads[15],
      whatsappNumber: payloads[16]
    });
    const i = /** @type {any} */ ({
      pager,
      tenant,
      cmd: 'upsert'
    });

    await UserTenants(i,
      /** @type {import('../jsApi.GEN').UserTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitTenant = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        tenants = o.tenants;
        notifier.showSuccess(`Tenant '${tenant.facilityName}' created !!`);

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
    heading="Add Tenant"
    FIELDS={fields}
    bind:isSubmitted={isSubmitTenant}
    OnSubmit={OnAddTenant}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-tenants">
    <h2>Tenant Management</h2>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={tenants}

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
      title="add tenant"
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
  .master-tenants {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-tenants h2 {
    margin: 0;
  }
</style>