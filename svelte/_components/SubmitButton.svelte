<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemLoaderLine } from '../node_modules/svelte-icons-pack/dist/ri';

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let isFullWidth  = /** @type {boolean} */ (false);
  export let label        = /** @type {string} */ ('Submit');
</script>

<button class={isFullWidth ? 'submit-btn full-width' : 'submit-btn'} on:click|preventDefault>
  {#if isSubmitted}
    <Icon className="spin" color="#FFF" size="15" src={RiSystemLoaderLine} />
  {/if}
  <span>{label}</span>
</button>

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

  .full-width {
    width: 100% !important;
    justify-content: center;
  }

  .submit-btn {
    border: none;
    width: fit-content;
    padding: 8px 15px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #fff;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    cursor: pointer;
  }

  .submit-btn:hover {
    background-color: var(--blue-005);
  }
</style>
