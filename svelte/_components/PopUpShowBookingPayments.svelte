<script>
  /** @typedef {import('../_types/property').Payment} Payment */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
	import { RiDesignBallPenLine, RiArrowsArrowGoBackLine, RiSystemDeleteBin5Line } from '../node_modules/svelte-icons-pack/dist/ri';
	import { dateISOFormat, formatPrice } from './xFormatter';
	import PopUpEditPayment from './PopUpEditPayment.svelte';
	import { AdminPayment } from '../jsApi.GEN';
	import { CmdDelete, CmdForm, CmdRestore, CmdUpsert } from './xConstant';
	import { notifier } from './xNotifier';

  let isShow = /** @type {boolean} */ (false);

  export let bookingId = /** @type {number} */ (0);
  export let payments = /** @type {Payment[]} */ ([]);

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;
  export const Reset = () => {}
	export let Refresh = async () => {}

	let popUpEditPayment = null;
	let isShowPopUpEditPayment = false;
	let isSubmitEditPayment = false;

	let editPaymentId;
	let editPaymentAt = dateISOFormat(0);
  let editTotalPaidIDR = 0;
  let editNote = '';

	function onPopUpEditPayment(/** @type {Payment} */ payment) {
		editPaymentId = payment.id;
		editPaymentAt =  payment.paymentAt;
		editTotalPaidIDR = payment.paidIDR;
		editNote = payment.note;
		popUpEditPayment.Show(
			payment.paymentMethod,
			payment.paymentStatus
		);
	}

	async function refreshPayments() { // @ts-ignore
		await AdminPayment({ 
      cmd: CmdForm,
      bookingId: bookingId
    }, /** @type {import('../jsApi.GEN').AdminPaymentCallback} */
    /** @returns {Promise<void>} */
      function(/** @type any */ o) {
      if (o.error) {
        console.log(o);
        notifier.showError(o.error);
        return
      }

      payments = o.paymentsByBooking;
    });
	}

	async function OnEditPayment(/** @type {Payment} */ payment) {
		payment.id = editPaymentId;
		payment.paymentMethod = payment.paymentMethod.value;
		payment.paymentStatus = payment.paymentStatus.value;
		const i = /** @type {any}*/ ({
      payment,
      cmd: CmdUpsert
    });
    await AdminPayment(i,
      /** @type {import('./jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

				refreshPayments();
				Refresh();

        notifier.showSuccess(`Payment #${payment.id} updated !!`);
				popUpEditPayment.Hide()
      }
    );

		isSubmitEditPayment = false;
	}

	async function OnDeletePayment(/** @type {Payment} */ payment) {
		payment.id = payment.id+'';
		const i = /** @type {any}*/ ({
			payment,
			cmd: CmdDelete
		});
		await AdminPayment(i,
			/** @type {import('./jsApi.GEN').AdminPaymentCallback} */
			/** @returns {Promise<void>} */
			function(/** @type any */ o) {
				if (o.error) {
					console.log(o);
					notifier.showError(o.error);
					return
				}

				refreshPayments();
				Refresh();

				notifier.showSuccess(`Payment #${payment.id} deleted !!`);
			}
		);
	}

	async function OnRestorePayment(/** @type {Payment} */ payment) {
		payment.id = payment.id+'';
		const i = /** @type {any}*/ ({
			payment,
			cmd: CmdRestore
		});
		await AdminPayment(i,
			/** @type {import('./jsApi.GEN').AdminPaymentCallback} */
			/** @returns {Promise<void>} */
			function(/** @type any */ o) {
				if (o.error) {
					console.log(o);
					notifier.showError(o.error);
					return
				}

				refreshPayments();
				Refresh();

				notifier.showSuccess(`Payment #${payment.id} restored !!`);
			}
		);
	}
</script>

<PopUpEditPayment 
	bind:isShow={isShowPopUpEditPayment}
	bind:this={popUpEditPayment}
	bind:paymentId={editPaymentId}
	bind:isSubmitted={isSubmitEditPayment}
	bind:paymentAt={editPaymentAt}
	bind:totalPaidIDR={editTotalPaidIDR}
	bind:note={editNote}
	OnSubmit={OnEditPayment}
/>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Payments for Booking {`#${bookingId}`}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
			<div class="table-root">
				<div class="table-container">
					<table>
						<thead>
							<tr>
								<th style="width: 40px;">No</th>
								<th class="a-row" style="width: 70px;">Actions</th>
								<th>ID</th>
								<th>Payment At</th>
								<th>Total Paid</th>
								<th>Payment Method</th>
								<th>Payment Status</th>
								<th class="textarea">Note</th>
							</tr>
						</thead>
						<tbody>
							{#each (payments || []) as py, idx}
								<tr class={py.deletedAt > 0 ? 'deleted' : ''}>
									<td class="num-row">{idx + 1}</td>
									<td class="a-row">
										<div class="actions">
											<button
												class="btn edit"
												title="Edit"
												on:click={() => onPopUpEditPayment(py)}
											>
												<Icon
													size="15"
													color="var(--gray-007)"
													src={RiDesignBallPenLine}
												/>
											</button>
											{#if py.deletedAt > 0}
												<button
													class="btn"
													title="Restore"
													on:click={() => OnRestorePayment(py)}
												>
													<Icon
														size="15"
														color="var(--gray-007)"
														src={RiArrowsArrowGoBackLine}
													/>
												</button>
											{:else}
												<button
													class="btn delete"
													title="Delete"
													on:click={() => OnDeletePayment(py)}
												>
													<Icon
														size="15"
														color="var(--gray-007)"
														src={RiSystemDeleteBin5Line}
													/>
												</button>
											{/if}
										</div>
									</td>
									<td>#{py.id}</td>
									<td>{py.paymentAt}</td>
									<td>{formatPrice(py.paidIDR, 'IDR')}</td>
									<td>{py.paymentMethod}</td>
									<td>{py.paymentStatus}</td>
									<td>{py.note}</td>
								</tr>
							{/each}
							{#if (payments || []).length === 0}
								<tr>
									<td class="num-row">0</td>
									<td colspan="6">no-data</td>
								</tr>
							{/if}
						</tbody>
					</table>
				</div>
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
		width: 80%;
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

	.table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-left: 1px solid var(--gray-003);
		border-right: 1px solid var(--gray-003);
    padding: 0;
    font-size: var(--font-base);
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: calc(100% - 20px);
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    border-top: 1px solid var(--gray-003);
    border-bottom: 1px solid var(--gray-003);
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
    font-size: var(--font-base);
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table-container table tbody tr.deleted {
    color: var(--red-005);
  }

	.table-root .table-container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

	.table-root .table-container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

	:global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--blue-005);
  }

  .table-root .table-container table tbody tr td .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

	.table-root .table-container table tbody tr td .actions .btn.delete:hover {
    background-color: var(--red-transparent);
  }

	:global(.table-root .table-container table tbody tr td .actions .btn.delete:hover svg) {
    fill: var(--red-005);
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

	.table-root .table-container table tbody tr td.num-row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
	}

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table thead tr th.textarea {
    min-width: 280px !important;
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