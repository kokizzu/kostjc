<script>
  /** @typedef {import('../_types/property').Facility} Facility */
  /**
   * @typedef {Object} FacilityJson
   * @property {string} facilityName
   * @property {number} extraChargeIDR
   */

  import Select from '../node_modules/svelte-select';
  import { onMount } from 'svelte';
  import { RandString } from './xGenerator';

  export let stringJson = '';
  export let facilities = /** @type {Facility[]} */ ([]);

  let facilitiesIds = /** @type {number[]} */([]);
  // let facilitiesPrice = 0;

  /**
   * @typedef {Object} SvelteSelectValue
   * @prop {any} value
   * @prop {any} label
   */
  let valuesSvelteSelectFacilities = /** @type {SvelteSelectValue[]} */ ([]);
  let valueSvelteSelectedFacilities = /** @type {SvelteSelectValue[]} */ ([]);

  /**
   * 
   * @param {string} str
   * @returns {FacilityJson}
   */
  function parseFacilityString(str) {
    const match = str.match(/^(.+?)\s*\((\d+)\)$/);
    
    if (!match) {
      throw new Error("Invalid format. Expected: 'Name (Number)'");
    }

    const facilityName = match[1].trim();
    const extraChargeIDR = parseInt(match[2], 10);

    return {
      facilityName,
      extraChargeIDR
    };
  }

  function resfreshStringJson() {
    let facilitiesArrOfObj = /** @type {FacilityJson[]} */ ([]);
    (valueSvelteSelectedFacilities || []).forEach(v => {
      if (parseInt(v.value)) {
        for (const f of facilities) {
          if (f.id == v.value) {
            facilitiesArrOfObj = [...facilitiesArrOfObj, {
              extraChargeIDR: f.extraChargeIDR,
              facilityName: f.facilityName
            }]
          }
        }
        return
      }
      const toObj = parseFacilityString(v.label)
      facilitiesArrOfObj = [...facilitiesArrOfObj, toObj]
    });

    stringJson = JSON.stringify(facilitiesArrOfObj)
  }

  $: if (valueSvelteSelectedFacilities) {
    resfreshStringJson();
  }

  let isFacilitiesReady = false;

  function handleSelectFacilities(/** @type {CustomEvent} */e) {
    console.log('handleSelectFacilities:',e.detail)
    facilitiesIds = [...facilitiesIds, e.detail.value];
    // facilities.forEach((f) => {
    //   if (f.id === e.detail.value) {
    //     facilitiesPrice += f.extraChargeIDR || 0;
    //   }
    // })
  }

  function handleClearFacilities(/** @type {CustomEvent} */e) {
    const index = facilitiesIds.indexOf(Number(e.detail.value));
    if (index !== -1) {
      facilitiesIds.splice(index, 1);
    }
    // facilities.forEach((f) => {
    //   if (f.id === e.detail.value) {
    //     facilitiesPrice -= f.extraChargeIDR || 0;
    //   }
    // })
  }
  
  let facilitiesJson = /** @type {FacilityJson[]} */ ([]);

  onMount(() => {
    (facilities || []).forEach((f) => {
      valuesSvelteSelectFacilities = [...valuesSvelteSelectFacilities, {
        value: f.id,
        label: `${f.facilityName} ${f.facilityType ? `(${f.facilityType})` : ''}`
      }]
    });
    if (stringJson) {
      try {
        facilitiesJson = JSON.parse(stringJson);
      } catch (err) {
        console.error(err)
      }
    }

    if (facilitiesJson && (facilitiesJson || []).length > 0) {
      facilitiesJson.forEach(f => {
        const valueRand = RandString(8);
        valuesSvelteSelectFacilities = [...valuesSvelteSelectFacilities, {
          value: valueRand,
          label: `${f.facilityName} (${f.extraChargeIDR})`
        }];

        valueSvelteSelectedFacilities = [...valueSvelteSelectedFacilities, {
          value: valueRand,
          label: `${f.facilityName} (${f.extraChargeIDR})`
        }]
      });
    }

    console.log('Facilities JSON ::', facilitiesJson)
    isFacilitiesReady = true;
  });
</script>

{#if isFacilitiesReady}
  <div class="input-box">
    <label class="label" for="facilities">Facilities</label>
    <Select
      multiple
      items={valuesSvelteSelectFacilities}
      bind:value={valueSvelteSelectedFacilities}
      on:select={handleSelectFacilities}
      on:clear={handleClearFacilities}
      placeholder="Facilities"
    />
  </div>
{/if}

<style>
  .input-box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input-box .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }
</style>