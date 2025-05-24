<script>
  /** @typedef {import('../_types/property').Booking} Booking */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import Select from '../node_modules/svelte-select';
  import { AdminBooking } from '../jsApi.GEN';

  export let isShow = /** @type {boolean} */ (false);
  export let isSubmitted  = /** @type {boolean} */ (false);
  export let tenants = /** @type {Record<number, string>} */ ({});
  export let rooms = /** @type {Record<number, string>} */ ({});

  export let bookingId = /** @type {number} */ (0);

  export let OnSubmit = async function(/** @type {Booking} */ booking) {
    console.log('OnSubmit :::', booking);
  }

  let roomId;
  let dateStart;
  let dateEnd;
  let tenantId;
  let basePriceIDR;
  let facilitiesObj;
  let totalPriceIDR;
  let paidAt;
  let extraTenants;

  export const Show = () => isShow = true;

  export async function GetBookingById(id) {
    console.log('GetBookingById id: ', id);
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

      console.log('Res: ',o);
    })
  }
  export const Hide = () => isShow = false;
  
  const cancel = () => isShow = false;
  
  function handleSubmit() {
    isSubmitted = true;
    OnSubmit();
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
      <InputBox
        bind:value={roomId}
        id="roomId"
        label="Room"
        type="combobox"
        values={rooms}
      />
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
      <InputBox
        bind:value={tenantId}
        id="tenantId"
        label="Tenant"
        type="combobox"
        values={tenants}
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