<script>
  /** @typedef {import('../_types/property').Room} Room */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './xFormatter';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let buildings = /** @type {Record<number, string>} */ ({});
  export let tenants = /** @type {Record<number, string>} */ ({});

  let roomName = '';
  let basePriceIDR = 0;
  let currentTenantId = 0;
  let buildingId = 0;
  let firstUseAt = dateISOFormat(0);
  let lastUseAt = dateISOFormat(30);
  let roomSize = '';
  let imageUrl = '';

  export let OnSubmit = async function(/** @type {Room} */ room) {
    console.log('OnSubmit :::', room);
  }

  async function submitAdd() {
    const room = /** @type {Room|any} */ ({
      roomName,
      basePriceIDR,
      currentTenantId,
      buildingId,
      firstUseAt,
      lastUseAt,
      roomSize,
      imageUrl
    });

    await OnSubmit(room);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    roomName = '';
    basePriceIDR = 0;
    currentTenantId = 0;
    buildingId = 0;
    firstUseAt = dateISOFormat(0);
    lastUseAt = dateISOFormat(30);
    roomSize = '';
    imageUrl = '';
  }
  
  const cancel = () => {
    isShow = false;
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Room</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="roomName"
        label="Room Name"
        bind:value={roomName}
        type="text"
        placeholder="D13"
      />
      <InputBox
        id="basePriceIDR"
        label="Base Price"
        bind:value={basePriceIDR}
        type="number"
        placeholder="0"
      />
      <InputBox
        id="currentTenantId"
        label="Current Tenant"
        bind:value={currentTenantId}
        type="combobox"
        values={tenants}
        isObject={true}
      />
      <InputBox
        id="buildingId"
        label="Building"
        bind:value={buildingId}
        type="combobox"
        values={buildings}
        isObject={true}
      />
      <InputBox
        id="roomSize"
        label="Room Size"
        bind:value={roomSize}
        type="text"
        placeholder="2x2.5"
      />
      <InputBox
        id="imageUrl"
        label="Image URL"
        bind:value={imageUrl}
        type="text"
        placeholder="img-D13.jpg"
      />
      <InputBox
        id="firstUseAt"
        label="First Use At"
        bind:value={firstUseAt}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
      <InputBox
        id="lastUseAt"
        label="Last Use At"
        bind:value={lastUseAt}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button class="ok" on:click|preventDefault={submitAdd} disabled={isSubmitted}>
          {#if !isSubmitted}
            <span>Submit</span>
          {/if}
          {#if isSubmitted}
            <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
          {/if}
        </button>
      </div>
    </div>
  </div>
</div>

<style>
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
		width: 500px;
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

  @media only screen and (max-width : 768px) {
    .popup-container {
      padding: 10px;
    }
  }
</style>