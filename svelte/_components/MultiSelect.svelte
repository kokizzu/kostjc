<script>
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
    valuesTarget = [...valuesTarget, e.detail.value];
  }

  let valuesSourceArrayObject = [];
  onMount(() => {
    for (const [k, v] of Object.entries(valuesSourceObj)) {
      valuesSourceArrayObject.push({
        value: k,
        label: v
      });
    }
  })
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
        on:select={handleSelect}
        on:clear={() => (valuesTarget = [])}
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

