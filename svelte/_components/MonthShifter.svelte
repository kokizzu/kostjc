<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiArrowsArrowRightSLine, RiArrowsArrowLeftSLine } from '../node_modules/svelte-icons-pack/dist/ri';

  /**
   * @description Format YYYY-MM from Date
   * @param {Date} date
   * @returns {string}
   */
  function formatYM(date) {
    const y = date.getFullYear();
    const m = String(date.getMonth() + 1).padStart(2, '0');
    return `${y}-${m}`;
  }

  /**
   * @description Parse YYYY-MM to Date
   * @param {string} ym
   * @returns {Date}
   */
  function parseYM(ym) {
    const [y, m] = ym.split('-').map(Number);
    return new Date(y, m - 1);
  }

  const now = new Date();
  const initialMonth = new Date(now.getFullYear(), now.getMonth());
  
  export let yearMonth = formatYM(initialMonth);
  export let OnChanges = async () => {};

  function formatMonthYear(/** @type {string} */ ym) {
    const date = new Date(ym + "-01");
    return date.toLocaleString('default', { month: 'long', year: 'numeric' });
  }

  async function shiftMonth(/** @type {number} */ direction) {
    const end = parseYM(yearMonth);
    end.setMonth(end.getMonth() + direction);
    yearMonth = formatYM(end);

    await OnChanges();
  }
</script>

<div class="month-shifter">
  <button on:click={() => shiftMonth(-1)} class="btn" aria-label="Previous">
    <Icon size="20" src={RiArrowsArrowLeftSLine} />
  </button>
  <span class="month-text">{formatMonthYear(yearMonth)}</span>
  <button on:click={() => shiftMonth(1)} class="btn" aria-label="Next">
    <Icon size="20" src={RiArrowsArrowRightSLine} />
  </button>
</div>

<style>
  .month-shifter {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 16px;
  }

  .month-shifter .btn {
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

  .month-shifter .btn:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-006);
  }

  .month-shifter .month-text {
    font-weight: 600;
    font-size: 18px;
    user-select: none;
  }
</style>