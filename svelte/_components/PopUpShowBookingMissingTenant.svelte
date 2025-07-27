<script>
  /** @typedef {import('../_types/property').PaymentOfBooking} PaymentOfBooking */
  /** @typedef {import('../_types/property').TenantBookingDetail} TenantBookingDetail */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';

  export let tenantName = '';
  export let bookings = /** @type {TenantBookingDetail[]} */ ([]);
  export let payments = /** @type {PaymentOfBooking[]} */ ([]);

  let isShow = false;

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  $: if (tenantName) {
    console.log('Bookings', bookings);
    console.log('Payments', payments);
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Booking History of {tenantName}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <h4>Cihuyy</h4>
    </div>
  </div>
</div>

<style>
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

  @media only screen and (max-width : 768px) {
    .popup-container {
      padding: 10px;
    }

    .popup-container .popup {
      width: 100%;
    }
  }
</style>