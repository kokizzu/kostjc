<script>
  /** @typedef {import('../_types/masters.js').Field} Field */
  /** @typedef {import('../_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/masters.js').ExtendedAction} ExtendedAction */
  /** @typedef {import('../_types/masters.js').ExtendedActionButton} ExtendedActionButton */

  import { onMount } from 'svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    CgChevronDoubleLeft, CgChevronLeft, CgChevronRight, CgChevronDoubleRight
  } from '../node_modules/svelte-icons-pack/dist/cg';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import Highlight from 'svelte-highlight/Highlight.svelte';
  import LineNumbers from  'svelte-highlight/LineNumbers.svelte'; 
  import json from 'svelte-highlight/languages/json';
  import atomOneDark from 'svelte-highlight/styles/atom-one-dark';
  
  export let renderFuncs  = /** @type {Record<string, Function>} */ ({});
  export let arrayOfArray = /** @type {boolean} */ (true);
  export let fields       = /** @type {Field[]} */ ([]);
  export let rows         = /** @type {any[] | Record<string, any>[]} */ ([]);
  export let pager        = /** @type {PagerOut} */ ({});
  export let REFS = {};
  export let EXTENDED_BUTTONS = /** @type {ExtendedActionButton[]} */ ([]);

  /**
   * @type {Record<string, number>}
   */
  export let COL_WIDTHS = {}

  function localeDatetime(dateStr) {
    if (!dateStr) return '';
    const dt = new Date(dateStr);
    const day = dt.toLocaleDateString('default', { weekday: 'long' });
    const date = dt.getDate();
    const month = dt.toLocaleDateString('default', { month: 'short' });
    const year = dt.getFullYear();
    let hh = /** @type {any} */ (dt.getHours());
    if (hh < 10) hh = '0' + hh;
    let mm = /** @type {any} */ (dt.getMinutes());
    if (mm < 10) mm = '0' + mm;
    const formattedDate = `${date} ${month} ${year} - ${hh}:${mm}`;
    return formattedDate;
  }

  let currentRows = pager.perPage;
  let totalRows = pager.countResult;
  // Rows per page options
  let rowsToShow = [5, 10, 20, 50, 70, 100, 200];
  // State for show rows options
  let showRowsNum = false;
  // Toggle show rows options
  function toggleRowsNum() {
    showRowsNum = !showRowsNum;
  }

  export let onRefreshTableView = function(/** @type {PagerIn} */ pager ) {
    console.log( 'TableView.onRefreshTableView', pager );
  };
  
  // Index of deletedAt field
  let deletedAtIdx = /** @type {number} */ (-1);

  let isPopupReady = false;

  onMount( () => {
    isPopupReady = true;
    for( let z = 0; z < fields.length; z++ ) {
      let field = fields[ z ];
      if( field.name==='deletedAt' ) {
        deletedAtIdx = z;
      }
    }
  } );
  
  
  function gotoPage(/** @type {number} */ page ) {
    onRefreshTableView({ ...pager, page });
  }
  
  function changePerPage(/** @type {number} */ perPage ) {
    showRowsNum = false;
    currentRows = perPage;
    onRefreshTableView({ ...pager, perPage});
  }
  
  /**
   * @param {number} row
   * @param {number} i
   * @param {Field} field
   */
  function cell( row, i, field ) {
    if( arrayOfArray ) return row[ i ] || '';
    return row[ field.name ] || '';
  }
  
  $: allowPrevPage = pager.page>1;
  $: allowNextPage = pager.page<pager.pages;
</script>

<svelte:head>
  {@html atomOneDark}
</svelte:head>

<section class="table-root">
  <div class="table-container">
    <table>
      <thead>
        <tr>
          <th class="no sticky">No</th>
          {#each (fields || []) as field}
            {#if field.name==='id'}
              <th class='a-row'>Action</th>
            {:else if !(field.name == 'beforeJson' || field.name == 'afterJson')}
              <th
                style="{COL_WIDTHS[field.name] ? `min-width: ${COL_WIDTHS[field.name]}px;` : ''}"
                class="
                {field.inputType === 'textarea' ? 'textarea' : ''}
                {field.inputType === 'datetime' ? 'datetime' : ''}
              ">
                {field.label}
              </th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#if rows && rows.length > 0}
          {#each (rows || []) as row, idx}
            <tr class:deleted={row[deletedAtIdx] > 0}>
              <td class="num-row sticky">{(pager.page -1) * pager.perPage + idx + 1}</td>
              {#each fields as field, i}
                {#if field.name === 'id'}
                  <td class="a-row">
                    <div class="actions">
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
                    </div>
                  </td>
                {:else if field.inputType === 'combobox' && REFS[field.name]}
                  <td>{REFS[field.name][row[field.name]]}</td>
                {:else if renderFuncs[ field.name ]}
                  <td>{renderFuncs[ field.name ]( cell( row, i, field ) ) }</td>
                {:else if field.inputType==='datetime' || field.name==='deletedAt' || field.name==='createdAt' || field.name==='updatedAt'}
                  <td>{localeDatetime( cell( row, i, field ) )}</td>
                {:else if !(field.name == 'beforeJson' || field.name == 'afterJson')}
                  <td>{cell( row, i, field )}</td>
                {/if}
              {/each}
            </tr>
          {/each}
        {:else}
          <tr>
            <td class="num-row">0</td>
            <td colspan={fields.length - 2}>No data</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
  <div class="pagination-container">
    <div class="filter">
      <div class="showing">
        <p>Showing <span class="text-blue">{(rows || []).length}</span> /</p>
      </div>
      <div class="row-to-show">
        {#if showRowsNum}
          <div class="rows">
            {#each rowsToShow as r}
              <button on:click={() => changePerPage(r)}>{r}</button>
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
      <button
        disabled={!allowPrevPage}
        class="btn to"
        title="Go to first page"
        on:click={() => gotoPage(1)}>
        <Icon size="16" src={CgChevronDoubleLeft} />
      </button>
      <button
        disabled={!allowPrevPage}
        class="btn to"
        title="Go to previous page"
        on:click={() => gotoPage(pager.page - 1)}
      >
        <Icon size="16" src={CgChevronLeft} />
      </button>
      <!-- {#each paginationShow as i}
        <button
          disabled={currentPage == i}
          class={currentPage === i ? 'btn active' : 'btn'}
          title={`Go to page ${i}`}
          on:click={() => goToPage(i)}>{i}</button
        >
      {/each} -->
      <button
      disabled={!allowNextPage}
        class="btn to"
        title="Go to next page"
        on:click={() => gotoPage(pager.page + 1)}
      >
        <Icon size="16" src={CgChevronRight} />
      </button>
      <button
        class="btn to"
        title="Go to last page"
        disabled={!allowNextPage}
        on:click={() => gotoPage(pager.pages)}
      >
        <Icon size="16" src={CgChevronDoubleRight} />
      </button>
    </div>
  </div>
</section>

<style>
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

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: calc(100% - 20px);
    border-radius: 10px;
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
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

  .table-root .table-container table thead tr th:first-child {
    border-top-left-radius: 10px;
  }

  .table-root .table-container table thead tr th:last-child {
    border-top-right-radius: 10px;
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

  :global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--blue-005);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn.delete:hover svg) {
    fill: var(--red-005);
  }

  .table-root .pagination-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 15px 0 15px;
  }

  .table-root .pagination-container .filter {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }

  .table-root .pagination-container .filter .row-to-show {
    position: relative;
    width: fit-content;
    height: fit-content;
  }

  .table-root .pagination-container .filter .row-to-show .btn {
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

  .table-root .pagination-container .filter .row-to-show .btn:hover {
    background-color: var(--blue-002);
  }

  .table-root .pagination-container .filter .row-to-show .rows {
    display: flex;
    flex-direction: column-reverse;
    position: absolute;
    width: 100%;
    top: -200px;
    border-radius: 5px;
    border: 1px solid var(--gray-004);
    background-color: #fff;
  }

  .table-root .pagination-container .filter .row-to-show .rows button {
    border: none;
    background-color: transparent;
    padding: 5px;
    cursor: pointer;
    color: var(--gray-007);
  }

  .table-root .pagination-container .filter .row-to-show .rows button:hover {
    background-color: var(--blue-transparent);
    color: var(--blue-007);
  }

  .table-root .pagination-container .pagination {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    overflow: hidden;
  }

  .table-root .pagination-container .pagination .btn {
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

  .table-root .pagination-container .pagination .btn:hover {
    border: 1px solid var(--gray-004);
  }

  .table-root .pagination-container .pagination .btn.to {
    background-color: var(--blue-006);
    color: #fff;
    font-weight: 600;
    border: none;
  }

  .table-root .pagination-container .pagination .btn.to:hover {
    background-color: var(--blue-005);
  }

  .table-root .pagination-container .pagination .btn.to:disabled {
    background-color: var(--gray-002);
    color: var(--gray-006);
    font-weight: 600;
    border: 1px solid var(--gray-004);
    cursor: not-allowed;
  }

  @media only screen and (max-width: 768px) {
    .table-root .table-container {
      overflow-x: scroll;
    }

    .table-root .pagination-container {
      flex-wrap: wrap;
      gap: 10px;
    }

    .table-root .pagination-container .pagination {
      gap: 2px;
      width: 100%;
      justify-content: center;
    }
  }
</style>
