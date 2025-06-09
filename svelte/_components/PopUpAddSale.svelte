<script>
  /** @typedef {import('../_types/cafe').Sale} Sale */
  /** @typedef {import('../_types/cafe').Menu} Menu */

  import { onMount } from 'svelte';
	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat, arrToArrNum } from './xFormatter';
  import MultiSelect from './MultiSelect.svelte';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let tenants = /** @type {Record<number, string>} */ ({});
  export let menus = /** @type {Record<number, string>} */ ({});

  let cashier = '';
  let tenantId = 0;
  let buyerName = '';
  let salesDate = dateISOFormat(0);
  let paidAt = dateISOFormat(0);
  let note = '';
  let qrisIDR = 0;
  let cashIDR = 0;
  let debtIDR = 0;
  let topupIDR = 0;
  let totalPriceIDR = 0;
  let menuIds = /** @type {number[]} */([]);

  export let OnSubmit = async function(/** @type {Sale} */ sale) {
    console.log('OnSubmit :::', sale);
  }

  async function submitAdd() {
      const sale = /** @type {Sale|any} */ ({
      cashier: cashier,
      tenantId: tenantId, 
      buyerName: buyerName,
      menuIds: menuIds,
      salesDate: salesDate,
      paidAt: paidAt,
      note: note,
      qrisIDR: qrisIDR,
      cashIDR: cashIDR,
      debtIDR: debtIDR,
      topupIDR: topupIDR,
    });


    await OnSubmit(sale);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
  cashier = '';
  tenantId = 0;
  menuIds = [];
  buyerName = '';
  salesDate = dateISOFormat(0);
  paidAt = dateISOFormat(0);
  note = '';
  qrisIDR = 0;
  cashIDR = 0;
  debtIDR = 0;
  topupIDR = 0;
  totalPriceIDR = 0;
  }
  
  const cancel = () => {
    isShow = false;
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Sale</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="cashier"
        label="Cashier"
        bind:value={cashier}
        type="text"
        placeholder="Cashier Name"
      />
      <InputBox
          id="tenantId"
          label="Tenant"
          isObject={true}
          bind:value={tenantId}
          type="combobox"
          values={tenants}
        />
      <InputBox
        id="buyerName"
        label="Buyer Name"
        bind:value={buyerName}
        type="text"
        placeholder="Jamal"
      />
      <MultiSelect
        id="menuIds"
        label="Menus"
        placeholder="Menus"
        valuesSourceType="object"
        bind:valuesTarget={menuIds}
        valuesSourceObj={menus}
      />
      <InputBox
      id="qrisIDR"
      label="QRIS IDR"
      bind:value={qrisIDR}
      type="number"
      placeholder="0"
      />
      <InputBox
      id="cashIDR"
      label="Cash IDR"
      bind:value={cashIDR}
      type="number"
      placeholder="0"
      />
      <InputBox
      id="debtIDR"
      label="Debt IDR"
      bind:value={debtIDR}
      type="number"
      placeholder="0"
      />
      <InputBox
      id="topupIDR"
      label="Top Up IDR"
      bind:value={topupIDR}
      type="number"
      placeholder="0"
      />
      <InputBox
      id="totalPriceIDR"
      label="Total Price IDR"
      bind:value={totalPriceIDR}
      type="number"
      placeholder="0"
      />
      <InputBox
        id="salesDate"
        label="Sales Date"
        bind:value={salesDate}
        type="datetime"
        placeholder="YYYY-MM-DD"
      />
       <InputBox
        id="paidAt"
        label="Paid At"
        bind:value={paidAt}
        type="datetime"
        placeholder="YYYY-MM-DD"
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
  }
</style>