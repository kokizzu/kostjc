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
  import {checkboxToDate} from './_components/xFormatter';
  import {CmdToggleWaAdded, CmdToggleTeleAdded} from './_components/xConstant';

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

  async function OnEditWaAddedAt(/** @type any */ id, /** @type any[]*/ payloads) {
    const i = /** @type {any}*/ ({
      tenant: {
        id: payloads[0]+'',
        waAddedAt: checkboxToDate(payloads[2]),
      },
      cmd: CmdToggleWaAdded
    });

    console.log('sending to AdminTenants', i);

    await AdminTenants(i,
      /** @type {import('./jsApi.GEN').AdminTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        notifier.showSuccess(`'${payloads[1]}' WaAddedAt updated !!`);
        RefreshData();
      }
    );
  }

  async function OnEditTeleAddedAt(/** @type any */ id, /** @type any[]*/ payloads) {
    const i = /** @type {any}*/ ({
      tenant: {
        id: payloads[0]+'',
        teleAddedAt: checkboxToDate(payloads[2]),
      },
      cmd: CmdToggleTeleAdded
    });

    console.log('sending to AdminTenants', i);

    await AdminTenants(i,
      /** @type {import('./jsApi.GEN').AdminTenantsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        notifier.showSuccess(`'${payloads[1]}' TeleAddedAt updated !!`);
        RefreshData();
      }
    );
  }

  /** @typedef {'showMissingAny' | 'showOnlyMissingTelegram' | 'showOnlyMissingWhatsapp' | 'showOnlyCompleted' | 'showAll' | 'showNotAddedToWhatsapp' | 'showNotAddedToTelegram'} FilterValues */
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
    },
    {
       id: 'showNotAddedToWhatsapp',
      name: filterName,
      value: 'showNotAddedToWhatsapp',
      label: 'Show Not Added To Whatsapp'
    },
    {
      id: 'showNotAddedToTelegram',
      name: filterName,
      value: 'showNotAddedToTelegram',
      label: 'Show Not Added To Telegram'
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
        case 'showNotAddedToWhatsapp': {
          if (!dt.tenantWaAddedAt) {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
         case 'showNotAddedToTelegram': {
          if (!dt.tenantTeleAddedAt) {
            missingDataFiltered = [...missingDataFiltered, dt];
          }
          break;
        }
      }
    }
    const collator = new Intl.Collator(undefined, { numeric: true, sensitivity: 'base' });
    missingDataFiltered.sort((a, b) => {
      const roomA = a.roomName?.toUpperCase() || '';
      const roomB = b.roomName?.toUpperCase() || '';
      return collator.compare(roomA, roomB);
    });
  }

  $: if (selectedFilter) {
    restrucureMissingData();
  }

  onMount(() => {
    isPopUpFormReady = true;
    restrucureMissingData();
  });

  function toggleCheckboxField(tenantId, tenantName, addedAt, fieldName) {
    let payloads = [tenantId, tenantName, addedAt];


    if (fieldName === 'waAddedAt' && OnEditWaAddedAt) {
      OnEditWaAddedAt(tenantId, payloads);
    }

    if (fieldName === 'teleAddedAt' && OnEditTeleAddedAt) {
      OnEditTeleAddedAt(tenantId, payloads);
    }
    
  }

  function handleToggleWaAddedAt(e, data) {
    const isChecked = /** @type {HTMLInputElement} */ (e.target).checked;
    toggleCheckboxField(data.tenantId, data.tenantName, isChecked, 'waAddedAt');
  }

  function handleToggleTeleAddedAt(e, data) {
    const isChecked = /** @type {HTMLInputElement} */ (e.target).checked;
    toggleCheckboxField(data.tenantId, data.tenantName, isChecked, 'teleAddedAt');
  }

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
            <th>Wa Added</th>
            <th>Tele Added</th>
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
              <td class="checkbox-cell">
                <input
                  type="checkbox"
                  checked={data.tenantWaAddedAt}
                  on:change={(e) => handleToggleWaAddedAt(e, data)}
                />
              </td>
              
              <td class="checkbox-cell">
                <input
                  type="checkbox"
                  checked={data.tenantTeleAddedAt}
                  on:change={(e) => handleToggleTeleAddedAt(e, data)}
                />
              </td>
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

  .checkbox-cell {
    text-align: center;
  }
  .checkbox-cell input[type="checkbox"] {
    transform: scale(1.2);
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