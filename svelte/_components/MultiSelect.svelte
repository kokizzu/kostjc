<script>
  // NOTE: valuesSourceType = 'array' is not yet implemented

  import { onMount } from 'svelte';
  import Select from '../node_modules/svelte-select';

  export let className = '';
  export let valuesSourceType = /** @type {'array' | 'object'} */ ('array')
  export let valuesTarget = [];
  export let valuesSource = [];
  export let valuesSourceObj = {};
  export let id;
  export let label;
  export let placeholder = '';

  function handleSelect(/** @type {CustomEvent} */e) {
    console.log('handleSelect:',e.detail)
    valuesTarget = [...valuesTarget, e.detail.value];
  }

  function handleClear(/** @type {CustomEvent} */e) {
    const index = valuesTarget.indexOf(Number(e.detail.value));
    if (index !== -1) {
      valuesTarget.splice(index, 1);
    }
  }

  /**
   * @typedef {Object} SvelteSelectValue
   * @prop {any} value
   * @prop {any} label
   */
  let valuesSourceArrayObject = /** @type {SvelteSelectValue[]} */ ([]);
  let valuesTargetArrayObject = [];

  function reStructureSvelteValues() {
    valuesSourceArrayObject = [];
    valuesTargetArrayObject = [];
    for (const [k, v] of Object.entries(valuesSourceObj)) {
      valuesSourceArrayObject.push({
        value: k,
        label: v
      });
    }
    if (valuesSourceType === 'object') {
      for (const kv of valuesSourceArrayObject) {
        for (const vt of valuesTarget) {
          if (kv.value == vt) {
            valuesTargetArrayObject = [...valuesTargetArrayObject, kv];
          }
        }
      }

      valuesSourceArrayObject.sort((a, b) => a.label.localeCompare(b.label));
    }
  }

  onMount(() => {
    reStructureSvelteValues();  
  });

  $: if (valuesTarget) {
    reStructureSvelteValues();
  }
</script>

<div class={className}>
  <div class="input-box">
    <label class="label" for={id}>{label}</label>
    {#if valuesSourceType == 'array'}
      <Select
        multiple
        items={valuesSource}
        on:select={handleSelect}
        on:clear={() => (valuesTarget = [])}
        placeholder={placeholder}
      />
    {/if}
    {#if valuesSourceType === 'object'}
      <Select
        multiple
        items={valuesSourceArrayObject}
        bind:value={valuesTargetArrayObject}
        on:select={handleSelect}
        on:clear={handleClear}
        placeholder={placeholder}
      />
    {/if}
  </div>
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

