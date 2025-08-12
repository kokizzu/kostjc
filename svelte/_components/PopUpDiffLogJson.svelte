<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
	import { RiArrowsArrowRightLine } from '../node_modules/svelte-icons-pack/dist/ri';
	import { notifier } from './xNotifier';

  let isShow = /** @type {boolean} */ (false);
	let differences = /** @type {ObjDifference[]} */ ([]);

  export const Show = () => {
		compareJson();
		isShow = true;
	}
  export const Hide = () => {
		isShow = false;
		differences = [];
	}

  export let beforeJson = '';
  export let afterJson = '';

	function compareJson() {
		try {
			const objBeforeJson = JSON.parse(beforeJson || '{}');
			const objAfterJson = JSON.parse(afterJson || '{}');

			differences = findDifferences(objBeforeJson, objAfterJson);
		} catch (error) {
			console.error(error);
			notifier.showError(error);
		}
	}

	/**
	 * @typedef {Object} ObjDifference
	 * @property {string} field
	 * @property {string} before
	 * @property {string} after
	 * @property {'added' | 'removed' | 'modified'} type
	 */

	/**
	 * @description Find differences
	 * @param {Object | Record<any, any>} before
	 * @param {Object | Record<any, any>} after
	 * @returns {ObjDifference[]}
	 */
	function findDifferences(before, after) {
		const differences = /** @type {ObjDifference[]} */ ([]);

		const allKeys = new Set([...Object.keys(before), ...Object.keys(after)]);

		(allKeys || []).forEach((/** @type {string} */ key) => {
			const beforeValue = before[key];
			const afterValue = after[key];

			if (beforeValue !== undefined && afterValue !== undefined) {
				if (JSON.stringify(beforeValue) !== JSON.stringify(afterValue)) {
					differences.push({
						field: key,
						before: beforeValue,
						after: afterValue,
						type: 'modified',
					});
				}
			} else if (beforeValue !== undefined && afterValue === undefined) {
				differences.push({
					field: key,
					before: beforeValue,
					after: null,
					type: 'removed',
				});
			} else if (beforeValue === undefined && afterValue !== undefined) {
				differences.push({
					field: key,
					before: null,
					after: afterValue,
					type: 'added'
				});
			}
		});

		return differences;
	}
</script>

<div class="popup-container {isShow ? 'show' : ''}">
  <div class="popup">
    <header class="header">
      <h2>Diff data</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <div class="data-object-container">
				{#each differences as dif}
					<div class="field-container">
						<span class="field">{dif.field}</span>
						<div class="values">
							<span class="before">{dif.before ? dif.before : '(not present)'}</span>
							<Icon
								src={RiArrowsArrowRightLine}
								size="20"
							/>
							<span class="after">{dif.after ? dif.after : '(removed)'}</span>
						</div>
					</div>
				{/each}
      </div>
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

  .popup-container .popup .forms .data-object-container {
    display: flex;
		flex-direction: column;
    gap: 10px;
  }

	.data-object-container .field-container {
		display: flex;
		flex-direction: column;
		gap: 5px;
		padding: 8px;
		background-color: var(--gray-002);
		border-radius: 8px;
	}

	.data-object-container .field-container .field {
		font-weight: 600;
		font-size: 15px;
	}

	.data-object-container .field-container .values {
		display: flex;
		flex-direction: row;
		gap: 5px;
	}

	.data-object-container .field-container .values .before {
		background-color: var(--red-transparent);
		padding: 3px 5px;
		border: 1px solid var(--red-005);
		border-radius: 5px;
		font-size: 12px;
		color: var(--red-006);
	}

	.data-object-container .field-container .values .after {
		background-color: var(--green-transparent);
		padding: 3px 5px;
		border: 1px solid var(--green-005);
		border-radius: 5px;
		font-size: 12px;
		color: var(--green-006);
	}

	.popup-container .popup .foot button {
		padding: 8px 18px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

	@media only screen and (max-width: 768px) {
		.popup-container {
      padding: 10px;
    }

    .popup-container .popup {
      width: 100%;
    }

		.popup-container .popup .forms .data-object-container {
			display: flex;
			flex-direction: column;
		}
	}
</style>