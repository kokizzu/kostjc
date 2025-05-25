<script>
  /**
    * @typedef { Array<string|number>
    * | Record<string|number, string>
    * | Record<string|number, any>
    * | string[]
    * | any[]
    * | number[]
    * } ComboboxValues
    */

  import { onMount } from 'svelte';
  import Select from '../node_modules/svelte-select';

  export let type         = /** @type {'combobox' | 'combobox-arr'} */ ('combobox');
  export let id           = /** @type {string} */ ('');
  export let label        = /** @type {string} */ ('');
  export let placeholder  = /** @type {string} */ ('');
  export let value        = /** @type {any} */ (null);
  export let values       = /** @type {ComboboxValues} */ ([]);

  let itemsArr    = /** @type {ComboboxValues} */ ([]);
  let itemsArrObj = /** @type {ComboboxValues} */ ([]);

  /**
   * @typedef {Object} SvelteSelectValue
   * @prop {any} value
   * @prop {any} label
   */
  let svelteSelectValue = /** @type {SvelteSelectValue} */ ({
    label: '',
    value: ''
  });

  onMount(() => {
    console.log('Value: ', value);
    switch (type) {
      case 'combobox-arr': {
        itemsArr = values;
        // @ts-ignore
        for (const v of values) {
          if (String(v) == String(value)) {
            svelteSelectValue = v;
          }
        }
      }
      case 'combobox': {
        for (const [k, v] of Object.entries(values)) {
          itemsArrObj = [...itemsArrObj, {
            value: k,
            label: v
          }];
          if (String(k) == String(value)) {
            svelteSelectValue = {
              value: k,
              label: v
            };
          }
        }
      }
    }
  });

  function handleSelect(/** @type {CustomEvent} */e) {
    value = e.detail.value;
  }

  function handleClear(/** @type {CustomEvent} */e) {
    value = '';
  }
</script>

<div class="input-box">
  <label class="label" for={id}>{label}</label>
  {#if type === 'combobox'}
    <Select
      items={itemsArrObj}
      bind:value={svelteSelectValue}
      on:select={handleSelect}
      on:clear={handleClear}
      placeholder={placeholder}
    />
  {:else if type === 'combobox-arr'}
    <Select
      items={itemsArr}
      bind:value={svelteSelectValue}
      on:select={handleSelect}
      on:clear={handleClear}
      placeholder={placeholder}
    />
  {/if}
</div>

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