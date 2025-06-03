<script>
  /** @typedef {import('../_types/masters.js').Field} Field */
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/masters.js').ExtendedAction} ExtendedAction */
  /** @typedef {import('../_types/masters.js').ExtendedActionButton} ExtendedActionButton */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiDesignBallPenLine,
    RiSystemDeleteBin5Line,
    RiSystemFilterLine,
    RiArrowsArrowGoBackLine,
    RiSystemInformationLine,
    RiArrowsExpandUpDownFill,
    RiArrowsArrowDownSFill,
    RiArrowsArrowUpSFill
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import {
    IoSearch,
    IoClose
  } from '../node_modules/svelte-icons-pack/dist/io';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import {
    CgChevronLeft,
    CgChevronRight,
    CgChevronDoubleRight,
    CgChevronDoubleLeft,
  } from '../node_modules/svelte-icons-pack/dist/cg';
  import InputBox from './InputBox.svelte';
  import { onMount } from 'svelte';
  import { datetime, formatPrice } from './xFormatter.js';
  import FilterTable from './FilterTable.svelte';
  import MultiSelect from './MultiSelect.svelte';

  export let FIELDS = /** @type Field[] */ ([]); // bind
  export let PAGER = /** @type PagerOut */ ({}); // bind
  export let MASTER_ROWS = /** @type any[][] */ ([]); // bind
  export let REFS = {};

  export let NAME = ''
  export let ACCESS = /** @type Access */ ({});
  export let ARRAY_OF_ARRAY = true;
  export let CAN_SEARCH_ROW = true;
  export let CAN_EDIT_ROW = true;
  export let CAN_DELETE_ROW = false;
  export let CAN_RESTORE_ROW = false;
  export let CAN_SHOW_INFO = false;
  export let CAN_OPEN_LINK = false;
  export let LINKS = /** @type ExtendedAction[] */ ([]);
  export let IS_CUSTOM_EDIT = false;
  export let UNSORTED_ROWS = [];
  export let EXTENDED_BUTTONS = /** @type {ExtendedActionButton[]} */ ([]);
  export let FIELD_TO_SEARCH = '';

  /**
   * @type {Record<string, number>}
   */
  export let COL_WIDTHS = {}

  // State for loading if hit ajax
  let isAjaxSubmitted = false;

  // Binding component FilterTable.svelte
  let filterTable = null;
  // For readiness of component FilterTable.svelte, prevent race condition
  let filterTableReady = false;
  // Key and label of column to filter
  let filterColumns = [];
  // Binding value of column, for payloaf
  let filtersMap = {};

  // Index of field 'deletedAt', for marker deleted rows
  let deletedIndex = -1;

  // PopUp for modify rows
  let showPopUp = false;

  // Pagination total, based on total pages
  let paginationTotal = 1;
  // Pagination all, based on total pages
  let paginationsAll = /** @type number[] */ ([]);
  // Pagination to show, based on total pagination
  let paginationShow = /** @type number[] */ ([]);
  // Current page describe which page is currently rendered
  let currentPage = 1;
  // State for sort, wheter is ascending or descending
  let isSortTableAsc = true;
  // State for sort field name
  let fieldNameToSort = (
    (PAGER.order || []).length ? PAGER.order[0]
      .substring(1)
      .replace('+', '')
      .replace('-', '')
      : ''
  )
  // Total Pages
  let totalPages = 0;
  // Total rows but rounded by current rows
  let totalRound = 0;
  // Rows per page
  let currentRows = PAGER.perPage;
  // Rows per page options
  let rowsToShow = [5, 10, 20, 50, 70, 100, 200];
  // State for show rows options
  let showRowsNum = false;
  // Total rows
  let totalRows = PAGER.countResult;
  // Total rows current
  let totalRowsCurrent = 0;
  // Payloads for modify rows
  let payloads = [];

  // Toggle show rows options
  function toggleRowsNum() {
    showRowsNum = !showRowsNum;
  }

  // Refresh Pagination
  function getPaginationShow() {
    totalRows = PAGER.countResult;
    totalPages = PAGER.pages;
    currentPage = PAGER.page;
    if (MASTER_ROWS && MASTER_ROWS.length) {
      totalRowsCurrent = MASTER_ROWS.length;
    } else {
      totalRowsCurrent = 0;
    }

    totalRound = Math.ceil(totalRows / currentRows) * currentRows;
    paginationTotal = totalRound / currentRows;
    paginationsAll = [];
    if (currentRows > PAGER.countResult) paginationTotal = 1;
    for (let i = 0; i < paginationTotal; i++) {
      paginationsAll = [...paginationsAll, i + 1];
    }

    let start = 0, end = 0;
    if (paginationTotal < 5) {
      (start = 0), (end = paginationTotal);
    } else if (currentPage < 5 && currentPage - 3 < 0) {
      (start = 0), (end = 5);
    } else if (currentPage > paginationTotal - 5 && currentPage + 3 < paginationTotal) {
      (start = currentPage - 3), (end = currentPage + 2);
    } else if (currentPage + 3 >= paginationTotal) {
      (start = paginationTotal - 5), (end = paginationTotal);
    } else {
      (start = currentPage - 3), (end = currentPage + 2);
    }

    paginationShow = paginationsAll.slice(start, end);
  }

  /**
   * @param {any} row
   * @param {number} i
   * @param {Field} field
   */
  function Cell(row, i, field) {
    if (ARRAY_OF_ARRAY) return row[i] || '';
    return row[field.name] || '';
  }

  // Trigger function "getPaginationShow()" if variable "PAGER" changed
  $: {
    if (PAGER) getPaginationShow();
  }

  onMount(() => {
    // FilterTable.svelte component is rendered
    filterTable = FilterTable;
    // FilterTable.svelte component is ready
    filterTableReady = true;
    // Total pages
    totalPages = PAGER.pages;
    // Loop column/fields, fill variable filterColumn for filters
    if (FIELDS && FIELDS.length > 0) {
      filterColumns = [];
      FIELDS.forEach((col, idx) => {
        if (col.name === 'deletedAt') deletedIndex = idx;
        filterColumns = [
          ...filterColumns,
          {
            key: col.name,
            label: col.label,
          },
        ];
      });
    }
    // Calculate pagination
    getPaginationShow();
    // Fill initial payloads
    if (FIELDS && FIELDS.length > 0) {
      FIELDS.forEach(() => (payloads = [...payloads, '']));
    }
  });

  // Export function, forward parameter to parent
  export let OnRestore = async function (/** @type any[]*/ row) {};
  export let OnDelete = async function (/** @type any[]*/ row) {};
  export let OnEdit = async function (/** @type any */ id, /** @type any[]*/ payloads) {};
  export let OnRefresh = async function (/** @type PagerIn */ pagerIn) {};
  export let OnInfo = async function (/** @type any[] */ row) {};

  function ApplyFilter() {
    // Hide FilterTable.svelte
    filterTable.Hide();
    // Make a 'filters' payload from variable filtersMap
    // Make it with format { key: [value, value] }
    let filters = /** @type {Record<string, string[]>} */ ({});
    for (let key in filtersMap) {
      let value = filtersMap[key];
      if (value) filters[key] = value.split('|');
    }
    OnRefresh({ ...PAGER, filters });
    // Refresh pagination view
    getPaginationShow();
  }

  // Apply row counts
  async function toRow(/** @type number */ perPage) {
    currentRows = perPage;
    showRowsNum = false;
    await OnRefresh({ ...PAGER, perPage });
    // Refresh pagination view
    getPaginationShow();
  }

  // Go to page, last page, first page
  async function goToPage(/** @type number */ page) {
    currentPage = page;
    await OnRefresh({ ...PAGER, page });
    // Refresh pagination view
    getPaginationShow();
  }

  // Restore row
  async function restoreRow(/** @type any[] */ row) {
    await OnRestore(row);
    // Refresh pagination view
    getPaginationShow();
  }

  // Delete row
  async function deleteRow(/** @type any[] */ row) {
    await OnDelete(row);
    // Refresh pagination view
    getPaginationShow();
  }

  // Row ID to modify
  let idToMod = '';

  function toggleShowPopUp(/** @type any */ id, /** @type any[]*/ row) {
    payloads = [];
    if (FIELDS && FIELDS.length > 0) {
      FIELDS.forEach((f, i) => {
        payloads = [...payloads, row[i]];
      });
    }

    showPopUp = true;
    idToMod = id;
  }

  function handleSubmitEdit() {
    OnEdit(idToMod, payloads);
    showPopUp = false;
  }

  function formatFacilities(facArrInt, facObjs) {
    if (!facArrInt.length) return '--';
    return facArrInt.map(fId => facObjs[Number(fId)]).join(', ');
  }

  async function OnSort(/** @type {Field} */ field) {
    isSortTableAsc = !isSortTableAsc;
    fieldNameToSort = field.name;
    if (isSortTableAsc) {
      PAGER.order = ['+' + field.name];
    } else {
      PAGER.order = ['-' + field.name];
    }
    await OnRefresh(PAGER);
    // Refresh pagination view
    getPaginationShow();
  }

  let searchValue = /** @type {string} */ ('');
  async function handleSearch() {
    PAGER.filters = {
      [FIELD_TO_SEARCH]: [searchValue]
    }
    await OnRefresh(PAGER);
  }
</script>

{#if filterTableReady}
  <FilterTable bind:this={filterTable} bind:filterColumns bind:filtersMap on:click={ApplyFilter} />
{/if}

{#if showPopUp}
  <div class="popup-container">
    <div class="popup">
      <header>
        <h2>Edit {NAME ? NAME : 'row'} {`#${idToMod}`}</h2>
        <button on:click={() => (showPopUp = false)}>
          <Icon size="22" color="var(--red-005)" src={IoClose} />
        </button>
      </header>
      <div class="forms">
        {#each (FIELDS || []) as field, idx}
          {#if field.name !== 'id'}
            {#if !field.readOnly}
              {#if field.inputType === 'combobox'}
                <InputBox
                  id={field.name}
                  label={field.label}
                  placeholder={field.description}
                  bind:value={payloads[idx]}
                  type={field.inputType}
                  values={REFS && REFS[field.name] ? REFS[field.name] : field.ref}
                  isObject={REFS && REFS[field.name] ? true : false}
                />
              {:else if field.inputType === 'combobox-arr'}
                <InputBox
                  id={field.name}
                  label={field.label}
                  placeholder={field.description}
                  bind:value={payloads[idx]}
                  type={field.inputType}
                  values={REFS && REFS[field.name] ? REFS[field.name] : field.ref}
                />
              {:else if field.inputType === 'multiselect'}
                <MultiSelect
                  id={field.name}
                  label={field.label}
                  placeholder={field.description}
                  bind:valuesTarget={payloads[idx]}
                  valuesSourceType="object"
                  valuesSourceObj={REFS && REFS[field.name] ? REFS[field.name] : field.ref}
                />
              {:else}
                <InputBox
                  id={field.name}
                  label={field.label}
                  placeholder={field.description}
                  bind:value={payloads[idx]}
                  type={field.inputType}
                />
              {/if}
            {/if}
          {/if}
        {/each}
      </div>
      <div class="foot">
        <button class="ok" on:click={handleSubmitEdit}>Submit</button>
      </div>
    </div>
  </div>
{/if}

<div class="table-root">
  <div class="actions-container">
    <div class="left">
      <div class="actions-button">
        <button class="btn" on:click={() => filterTable.Show()} title="filter table">
          <Icon color="var(--gray-007)" size="18" src={RiSystemFilterLine} />
        </button>
        <!-- Action buttons -->
        <slot />
      </div>
      {#if isAjaxSubmitted}
        <div class="loader">
          <Icon className="spin" color="var(--blue-007)" size="28" src={FiLoader} />
        </div>
      {/if}
    </div>
    <div class="right">
      {#if CAN_SEARCH_ROW}
        <div class="search-handler">
          <button class="search-btn" title="Search" on:click={handleSearch}>
            <Icon color="var(--gray-007)" size="16" src={IoSearch} />
          </button>
          <input
            placeholder="Search..."
            type="text"
            name="searchRow"
            id="searchRow"
            class="search"
            bind:value={searchValue}
          />
        </div>
      {/if}
    </div>
  </div>
  <div class="table-container">
    <table>
      <thead>
        <tr>
          <th class="no sticky">No</th>
          {#each FIELDS || [] as f, _ (f.name)}
            {#if f.name === 'id'}
              <th class="a-row">Actions</th>
            {:else}
              <th
                style="{COL_WIDTHS[f.name] ? `min-width: ${COL_WIDTHS[f.name]}px;` : ''}"
                class="
                  {f.inputType === 'textarea' ? 'textarea' : ''}
                  {f.inputType === 'datetime' ? 'datetime' : ''}
                ">
                <button class="heading" on:click={() => OnSort(f)} disabled={UNSORTED_ROWS.includes(f.name)}>
                  <span>{f.label}</span>
                  {#if !UNSORTED_ROWS.includes(f.name)}
                    {#if isSortTableAsc && f.name === fieldNameToSort}
                      <Icon
                        className="sort-icon"
                        size="13"
                        color="var(--gray-007)"
                        src={RiArrowsArrowDownSFill}
                      />
                    {/if}
                    {#if !isSortTableAsc && f.name === fieldNameToSort}
                      <Icon
                        className="sort-icon"
                        size="13"
                        color="var(--gray-007)"
                        src={RiArrowsArrowUpSFill}
                      />
                    {/if}
                    {#if f.name !== fieldNameToSort}
                      <Icon
                        className="sort-icon"
                        size="13"
                        color="var(--gray-007)"
                        src={RiArrowsExpandUpDownFill}
                      />
                    {/if}
                  {/if}
                </button>
              </th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#if MASTER_ROWS && MASTER_ROWS.length > 0}
          {#each MASTER_ROWS as row}
            <tr class={CAN_DELETE_ROW && (row[deletedIndex] > 0 || row[deletedIndex] === 'terminated') ? 'deleted' : ''}>
              <td class="num-row sticky">{(PAGER.page -1) * PAGER.perPage + MASTER_ROWS.indexOf(row) + 1}</td>
              {#each FIELDS || [] as f, idx}
                {#if f.name === 'id'}
                  <td class="a-row">
                    {#if ACCESS.admin
                      || ACCESS.staff
                      || ACCESS.user
                    }
                      <div class="actions">
                        {#if CAN_SHOW_INFO}
                          <button class="btn info" title="Info" on:click={() => OnInfo(row)}>
                            <Icon size="15" color="var(--gray-007)" src={RiSystemInformationLine} />
                          </button>
                        {/if}
                        {#if !IS_CUSTOM_EDIT}
                          {#if CAN_EDIT_ROW}
                            <button
                              class="btn edit"
                              title="Edit"
                              on:click={() => toggleShowPopUp(Cell(row, idx, f), row)}
                            >
                              <Icon size="15" color="var(--gray-007)" src={RiDesignBallPenLine} />
                            </button>
                          {/if}
                        {/if}
                        {#if CAN_DELETE_ROW || CAN_RESTORE_ROW}
                          {#if row[deletedIndex] > 0 || row[deletedIndex] === 'terminated'}
                            <button class="btn info" title="Restore" on:click={() => restoreRow(row)}>
                              <Icon size="15" color="var(--gray-007)" src={RiArrowsArrowGoBackLine} />
                            </button>
                          {:else}
                            <button class="btn delete" title="Delete" on:click={() => deleteRow(row)}>
                              <Icon size="15" color="var(--gray-007)" src={RiSystemDeleteBin5Line} />
                            </button>
                          {/if}
                        {/if}
                        {#if CAN_OPEN_LINK}
                          {#each (LINKS || []) as l}
                            <a
                              title="{l.tooltip}"
                              class="btn link"
                              href={l.link(row)}
                              target={l.isTargetBlank ? '_blank' : ''}
                            >
                              <Icon
                                size="15"
                                color="var(--gray-007)"
                                src={l.icon}
                              />
                            </a>
                          {/each}
                        {/if}
                        {#if (EXTENDED_BUTTONS || []).length > 0}
                          {#each (EXTENDED_BUTTONS || []) as b}
                            <button
                              class="btn"
                              title="{b.tooltip}"
                              on:click={() => b.action(row)}
                            >
                              <Icon
                                size="15"
                                color="var(--gray-007)"
                                src={b.icon}
                              />
                            </button>
                          {/each}
                        {/if}
                      </div>
                    {:else}
                      <span>--</span>
                    {/if}
                  </td>
                {:else if f.name === 'facilities'}
                  <td class="textarea">{formatFacilities(row[idx], REFS['facilities'])}</td>
                {:else if f.type == 'intArr'}  
                  <td class="intArr">{row[idx] ? (
                    Object.entries(REFS[f.name]).map(([k, v]) => `${v}`).join(', ')
                  ) : '--'}</td>
                {:else if f.type == 'currency'}
                  <td class="currency">{row[idx] ? formatPrice(row[idx], (f.mapping || 'IDR')) : '--'}</td>
                {:else if f.inputType === 'datetime'}
                  <td class="datetime">{row[idx] ? datetime(row[idx]) : '--'}</td>
                {:else if f.inputType === 'combobox' && REFS[f.name]}
                  <td class="combobox">{REFS[f.name][row[idx]] || '--'}</td>
                {:else if f.inputType === 'percentage'}
                  <td class="percentage">{row[idx] || '0'}%</td>
                {:else}
                  <td class={f.type}>
                    {typeof row[idx] === 'boolean' ? (row[idx] ? 'Yes' : 'No') : (row[idx] || '--')}
                  </td>
                {/if}
              {/each}
            </tr>
          {/each}
        {:else}
          <tr>
            <td class="num-row">0</td>
            <td>no-data</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
  <div class="pagination_container">
    <div class="filter">
      <div class="showing">
        <p>Showing <span class="text-blue">{totalRowsCurrent}</span> /</p>
      </div>
      <div class="row_to_show">
        {#if showRowsNum}
          <div class="rows">
            {#each rowsToShow as r}
              <button on:click={() => toRow(r)}>{r}</button>
            {/each}
          </div>
        {/if}
        <button class="btn" on:click={toggleRowsNum}>
          <span>{currentRows}</span>
          <Icon className={showRowsNum ? 'rotate_right' : 'dropdown'} size="13" src={CgChevronRight} />
        </button>
      </div>
      <p>record(s)</p>
    </div>
    <div>
      <p>Total:<span class="text-blue">{totalRows}</span></p>
    </div>
    <div class="pagination">
      <button disabled={currentPage == 1} class="btn to" title="Go to first page" on:click={() => goToPage(1)}>
        <Icon size="16" src={CgChevronDoubleLeft} />
      </button>
      <button
        disabled={currentPage == 1}
        class="btn to"
        title="Go to previous page"
        on:click={() => goToPage(currentPage - 1)}
      >
        <Icon size="16" src={CgChevronLeft} />
      </button>
      {#each paginationShow as i}
        <button
          disabled={currentPage == i}
          class={currentPage === i ? 'btn active' : 'btn'}
          title={`Go to page ${i}`}
          on:click={() => goToPage(i)}>{i}</button
        >
      {/each}
      <button
        disabled={currentPage == paginationTotal}
        class="btn to"
        title="Go to next page"
        on:click={() => goToPage(currentPage + 1)}
      >
        <Icon size="16" src={CgChevronRight} />
      </button>
      <button
        disabled={currentPage == paginationTotal}
        class="btn to"
        title="Go to last page"
        on:click={() => goToPage(paginationTotal)}
      >
        <Icon size="16" src={CgChevronDoubleRight} />
      </button>
    </div>
  </div>
</div>

<style>
  .popup-container {
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    z-index: 2000;
    background-color: rgba(0 0 0 / 40%);
    backdrop-filter: blur(1px);
    display: flex;
    justify-content: center;
    padding: 50px;
    overflow: auto;
  }

  .popup-container .popup {
    border-radius: 8px;
    background-color: #fff;
    height: fit-content;
    width: 500px;
    display: flex;
    flex-direction: column;
  }

  .popup-container .popup header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 20px;
    border-bottom: 1px solid var(--gray-004);
  }

  .popup-container .popup header h2 {
    margin: 0;
  }

  .popup-container .popup header button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px;
    border-radius: 50%;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }

  .popup-container .popup header button:hover {
    background-color: #ef444420;
  }

  .popup-container .popup .forms {
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .popup-container .popup .foot {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
    align-items: center;
    padding: 10px 20px;
    border-top: 1px solid var(--gray-004);
  }

  .popup-container .popup .foot button {
    padding: 8px 13px;
    border-radius: 9999px;
    border: none;
    color: #fff;
    cursor: pointer;
    font-weight: 600;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .popup-container .popup .foot button.ok {
    background-color: var(--green-006);
    border: 1px solid var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .popup-container .popup .foot button.ok:hover {
    background-color: var(--green-005);
  }

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.action-btn:hover svg) {
    fill: var(--blue-005);
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  :global(.dropdown) {
    transition: all 0.2s ease-in-out;
  }

  :global(.rotate) {
    transition: all 0.2s ease-in-out;
    transform: rotate(180deg);
  }

  :global(.rotate_right) {
    transition: all 0.2s ease-in-out;
    transform: rotate(-90deg);
  }

  :global(.sort_icon) {
    margin-bottom: -5px;
  }

  .table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    border: 1px solid var(--gray-003);
    padding: 0 0 20px 0;
    font-size: var(--font-base);
  }

  .table-root .text-blue {
    color: var(--blue-005);
    font-weight: 600;
    padding: 5px;
  }

  .table-root p {
    margin: 0;
  }

  .table-root .actions-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 10px 15px;
    background-color: #fff;
    border-radius: 10px;
  }

  .table-root .actions-container .left,
  .table-root .actions-container .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

  .table-root .actions-container .left .debug .btn {
    border: none;
    background-color: var(--blue-006);
    color: #fff;
    width: fit-content;
    padding: 4px 10px;
    border-radius: 9999px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 3px;
    cursor: pointer;
  }

  .table-root .actions-container .left .debug .btn:hover {
    background-color: var(--blue-005);
  }

  .table-root .actions-container .right .search-handler {
    display: flex;
    flex-direction: row;
    width: fit-content;
    height: fit-content;
    position: relative;
  }

  .table-root .actions-container .right .search-handler input.search {
    padding: 12px 40px 12px 15px;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    background-color: transparent;
    width: 370px;
  }

  .table-root .actions-container .right .search-handler input.search:focus {
    border-color: var(--blue-005);
    outline: 1px solid var(--blue-005);
  }

  .table-root .actions-container .right .search-handler .search-btn {
    position: absolute;
    background-color: transparent;
    padding: 8px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    right: 5px;
    top: 4px;
  }

  .table-root .actions-container .right .search-handler .search-btn:hover {
    background-color: var(--blue-transparent);
  }

  :global(.table-root .actions-container .right .search-handler .search-btn:hover svg) {
    fill: var(--blue-005);
  }

  .table-root .actions-container .actions-button {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: calc(100% - 20px);
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    border-top: 1px solid var(--gray-003);
    border-bottom: 1px solid var(--gray-003);
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    font-size: var(--font-base);
    position: relative;
  }

  .table-root .table-container table .sticky {
    position: sticky;
    left: 0;
    z-index: 10;
    background-color: var(--gray-001);
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
    position: relative;
    z-index: 1;
  }

  .table-root .table-container table thead tr th .heading {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    background-color: transparent;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    text-transform: capitalize;
    font-weight: 600;
    color: var(--gray-008);
  }

  .table-root .table-container table thead tr th .heading:focus {
    outline: none;
  }

  .table-root .table-container table thead tr th .heading:hover {
    background-color: var(--gray-002);
  }

  .table-root .table-container table thead tr th .heading:disabled {
    color: var(--gray-008);
    background-color: transparent;
    cursor: default;
  }

  .table-root .table-container table thead tr th .heading:disabled:hover {
    background-color: transparent;
  }

  .table-root .table-container table thead tr th.textarea {
    min-width: 280px !important;
  }

  .table-root .table-container table thead tr th.datetime {
    min-width: 140px !important;
  }

  .table-root .table-container table tbody tr.deleted {
    color: var(--red-005);
  }

  .table-root .table-container table thead tr th.no {
    width: 30px;
    cursor: default;
  }

  .table-root .table-container table thead tr th.a-row {
    max-width: fit-content;
    min-width: fit-content;
    width: fit-content;
    cursor: default;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table-container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

	.table-root .table-container table tbody tr td.num-row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
	}

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .table-root .table-container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  .table-root .table-container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .table-root .table-container table tbody tr td .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

  .table-root .table-container table tbody tr td .actions .btn.delete:hover {
    background-color: var(--red-transparent);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--blue-005);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn.delete:hover svg) {
    fill: var(--red-005);
  }

  .table-root .pagination_container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 15px 0 15px;
  }

  .table-root .pagination_container .filter {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }

  .table-root .pagination_container .filter .row_to_show {
    position: relative;
    width: fit-content;
    height: fit-content;
  }

  .table-root .pagination_container .filter .row_to_show .btn {
    border: none;
    background-color: var(--blue-transparent);
    color: var(--blue-005);
    width: fit-content;
    padding: 3px 3px 3px 6px;
    font-weight: 600;
    border: 1px solid var(--blue-004);
    border-radius: 9999px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 1px;
    cursor: pointer;
  }

  .table-root .pagination_container .filter .row_to_show .btn:hover {
    background-color: var(--blue-002);
  }

  .table-root .pagination_container .filter .row_to_show .rows {
    display: flex;
    flex-direction: column-reverse;
    position: absolute;
    width: 100%;
    top: -200px;
    border-radius: 5px;
    border: 1px solid var(--gray-004);
    background-color: #fff;
  }

  .table-root .pagination_container .filter .row_to_show .rows button {
    border: none;
    background-color: transparent;
    padding: 5px;
    cursor: pointer;
    color: var(--gray-007);
  }

  .table-root .pagination_container .filter .row_to_show .rows button:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-007);
  }

  .table-root .pagination_container .pagination {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    overflow: hidden;
  }

  .table-root .pagination_container .pagination .btn {
    border: none;
    background-color: transparent;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 6px 10px;
    border-radius: 9999px;
    cursor: pointer;
    gap: 5px;
    color: var(--gray-007);
    border: 1px solid transparent;
  }

  .table-root .pagination_container .pagination .btn:hover {
    border: 1px solid var(--gray-004);
  }

  .table-root .pagination_container .pagination .btn.active {
    background-color: var(--blue-transparent);
    color: var(--blue-006);
    font-weight: 600;
    border: 1px solid var(--blue-004);
  }

  .table-root .pagination_container .pagination .btn.to {
    background-color: var(--blue-006);
    color: #fff;
    font-weight: 600;
    border: none;
  }

  .table-root .pagination_container .pagination .btn.to:hover {
    background-color: var(--blue-005);
  }

  .table-root .pagination_container .pagination .btn.to:disabled {
    background-color: var(--gray-002);
    color: var(--gray-006);
    font-weight: 600;
    border: 1px solid var(--gray-004);
    cursor: not-allowed;
  }

  @media only screen and (max-width: 768px) {
    .table-root .actions-container {
      flex-wrap: wrap;
      gap: 10px;
    }

    .table-root .actions-container .left {
      flex-wrap: wrap;
    }

    .table-root .table-container {
      overflow-x: scroll;
    }

    .table-root .pagination_container {
      flex-wrap: wrap;
      gap: 10px;
    }

    .table-root .pagination_container .pagination {
      gap: 2px;
    }
  }
</style>
