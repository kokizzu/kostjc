<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiArrowsArrowRightSLine, RiArrowsArrowLeftSLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import { localeDateFromYYYYMMDD } from './xFormatter';
  
  export let date = ''; // format YYYY-MM-DD
  export let onChanges = async () => {};
  export let isLoading = false;

  $: dateObj = date ? new Date(date) : new Date();
  $: dateToShow = localeDateFromYYYYMMDD(dateObj.toISOString());

  async function shiftDate(/** @type {number} */ direction) {
    const shifted = new Date(dateObj);
    shifted.setDate(shifted.getDate() + direction);

    date = shifted.toISOString().split("T")[0];

    await onChanges();
  }
</script>

<div class="date-shifter">
  <button disabled={isLoading} on:click={() => shiftDate(-1)} class="btn" aria-label="Previous">
    <Icon size="20" src={RiArrowsArrowLeftSLine} />
  </button>
  <span class="date-text">{dateToShow}</span>
  <button disabled={isLoading} on:click={() => shiftDate(1)} class="btn" aria-label="Next">
    <Icon size="20" src={RiArrowsArrowRightSLine} />
  </button>
</div>

<style>
  .date-shifter {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
  }

  .date-shifter .btn {
    background-color: var(--gray-002);
    border: none;
    border-radius: 9999px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--gray-008);
    padding: 8px;
  }

  .date-shifter .btn:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }

  .date-shifter .date-text {
    font-weight: 600;
    font-size: 18px;
    user-select: none;
  }
</style>