<script>
  import { onMount } from 'svelte';
  import Select from '../node_modules/svelte-select';
  export let id;
  export let label;
  export let placeholder = '';
  export let valuesTarget = [];
  export let valuesSourceObj = {};

  let selected = null;

  $: valuesSourceArrayObject = Object.entries(valuesSourceObj).map(([k, v]) => ({
    value: Number(k),
    label: v
  }));

  function handleSelect(e) {
    const val = Number(e.detail.value);
    valuesTarget = [...valuesTarget, val];
    selected = null;
  }

  function removeItem(index) {
    valuesTarget.splice(index, 1);
    valuesTarget = [...valuesTarget];
  }
</script>

<div class="input-box">
  <label class="label" for={id}>{label}</label>
  <Select
    items={valuesSourceArrayObject}
    bind:value={selected}
    on:select={handleSelect}
    placeholder={placeholder}
    clearable={false}
  />

  <div class="selected-list">
    {#each valuesTarget as val, index}
      <div class="selected-item">
        {valuesSourceObj[val]} 
        <button on:click={() => removeItem(index)}>x</button>
      </div>
    {/each}
  </div>
</div>

<style>
  .input-box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    gap: 6px;
    margin-bottom: 1rem;
  }

  .input-box .label {
    font-size: var(--font-base);
    margin-left: 10px;
  }

  .selected-list {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    margin-top: 6px;
  }

  .selected-item {
    background-color: var(--gray-002, #f2f2f2);
    padding: 4px 8px;
    border-radius: 6px;
    font-size: var(--font-base);
    display: flex;
    align-items: center;
    gap: 4px;
    border: 1px solid #ccc;
  }

  .selected-item button {
    background: transparent;
    border: none;
    cursor: pointer;
    font-size: large;
    font-weight: bolder;
    color: red;
    padding: 0;
  }
</style>
