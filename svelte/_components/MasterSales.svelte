<script>
  /** @typedef {import('../_types/cafe').Sale} Sale */
  /** @typedef {import('../_types/cafe').Payment} Payment */
  /** @typedef {import('../_types/cafe').UnpaidItem} UnpaidItem */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import SaleForm from './StaffSaleForm.svelte';
  import PaymentForm from './StaffPaymentForm.svelte';
  import { notifier } from '../_components/xNotifier';
  import {
    FiShoppingCart,
    FiCreditCard,
    FiAlertCircle,
    FiTrendingUp,
    FiPlus,
    FiClock,
    FiCheckCircle,
    FiX,
  } from '../node_modules/svelte-icons-pack/dist/fi';

  import { FaSolidUtensils, FaSolidShirt, FaSolidGamepad } from '../node_modules/svelte-icons-pack/dist/fa';
  import { convertMenuChoicesToMenuOptions } from '../_helper/sale';

  let menuChoices = /** @type {Record<Number, string>} */ ({/* menuChoices */});

  let menus = convertMenuChoicesToMenuOptions(menuChoices);

  function convertMenusToObject(menus) {
    const obj = {};
    for (const menu of menus) {
      obj[menu.id] = menu;
    }
    return obj;
  }
  let menusAsObject = convertMenusToObject(menus);


  export let sales;
  export let tenants;

  if (sales == undefined || sales == null) {
    sales = [];
  }

  function parseSalesPerItem(sales) {
    const result = {};

    for (const sale of sales) {
      const menuIds = sale[4];
      const totalPrice = sale[11];

      const perItemPrice = Math.floor(totalPrice / menuIds.length);

      for (const id of menuIds) {
        if (!result[id]) {
          result[id] = {
            name: menusAsObject[id]?.name,
            qty: 0,
            total: 0,
            times: '',
          };
        }

        result[id].qty += 1;
        result[id].total += menusAsObject[id]?.price || perItemPrice;
        result[id].times = new Date(sale[17] * 1000).toLocaleTimeString("id-ID", {
            hour: "2-digit",
            minute: "2-digit",
            hour12: false
          });
      }
    }
    return Object.values(result);
  }

  $: parsedItems = parseSalesPerItem(sales);

  let todayPayments = /** @type {Payment[]} */ ([]);

  let unpaidItems = /** @type {UnpaidItem[]} */ ([]);

  function addUnpaidItem(/** @type {Sale|any} */ sale) {
    unpaidItems = [
      ...unpaidItems,
      {
        id: unpaidItems.length + 1,
        customer: sale.buyerName,
        item: sale.menuIds,
        amount: sale.totalPriceIDR,
        date: sale.salesDate,
      },
    ];

    console.log('Unpaid items:', unpaidItems);
  }

  let overpaidItems = [
    { id: 1, customer: 'Rini', excess: 5000, date: '2024-01-10' },
    { id: 2, customer: 'Doni', excess: 10000, date: '2024-01-08' },
  ];

  let borrowedUtensils = [
    { id: 1, customer: 'Tono', item: 'Piring', qty: 2, date: '2024-01-10' },
    { id: 2, customer: 'Lisa', item: 'Gelas', qty: 1, date: '2024-01-09' },
  ];

  let laundryItems = [
    { id: 1, customer: 'Eko', status: 'Cuci', items: 'Kemeja, Celana', date: '2024-01-10' },
    { id: 2, customer: 'Nina', status: 'Jemur', items: 'Dress', date: '2024-01-09' },
  ];

  let gameItems = [
    { id: 1, customer: 'Rio', game: 'Billiard', duration: '2 jam', status: 'Playing' },
    { id: 2, customer: 'Lia', game: 'PS5', duration: '1 jam', status: 'Waiting' },
  ];

  // Computed values
  let totalSales = 0;
  let totalPayments = 0;

console.log("Today Sales:", sales);
console.log("Today Payments:", todayPayments);

if (sales && Array.isArray(sales) && sales.length > 0) {
  totalSales = sales.reduce((sum, sale) => {
    const saleAmount = sale[13] || 0;
    return sum + (typeof saleAmount === 'number' ? saleAmount : 0);
  }, 0);
  
  console.log("Calculated Total Sales:", totalSales);
}

if (todayPayments && Array.isArray(todayPayments) && todayPayments.length > 0) {
  totalPayments = todayPayments.reduce((sum, payment) => {
    const paymentAmount = payment.amount || 0;
    return sum + (typeof paymentAmount === 'number' ? paymentAmount : 0);
  }, 0);
  
  console.log("Calculated Total Payments:", totalPayments);
}

console.log("Final Total Sales:", totalSales);
console.log("Final Total Payments:", totalPayments);

  // State for modals
  let showSaleForm = false;
  let showPaymentForm = false;

  // Current sale data (shared between forms)
  let currentSale = /** @type {Sale|any} */ ({
    cashier: '',
    tenantId: '0',
    buyerName: '',
    menuIds: [],
    paymentMethod: '',
    paymentStatus: '',
    transferIDR: 0,
    qrisIDR: 0,
    cashIDR: 0,
    debtIDR: 0,
    topupIDR: 0,
    totalPriceIDR: 0,
    donation: 0,
    salesDate: '',
    paidAt: '',
    note: '',
  });

  // Event handlers
  export function addSale() {
    showSaleForm = true;
  }

  export function addPayment() {
    showPaymentForm = true;
  }

  function closeSaleForm() {
    showSaleForm = false;
  }

  function closePaymentForm() {
    showPaymentForm = false;
  }

  export let OnSubmit = async function (/** @type {Sale} */ sale) {
    console.log('OnSubmit :::', sale);
  };

  async function handleSaleSubmit(event) {
    event.preventDefault();
    const saleData = event.detail;
    // Merge with current sale data
    currentSale = { ...currentSale, ...saleData };
    console.log('Sale data submitted:', currentSale);

    await OnSubmit(currentSale);

    closeSaleForm();

    // showPaymentForm = true;
  }

  async function handlePaymentSubmit(event) {
    event.preventDefault();
    const paymentData = event.detail;
    // Merge with current sale data
    currentSale = { ...currentSale, ...paymentData, paidAt: new Date().toISOString() };
    console.log('Complete sale data to submit to API:', currentSale);

    // Here you would submit the complete data to your API

    // submitToAPI(currentSale);

    closePaymentForm();

    // Reset current sale after submission
    currentSale = /** @type {Sale|any} */ ({
      id: '',
      cashier: '',
      tenantId: '0',
      buyerName: '',
      menuIds: [],
      transferIDR: 0,
      qrisIDR: 0,
      cashIDR: 0,
      debtIDR: 0,
      topupIDR: 0,
      totalPriceIDR: 0,
      donation: 0,
      salesDate: '',
      paidAt: '',
      note: '',
    });
  }

  function formatCurrency(amount) {
    return new Intl.NumberFormat('id-ID').format(amount);
  }
</script>

<div class="dashboard">
  <div class="container">
    <h1 class="title">Dashboard Kasir</h1>

    <div class="main-grid">
      <!-- Kolom 1 - Penjualan & Pembayaran -->
      <div class="column-1">
        <!-- Penjualan Hari Ini -->
        <div class="card">
          <div class="card-header">
            <div class="header-title">
              <Icon src={FiShoppingCart} className="icon" color="color: #2563eb;" />
              <h2>Penjualan Hari Ini</h2>
            </div>
            <button class="btn btn-blue" on:click={addSale}>
              <Icon src={FiPlus} className="icon-small" />
              Add
            </button>
          </div>
          <div class="card-content">
            <div class="items-list">
              {#each parsedItems as sale (sale.name)}
                <div class="item-row item-blue">
                  <div class="item-info">
                    <p class="item-name">{sale.name}</p>
                    <p class="item-details">
                      Qty: {sale.qty} • {sale.times}
                    </p>
                  </div>
                  <p class="item-price price-blue">Rp {formatCurrency(sale.total)}</p>
                </div>
              {/each}
              <div class="total-section">
                <div class="total-row">
                  <span>Total Hari Ini:</span>
                  <span class="total-amount total-blue">Rp {formatCurrency(totalSales)}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Pembayaran Hari Ini -->
        <div class="card">
          <div class="card-header">
            <div class="header-title">
              <Icon src={FiCreditCard} className="icon" color="#059669" />
              <h2>Pembayaran Hari Ini</h2>
            </div>
            <button class="btn btn-green" on:click={addPayment}>
              <Icon src={FiPlus} className="icon-small" />
              Add
            </button>
          </div>
          <div class="card-content">
            <div class="items-list">
              {#each todayPayments as payment (payment.id)}
                <div class="item-row item-green">
                  <div class="item-info">
                    <p class="item-name">{payment.customer}</p>
                    <p class="item-details">
                      {payment.method} • {payment.time}
                    </p>
                  </div>
                  <p class="item-price price-green">Rp {formatCurrency(payment.amount)}</p>
                </div>
              {/each}
              <div class="total-section">
                <div class="total-row">
                  <span>Total Pembayaran:</span>
                  <span class="total-amount total-green">Rp {formatCurrency(totalPayments)}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Kolom 2 - 5 Box Tracking -->
      <div class="column-2">
        <div class="tracking-grid">
          <!-- Belum Dibayar/Hutang -->
          <div class="card">
            <div class="card-header">
              <div class="header-title">
                <Icon src={FiAlertCircle} className="icon" color="#dc2626" />
                <h2>Belum Dibayar</h2>
                <span class="badge badge-red">{unpaidItems.length}</span>
              </div>
            </div>
            <div class="card-content">
              <div class="items-list">
                {#each unpaidItems as item (item.id)}
                  <div class="item-row item-red">
                    <div class="item-info">
                      <p class="item-name">{item.customer}</p>
                      <p class="item-details">{item.item}</p>
                      <p class="item-date">{item.date}</p>
                    </div>
                    <p class="item-price price-red">Rp {formatCurrency(item.amount)}</p>
                  </div>
                {/each}
              </div>
            </div>
          </div>

          <!-- Lebih Bayar -->
          <div class="card">
            <div class="card-header">
              <div class="header-title">
                <Icon src={FiTrendingUp} className="icon" color="#ea580c" />
                <h2>Lebih Bayar</h2>
                <span class="badge badge-orange">{overpaidItems.length}</span>
              </div>
            </div>
            <div class="card-content">
              <div class="items-list">
                {#each overpaidItems as item (item.id)}
                  <div class="item-row item-orange">
                    <div class="item-info">
                      <p class="item-name">{item.customer}</p>
                      <p class="item-date">{item.date}</p>
                    </div>
                    <p class="item-price price-orange">+Rp {formatCurrency(item.excess)}</p>
                  </div>
                {/each}
              </div>
            </div>
          </div>

          <!-- Utensils Dipinjam -->
          <div class="card">
            <div class="card-header">
              <div class="header-title">
                <Icon src={FaSolidUtensils} className="icon" color="#9333ea" />
                <h2>Utensils Dipinjam</h2>
                <span class="badge badge-purple">{borrowedUtensils.length}</span>
              </div>
            </div>
            <div class="card-content">
              <div class="items-list">
                {#each borrowedUtensils as item (item.id)}
                  <div class="item-row item-purple">
                    <div class="item-info">
                      <p class="item-name">{item.customer}</p>
                      <p class="item-details">
                        {item.item} ({item.qty}x)
                      </p>
                      <p class="item-date">{item.date}</p>
                    </div>
                    <span class="status-badge status-purple">
                      <Icon src={FiClock} className="icon-tiny" />
                      Dipinjam
                    </span>
                  </div>
                {/each}
              </div>
            </div>
          </div>

          <!-- Laundry -->
          <div class="card">
            <div class="card-header">
              <div class="header-title">
                <Icon src={FaSolidShirt} className="icon" color="#0891b2;" />
                <h2>Laundry</h2>
                <span class="badge badge-cyan">{laundryItems.length}</span>
              </div>
            </div>
            <div class="card-content">
              <div class="items-list">
                {#each laundryItems as item (item.id)}
                  <div class="item-row item-cyan">
                    <div class="item-info">
                      <p class="item-name">{item.customer}</p>
                      <p class="item-details">{item.items}</p>
                      <p class="item-date">{item.date}</p>
                    </div>
                    <span class="status-badge {item.status === 'Cuci' ? 'status-blue' : 'status-yellow'}">
                      {item.status}
                    </span>
                  </div>
                {/each}
              </div>
            </div>
          </div>
        </div>

        <!-- Main Dingdong - Full Width -->
        <div class="card">
          <div class="card-header">
            <div class="header-title">
              <Icon src={FaSolidGamepad} className="icon" color="#4f46e5;" />
              <h2>Main Dingdong</h2>
              <span class="badge badge-indigo">{gameItems.length}</span>
            </div>
          </div>
          <div class="card-content">
            <div class="game-grid">
              {#each gameItems as item (item.id)}
                <div class="game-item">
                  <div class="game-info">
                    <p class="item-name">{item.customer}</p>
                    <p class="item-details">
                      {item.game} • {item.duration}
                    </p>
                  </div>
                  <span class="status-badge {item.status === 'Playing' ? 'status-green' : 'status-yellow'}">
                    {#if item.status === 'Playing'}
                      <Icon src={FiCheckCircle} className="icon-tiny" />
                    {:else}
                      <Icon src={FiClock} className="icon-tiny" />
                    {/if}
                    {item.status}
                  </span>
                </div>
              {/each}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Sale Form Modal -->
  {#if showSaleForm}
    <div class="modal-overlay">
      <div class="modal">
        <div class="modal-header">
          <h2>Tambah Penjualan</h2>
          <button class="close-btn" on:click={closeSaleForm}>
            <Icon src={FiX} size={18} />
          </button>
        </div>
        <SaleForm on:saleSubmit={handleSaleSubmit} tenants={tenants} />
      </div>
    </div>
  {/if}

  <!-- Payment Form Modal -->
  {#if showPaymentForm}
    <div class="modal-overlay">
      <div class="modal">
        <div class="modal-header">
          <h2>Tambah Pembayaran</h2>
          <button class="close-btn" on:click={closePaymentForm}>
            <Icon src={FiX} size={18} />
          </button>
        </div>
        <PaymentForm
          on:submit={handlePaymentSubmit}
          buyerName={currentSale.buyerName}
          totalPriceIDR={currentSale.totalPriceIDR}
        />
      </div>
    </div>
  {/if}
</div>

<style>
  /* Reset and base styles */
  * {
    box-sizing: border-box;
  }

  .dashboard {
    min-height: 100vh;
    background-color: #f9fafb;
    padding: 1rem;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  }

  .container {
    max-width: 1280px;
    margin: 0 auto;
  }

  .title {
    font-size: 1.875rem;
    font-weight: 700;
    color: #111827;
    margin-bottom: 2rem;
    margin-top: 0;
  }

  /* Layout Grid */
  .main-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  @media (min-width: 1024px) {
    .main-grid {
      grid-template-columns: 1fr 2fr;
    }
  }

  .column-1 {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .column-2 {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .tracking-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  @media (min-width: 768px) {
    .tracking-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  /* Card Components */
  .card {
    background: white;
    border-radius: 0.5rem;
    box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
    border: 1px solid #e5e7eb;
    overflow: hidden;
  }

  .card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1rem;
    border-bottom: 1px solid #e5e7eb;
  }

  .header-title {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    flex: 1;
  }

  .header-title h2 {
    font-size: 1.125rem;
    font-weight: 600;
    margin: 0;
    color: #111827;
  }

  .card-content {
    padding: 1rem;
  }

  /* Buttons */
  .btn {
    display: inline-flex;
    align-items: center;
    padding: 0.375rem 0.75rem;
    font-size: 0.875rem;
    font-weight: 500;
    color: white;
    border: none;
    border-radius: 0.375rem;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .btn-blue {
    background-color: #2563eb;
  }

  .btn-blue:hover {
    background-color: #1d4ed8;
  }

  .btn-green {
    background-color: #059669;
  }

  .btn-green:hover {
    background-color: #047857;
  }

  /* Badges */
  .badge {
    display: inline-flex;
    align-items: center;
    padding: 0.25rem 0.5rem;
    font-size: 0.75rem;
    font-weight: 500;
    border-radius: 9999px;
    margin-left: auto;
  }

  .badge-red {
    background-color: #fef2f2;
    color: #991b1b;
  }

  .badge-orange {
    background-color: #fff7ed;
    color: #9a3412;
  }

  .badge-purple {
    background-color: #faf5ff;
    color: #7c2d12;
  }

  .badge-cyan {
    background-color: #ecfeff;
    color: #155e75;
  }

  .badge-indigo {
    background-color: #eef2ff;
    color: #3730a3;
  }

  /* Status Badges */
  .status-badge {
    display: inline-flex;
    align-items: center;
    padding: 0.25rem 0.5rem;
    font-size: 0.75rem;
    font-weight: 500;
    border: 1px solid;
    border-radius: 0.375rem;
  }

  .status-purple {
    color: #9333ea;
    border-color: #d8b4fe;
  }

  .status-blue {
    color: #2563eb;
    border-color: #93c5fd;
  }

  .status-yellow {
    color: #d97706;
    border-color: #fcd34d;
  }

  .status-green {
    color: #059669;
    border-color: #86efac;
  }

  /* Item Lists */
  .items-list {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .item-row {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    padding: 0.5rem;
    border-radius: 0.5rem;
  }

  .item-blue {
    background-color: #eff6ff;
  }
  .item-green {
    background-color: #f0fdf4;
  }
  .item-red {
    background-color: #fef2f2;
  }
  .item-orange {
    background-color: #fff7ed;
  }
  .item-purple {
    background-color: #faf5ff;
  }
  .item-cyan {
    background-color: #ecfeff;
  }

  .item-info {
    flex: 1;
  }

  .item-name {
    font-weight: 500;
    font-size: 0.875rem;
    margin: 0 0 0.25rem 0;
    color: #111827;
  }

  .item-details {
    font-size: 0.75rem;
    color: #6b7280;
    margin: 0 0 0.25rem 0;
  }

  .item-date {
    font-size: 0.75rem;
    color: #9ca3af;
    margin: 0;
  }

  .item-price {
    font-weight: 600;
    font-size: 0.875rem;
    margin: 0;
  }

  .price-blue {
    color: #2563eb;
  }
  .price-green {
    color: #059669;
  }
  .price-red {
    color: #dc2626;
  }
  .price-orange {
    color: #ea580c;
  }

  /* Total Section */
  .total-section {
    border-top: 1px solid #e5e7eb;
    padding-top: 0.5rem;
    margin-top: 0.75rem;
  }

  .total-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: 700;
  }

  .total-amount {
    font-weight: 600;
  }

  .total-blue {
    color: #2563eb;
  }
  .total-green {
    color: #059669;
  }

  /* Game Grid */
  .game-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 0.75rem;
  }

  @media (min-width: 768px) {
    .game-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .game-item {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    padding: 0.75rem;
    background-color: #eef2ff;
    border-radius: 0.5rem;
  }

  .game-info {
    flex: 1;
  }

  /* Modal Styles */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  .modal {
    background-color: white;
    border-radius: 0.5rem;
    width: 90%;
    max-width: 500px;
    max-height: 90vh;
    overflow-y: auto;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1);
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    border-bottom: 1px solid #e5e7eb;
  }

  .modal-header h2 {
    font-size: 1.25rem;
    font-weight: 600;
    margin: 0;
  }

  .close-btn {
    background: none;
    border: none;
    cursor: pointer;
    color: #6b7280;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.25rem;
    border-radius: 0.25rem;
  }

  .close-btn:hover {
    background-color: #f3f4f6;
    color: #111827;
  }

  /* Responsive adjustments */
  @media (max-width: 768px) {
    .dashboard {
      padding: 0.5rem;
    }

    .title {
      font-size: 1.5rem;
      margin-bottom: 1.5rem;
    }

    .card-header {
      flex-direction: column;
      align-items: flex-start;
      gap: 0.5rem;
    }

    .header-title {
      width: 100%;
    }

    .btn {
      align-self: flex-end;
    }

    .modal {
      width: 95%;
      max-height: 80vh;
    }
  }
</style>
