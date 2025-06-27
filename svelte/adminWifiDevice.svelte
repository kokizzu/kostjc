<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').WifiDevice} WifiDevice */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminWifiDevice } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';
  import PopUpAddWifiDevice from './_components/PopUpAddWifiDevice.svelte';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let wifiDevice    = /** @type {WifiDevice} */ ({/* wifiDevice */});
  let wifiDevices  = /** @type {any[][]} */([/* wifiDevices */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});
  const rooms  = /** @type {Record<Number, string>} */ ({/* rooms */});
  const tenants = /** @type {Record<Number, string>} */ ({/* tenants */});

  let isPopUpAddWifiDeviceReady = /** @type boolean */ (false);
  let popUpWifiDevice = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddWifiDevice |any
  } */ (null);
  let isSubmitAddWifiDevice = /** @type boolean */ (false);

  onMount(() => isPopUpAddWifiDeviceReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminWifiDevice( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminWifiDeviceCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddWifiDevice = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        wifiDevices = o.wifiDevices;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      wifiDevice: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminWifiDevice(i,
      /** @type {import('./jsApi.GEN').AdminWifiDeviceCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        wifiDevices = o.wifiDevices;
        notifier.showSuccess(`Wifi Device #${row[0]} restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      wifiDevice: {
        id: row[0]
      },
      cmd: CmdDelete
    });
    await AdminWifiDevice(i,
      /** @type {import('./jsApi.GEN').AdminWifiDeviceCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        wifiDevices = o.wifiDevices;
        notifier.showSuccess(`Wifi Device #${row[0]} deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const wifiDevice = {
      id: id
    }
    const i = /** @type {any}*/ ({
      pager,
      wifiDevice,
      cmd: CmdUpsert
    });
    await AdminWifiDevice(i,
      /** @type {import('./jsApi.GEN').AdminWifiDeviceCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        wifiDevices = o.wifiDevices;
        notifier.showSuccess(`Wifi Device #${wifiDevice.id} updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAdd(/** @type {WifiDevice} */ wifiDevice) {
    isSubmitAddWifiDevice = true;
    const i = /** @type {any} */ ({
      pager,
      wifiDevice,
      cmd: CmdUpsert
    });

    await AdminWifiDevice(i,
      /** @type {import('../jsApi.GEN').AdminWifiDeviceCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddWifiDevice = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        wifiDevices = o.wifiDevices;
        notifier.showSuccess(`Wifi Device created !!`);
        
        OnRefresh(pager);
        popUpWifiDevice.Reset();
        popUpWifiDevice.Hide();
      }
    );
  }
</script>

{#if isPopUpAddWifiDeviceReady}
  <PopUpAddWifiDevice
    bind:this={popUpWifiDevice}
    bind:isSubmitted={isSubmitAddWifiDevice}
    rooms={rooms}
    tenants={tenants}
    OnSubmit={OnAdd}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-wifidevice">
    <h2>Wifi Device Management</h2>
    <MasterTable
      NAME="Wifi Device"
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={wifiDevices}
      REFS={{
        'tenantId': tenants,
        'roomId': rooms
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
      on:click={() => popUpWifiDevice.Show()}
      title="Add wifi device"
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
  .master-wifidevice {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-wifidevice h2 {
    margin: 0;
  }
</style>