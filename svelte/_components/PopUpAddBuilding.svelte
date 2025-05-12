<script>
    import { onMount } from 'svelte';
  /** @typedef {import('../_types/property').Building} Building */
  /** @typedef {import('../_types/property').Facility} Facility */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import { RiArrowsArrowDropRightLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import InputBox from './InputBox.svelte';
    import MultiSelect from './MultiSelect.svelte';
    import { arrToArrNum } from './xFormatter';

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
    const facilitiesArrNums = arrToArrNum(facilitiesNums);
    const building = /** @type {Building|any} */ ({
      buildingName: buildingName,
      locationId: locationId,
      facilities: facilitiesArrNums
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

  let facilitiesObj = {};
  let isFacilitiesReady = false;
  onMount(() => {
    (facilities || []).forEach(f => {
      facilitiesObj[f.id] = `${f.facilityName} (${f.facilityType})`
    });
    isFacilitiesReady = true;
  })
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
      {#if isFacilitiesReady}
        <MultiSelect
          id="facilities"
          label="Facilities"
          placeholder="Facilities"
          valuesSourceType="object"
          bind:valuesTarget={facilitiesNums}
          valuesSourceObj={facilitiesObj}
        />
      {/if}
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
</style>