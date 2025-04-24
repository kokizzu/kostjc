<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Room} Room */

  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminRoom } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine, RiDocumentFileCodeLine } from './node_modules/svelte-icons-pack/dist/ri';
  import PopUpAddRoom from './_components/PopUpAddRoom.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let room  = /** @type {Room} */ ({/* room */});
  let rooms = /** @type {any[][]} */([/* rooms */]);
  let tenants = /** @type {Record<Number, string>} */({/* tenants */});
  let buildings = /** @type {Record<Number, string>} */({/* buildings */});
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddRoom |any
  } */ (null);
  let isSubmitAddRoom = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true;
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await AdminRoom( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminRoomCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddRoom = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        rooms = o.rooms;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      room: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await AdminRoom(i,
      /** @type {import('./jsApi.GEN').AdminRoomCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        rooms = o.rooms;
        notifier.showSuccess(`Room '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      room: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminRoom(i,
      /** @type {import('./jsApi.GEN').AdminRoomCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        rooms = o.rooms;
        notifier.showSuccess(`Room '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    console.log('Room ID to Edit: ' + String(id));
    const room = {
      id: payloads[0],
      roomName: String(payloads[1]),
      basePriceIDR: Number(payloads[2]),
      currentTenantId: payloads[3],
      buildingId: payloads[4],
      firstUseAt: String(payloads[5]),
      lastUseAt: String(payloads[6])
    }
    const i = /** @type {any}*/ ({
      pager,
      room,
      cmd: 'upsert'
    });
    await AdminRoom(i,
      /** @type {import('./jsApi.GEN').AdminRoomCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        rooms = o.rooms;
        notifier.showSuccess(`Room '${room.buildingName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddRoom(/** @type {Room} */ room) {
    isSubmitAddRoom = true;
    const i = /** @type {any} */ ({
      pager,
      room,
      cmd: 'upsert'
    });

    await AdminRoom(i,
      /** @type {import('../jsApi.GEN').AdminRoomCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddRoom = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        rooms = o.rooms;
        notifier.showSuccess(`Room '${room.roomName}' created !!`);

        popUpForms.Reset();

        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddRoom
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddRoom}
    tenants={tenants}
    buildings={buildings}
    OnSubmit={OnAddRoom}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-room">
    <h2>Room Management</h2>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'currentTenantId': tenants,
        'buildingId': buildings
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={rooms}

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
      title="Add Room"
    >
      <Icon
        color="var(--gray-007)"
        size="18"
        src={RiSystemAddBoxLine}
      />
    </button>
    <button
      class="btn"
      title="Export to JSON"
    >
      <Icon
        color="var(--gray-007)"
        size="18"
        src={RiDocumentFileCodeLine}
      />
    </button>
    </MasterTable>
  </div>
</LayoutMain>

<style>
  .master-room {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-room h2 {
    margin: 0;
  }
</style>