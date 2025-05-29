<script>
  /** @typedef {import('../_types/masters').RadioOption<any>} RadioOption */
  
  export let selected = /** @type {string} */ ('');
  export let options = /** @type {RadioOption[]} */ ([]);
  export let className = '';
</script>

<div class="{className}">
  {#each options as option}
    <div class="switcher-container">
      <label class="switcher" for={option.id}>
        <input
          type="radio"
          id={option.id}
          name={option.name}
          bind:group={selected}
          value={option.value}
        />
        <span class="slider"></span>
      </label>
      <label for="{option.id}" class="label">{option.label}</label>
    </div>
  {/each}
</div>

<style>
  .switcher-container {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
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
    background-color: var(--blue-005) !important;
  }

  .switcher input:focus + .slider {
    box-shadow: 0 0 1px var(--blue-005) !important;
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
</style>