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
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';
  import PopUpAddRoom from './_components/PopUpAddRoom.svelte';
  import Highlight from 'svelte-highlight/Highlight.svelte';
  import LineNumbers from  'svelte-highlight/LineNumbers.svelte'; 
  import json from 'svelte-highlight/languages/json';
  import atomOneDark from 'svelte-highlight/styles/atom-one-dark';
  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let room      = /** @type {Room} */ ({/* room */});
  let rooms     = /** @type {any[][]} */([/* rooms */]);
  let tenants   = /** @type {Record<Number, string>} */({/* tenants */});
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
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
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
      cmd: CmdRestore
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
      cmd: CmdDelete
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
    const room = {
      id: payloads[0],
      roomName: String(payloads[1]),
      roomSize: String(payloads[2]),
      basePriceIDR: Number(payloads[3]),
      currentTenantId: payloads[4],
      buildingId: payloads[5],
      firstUseAt: String(payloads[6]),
      lastUseAt: String(payloads[7]),
      imageUrl: String(payloads[8])
    }
    const i = /** @type {any}*/ ({
      pager,
      room,
      cmd: CmdUpsert
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
        notifier.showSuccess(`Room '${room.roomName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddRoom(/** @type {Room} */ room) {
    isSubmitAddRoom = true;
    const i = /** @type {any} */ ({
      pager,
      room,
      cmd: CmdUpsert
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

        OnRefresh(pager);
        popUpForms.Hide();
      }
    );
  }

  let isShowPopUpExportRoom = false;
  let roomsObjContainer;
  const hidePopUpExportRoom = () => {
    isShowPopUpExportRoom = false;
  }

  let roomsObjJson = '';
  const showPopUpExportRoom = () => {
    const firstChild = roomsObjContainer.firstElementChild;
    firstChild.style.width = '100%';
    let roomsObj = [];
    (rooms || []).forEach((r) => {
      let roomObj = {};
      (fields || []).forEach((f, fId) => {
        switch (f.name) {
          case 'roomName':
            roomObj.name = r[fId];
          case 'roomSize':
            roomObj.size = r[fId];
          case 'basePriceIDR':
            roomObj.normalPrice = r[fId];
          case 'lastUseAt':
            roomObj.availableAt = r[fId];
          case 'imageUrl':
            roomObj.image_url = r[fId];
        }
      });
      roomsObj = [...roomsObj, roomObj];
    })
    console.log('Rooms Objects: ', roomsObj);
    roomsObjJson = JSON.stringify(roomsObj, null, 2);
    isShowPopUpExportRoom = true;
  }
  const exportRoomToJson = () => {
    const blob = new Blob([roomsObjJson], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'rooms.json';
    a.click();
    URL.revokeObjectURL(url);
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

<div class={`popup-container ${isShowPopUpExportRoom ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Export Room</h2>
      <button on:click={hidePopUpExportRoom}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <div class="rooms-object-container" bind:this={roomsObjContainer}>
        <Highlight language={json} code={roomsObjJson} let:highlighted>
          <LineNumbers {highlighted}/>
        </Highlight>
      </div>
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={hidePopUpExportRoom}>Cancel</button>
        <button class="ok" on:click|preventDefault={exportRoomToJson}>
            <Icon color="#FFF" size="16" src={RiDocumentFileCodeLine} />
            <span>Export</span>
        </button>
      </div>
    </div>
  </div>
</div>

<svelte:head>
  {@html atomOneDark}
</svelte:head>

<LayoutMain access={segments} user={user}>
  <div class="master-room">
    <h2>Room Management</h2>
    <MasterTable
      NAME="Room"
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
      on:click={showPopUpExportRoom}
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

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .popup-container {
    display: none;
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup-container.show {
    display: flex;
  }

	.popup-container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 600px;
		display: flex;
		flex-direction: column;
	}

  .popup-container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 15px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup-container .popup header h2 {
		margin: 0;
	}

	.popup-container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup-container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup-container .popup header button:active {
		background-color: #ef444430;
	}

	.popup-container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup-container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 15px 20px;
		border-top: 1px solid var(--gray-004);
	}

  .popup-container .popup .foot .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

	.popup-container .popup .foot button {
		padding: 8px 18px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

	.popup-container .popup .foot button.ok {
		background-color: var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 10px;
	}

	.popup-container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup-container .popup .foot button.ok:disabled {
		cursor: not-allowed;
		background-color: var(--gray-003);
		color: var(--gray-007);
	}

	.popup-container .popup .foot button.cancel {
		background-color: transparent;
    color: var(--gray-007);
	}

	.popup-container .popup .foot button.cancel:hover {
		background-color: var(--gray-001);
	}

  .facilities-form {
    display: flex;
    flex-direction: column;
    gap: 3px;
  }

  .rooms-object-container {
    display: flex;
    width: 100%;
    height: 400px;
    overflow-y: auto;
    background-color: #282c34;
    /* color: var(--green-006); */
    padding: 10px 2px 10px 0;
    border-radius: 10px;
  }

  .object-highlighter {
    width: 100%;
  }

  @media only screen and (max-width : 768px) {
    .popup-container {
      padding: 10px;
    }
  }
</style>