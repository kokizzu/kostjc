<script>
  /** @typedef {import('../_types/masters').InputType} InputType*/

  import { onMount } from 'svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { AiOutlineEye, AiOutlineEyeInvisible } from '../node_modules/svelte-icons-pack/dist/ai';
  import { dateISOFormat } from './xFormatter';
  import Select from '../node_modules/svelte-select';

  export let className = '';
  export let type = /** @type {InputType | any} */ ('text');
  export let id;
  export let value;
  export let autocomplete = 'off';

  export let values = /** @type {
    Array<string|number> | Record<string|number, string> | any
  } */ ({});
  export let label;
  export let placeholder = '';
  export let isObject = false;

  if (isObject) value = value+'';
  let isShowPassword = false;
  let inputElm;
  
  onMount(() => {
    if (type === 'datetime' && value == '') {
      value = dateISOFormat(0);
      if (id == 'ktpDateBirth') value = dateISOFormat(-6209.25);
    }
    if (type === 'password') inputElm.type = type;
    // Boolean input must be use random id, because it's a checkbox
    if (type === 'bool') id = id + Math.random();
  });

  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text';
    else inputElm.type = 'password';
  }

  let itemsArr = [];
  let itemsArrObj = [];

  let svelteSelectValue;

  onMount(() => {
    if (type === 'combobox-arr') {
      itemsArr = values;
      for (const v of values) {
        if (v == value) {
          svelteSelectValue = v;
        }
      }
    } else if (type === 'combobox') {
      for (const [k, v] of Object.entries(values)) {
        itemsArrObj = [...itemsArrObj, {
          value: k,
          label: v
        }];
        if (k == value) {
          svelteSelectValue = {
            value: k,
            label: v
          };
        }
      }
    }
  });

  function handleSelect(/** @type {CustomEvent} */e) {
    value = e.detail.value;
  }
</script>

<div class={className}>
  <div class="input-box {type == 'bool' ? 'bool' : ''} {type == 'password' ? 'with-password' : ''}">
    {#if type === 'bool' || type === 'checkbox'}
      <label class="label" for={id}>{label}</label>
      <label class="switcher" for={id}>
        <input type="checkbox" id={id} bind:checked={value}>
        <span class="slider"></span>
      </label>
    {:else if type === 'combobox-arr' || type === 'select'}
      <label class="label" for={id}>{label}</label>
      <Select
        items={itemsArr}
        bind:value={svelteSelectValue}
        on:select={handleSelect}
        placeholder={placeholder}
      />
    {:else if type === 'combobox'}
      {#if isObject}
        <label class="label" for={id}>{label}</label>
        <Select
          items={itemsArrObj}
          bind:value={svelteSelectValue}
          on:select={handleSelect}
          placeholder={placeholder}
        />
      {:else}
        <label class="label" for={id}>{label}</label>
        <Select
          items={itemsArr}
          bind:value={svelteSelectValue}
          on:select={handleSelect}
          placeholder={placeholder}
        />
      {/if}
    {:else if type === 'number'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder} {autocomplete}/>
    {:else if type === 'datetime'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder} />
    {:else if type === 'percentage'}
      <label class="label" for={id}>{label}</label>
      <div class="input-percentage">
        <input type="number" bind:value={value} {id} {placeholder} />
        <span>%</span>
      </div>
    {:else if type === 'float'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder} {autocomplete}/>
    {:else if type === 'textarea'}
      <label class="label" for={id}>{label}</label>
      <textarea bind:value={value} {id} {placeholder}></textarea>
    {:else if type === 'text'}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder} {autocomplete}/>
    {:else if type === 'email'}
      <label class="label" for={id}>{label}</label>
      <input type="email" bind:value={value} {id} {placeholder} {autocomplete}/>
    {:else if type === 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'password'}
      <label class="label" for={id}>{label}</label>
      <input bind:value={value} {id} bind:this={inputElm} {placeholder} autocomplete="off"/>
      {#if type === 'password'}
        <button class="eye" on:click={toggleShowPassword}>
          {#if !isShowPassword}
            <Icon color="#495057" size="20" src={AiOutlineEye}/>
          {/if}
          {#if isShowPassword}
            <Icon color="#495057" size="20" src={AiOutlineEyeInvisible}/>
          {/if}
        </button>
      {/if}
    {:else if type === 'color'}
      <label class="label" for={id}>{label}</label>
      <div class="color_box">
        <input type="color" bind:value={value} {id} class="color-input"/>
      </div>
    {:else}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder} {autocomplete}/>
    {/if}
  </div>
</div>

<style>
  .show {
    display: block;
  }

  .hidden {
    display: none;
  }

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

  .input-box.with-password input{
    padding-right: 40px !important;
  }

  .input-box.bool {
    width: fit-content;
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

  .input-box input,
  .input-box textarea {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input-box input:hover,
  .input-box textarea:hover {
    border: 1px solid var(--gray-005);
  }

  .input-box input:focus,
  .input-box textarea:focus {
    border-color: var(--blue-005);
    outline: 1px solid var(--blue-005);
  }

  .color_box {
    display: flex;
    flex-direction: column;
    justify-content: start;
  }

  .color-input {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    width: 100px !important;
    height: 30px;
    background-color: transparent;
    border: none;
    cursor: pointer;
    padding: 0 !important;
  }
  .color-input::-webkit-color-swatch {
    border-radius: 5px;
    border: none;
  }
  .color-input::-moz-color-swatch {
    border-radius: 5px;
    border: none;
  }

  .input-box textarea {
    resize: vertical;
    height: 90px;
    min-height: 50px;
    max-height: 300px;
  }
  
  .switcher {
    position: relative;
    display: inline-block;
    width: 43px;
    height: 24px;
    margin-left: 0 !important;
  }
  
  .switcher input {
    opacity: 0 !important;
    width: 0 !important;
    height: 0 !important;
  }
  
  .switcher .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: var(--gray-004);
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher .slider:before {
    position: absolute;
    content: "";
    height: 16px;
    width: 16px;
    left: 4px;
    bottom: 4px;
    background-color: white;
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher input:checked + .slider {
    background-color: var(--blue-006) !important;
  }

  .switcher input:focus + .slider {
    box-shadow: 0 0 1px var(--blue-006) !important;
  }

  .switcher input:checked + .slider:before {
    -webkit-transform: translateX(19px);
    -ms-transform: translateX(19px);
    transform: translateX(19px);
  }
  
  .switcher .slider {
    border-radius: 34px;
  }

  .switcher .slider:before {
    border-radius: 50%;
  }

  .input-box .eye {
    position: absolute;
    height: fit-content;
    width: fit-content;
    background-color: transparent;
    padding: 0;
    top: 31px;
    bottom: auto;
    right: 10px;
    border: none;
    cursor: pointer;
  }

  .input-box .eye:focus {
    outline: none;
  }

  :global(.input-box .eye:hover svg) {
    fill: var(--blue-005);
  }

  .input-percentage {
    display: flex;
    position: relative;
  }

  .input-percentage input {
    padding-right: 30px !important;
  }

  .input-percentage span {
    position: absolute;
    right: 10px;
    bottom: 10px;
    font-weight: 700;
  }

  .input-box .options-container .input-container input {
    padding-right: 30px !important;
  }

  :global(.input-box .options-container .input-container .arrow .dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.input-box .options-container .input-container .arrow .rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(90deg);
	}
</style>