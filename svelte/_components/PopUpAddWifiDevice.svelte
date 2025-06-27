<script>
  /** @typedef {import('../_types/property').WifiDevice} WifiDevice */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './xFormatter';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let rooms = /** @type {Record<number, string>} */ ({});
  export let tenants = /** @type {Record<number, string>} */ ({});

  let startAt = dateISOFormat(0);
  let endAt = dateISOFormat(30);
  let paidAt = dateISOFormat(0);
  let priceIDR = 0;
  let macAddress = '';
  let tenantId = 0;
  let roomId = 0

  export let OnSubmit = async function(/** @type {WifiDevice} */ wifiDevice) {
    console.log('OnSubmit :::', wifiDevice);
  }

  async function submitAdd() {
    const room = /** @type {WifiDevice|any} */ ({
      startAt,
      endAt,
      paidAt,
      priceIDR,
      macAddress,
      tenantId,
      roomId
    });

    await OnSubmit(room);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    startAt = dateISOFormat(0);
    endAt = dateISOFormat(30);
    paidAt = dateISOFormat(0);
    priceIDR = 0;
    macAddress = '';
    tenantId = 0;
    roomId = 0
  }
  
  const cancel = () => {
    isShow = false;
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Wifi Device</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="startAt"
        label="Start At"
        bind:value={startAt}
        type="datetime"
      />
      <InputBox
        id="endAt"
        label="End At"
        bind:value={endAt}
        type="datetime"
      />
      <InputBox
        id="paidAt"
        label="Paid At"
        bind:value={paidAt}
        type="datetime"
      />
      <InputBox
        id="priceIDR"
        label="Price IDR"
        bind:value={priceIDR}
        type="number"
      />
      <InputBox
        id="macAddress"
        label="Mac Address"
        bind:value={macAddress}
        type="text"
      />
      <InputBox
        id="tenantId"
        label="Tenant"
        bind:value={tenantId}
        type="combobox"
        values={tenants}
        isObject={true}
      />
      <InputBox
        id="roomId"
        label="Room"
        bind:value={roomId}
        type="combobox"
        values={rooms}
        isObject={true}
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