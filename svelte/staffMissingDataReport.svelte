<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/users.js').Tenant} Tenant */
  /** @typedef {import('./_types/property.js').MissingTenantData} MissingTenantData */

  import LayoutMain from './_layouts/main.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiDesignBallPenLine } from './node_modules/svelte-icons-pack/dist/ri';
  import PopUpEditMissingTenant from './_components/PopUpEditMissingTenant.svelte';
  import { AdminTenants, StaffMissingDataReport } from './jsApi.GEN';
  import { CmdUpsert } from './_components/xConstant';
  import { notifier } from './_components/xNotifier';
  import Radio from './_components/Radio.svelte';
  import { onMount } from 'svelte';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let missingData = /** @type {MissingTenantData[]} */ ([/* missingData */]);

  async function RefreshData() {
    await StaffMissingDataReport({},
    /** @type {import('./jsApi.GEN').StaffMissingDataReportCallback} */
    /** @returns {Promise<void>} */
    function(/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        notifier.showError(o.error || 'something went wrong');
        return
      }

      missingData = o.missingData;
      restrucureMissingData();

      return
    })
  }

  let popUpEditMissingTenant = /** @type {import('svelte').SvelteComponent | HTMLElement | PopUpEditMissingTenant | any} */ (null);
  let isSubmitEditMissingTenant = /** @type {boolean} */ (false);
  let isPopUpFormReady = /** @type {boolean} */ (false);

  async function  showPopUpEditTenant(/** @type {number} */ tenantId) {
    popUpEditMissingTenant.Show();
    popUpEditMissingTenant.GetTenantById(tenantId);
  }

  async function submitEditMissingTenant(/** @type {Tenant} */ tenant) {
    isSubmitEditMissingTenant = true;
    tenant.id = tenant.id+'';
    await AdminTenants({ // @ts-ignore
      tenant: tenant,
      cmd: CmdUpsert
    }, 
    /** @type {import('./jsApi.GEN').AdminTenantsCallback} */
    /** @returns {Promise<void>} */
    function(/** @type any */ o) {
      isSubmitEditMissingTenant = false;
      if (o.error) {
        console.log(o);
        notifier.showError(o.error || `failed to update tenant #${tenant.id}`);
        return
      }

      notifier.showSuccess(`Tenant '${tenant.tenantName}' updated !!`);

      popUpEditMissingTenant.Hide();
      RefreshData();
    });
  }
  
  /** @typedef {'showMissingAny' | 'showOnlyMissingTelegram' | 'showOnlyMissingWhatsapp' | 'showOnlyCompleted' | 'showAll'} FilterValues */
  /** @typedef {import('./_types/masters.js').RadioOption<FilterValues>} RadioOption<FilterValues> */

  let selectedFilter = /** @type {FilterValues} */ ('showAll');
  const filterName = /** @type {string} */ ('show-filter');
  const filterOptions = /** @type {RadioOption[]} */ ([
    {
      id: 'showMissingAny',
      name: filterName,
      value: 'showMissingAny',
      label: 'Show Missing Any'
    },
    {
      id: 'showOnlyMissingTelegram',
      name: filterName,
      value: 'showOnlyMissingTelegram',
      label: 'Show Only Missing Telegram'
    },
    {
      id: 'showOnlyMissingWhatsapp',
      name: filterName,
      value: 'showOnlyMissingWhatsapp',
      label: 'Show Only Missing Whatsapp'
    },
    {
      id: 'showOnlyCompleted',
      name: filterName,
      value: 'showOnlyCompleted',
      label: 'Show Only Completed'
    },
    {
      id: 'showAll',
      name: filterName,
      value: 'showAll',
      label: 'Show All'
    }
  ]);

  let missingDataFiltered = /** @type {MissingTenantData[]} */ ([]);

  function restrucureMissingData() {
    missingDataFiltered = [];
    console.log('missingDataFiltered [BEFORE]', missingDataFiltered);
    for (const dt of (missingData || [])) {
      switch (selectedFilter) {
        case 'showAll': {
          missingDataFiltered = [...missingDataFiltered, dt];
          break;
        }
        case 'showMissingAny': {
          if (!dt.tenantTelegramUsername || !dt.tenantWhatsappNumber) {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
        case 'showOnlyCompleted': {
          if (dt.tenantTelegramUsername != '' && dt.tenantWhatsappNumber != '') {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
        case 'showOnlyMissingTelegram': {
          if (!dt.tenantTelegramUsername) {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
        case 'showOnlyMissingWhatsapp': {
          if (!dt.tenantWhatsappNumber) {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
      }
    }
  }

  $: if (selectedFilter) {
    restrucureMissingData();
  }

  onMount(() => {
    isPopUpFormReady = true;
    restrucureMissingData();
  });
</script>

{#if isPopUpFormReady}
  <PopUpEditMissingTenant
    bind:this={popUpEditMissingTenant}
    bind:isSubmitted={isSubmitEditMissingTenant}
    OnSubmit={submitEditMissingTenant}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <Radio
      className="filters"
      options={filterOptions}
      bind:selected={selectedFilter}
    />
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Actions</th>
            <th>Room</th>
            <th style="min-width: 170px;">Tenant Name</th>
            <th>Telegram</th>
            <th style="min-width: 180px;">WhatsApp</th>
            <th style="min-width: 140px;">Last Use At</th>
          </tr>
        </thead>
        <tbody>
          {#each (missingDataFiltered || []) as data}
            <tr>
              <td>
                <div class="actions">
                  <button
                    class="btn"
                    title="Edit"
                    on:click={() => showPopUpEditTenant(data.tenantId)}
                  >
                    <Icon
                      size="15"
                      color="var(--gray-007)"
                      src={RiDesignBallPenLine}
                    />
                  </button>
                </div>
              </td>
              <td>{data.roomName}</td>
              <td>{data.tenantName || '--'}</td>
              <td>{data.tenantTelegramUsername || '--'}</td>
              <td>{data.tenantWhatsappNumber || '--'}</td>
              <td>{data.lastUseAt || '--'}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</LayoutMain>

<style>
  .report-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
    padding: 20px;
  }

  :global(.filters) {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 10px;
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

  table tbody tr td {
    padding: 8px 12px;
  }

  table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  table tbody tr td .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

  :global(table tbody tr td .actions .btn:hover svg) {
    fill: var(--blue-005);
  }

  @media only screen and (max-width : 768px) {

  }
</style>