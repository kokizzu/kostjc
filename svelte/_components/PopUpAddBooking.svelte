<script>
  /** @typedef {import('../_types/property').Booking} Booking */
  /** @typedef {import('../_types/property').Facility} Facility */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { arrToArrNum, dateISOFormat } from './xFormatter';
  import MultiSelect from './MultiSelect.svelte';
  import { onMount } from 'svelte';
  import Select from '../node_modules/svelte-select';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let tenants = /** @type {Record<number, string>} */ ({});
  export let facilities = /** @type {Facility[]} */ ([]);
  export let rooms = /** @type {Record<number, string>} */ ({});

  let dateStart = dateISOFormat(0);
  let dateEnd = dateISOFormat(30);
  let basePriceIDR = 0;
  let facilitiesIds = /** @type {number[]} */([]);
  let totalPriceIDR = 0;
  let paidAt = dateISOFormat(0);
  let tenantId = null;
  let extraTenantsIds = /** @type {number[]} */([]);
  let roomId = null;

  export let OnSubmit = async function(/** @type {Booking} */ booking, /** @type {number[]} */ facilities) {
    console.log('OnSubmit :::', booking, facilities);
  }

  let facilitiesPrice = 0;

  async function submitAdd() {
    totalPriceIDR = basePriceIDR + facilitiesPrice;
    const booking = /** @type {Booking|any} */ ({
      dateStart,
      dateEnd,
      basePriceIDR,
      totalPriceIDR,
      paidAt,
      tenantId,
      extraTenants: extraTenantsIds,
      roomId
    });

    facilitiesIds = arrToArrNum(facilitiesIds);
    await OnSubmit(booking, facilitiesIds);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    dateStart = dateISOFormat(0);
    dateEnd = dateISOFormat(30);
    basePriceIDR = 0;
    facilitiesIds = [];
    totalPriceIDR = 0;
    paidAt = dateISOFormat(0);
    tenantId = 0;
    extraTenantsIds = [];
    roomId = 0;
    velueSvelteSelectFacilities = [];
    facilitiesPrice = 0;
  }
  
  const cancel = () => isShow = false;

  /**
   * @typedef {Object} SvelteSelectValue
   * @prop {any} value
   * @prop {any} label
   */
  let valuesSvelteSelectFacilities = /** @type {SvelteSelectValue[]} */ ([]);
  let velueSvelteSelectFacilities = [];
  let isFacilitiesReady = false;

  function handleSelectFacilities(/** @type {CustomEvent} */e) {
    console.log('handleSelectFacilities:',e.detail)
    facilitiesIds = [...facilitiesIds, e.detail.value];
    facilities.forEach((f) => {
      if (f.id === e.detail.value) {
        facilitiesPrice += f.extraChargeIDR || 0;
      }
    })
  }

  function handleClearFacilities(/** @type {CustomEvent} */e) {
    const index = facilitiesIds.indexOf(Number(e.detail.value));
    if (index !== -1) {
      facilitiesIds.splice(index, 1);
    }
    facilities.forEach((f) => {
      if (f.id === e.detail.value) {
        facilitiesPrice -= f.extraChargeIDR || 0;
      }
    })
  }

  onMount(() => {
    (facilities || []).forEach((f) => {
      valuesSvelteSelectFacilities = [...valuesSvelteSelectFacilities, {
        value: f.id,
        label: `${f.facilityName} ${f.facilityType ? `(${f.facilityType})` : ''}`
      }]
    });
    isFacilitiesReady = true;
  });
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Booking</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="dateStart"
        label="Date Start"
        bind:value={dateStart}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
      <InputBox
        id="dateEnd"
        label="Date End"
        bind:value={dateEnd}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
      <InputBox
        id="basePriceIDR"
        label="Base Price IDR"
        bind:value={basePriceIDR}
        type="number"
        placeholder="0"
      />
      {#if isFacilitiesReady}
      <div class="input-box">
        <label class="label" for="facilities">Facilities</label>
        <Select
          multiple
          items={valuesSvelteSelectFacilities}
          bind:value={velueSvelteSelectFacilities}
          on:select={handleSelectFacilities}
          on:clear={handleClearFacilities}
          placeholder="Facilities"
        />
      </div>
      {/if}
      <div class="input-totalprice">
        <label class="label" for="totalPriceIDR">Total Price</label>
        <div class="totalprice">
          <span class="baseprice">{basePriceIDR} + </span>
          <span class="facilityprice">{facilitiesPrice}</span>
          <span class="total"> = {basePriceIDR + facilitiesPrice}</span>
        </div>
      </div>
      <InputBox
        id="paidAt"
        label="Paid At"
        bind:value={paidAt}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
      <InputBox
        id="tenantId"
        label="Tenant"
        bind:value={tenantId}
        type="combobox"
        values={tenants}
        isObject={true}
      />
      <MultiSelect
        id="extraTenants"
        label="Extra Tenants"
        bind:valuesTarget={extraTenantsIds}
        valuesSourceObj={tenants}
        valuesSourceType="object"
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
  :global(.rotate-b) {
    transition: all 0.2s ease-in-out;
    transform: rotate(90deg);
  }

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(180deg);
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

  .input-totalprice {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input-totalprice label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .input-totalprice .totalprice {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: var(--gray-001);
    padding: 10px 12px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .input-box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input-box .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  @media only screen and (max-width : 768px) {
    .popup-container {
      padding: 10px;
    }
  }
</style>