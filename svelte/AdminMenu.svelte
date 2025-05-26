<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/cafe.js').Menu} Menu */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminMenu } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine, RiDocumentFileCodeLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';
  import Highlight from 'svelte-highlight/Highlight.svelte';
  import LineNumbers from  'svelte-highlight/LineNumbers.svelte'; 
  import json from 'svelte-highlight/languages/json';
  import atomOneDark from 'svelte-highlight/styles/atom-one-dark';
  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';
  import PopUpAddMenu from './_components/PopUpAddMenu.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let menu      = /** @type {Menu} */ ({/* menu */});
  let menus     = /** @type {any[][]} */([/* menus */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddMenu |any
  } */ (null);
  let isSubmitAddMenu = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true;
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminMenu( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminMenuCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddMenu = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        menus = o.menus;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      menu: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminMenu(i,
      /** @type {import('./jsApi.GEN').AdminMenuCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        menus = o.menus;
        notifier.showSuccess(`Menu '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      menu: {
        id: row[0]
      },
      cmd: CmdDelete
    });
    await AdminMenu(i,
      /** @type {import('./jsApi.GEN').AdminMenuCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        menus = o.menus;
        notifier.showSuccess(`Menu '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const menu = {
      id: payloads[0],
      name: String(payloads[1]),
      hppIDR: Number(payloads[2]),
      salePriceIDR: Number(payloads[3]),
      detail: String(payloads[4]),
      imageUrl: String(payloads[5])
    }

    const i = /** @type {any}*/ ({
      pager,
      menu,
      cmd: CmdUpsert
    });
    await AdminMenu(i,
      /** @type {import('./jsApi.GEN').AdminMenuCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        menus = o.menus;
        notifier.showSuccess(`Menu '${menu.roomName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddMenu(/** @type {Menu} */ menu) {
    isSubmitAddMenu = true;
    const i = /** @type {any} */ ({
      pager: pager,
      menu: menu,
      cmd: CmdUpsert
    });

    await AdminMenu(i,
      /** @type {import('../jsApi.GEN').AdminMenuCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddMenu = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        menus = o.menus;
        notifier.showSuccess(`Menu '${menu.name}' created !!`);

        OnRefresh(pager);
        popUpForms.Hide();
      }
    );
  }

  let isShowPopUpExportMenu = false;
  let menusObjContainer;
  const hidePopUpExportMenu = () => {
    isShowPopUpExportMenu = false;
  }

  let menusObjJson = '';
  const showPopUpExportRoom = () => {
    const firstChild = menusObjContainer.firstElementChild;
    firstChild.style.width = '100%';
    let menusObj = [];
    (menus || []).forEach((r) => {
      let menuObj = {};
      (fields || []).forEach((f, fId) => {
        switch (f.name) {
          case 'name':
            menuObj.name = r[fId];
          case 'hppIDR':
            menuObj.hppIDR = r[fId];
          case 'salePriceIDR':
            menuObj.salePriceIDR = r[fId];
          case 'detail':
            menuObj.detail = r[fId];
          case 'imageUrl':
            menuObj.image_url = r[fId];
        }
      });
      menusObj = [...menusObj, menuObj];
    })
    console.log('Menus Objects: ', menusObj);
    menusObjJson = JSON.stringify(menusObj, null, 2);
    isShowPopUpExportMenu = true;
  }
  const exportRoomToJson = () => {
    const blob = new Blob([menusObjJson], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'menus.json';
    a.click();
    URL.revokeObjectURL(url);
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddMenu
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddMenu}
    OnSubmit={OnAddMenu}
  />
{/if}

<div class={`popup-container ${isShowPopUpExportMenu ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Export Room</h2>
      <button on:click={hidePopUpExportMenu}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <div class="rooms-object-container" bind:this={menusObjContainer}>
        <Highlight language={json} code={menusObjJson} let:highlighted>
          <LineNumbers {highlighted}/>
        </Highlight>
      </div>
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={hidePopUpExportMenu}>Cancel</button>
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
    <h2>Menu Management</h2>
    <MasterTable
      NAME="Menu"
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={menus}

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