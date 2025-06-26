<script>
  /** @typedef {import('../_types/property').Payment} Payment */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './xFormatter';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let bookings = /** @type {Record<number, string>} */ ({});
  export let bookingId = 0;
  export let isBookingReadOnly = false;
  export let paymentAt = dateISOFormat(0);

  const PaymentMethods = [
    'Cash',
    'QRIS',
    'Transfer',
    'Donation'
  ];
  const PaymentStatuses = [
    'Paid',
    'Unpaid',
    'Refunded'
  ];

  let paidIDR = 0;
  let paymentMethod = PaymentMethods[0];
  let paymentStatus = PaymentStatuses[0];
  let note = '';

  export let OnSubmit = async function(/** @type {Payment} */ payment) {
    console.log('OnSubmit :::', payment);
  }

  async function submitAdd() {
    const payment = /** @type {Payment|any} */ ({
      bookingId: bookingId,
      paymentAt: paymentAt,
      paidIDR: paidIDR,
      paymentMethod: paymentMethod,
      paymentStatus: paymentStatus,
      note: note
    });

    await OnSubmit(payment);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    bookingId = 0;
    paymentAt = dateISOFormat(0);
    paidIDR = 0;
    paymentMethod = PaymentMethods[0];
    paymentStatus = PaymentStatuses[0];
    note = '';
  }
  
  const cancel = () => {
    isShow = false;
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Payment {isBookingReadOnly ? `for Booking #${bookingId}` : ''}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      {#if !isBookingReadOnly}
        <InputBox
          id="booking"
          label="Booking"
          isObject={true}
          bind:value={bookingId}
          type="combobox"
          values={bookings}
        />
      {/if}
      <InputBox
        id="paymentAt"
        label="Payment At"
        bind:value={paymentAt}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
      <InputBox
        id="paidIDR"
        label="Total Paid"
        bind:value={paidIDR}
        type="number"
        placeholder="0"
      />
      <InputBox
        id="paymentMethod"
        label="Payment Method"
        bind:value={paymentMethod}
        type="combobox-arr"
        values={PaymentMethods}
      />
      <InputBox
        id="paymentStatus"
        label="Payment Status"
        bind:value={paymentStatus}
        type="combobox-arr"
        placeholder="Paid"
        values={PaymentStatuses}
      />
      <InputBox
        id="note"
        label="Note"
        bind:value={note}
        type="textarea"
        placeholder="Note"
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

    .popup-container .popup {
      width: 100%;
    }
  }
</style>