<script>
  /** @typedef {import('../_types/property').Booking} Booking */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import Select from '../node_modules/svelte-select';
  import { AdminBooking } from '../jsApi.GEN';
  import MultiSelect from './MultiSelect.svelte';
    import { dateISOFormat } from './xFormatter';

  export let isShow = /** @type {boolean} */ (false);
  export let isSubmitted  = /** @type {boolean} */ (false);
  export let tenants = /** @type {Record<number, string>} */ ({});
  export let rooms = /** @type {Record<number, string>} */ ({});

  export let bookingId = /** @type {number} */ (0);

  export let OnSubmit = async function(/** @type {Booking} */ booking) {
    console.log('OnSubmit :::', booking);
  }

  let roomId = 0;
  let dateStart;
  let dateEnd;
  let tenantId = 0;
  let basePriceIDR;
  let facilitiesObj;
  let totalPriceIDR;
  let paidAt;
  let extraTenants;

  export const Show = () => isShow = true;

  /**
   * @description 
   * @param {number} id
   * @returns {Promise<void>}
   */
  export async function GetBookingById(id) {
    await AdminBooking({
      cmd: 'form', // @ts-ignore
      booking: {
        id: id+''
      }
    }, function (/** @type {any} */ o) {
      if (o.error) {
        return;
      }
      const booking = o.booking;

      roomId = booking.roomId;
      dateStart = booking.dateStart;
      dateEnd = booking.dateEnd;
      tenantId = booking.tenantId;
      basePriceIDR = booking.basePriceIDR;
      facilitiesObj = booking.facilitiesObj;
      totalPriceIDR = booking.totalPriceIDR;
      paidAt = booking.paidAt;
      extraTenants = booking.extraTenants;

      reStructureSvelteValues(tenantId, roomId);

      console.log('Res: ',o);
    })
  }
  export const Hide = () => {
    isShow = false;
    reset();
  }

  function reset() {
    roomId = 0;
    dateStart = dateISOFormat(0);
    dateEnd = dateISOFormat(30);
    tenantId = 0;
    basePriceIDR = 0;
    facilitiesObj = {};
    totalPriceIDR = 0
    paidAt = dateISOFormat(0);
    extraTenants = [];
  }
  
  const cancel = () => {
    isShow = false;
    reset();
  }
  
  async function handleSubmit() {
    isSubmitted = true;
    // @ts-ignore
    const booking = /** @type {Booking} */ ({
      id: bookingId+'',
      roomId: roomId,
      dateStart: dateStart,
      dateEnd: dateEnd,
      tenantId: tenantId,
      basePriceIDR: basePriceIDR,
      facilitiesObj: facilitiesObj,
      totalPriceIDR: totalPriceIDR,
      paidAt: paidAt,
      extraTenants: extraTenants
    });
    await OnSubmit(booking);
  }

  let itemsArrObjTenant = /** @type {Record<string|number, string>[]} */ ([]);
  /**
   * @typedef {Object} SvelteSelectValue
   * @prop {any} value
   * @prop {any} label
   */
   let svelteSelectValueTenant = /** @type {SvelteSelectValue} */ ({
    label: '',
    value: ''
  });

  function handleSelectTenant(/** @type {CustomEvent} */e) {
    tenantId = e.detail.value;
  }

  function handleClearTenant(/** @type {CustomEvent} */e) {
    tenantId = 0;
  }

  let itemsArrObjRoom = /** @type {Record<string|number, string>[]} */ ([]);
  let svelteSelectValueRoom = /** @type {SvelteSelectValue} */ ({
    label: '',
    value: ''
  });

  function handleSelectRoom(/** @type {CustomEvent} */e) {
    roomId = e.detail.value;
  }

  function handleClearRoom(/** @type {CustomEvent} */e) {
    roomId = 0;
  }

  /**
   * @description make sure it can use svelte-select 
   * @param {number} tenantId
   * @param {number} roomId
   */
  function reStructureSvelteValues(tenantId, roomId) {
    for (const [k, v] of Object.entries(rooms)) {
      itemsArrObjRoom = [...itemsArrObjRoom, {
        value: k,
        label: v
      }];
      if (String(k) == String(roomId)) {
        svelteSelectValueRoom = {
          value: k,
          label: v
        };
      }
    }
    for (const [k, v] of Object.entries(tenants)) {
      itemsArrObjTenant = [...itemsArrObjTenant, {
        value: k,
        label: v
      }];
      if (String(k) == String(tenantId)) {
        svelteSelectValueTenant = {
          value: k,
          label: v
        };
      }
    }
  }

  $: if (extraTenants) {
    reStructureSvelteValues(tenantId, roomId);
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Edit Booking {`#${bookingId}`}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <div class="input-box">
        <label class="label" for="paymentMethod">Room</label>
        <Select
          items={itemsArrObjRoom}
          bind:value={svelteSelectValueRoom}
          on:select={handleSelectRoom}
          on:clear={handleClearRoom}
        />
      </div>
      <InputBox
        bind:value={dateStart}
        id="dateStart"
        label="Date Start"
        type="date"
      />
      <InputBox
        bind:value={dateEnd}
        id="dateEnd"
        label="Date End"
        type="date"
      />
      <div class="input-box">
        <label class="label" for="paymentMethod">Tenant</label>
        <Select
          items={itemsArrObjTenant}
          bind:value={svelteSelectValueTenant}
          on:select={handleSelectTenant}
          on:clear={handleClearTenant}
        />
      </div>
      <InputBox
        bind:value={basePriceIDR}
        id="basePriceIDR"
        label="Base Price"
        type="number"
      />
      <InputBox
        bind:value={facilitiesObj}
        id="facilitiesObj"
        label="Facilities"
        type="textarea"
      />
      <InputBox
        bind:value={totalPriceIDR}
        id="totalPriceIDR"
        label="Total Price"
        type="number"
      />
      <InputBox
        bind:value={paidAt}
        id="paidAt"
        label="Paid At"
        type="date"
      />
      <MultiSelect
        id="extraTenants"
        label="Extra Tenants"
        placeholder="Extra Tenants"
        bind:valuesTarget={extraTenants}
        valuesSourceObj={tenants}
        valuesSourceType="object"
      />
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button class="ok" on:click|preventDefault={handleSubmit} disabled={isSubmitted}>
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
		z-index: 2001;
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
</style>