<script>
  /** @typedef {import('../_types/masters.js').Field} Field */
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/masters.js').ExtendedAction} ExtendedAction */
  /** @typedef {import('../_types/property.js').Booking} Booking */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiDesignBallPenLine,
    RiSystemDeleteBin5Line,
    RiSystemFilterLine,
    RiArrowsArrowGoBackLine,
    RiSystemInformationLine,
    RiArrowsArrowDropRightLine,
    RiFinanceBankCardLine,
    RiArrowsExpandUpDownFill,
    RiArrowsArrowDownSFill,
    RiArrowsArrowUpSFill
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { IoSearch, IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import {
    CgChevronLeft,
    CgChevronRight,
    CgChevronDoubleRight,
    CgChevronDoubleLeft,
  } from '../node_modules/svelte-icons-pack/dist/cg';
  import InputBox from './InputBox.svelte';
  import { onMount } from 'svelte';
  import { dateISOFormat, datetime, formatPrice } from './xFormatter.js';
  import FilterTable from './FilterTable.svelte';
  import {AdminBooking, AdminPayment} from '../jsApi.GEN'
  import { notifier } from './xNotifier';
  import PopUpAddPayment from './PopUpAddPayment.svelte';

  export let tenants = /** @type {Record<number, string>} */ ({});

  /**
   * @description Get extra tenants
   * @param {number[]} tenantIds
   * @returns {string}
   * */
  function getExtraTenants(tenantIds) {
    if (!tenantIds.length) return '--';
    return tenantIds.map(tenantId => tenants[tenantId]).join(', ');
  }

  function formatFacilities(facilities) {
    let formattedFacilities = '';
    let arrFacilities = [];

    try {
      arrFacilities = JSON.parse(facilities);
    } catch (e) {
      return facilities;
    }
    
    try {
      formattedFacilities = arrFacilities.map(facility => 
          `${facility.facilityName} (IDR ${facility.extraChargeIDR.toLocaleString('en-ID')})`
        ).join(', ');
    } catch (e) {
      return formattedFacilities;
    }
  }


  export let FIELDS = /** @type Field[] */ ([]); // bind
  export let PAGER = /** @type PagerOut */ ({}); // bind
  export let MASTER_ROWS = /** @type any[][] */ ([]); // bind
  export let REFS = {};

  export let CAN_SHOW_INFO = false;

  export let ACCESS = /** @type Access */ ({});
  export let ARRAY_OF_ARRAY = true;
  export let CAN_SEARCH_ROW = true;

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

  // Pagination total, based on total pages
  let paginationTotal = 1;
  // Pagination all, based on total pages
  let paginationsAll = /** @type number[] */ ([]);
  // Pagination to show, based on total pagination
  let paginationShow = /** @type number[] */ ([]);
  // Current page describe which page is currently rendered
  let currentPage = 1;
  // State for sort, wheter is ascending or descending
  let isSortTableAsc = false;
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

  // PopUp for modify rows
  let showPopUpEdit = false;

  // Row ID to modify
  let idToMod = 0;

  async function toggleShowPopUpEdit(/** @type any */ id, /** @type any[]*/ row) {
    let booking = /** @type {Booking} */ ({});
    idToMod = id;
    let isErrAjax = false;
    await AdminBooking({
      cmd: 'form', // @ts-ignore
      booking: {
        id: id
      }
    }, function (/** @type {any} */ o) {
      if (o.error) {
        isErrAjax = true;
        return;
      }
      booking = o.booking;
    })
    fillExtraTenants();
    payloads = [];
    if (FIELDS && FIELDS.length > 0) {
      FIELDS.forEach((f, i) => {
        if (f.name === 'extraTenants') {
          extraTenantsIds = row[i];
          (extraTenantsIds || []).forEach((id) => {
            (extraTenants || []).forEach((extraTenant) => {
              if (extraTenant.id === id) {
                selectedExtraTenants = extraTenant.name;
                extraTenantsToShow = [...extraTenantsToShow, extraTenant];
                extraTenants = extraTenants.filter(f => f.id !== extraTenant.id);
              }
            })
          })
        }
        payloads = [...payloads, row[i]];
      });
    }

    if (!isErrAjax) {
      payloads[1] = booking.roomId;
      payloads[3] = booking.dateStart;
      payloads[4] = booking.dateEnd;
      payloads[5] = booking.tenantId;
      payloads[6] = booking.basePriceIDR;
      payloads[7] = booking.facilitiesObj;
      payloads[8] = booking.totalPriceIDR;
      payloads[9] = booking.paidAt;
      payloads[10] = booking.extraTenants;
    }

    showPopUpEdit = true;
    idToMod = id;
  }

  function handleSubmitEdit() {
    (FIELDS || []).forEach((f, i) => {
      if (f.name === 'extraTenants') {
        payloads[i] = extraTenantsIds;
      }
    })
    OnEdit(idToMod, payloads);
    closePopUpEdit();
  }

  let extraTenantsIds = [];

  /**
   * @typedef {Object} ExtraTenant
   * @property {number} id
   * @property {string} name
   */

  let extraTenants = /** @type {ExtraTenant[]} */ ([]);

  function fillExtraTenants() {
    for (const [k, v] of Object.entries(tenants)) extraTenants.push({id: Number(k), name: v});
  }

  let extraTenantsToShow = /** @type {ExtraTenant[]} */ ([]);
  let showExtraTenants = false;
  let selectedExtraTenants = 'Tenant....';

  function handleExtraTenantsClose() {
		showExtraTenants = false;
		document.body.removeEventListener('click', handleExtraTenantsClose)
	}
	function toggleExtraTenants() {
		showExtraTenants = !showExtraTenants;
		if (showExtraTenants) document.body.addEventListener('click', handleExtraTenantsClose);
		else document.body.removeEventListener('click', handleExtraTenantsClose);
	}

	function choseExtraTenant(/** @type {ExtraTenant} */ extraTenant) {
    selectedExtraTenants = extraTenant.name;
    showExtraTenants = false;
    extraTenantsIds = [...extraTenantsIds, Number(extraTenant.id)];
		extraTenantsToShow = [...extraTenantsToShow, extraTenant];
    extraTenants = extraTenants.filter(f => f.id !== extraTenant.id);
	}

	const removeExtraTenant = (/** @type {number} */ idx) => {
    extraTenantsIds = extraTenantsIds.filter(( _, i ) => i !== idx);
    extraTenants = [...extraTenants, extraTenantsToShow[idx]];
		extraTenantsToShow = extraTenantsToShow.filter(( _, i ) => i !== idx);
	}

  function closePopUpEdit() {
    extraTenantsIds = [];
    extraTenants = []
    selectedExtraTenants = 'Tenant....';
    showExtraTenants = false;
    extraTenantsToShow = [];
    showPopUpEdit = false;
  }

  let popUpAddPayment = null;
  let isSubmitAddPayment = false;
  let isPopUpAddPaymentReady = false;
  let paymentAt = dateISOFormat(0);

  onMount(() => {
    isPopUpAddPaymentReady = true;
  })

  function showPopUpAddPayment(row) {
    idToMod = row[0];
    paymentAt = row[3];
    popUpAddPayment.Show();
  }

  async function submitAddPayment(payment) {
    isSubmitAddPayment = true;
    const i = /** @type {any} */ ({
      payment,
      cmd: 'upsert'
    });

    await AdminPayment(i,
      /** @type {import('../jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddPayment = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        notifier.showSuccess(`Payment created !!`);

        popUpAddPayment.Reset();
        popUpAddPayment.Hide();
        OnRefresh(PAGER);
      }
    );
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
</script>

{#if isPopUpAddPaymentReady}
  <PopUpAddPayment
    isBookingReadOnly
    bind:bookingId={idToMod}
    bind:this={popUpAddPayment}
    bind:isSubmitted={isSubmitAddPayment}
    OnSubmit={submitAddPayment}
  />
{/if}

{#if filterTableReady}
  <FilterTable bind:this={filterTable} bind:filterColumns bind:filtersMap on:click={ApplyFilter} />
{/if}

{#if showPopUpEdit}
  <div class="popup_container">
    <div class="popup">
      <header>
        <h2>Edit Booking {`#${idToMod}`}</h2>
        <button on:click={closePopUpEdit}>
          <Icon size="22" color="var(--red-005)" src={IoClose} />
        </button>
      </header>
      <div class="forms">
        {#each (FIELDS || []) as field, idx}
          {#if field.name !== 'id' && !field.readOnly}
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
            {:else if field.name === 'extraTenants'}
              <div class="multiselect-form">
                <label class="label" for="extraTenants">Extra Tenants</label>
                <div class="multiselect-selected">
                  {#each extraTenantsToShow as ext, idx}
                    <div class="facility-show">
                      <span>{ext.name}</span>
                      <button on:click={() => removeExtraTenant(idx)}>
                        <Icon size="18" color="var(--sky-001)" src={IoClose}/>
                      </button>
                    </div>
                  {/each}
                </div>
                <div class="dropdown-multiselect">
                  <div class="dropdown-item">
                    <button id="extraTenants" class="dropdown-btn" on:click|stopPropagation={toggleExtraTenants}>
                      <span>{selectedExtraTenants}</span>
                      <Icon
                        className={showExtraTenants ? 'rotate-b' : 'dropdown'}
                        size="25"
                        src={RiArrowsArrowDropRightLine}
                      />
                    </button>
                    {#if showExtraTenants}
                      <div class="dropdown-list">
                        {#if extraTenants && extraTenants.length}
                          {#each extraTenants as ext}
                            <button
                              class="facility"
                              on:click|stopPropagation={() => choseExtraTenant(ext)}>
                              <span>{ext.name}</span>
                            </button>
                          {/each}
                        {:else}
                          <p>No Tenants yet, please add it in master tenants</p>
                        {/if}
                      </div>
                    {/if}
                  </div>
                </div>
              </div>
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
          <button class="search-btn" title="Search">
            <Icon color="var(--gray-007)" size="16" src={IoSearch} />
          </button>
          <input placeholder="Search..." type="text" name="searchRow" id="searchRow" class="search" />
        </div>
      {/if}
    </div>
  </div>
  <div class="table_container">
    <table>
      <thead>
        <tr>
          <th class="no">No</th>
          {#each FIELDS || [] as f, _ (f.name)}
            {#if f.name === 'id'}
              <th class="a_row">Actions</th>
            {:else}
              <th
                style="
                  {f.name === 'facilitiesObj' ? 'min-width: 300px;' : ''}
                  {f.name === 'extraTenants' ? 'min-width: 250px;' : ''}
                  {f.name === 'tenantId' ? 'min-width: 200px;' : ''}
                "
                class="
								{f.inputType === 'textarea' ? 'textarea' : ''}
								{f.inputType === 'datetime' ? 'datetime' : ''}
							">
                <button
                  class="heading"
                  on:click={() => OnSort(f)}
                  disabled={f.name === 'totalPaidIDR'}
                >
                  <span>{f.label}</span>
                  {#if f.name !== 'totalPaidIDR'}
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
            <tr
              class={(row[deletedIndex] > 0 || row[deletedIndex] === 'terminated') ? 'deleted' : ''}
            >
              <td class="num_row">{(PAGER.page -1) * PAGER.perPage + MASTER_ROWS.indexOf(row) + 1}</td>
              {#each FIELDS || [] as f, idx}
                {#if f.name === 'id'}
                  <td class="a_row">
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
                        <button
                          class="btn edit"
                          title="Edit"
                          on:click={() => toggleShowPopUpEdit(Cell(row, idx, f), row)}
                        >
                          <Icon size="15" color="var(--gray-007)" src={RiDesignBallPenLine} />
                        </button>
                        <button class="btn payment" title="Input Payment" on:click={() => showPopUpAddPayment(row)}>
                          <Icon size="15" color="var(--gray-007)" src={RiFinanceBankCardLine} />
                        </button>
                        {#if row[deletedIndex] > 0 || row[deletedIndex] === 'terminated'}
                          <button class="btn info" title="Restore" on:click={() => restoreRow(row)}>
                            <Icon size="15" color="var(--gray-007)" src={RiArrowsArrowGoBackLine} />
                          </button>
                        {:else}
                          <button class="btn delete" title="Delete" on:click={() => deleteRow(row)}>
                            <Icon size="15" color="var(--gray-007)" src={RiSystemDeleteBin5Line} />
                          </button>
                        {/if}
                      </div>
                    {:else}
                      <span>--</span>
                    {/if}
                  </td>
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
                {:else if f.name === 'facilitiesObj'}
                    <td class="textarea">{formatFacilities(row[idx])}</td>
                {:else if f.name === 'extraTenants'}
                  <td class="textarea">{getExtraTenants(row[idx])}</td>
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
            <td class="num_row">0</td>
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
  .popup_container {
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

  .popup_container .popup {
    border-radius: 8px;
    background-color: #fff;
    height: fit-content;
    width: 500px;
    display: flex;
    flex-direction: column;
  }

  .popup_container .popup header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 20px;
    border-bottom: 1px solid var(--gray-004);
  }

  .popup_container .popup header h2 {
    margin: 0;
  }

  .popup_container .popup header button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px;
    border-radius: 50%;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }

  .popup_container .popup header button:hover {
    background-color: #ef444420;
  }

  .popup_container .popup .forms {
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .popup_container .popup .foot {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
    align-items: center;
    padding: 10px 20px;
    border-top: 1px solid var(--gray-004);
  }

  .popup_container .popup .foot button {
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

  .popup_container .popup .foot button.ok {
    background-color: var(--green-006);
    border: 1px solid var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .popup_container .popup .foot button.ok:hover {
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

  .table-root .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table-root .table_container table {
    width: 100%;
    background: #fff;
    border-top: 1px solid var(--gray-003);
    border-bottom: 1px solid var(--gray-003);
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
    font-size: var(--font-base);
  }

  .table-root .table_container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table_container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table_container table thead tr th .heading {
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
  }

  .table-root .table_container table thead tr th .heading:focus {
    outline: none;
    background-color: var(--gray-002);
  }

  .table-root .table_container table thead tr th .heading:hover {
    background-color: var(--gray-002);
  }

  .table-root .table_container table thead tr th .heading:disabled {
    color: var(--gray-008);
    background-color: transparent;
  }

  .table-root .table_container table thead tr th .heading:disabled:hover {
    background-color: transparent;
  }

  .table-root .table_container table thead tr th.textarea {
    min-width: 280px !important;
  }

  .table-root .table_container table thead tr th.datetime {
    min-width: 140px !important;
  }

  .table-root .table_container table tbody tr.deleted {
    color: var(--red-005);
  }

  .table-root .table_container table thead tr th.no {
    width: 30px;
  }

  .table-root .table_container table thead tr th.a_row {
    max-width: fit-content;
    min-width: fit-content;
    width: fit-content;
  }

  .table-root .table_container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table_container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table_container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table_container table tbody tr:last-child td,
	.table-root .table_container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

	.table-root .table_container table tbody tr td.num_row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
	}

  .table-root .table_container table tbody tr:last-child td,
  .table-root .table_container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table_container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table_container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .table-root .table_container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  .table-root .table_container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .table-root .table_container table tbody tr td .actions .btn:hover {
    background-color: var(--blue-transparent);
  }

  .table-root .table_container table tbody tr td .actions .btn.delete:hover {
    background-color: var(--red-transparent);
  }

  :global(.table-root .table_container table tbody tr td .actions .btn:hover svg) {
    fill: var(--blue-005);
  }

  :global(.table-root .table_container table tbody tr td .actions .btn.delete:hover svg) {
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

  .multiselect-form {
    display: flex;
    flex-direction: column;
    gap: 3px;
  }

  .multiselect-form label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .multiselect-form .multiselect-selected {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 5px;
  }

  .multiselect-form .multiselect-selected .facility-show {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    background-color: var(--sky-006);
    color: #FFF;
    padding: 5px 5px 5px 15px;
    border-radius: 5px;
  }

  .multiselect-form .multiselect-selected .facility-show button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px;
    border-radius: 9999px;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }

  .multiselect-form .multiselect-selected .facility-show button:hover {
    background-color: var(--sky-005);
  }

  .dropdown-multiselect {
		position: relative;
		flex-grow: 1;
		display: flex;
  }

  .dropdown-multiselect {
		position: relative;
		flex-grow: 1;
		display: flex;
	}

	.dropdown-multiselect .dropdown-item {
		width: 100%;
	}

	.dropdown-multiselect .dropdown-item .dropdown-btn {
		width: 100% !important;
	}

	.dropdown-item {
		flex-grow: 1;
		position: relative;
	}

	.dropdown-item .dropdown-btn {
		width: 100%;
		height: fit-content;
		background-color: #FFF;
		padding: 10px 15px;
		border-radius: 5px;
		display: flex;
		flex-direction: row;
		border: 1px solid var(--gray-003);
		align-items: center;
		cursor: pointer;
		color: var(--gray-007);
		font-weight: 600;
	}

	.dropdown-item .dropdown-btn:hover {
		background-color: var(--gray-001);
	}

	.dropdown-item .dropdown-btn span {
		flex-grow: 1;
		text-align: left;
	}

	.dropdown-item .dropdown-list {
		background-color: #FFF;
		max-height: 300px;
		overflow-y: auto;
		width: 100%;
    margin-top: 10px;
		border: 1px solid var(--sky-004);
		border-radius: 5px;
		display: flex;
		flex-direction: column;
		position: absolute;
		top: 40px;
		z-index: 99999;
	}

	.dropdown-item .dropdown-list::-webkit-scrollbar-thumb {
    background-color : var(--gray-003);
    border-radius    : 8px;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-thumb:hover {
    background-color : var(--gray-004);
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar {
    width : 8px;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-track {
    background-color : transparent;
  }

  .dropdown-item .dropdown-list::-webkit-scrollbar-button {
    display: none;
  }

	.dropdown-item .dropdown-list .facility {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 10px;
		padding: 10px 15px;
		background-color: #FFF;
		border: none;
		color: var(--gray-007);
		width: 100%;
		cursor: pointer;
	}

	.dropdown-item .dropdown-list .facility:hover {
		color: var(--sky-009);
    background-color: #0ea5e920;
	}

  @media only screen and (max-width: 768px) {
    .table-root .actions-container {
      flex-wrap: wrap;
      gap: 10px;
    }

    .table-root .actions-container .left {
      flex-wrap: wrap;
    }

    .table-root .table_container {
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
