<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
	import { RiArrowsArrowRightLine } from '../node_modules/svelte-icons-pack/dist/ri';
	import { localeDatetime } from './xFormatter';
	import { notifier } from './xNotifier';

  let isShow = /** @type {boolean} */ (false);
	let differences = /** @type {ObjDifference[]} */ ([]);

	export let users = /** @type {Record<number, string>} */ ({});
	export let tenants = /** @type {Record<number, string>} */ ({});

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

	// Obtain it from file:
	// auth_tables.go
	// property_tables.go
	// cafe_tables.go
	const KeyMap = {
		'id': 'ID',
		'createdAt': 'Created At',
		'createdBy': 'Created By',
		'updatedAt': 'Updated At',
		'updatedBy': 'Updated By',
		'deletedAt': 'Deleted At',
		'deletedBy': 'Deleted By',
		'restoredBy': 'Restored By',
		'name': 'Name',
		'address': 'Address',
		'gmapLocation': 'Google Map Location',
		'facilityName': 'Facility Name',
		'extraChargeIDR': 'Extra Charge IDR',
		'facilityType': 'Facility Type',
		'descriptionEN': 'Description (EN)',
		'buildingName': 'Building Name',
		'locationId': 'Location',
		'facilities': 'Facilities',
		'roomName': 'Room Name',
		'currentTenantId': 'Current Tenant',
		'firstUseAt': 'First Use At',
		'lastUseAt': 'Last Use At',
		'buildingId': 'Building',
		'roomSize': 'Room Size',
		'imageUrl': 'Image URL',
		'dateStart': 'Date Start',
		'dateEnd': 'Date End',
		'basePriceIDR': 'Base Price IDR',
		'totalPriceIDR': 'Total Price IDR',
		'paidAt': 'Paid At',
		'facilitiesObj': 'Facilities',
		'tenantId': 'Tenant',
		'extraTenants': 'Extra Tenants',
		'roomId': 'Room',
		'totalPaidIDR': 'Total Paid IDR',
		'bookingId': 'Booking',
		'paymentAt': 'Payment At',
		'paidIDR': 'Paid IDR',
		'paymentMethod': 'Payment Method',
		'paymentStatus': 'Payment Status',
		'note': 'Note',
		'stockName': 'Stock Name',
		'stockAddedAt': 'Stock Added At',
		'quantity': 'Quantity',
		'priceIDR': 'Price IDR',
		'startAt': 'Start At',
		'endAt': 'End At',
		'macAddress': 'MAC Address',
		'hppIDR': 'HPP IDR',
		'salePriceIDR': 'Sale Price IDR',
		'detail': 'Detail',
		'cashier': 'Cashier',
		'buyerName': 'Buyer Name',
		'menuIds': 'Menus',
		'qrisIDR': 'QRIS IDR',
		'cashIDR': 'Cash IDR',
		'debtIDR': 'Debt IDR',
		'topupIDR': 'Topup IDR',
		'salesDate': 'Sales Date',
		'donation': 'Donation',
		'transferIDR': 'Transfer IDR',
		'email': 'Email',
		'password': 'Password',
		'passwordSetAt': 'Password Set At',
		'secretCode': 'Secret Code',
		'secretCodeAt': 'Secret Code At',
		'verifiedAt': 'Verified At',
		'lastLoginAt': 'Last Login At',
		'fullName': 'Full Name',
		'userName': 'Username',
		'role': 'Role',
		'tenantName': 'Tenant Name',
		'ktpRegion': 'KTP Region',
		'ktpNumber': 'KTP Number',
		'ktpName': 'KTP Name',
		'ktpPlaceBirth': 'KTP Place of Birth',
		'ktpDateBirth': 'KTP Date of Birth',
		'ktpGender': 'KTP Gender',
		'ktpAddress': 'KTP Address',
		'ktpRtRw': 'KTP RT/RW',
		'ktpKelurahanDesa': 'KTP Kelurahan/Desa',
		'ktpKecamatan': 'KTP Kecamatan',
		'ktpReligion': 'KTP Religion',
		'ktpMaritalStatus': 'KTP Marital Status',
		'ktpCitizenship': 'KTP Citizenship',
		'ktpOccupation': 'KTP Occupation',
		'telegramUsername': 'Telegram Username',
		'whatsappNumber': 'WhatsApp Number',
		'waAddedAt': 'WhatsApp Added At',
		'teleAddedAt': 'Telegram Added At',
		'userId': 'User',
	}

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

			// +====== BEFORE VALUE ======+
			let beforeValueFinal = beforeValue;
			let afterValueFinal = afterValue;

			if (String(key).includes('At')) {
				console.log(key, beforeValue, afterValue);
				beforeValueFinal = Number(beforeValue) == 0 ? '--' : localeDatetime(beforeValue);
				afterValueFinal =  Number(afterValue) == 0 ? '--' : localeDatetime(afterValue);
			}

			if (String(key) == 'extraTenants') {
				beforeValueFinal = (beforeValue || []).map((/** @type {string | number} */ id) => tenants[id]).join(', ');
				afterValueFinal = (afterValue || []).map((/** @type {string | number} */ id) => tenants[id]).join(', ');
			}

			if (String(key) == 'tenantId' || String(key) == 'currentTenantId') {
				beforeValueFinal = tenants[beforeValue] || beforeValue;
				afterValueFinal = tenants[afterValue] || afterValue;
			}

			if (String(key) == 'userId') {
				beforeValueFinal = users[beforeValue] || beforeValue;
				afterValueFinal = users[afterValue] || afterValue;
			}

			if (String(key).includes('By')) {
				beforeValueFinal = users[beforeValue] || beforeValue;
				afterValueFinal = users[afterValue] || afterValue;
			}

			if (beforeValue !== undefined && afterValue !== undefined) {
				if (JSON.stringify(beforeValue) !== JSON.stringify(afterValue)) {
					differences.push({
						field: KeyMap[key] || key,
						before: beforeValueFinal,
						after: afterValueFinal,
						type: 'modified',
					});
				}
			} else if (beforeValue !== undefined && afterValue === undefined) {
				differences.push({
					field: KeyMap[key] || key,
					before: beforeValueFinal,
					after: null,
					type: 'removed',
				});
			} else if (beforeValue === undefined && afterValue !== undefined) {
				differences.push({
					field: KeyMap[key] || key,
					before: null,
					after: afterValueFinal,
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
		width: 700px;
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