<script>
  /** @typedef {import('../_types/masters').InputType} InputType*/

  import { onMount } from 'svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { AiOutlineEye, AiOutlineEyeInvisible } from '../node_modules/svelte-icons-pack/dist/ai';
  import { RandString } from './xGenerator';
  import { RiArrowsArrowRightSLine } from '../node_modules/svelte-icons-pack/dist/ri';

  export let className = '';
  export let type = /** @type {InputType | any} */ ('text');
  export let id;
  export let value;

  export let values = /** @type {
    Array<string|number> | Record<string|number, string> | any
  } */ ({});
  export let label;
  export let placeholder = '';
  export let isObject = false;

  if (isObject) value = value+'';
  let isShowPassword = false;
  let inputElm;

  let valueToShowFromObj = value+'';

  const randStr = RandString(5);
  
  onMount(() => {
    if (type === 'password') inputElm.type = type;
    // Boolean input must be use random id, because it's a checkbox
    if (type === 'bool') id = id + Math.random();
    if (isObject) {
      if (type === 'combobox' || type === 'select') {
        try {
          const valuesArr = Object.entries(values);
          value = valuesArr[0][0];
        } catch (error) {}
      }
      if (values && values[value]) valueToShowFromObj = values[value];
      else valueToShowFromObj = '';
    } else {
      if (type === 'combobox' || type === 'select') {
        value = values[0];
      }
      valueToShowFromObj = value+'';
    }
  });

  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text';
    else inputElm.type = 'password';
  }

  let isShowOptions = false;
  let currentFocus = -1;

  function filterOptions(options) {
    if (valueToShowFromObj === '') {
      options.forEach((option) => {
        option.style.display = 'block';
      });
      currentFocus = -1;
      return;
    } else {
      options.forEach((option) => {
        const textValue = option.textContent; // @ts-ignore
        if (textValue.toUpperCase().indexOf(valueToShowFromObj.toUpperCase()) > -1) { // @ts-ignore
          option.style.display = 'block';        
        } else { // @ts-ignore
          option.style.display = 'none'; 
        }
      });
    }
  }

  function highlightOption(options, isIncreased) {
    if (!options.length) return;

    if (isIncreased) {
      currentFocus++;
    } else {
      currentFocus--;
    }

    if (currentFocus < 0) currentFocus = options.length - 1;
    if (currentFocus > options.length - 1) currentFocus = 0;

    removeActive(options);

    if(options[currentFocus] && options[currentFocus].style.display === 'none') {
      highlightOption(options, isIncreased);
    };

    if (options[currentFocus]) {
      options[currentFocus].classList.add('active');
      options[currentFocus].scrollIntoView({block: 'nearest'});
    }
  }

  function removeActive(options) {
    if (!options.length) {
      options = document.querySelectorAll('.option.'+randStr);
    };
    options.forEach(option => option.classList.remove('active'));
  }

  function handleKey(/** @type {KeyboardEvent} */e) {
    const options = document.querySelectorAll('.option.'+randStr);

    switch (e.key) {
      case 'ArrowDown':
        highlightOption(options, true);
        break;
      case 'ArrowUp':
        highlightOption(options, false);
        break;
      case 'Enter':
        if (currentFocus > -1) { // @ts-ignore
          options[currentFocus].click();
          removeActive(options);
          optionClicked = false;
        }
        break;
      default:
        filterOptions(options);
        break;
    }
  }

  let optionClicked = false;
</script>

<div class={className}>
  <div class="input_box {type == 'bool' ? 'bool' : ''} {type == 'password' ? 'with_password' : ''}">
    {#if type === 'bool' || type === 'checkbox'}
      <label class="label" for={id}>{label}</label>
      <label class="switcher" for={id}>
        <input type="checkbox" id={id} bind:checked={value}>
        <span class="slider"></span>
      </label>
    {:else if type == 'select' || type === 'combobox'}
      {#if isObject}
        <label class="label" for={id}>{label}</label>
        <div class={`options_container ${randStr}`}>
          <div class="input_container">
            <input
              type="text"
              bind:value={valueToShowFromObj}
              on:focus={() => isShowOptions = true}
              on:blur|preventDefault={() => {
                console.log('isOptionClicked', optionClicked);
                currentFocus = -1;
                const options = document.querySelectorAll('.option .'+randStr);
                removeActive(options);
                valueToShowFromObj = values[value];
                setTimeout(() => {
                  if (optionClicked) {
                    isShowOptions = false;
                  } else {
                    setTimeout(() => {
                      isShowOptions = false;
                      optionClicked = false;
                    }, 100);
                  }
                }, 200);
              }}
              on:keyup={handleKey}
            />
            <button class="arrow" on:click={() => isShowOptions = !isShowOptions}>
              <Icon
                className="icon {isShowOptions ? 'rotate' : 'dropdown'}"
                color="var(--gray-007)"
                src={RiArrowsArrowRightSLine}
                size="17"
              />
            </button>
          </div>
          <div class="options_list {isShowOptions ? 'show' : 'hidden'}">
            {#each Object.entries(values) as [k, v]}
              <button class="option {randStr}" on:click={() => {
                value = k;
                valueToShowFromObj = v;
                isShowOptions = false;
                optionClicked = true;
              }}>
                {v}
              </button>
            {/each}
          </div>
        </div>
      {:else}
        <label class="label" for={id}>{label}</label>
        <div class="options_container {randStr}">
          <div class="input_container">
            <input
              type="text"
              bind:value={valueToShowFromObj}
              on:focus={() => isShowOptions = true}
              on:blur={() => {
                currentFocus = -1;
                const options = document.querySelectorAll('.option .'+randStr);
                removeActive(options);
                valueToShowFromObj = value || values[value] || values[0] || options[0].textContent; ;
                if (isShowOptions) optionClicked = true;
                setTimeout(() => {
                  if (optionClicked) {
                    isShowOptions = false;
                  } else {
                    setTimeout(() => {
                      isShowOptions = false;
                      optionClicked = false;
                    }, 100);
                  }
                }, 200);
              }}
              on:keyup={handleKey}
            />
            <button class="arrow" on:click={() => isShowOptions = !isShowOptions}>
              <Icon
                className="icon {isShowOptions ? 'rotate' : 'dropdown'}"
                color="var(--gray-007)"
                src={RiArrowsArrowRightSLine}
                size="17"
              />
            </button>
          </div>
          <div class="options_list {isShowOptions ? 'show' : 'hidden'}">
            {#each values as v}
              <button class="option {randStr}" on:click|preventDefault={() => {
                value = v;
                valueToShowFromObj = v;
                isShowOptions = false;
              }}>
                {v}
              </button>
            {/each}
          </div>
        </div>
      {/if}
    {:else if type === 'number'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder}/>
    {:else if type === 'datetime'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'percentage'}
      <label class="label" for={id}>{label}</label>
      <div class="input_percentage">
        <input type="number" bind:value={value} {id} {placeholder} />
        <span>%</span>
      </div>
    {:else if type === 'float'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder}/>
    {:else if type === 'textarea'}
      <label class="label" for={id}>{label}</label>
      <textarea bind:value={value} {id} {placeholder}></textarea>
    {:else if type === 'text'}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder} autocomplete="off" />
    {:else if type === 'email'}
      <label class="label" for={id}>{label}</label>
      <input type="email" bind:value={value} {id} {placeholder}/>
    {:else if type === 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'password'}
      <label class="label" for={id}>{label}</label>
      <input bind:value={value} {id} bind:this={inputElm} {placeholder}/>
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
      <input type="text" bind:value={value} {id} {placeholder}/>
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

  .input_box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input_box.with_password input{
    padding-right: 40px !important;
  }

  .input_box.bool {
    width: fit-content;
  }

  .input_box .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .input_box input,
  .input_box textarea {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input_box input:focus,
  .input_box textarea:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
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

  .input_box textarea {
    resize: vertical;
    height: 90px;
    min-height: 50px;
    max-height: 300px;
  }
  
  /* The switch - the box around the slider */
  .switcher {
    position: relative;
    display: inline-block;
    width: 50px;
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
    background-color: var(--gray-005);
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
    background-color: var(--sky-006) !important;
  }

  .switcher input:focus + .slider {
    box-shadow: 0 0 1px var(--sky-006) !important;
  }

  .switcher input:checked + .slider:before {
    -webkit-transform: translateX(26px);
    -ms-transform: translateX(26px);
    transform: translateX(26px);
  }
  
  .switcher .slider {
    border-radius: 34px;
  }

  .switcher .slider:before {
    border-radius: 50%;
  }

  .input_box .eye {
    position: absolute;
    height: fit-content;
    width: fit-content;
    background-color: transparent;
    padding: 0;
    top: 30px;
    bottom: auto;
    right: 10px;
    border: none;
    cursor: pointer;
  }

  :global(.input_box .eye:hover svg) {
    fill: var(--blue-005);
  }

  .input_percentage {
    display: flex;
    position: relative;
  }

  .input_percentage input {
    padding-right: 30px !important;
  }

  .input_percentage span {
    position: absolute;
    right: 10px;
    bottom: 10px;
    font-weight: 700;
  }

  .input_box .options_container {
    display: block;
    position: relative;
    width: 100%;
  }

  .input_box .options_container .input_container {
    position: relative;
    width: 100%;
  }

  .input_box .options_container .input_container input {
    padding-right: 30px !important;
  }

  .input_box .options_container .input_container .arrow {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    right: 0;
    top: 0;
    bottom: 0;
    background-color: transparent;
    border-top: 2px solid transparent;
    border-bottom: 2px solid transparent;
    border-left: 1px solid transparent;
    border-right: 2px solid transparent;
    cursor: pointer;
    padding: 8px;
    height: 100%;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
  }

  .input_box .options_container .input_container .arrow:hover {
    background-color: var(--gray-001);
    border-top: 2px solid var(--gray-003);
    border-bottom: 2px solid var(--gray-003);
    border-left: 1px solid var(--gray-003);
    border-right: 2px solid var(--gray-003);
    padding: 12px;
    outline: 1px solid var(--gray-003);
  }

  :global(.input_box .options_container .input_container .arrow .dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.input_box .options_container .input_container .arrow .rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(90deg);
	}

  .input_box .options_container .options_list {
    margin-top: 5px;
    position: absolute;
    top: 100%;
    left: 0;
    width: 100%;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    height: fit-content;
    max-height: 250px;
    overflow-y: auto;
    z-index: 1999;
  }

  .input_box .options_container .options_list .option {
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
    padding: 5px 10px;
    cursor: pointer;
    background-color: #FFF;
    border: none;
    height: fit-content;
    width: 100%;
    text-align: left;
  }

  .input_box .options_container .options_list .option:hover {
    background-color: var(--gray-002);
  }

  :global(.input_box .options_container .options_list .option.active) {
    background-color: var(--gray-002) !important;
  }
</style>