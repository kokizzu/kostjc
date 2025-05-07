<script>
  /** @typedef {import('../_types/property').Building} Building */
  /** @typedef {import('../_types/property').Facility} Facility */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import { RiArrowsArrowDropRightLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import InputBox from './InputBox.svelte';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let locations = /** @type {Record<number, string>} */ ({});
  export let facilities = /** @type {Facility[]} */ ([]);

  let buildingName = '';
  let locationId = 0;
  let facilitiesNums = /** @type {number[]} */ ([]);

  export let OnSubmit = async function(/** @type {Building} */ building) {
    console.log('OnSubmit :::', building);
  }

  async function submitAdd() {
    const building = /** @type {Building|any} */ ({
      buildingName: buildingName,
      locationId: locationId,
      facilities: facilitiesNums
    });

    await OnSubmit(building);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    buildingName = '';
    locationId = 0;
    facilitiesNums = [];
  }
  
  const cancel = () => {
    isShow = false;
  }

  let facilitiesToShow = /** @type {Facility[]} */ ([]);
  let showFacilities = false;
  let selectedFacility = 'Facility....';

  function handleStationListClose() {
		showFacilities = false;
		document.body.removeEventListener('click', handleStationListClose)
	}
	function toggleStationList() {
		showFacilities = !showFacilities;
		if (showFacilities) document.body.addEventListener('click', handleStationListClose);
		else document.body.removeEventListener('click', handleStationListClose);
	}

	function choseFacility(/** @type {Facility} */ facility) {
    selectedFacility = facility.facilityName
    showFacilities = false;
    facilitiesNums = [...facilitiesNums, Number(facility.id)];
		facilitiesToShow = [...facilitiesToShow, facility];
    facilities = facilities.filter(f => f.id !== facility.id);
	}

	const removeFacility = (/** @type {number} */ idx) => {
    facilitiesNums = facilitiesNums.filter(( _, i ) => i !== idx);
		facilitiesToShow = facilitiesToShow.filter(( _, i ) => i !== idx);
    facilities = [...facilities, facilitiesToShow[idx]];
	}
</script>

<div class={`popup_container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Add Building</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="buildingName"
        label="Building Name"
        bind:value={buildingName}
        type="text"
        placeholder="Kost JC"
      />
      <InputBox
        id="location"
        label="Location"
        isObject={true}
        bind:value={locationId}
        type="combobox"
        values={locations}
      />
      <div class="facilities-form">
        <label class="label" for="facilities">Facilities</label>
        <div class="facilities-selected">
          {#each facilitiesToShow as fac, idx}
            <div class="facility-show">
              <span>{fac.facilityName}</span>
              <button on:click={() => removeFacility(idx)}>
                <Icon size="18" color="var(--sky-001)" src={IoClose}/>
              </button>
            </div>
          {/each}
        </div>
        <div class="dropdown-facilities">
          <div class="dropdown-item">
            <button id="facilities" class="dropdown-btn" on:click|stopPropagation={toggleStationList}>
              <span>{selectedFacility}</span>
              <Icon
                className={showFacilities ? 'rotate' : 'dropdown'}
                size="20"
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
                      <span>{fac.facilityName}</span>
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

  .facilities-form {
    display: flex;
    flex-direction: column;
    gap: 3px;
  }

  .facilities-form label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .facilities-form .facilities-selected {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 5px;
  }

  .facilities-form .facilities-selected .facility-show {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    background-color: var(--sky-006);
    color: #FFF;
    padding: 5px 5px 5px 15px;
    border-radius: 5px;
  }

  .facilities-form .facilities-selected .facility-show button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px;
    border-radius: 9999px;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }

  .facilities-form .facilities-selected .facility-show button:hover {
    background-color: var(--sky-005);
  }

  .dropdown-facilities {
		position: relative;
		flex-grow: 1;
		display: flex;
  }

  .dropdown-facilities {
		position: relative;
		flex-grow: 1;
		display: flex;
	}

	.dropdown-facilities .dropdown-item {
		width: 100%;
	}

	.dropdown-facilities .dropdown-item .dropdown-btn {
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
		min-height: fit-content;
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

  @media only screen and (max-width : 768px) {
    .popup_container {
      padding: 10px;
    }
  }
</style>