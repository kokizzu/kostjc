<script>
  /** @typedef {import('../_types/property').PaymentPromise} PaymentPromise */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import { dateISOFormat } from './xFormatter';

  let isShow = false;

  export let isSubmitted = false;
  export let tenantName = '';
  let promisedAt = dateISOFormat(0);
  let promisePayAt = dateISOFormat(0);
  let note = '';

  export let OnSubmit = async function(/** @type {PaymentPromise} */ promise) {
    console.log('OnSubmit :::', promise);
  }

  async function submitAdd() {
    await OnSubmit(/** @type {PaymentPromise|any} */ ({
      promisedAt,
      promisePayAt,
      note,
    }));
  }

  function reset() {
    promisedAt = dateISOFormat(0);
    promisePayAt = dateISOFormat(0);
    note = '';
  }

  export const Show = () => {
    if (!promisedAt) promisedAt = dateISOFormat(0);
    if (!promisePayAt) promisePayAt = dateISOFormat(0);
    isShow = true;
  }
  export const Hide = () => {
    isShow = false;
    reset();
  }
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header>
      <h2>Payment Promise <span>{tenantName}</span></h2>
      <button class="close" on:click={Hide} aria-label="Close">
        <Icon size="20" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <label>
        <span>Promised</span>
        <input type="date" bind:value={promisedAt} />
      </label>
      <label>
        <span>Will Pay</span>
        <input type="date" bind:value={promisePayAt} />
      </label>
      <label class="note-field">
        <span>Note</span>
        <input type="text" bind:value={note} placeholder="Optional" />
      </label>
    </div>
    <div class="foot">
      <button class="cancel" on:click|preventDefault={Hide}>Cancel</button>
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
		padding: 8px;
    overflow: auto;
	}

  .popup-container.show {
    display: flex;
  }

	.popup {
		border-radius: 8px;
		background-color: #fff;
		height: fit-content;
		width: 360px;
    max-width: 100%;
		display: flex;
		flex-direction: column;
	}

  header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 4px;
		border-bottom: 1px solid var(--gray-004);
    gap: 4px;
	}

	header h2 {
		margin: 0;
    font-size: var(--font-lg);
    display: flex;
    flex-direction: column;
    gap: 1px;
	}

  header h2 span {
    color: var(--gray-006);
    font-size: var(--font-sm);
    font-weight: 600;
  }

	header button.close {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 3px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	header button.close:hover {
		background-color: #ef444420;
	}

	.forms {
		padding: 4px;
		display: grid;
    grid-template-columns: 1fr 1fr;
		gap: 4px;
	}

  .note-field {
    grid-column: 1 / -1;
  }

  label {
    display: flex;
    flex-direction: column;
    gap: 2px;
    color: var(--gray-007);
    font-size: var(--font-sm);
    font-weight: 600;
  }

  input {
    border: 1px solid var(--gray-004);
    border-radius: 6px;
    padding: 4px;
    font-size: var(--font-base);
  }

	.foot {
		display: flex;
    justify-content: flex-end;
		gap: 4px;
		align-items: center;
		padding: 4px;
		border-top: 1px solid var(--gray-004);
	}

	.foot button {
		padding: 4px 10px;
		border-radius: 9999px;
		border: none;
		cursor: pointer;
		font-weight: 700;
	}

	.foot button.ok {
		background-color: var(--green-006);
    color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
	}

	.foot button.ok:hover {
		background-color: var(--green-005);
	}

	.foot button.ok:disabled {
		cursor: not-allowed;
		background-color: var(--gray-003);
		color: var(--gray-007);
	}

	.foot button.cancel {
		background-color: transparent;
    color: var(--gray-007);
	}

	.foot button.cancel:hover {
		background-color: var(--gray-001);
	}

  @media only screen and (max-width : 480px) {
    .forms {
      grid-template-columns: 1fr;
    }
  }
</style>
