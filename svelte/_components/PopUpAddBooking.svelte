<script>
  /** @typedef {import('../_types/property').Booking} Booking */
  /** @typedef {import('../_types/property').Facility} Facility */


	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { RiArrowsArrowDropRightLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './xFormatter';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let tenants = /** @type {Record<number, string>} */ ({});
  export let facilities = /** @type {Facility[]} */ ([]);
  export let rooms = /** @type {Record<number, string>} */ ({});

  let dateStart = dateISOFormat(0);
  let dateEnd = dateISOFormat(30);
  let basePriceIDR = 0;
  let facilitiesNums = /** @type {number[]} */([]);
  let totalPriceIDR = 0;
  let paidAt = dateISOFormat(0);
  let tenantId = 0;
  let extraTenantsIds = /** @type {number[]} */([]);
  let roomId = 0;

  export let OnSubmit = async function(/** @type {Booking} */ booking, /** @type {number[]} */ facilities) {
    console.log('OnSubmit :::', booking, facilities);
  }

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

    await OnSubmit(booking, facilitiesNums);
    extraTenantsIds = [];
    extraTenantsToShow = [];
  }

  export const Show = () => {
    isShow = true;
    fillExtraTenants();
  }
  export const Hide = () => isShow = false;

  export const Reset = () => {
    dateStart = dateISOFormat(0);
    dateEnd = dateISOFormat(30);
    basePriceIDR = 0;
    facilitiesNums = [];
    totalPriceIDR = 0;
    paidAt = dateISOFormat(0);
    tenantId = 0;
    facilitiesToShow = [];
    selectedFacility = 'Facility....';
    facilitiesPrice = 0;
    extraTenantsIds = [];
    roomId = 0;
    extraTenantsIds = [];
    extraTenants = []
    selectedExtraTenants = 'Tenant....';
    showExtraTenants = false;
    extraTenantsToShow = [];
  }
  
  const cancel = () => {
    isShow = false;
  }

  // Facilities

  let facilitiesToShow = /** @type {Facility[]} */ ([]);
  let showFacilities = false;
  let selectedFacility = 'Facility....';
  let facilitiesPrice = 0;

  function handleFacilitiesClose() {
		showFacilities = false;
		document.body.removeEventListener('click', handleFacilitiesClose)
	}
	function toggleFacilities() {
		showFacilities = !showFacilities;
		if (showFacilities) document.body.addEventListener('click', handleFacilitiesClose);
		else document.body.removeEventListener('click', handleFacilitiesClose);
	}

	function choseFacility(/** @type {Facility} */ facility) {
    selectedFacility = facility.facilityName + ' (' + facility.facilityType + ')';
    showFacilities = false;
    facilitiesPrice += facility.extraChargeIDR;
    facilitiesNums = [...facilitiesNums, Number(facility.id)];
		facilitiesToShow = [...facilitiesToShow, facility];
    facilities = facilities.filter(f => f.id !== facility.id);
	}

	const removeFacility = (/** @type {number} */ idx) => {
    facilitiesPrice -= facilitiesToShow[idx].extraChargeIDR;
    facilitiesNums = facilitiesNums.filter(( _, i ) => i !== idx);
    facilities = [...facilities, facilitiesToShow[idx]];
		facilitiesToShow = facilitiesToShow.filter(( _, i ) => i !== idx);
	}

  // Extra Tenants

  /**
   * @typedef {Object} ExtraTenant
   * @property {number} id
   * @property {string} name
   */

  let extraTenants = /** @type {ExtraTenant[]} */ ([]);
  function fillExtraTenants() {
    for (const [k, v] of Object.entries(tenants)) extraTenants.push({id: Number(k), name: v});
  }

  let extraTenantsToShow = /** @type {ExtraTenant[]} */ ([]);
  let showExtraTenants = false;
  let selectedExtraTenants = 'Tenant....';

  function handleExtraTenantsClose() {
		showExtraTenants = false;
		document.body.removeEventListener('click', handleExtraTenantsClose)
	}
	function toggleExtraTenants() {
		showExtraTenants = !showExtraTenants;
		if (showExtraTenants) document.body.addEventListener('click', handleExtraTenantsClose);
		else document.body.removeEventListener('click', handleExtraTenantsClose);
	}

	function choseExtraTenant(/** @type {ExtraTenant} */ extraTenant) {
    selectedExtraTenants = extraTenant.name;
    showExtraTenants = false;
    extraTenantsIds = [...extraTenantsIds, Number(extraTenant.id)];
		extraTenantsToShow = [...extraTenantsToShow, extraTenant];
    extraTenants = extraTenants.filter(f => f.id !== extraTenant.id);
	}

	const removeExtraTenant = (/** @type {number} */ idx) => {
    extraTenantsIds = extraTenantsIds.filter(( _, i ) => i !== idx);
    extraTenants = [...extraTenants, extraTenantsToShow[idx]];
		extraTenantsToShow = extraTenantsToShow.filter(( _, i ) => i !== idx);
	}
</script>

<div class={`popup_container ${isShow ? 'show' : ''}`}>
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
      <div class="multiselect-form">
        <label class="label" for="facilities">Facilities</label>
        <div class="multiselect-selected">
          {#each facilitiesToShow as fac, idx}
            <div class="facility-show">
              <span>{fac.facilityName} ({fac.facilityType})</span>
              <button on:click={() => removeFacility(idx)}>
                <Icon size="18" color="var(--sky-001)" src={IoClose}/>
              </button>
            </div>
          {/each}
        </div>
        <div class="dropdown-multiselect">
          <div class="dropdown-item">
            <button id="facilities" class="dropdown-btn" on:click|stopPropagation={toggleFacilities}>
              <span>{selectedFacility}</span>
              <Icon
                className={showFacilities ? 'rotate-b' : 'dropdown'}
                size="25"
                src={RiArrowsArrowDropRightLine}
              />
            </button>
            {#if showFacilities}
              <div class="dropdown-list">
                {#if facilities && facilities.length}
                  {#each facilities as fac}
                    <button
                      class="facility"
                      on:click|stopPropagation={() => choseFacility(fac)}>
                      <span>{fac.facilityName} ({fac.facilityType})</span>
                    </button>
                  {/each}
                {:else}
                  <p>No Facilities yet, please add it in master facility</p>
                {/if}
              </div>
            {/if}
          </div>
        </div>
      </div>
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
      <div class="multiselect-form">
        <label class="label" for="extraTenants">Extra Tenants</label>
        <div class="multiselect-selected">
          {#each extraTenantsToShow as ext, idx}
            <div class="facility-show">
              <span>{ext.name}</span>
              <button on:click={() => removeExtraTenant(idx)}>
                <Icon size="18" color="var(--sky-001)" src={IoClose}/>
              </button>
            </div>
          {/each}
        </div>
        <div class="dropdown-multiselect">
          <div class="dropdown-item">
            <button id="extraTenants" class="dropdown-btn" on:click|stopPropagation={toggleExtraTenants}>
              <span>{selectedExtraTenants}</span>
              <Icon
                className={showExtraTenants ? 'rotate-b' : 'dropdown'}
                size="25"
                src={RiArrowsArrowDropRightLine}
              />
            </button>
            {#if showExtraTenants}
              <div class="dropdown-list">
                {#if extraTenants && extraTenants.length}
                  {#each extraTenants as ext}
                    <button
                      class="facility"
                      on:click|stopPropagation={() => choseExtraTenant(ext)}>
                      <span>{ext.name}</span>
                    </button>
                  {/each}
                {:else}
                  <p>No Tenants yet, please add it in master tenants</p>
                {/if}
              </div>
            {/if}
          </div>
        </div>
      </div>
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

  .popup_container {
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

  .popup_container.show {
    display: flex;
  }

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 15px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header h2 {
		margin: 0;
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup header button:active {
		background-color: #ef444430;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup_container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 15px 20px;
		border-top: 1px solid var(--gray-004);
	}

  .popup_container .popup .foot .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

	.popup_container .popup .foot button {
		padding: 8px 18px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

	.popup_container .popup .foot button.ok {
		background-color: var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
	}

	.popup_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup_container .popup .foot button.ok:disabled {
		cursor: not-allowed;
		background-color: var(--gray-003);
		color: var(--gray-007);
	}

	.popup_container .popup .foot button.cancel {
		background-color: transparent;
    color: var(--gray-007);
	}

	.popup_container .popup .foot button.cancel:hover {
		background-color: var(--gray-001);
	}

  .multiselect-form {
    display: flex;
    flex-direction: column;
    gap: 3px;
  }

  .multiselect-form label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .multiselect-form .multiselect-selected {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 5px;
  }

  .multiselect-form .multiselect-selected .facility-show {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    background-color: var(--sky-006);
    color: #FFF;
    padding: 5px 5px 5px 15px;
    border-radius: 5px;
  }

  .multiselect-form .multiselect-selected .facility-show button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px;
    border-radius: 9999px;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }

  .multiselect-form .multiselect-selected .facility-show button:hover {
    background-color: var(--sky-005);
  }

  .dropdown-multiselect {
		position: relative;
		flex-grow: 1;
		display: flex;
  }

  .dropdown-multiselect {
		position: relative;
		flex-grow: 1;
		display: flex;
	}

	.dropdown-multiselect .dropdown-item {
		width: 100%;
	}

	.dropdown-multiselect .dropdown-item .dropdown-btn {
		width: 100% !important;
	}

	.dropdown-item {
		flex-grow: 1;
		position: relative;
	}

	.dropdown-item .dropdown-btn {
		width: 100%;
		height: fit-content;
		background-color: #FFF;
		padding: 10px 15px;
		border-radius: 5px;
		display: flex;
		flex-direction: row;
		border: 1px solid var(--gray-003);
		align-items: center;
		cursor: pointer;
		color: var(--gray-007);
		font-weight: 600;
	}

	.dropdown-item .dropdown-btn:hover {
		background-color: var(--gray-001);
	}

	.dropdown-item .dropdown-btn span {
		flex-grow: 1;
		text-align: left;
	}

	.dropdown-item .dropdown-list {
		background-color: #FFF;
		max-height: 300px;
		overflow-y: auto;
		width: 100%;
    margin-top: 10px;
		border: 1px solid var(--sky-004);
		border-radius: 5px;
		display: flex;
		flex-direction: column;
		position: absolute;
		top: 40px;
		z-index: 99999;
	}

	.dropdown-item .dropdown-list::-webkit-scrollbar-thumb {
    background-color : var(--gray-003);
    border-radius    : 8px;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-thumb:hover {
    background-color : var(--gray-004);
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar {
    width : 8px;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-track {
    background-color : transparent;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-button {
    display: none;
  }

	.dropdown-item .dropdown-list .facility {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 10px;
		padding: 10px 15px;
		background-color: #FFF;
		border: none;
		color: var(--gray-007);
		width: 100%;
		cursor: pointer;
	}

	.dropdown-item .dropdown-list .facility:hover {
		color: var(--sky-009);
    background-color: #0ea5e920;
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

  @media only screen and (max-width : 768px) {
    .popup_container {
      padding: 10px;
    }
  }
</style>