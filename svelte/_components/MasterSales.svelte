<script>
  /** @typedef {import('../_types/cafe').Sale} Sale */
  /** @typedef {import('../_types/cafe').Payment} Payment */
  /** @typedef {import('../_types/cafe').Unpaid} Unpaid */
  /** @typedef {import('../_types/cafe').Overpaid} Overpaid */
    /** @typedef {import('../_types/cafe').TodaySales} TodaySales */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import SaleForm from './StaffSaleForm.svelte';
  import PaymentForm from './StaffPaymentForm.svelte';
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
  import { parseMenuMapToOptions, convertMenusToObject, parseSalesToTodaySales, parseSalesToTodayPayment, parseSalesToOverpaid, parseSalesToUnpaid} from '../_helper/sale.js';

  let menuChoices = /** @type {Record<Number, string>} */ ({/* menuChoices */});

  export let sales;
  console.log("Sales", sales);
  export let tenants;
  export let user; // logged-in user passed from page

  if (sales == undefined || sales == null) {
    sales = [];
  }

  let menus = parseMenuMapToOptions(menuChoices);
  let menusAsObject = convertMenusToObject(menus);

  console.log("Menus", menus);
  console.log("Menus As Object", menusAsObject);

  $: todaySales = parseSalesToTodaySales(sales, menusAsObject);
  $: console.log("Today Sales", todaySales);
 
  let todayPayments = /** @type {Payment[]} */ ([]);
  $: todayPayments = parseSalesToTodayPayment(sales, tenants)
  $: console.log("Today Payments", todayPayments);

  let unpaidItems = /** @type {Unpaid[]} */ ([]);
  $: unpaidItems = parseSalesToUnpaid(sales, menusAsObject, tenants)
  $: console.log("Unpaid Items", unpaidItems);

  let overpaidItems = /** @type {Overpaid[]} */ ([]);
  $: overpaidItems = parseSalesToOverpaid(sales, tenants)
  $: console.log("Overpaid Items", overpaidItems);

  let totalSales = 0;
  let totalPayments = 0;

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

//ini reaktif
// $: totalSales =
//   Array.isArray(sales)
//     ? sales.reduce((sum, s) => {
//         const v = s?.[14] ?? 0;
//         return sum + (typeof v === 'number' ? v : 0);
//       }, 0)
//     : 0;

$: totalSales =
  Array.isArray(sales)
    ? sales.reduce((sum, s) => {
        const salesDate = s?.[15]; 
        const today = new Date().toISOString().split('T')[0];

        if (salesDate !== today) return sum;

        const v = s?.[14] ?? 0;
        return sum + (typeof v === 'number' ? v : 0);
      }, 0)
    : 0;

$: console.log("Final Total Sales:", totalSales);

// reaktif how to sum totalPayments
$: totalPayments =
  Array.isArray(todayPayments)
    ? todayPayments.reduce((sum, p) => {
        const v = p?.amount ?? 0;
        return sum + (typeof v === 'number' ? v : 0);
      }, 0)
    : 0;

$: console.log("Final Total Payments:", totalPayments);

  // State for modals
  let showSaleForm = false;
  let showPaymentForm = false;

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

  export let OnEdit = async function (/** @type {Sale} */ saleData) {
    console.log('OnEdit :::', saleData);
  };

  async function handleSaleSubmit(event) {
    event.preventDefault();
    const saleData = event.detail;
    await OnSubmit(saleData);
    closeSaleForm();
  }

  async function handlePaymentSubmit(event) {
    event.preventDefault();
    const paymentData = event.detail;
    await OnEdit(paymentData);
    closePaymentForm();
  }

  function formatCurrency(amount) {
    return new Intl.NumberFormat('id-ID').format(amount);
  }

  // Helpers: find sale by payment id and convert to typed Sale object
  function findSaleRowByPaymentId(paymentId) {
    return Array.isArray(sales)
      ? sales.find(s => String(s?.[0]) === String(paymentId))
      : null;
  }

  /** @returns {import('../_types/cafe').Sale} */
  function mapSaleRowToSale(row) {
    if (!Array.isArray(row)) return /** @type {any} */ ({});
    return /** @type {any} */ ({
    id: row[0],                              // Direct assignment (likely string/number)
    cashier: row[1] ?? '',                   // String with fallback
    tenantId: row[2] ?? 0,                   // Number with fallback
    buyerName: row[3] ?? '',                 // String with fallback
    menuIds: Array.isArray(row[4]) ? row[4] : [], // Safe array handling
    paymentMethod: row[5] ?? '',             // String with fallback
    paymentStatus: row[6] ?? '',             // String with fallback
    
    // Proper Number() casting for currency fields
    transferIDR: Number(row[7] ?? 0),
    qrisIDR: Number(row[8] ?? 0), 
    cashIDR: Number(row[9] ?? 0),
    changeIDR: Number(row[10] ?? 0),
    debtIDR: Number(row[11] ?? 0),
    topupIDR: Number(row[12] ?? 0),
    donation: Number(row[13] ?? 0),
    totalPriceIDR: Number(row[14] ?? 0),
    
    // String casting for date fields
    salesDate: String(row[15] ?? ''),
    paidAt: String(row[15] ?? ''),           // Note: Same index as salesDate
    note: String(row[17] ?? ''),
    
    // Number casting for timestamps
    createdAt: Number(row[18] ?? 0),
    updatedAt: Number(row[19] ?? 0),
    deletedAt: Number(row[20] ?? 0),
    });

  }

  let showDetailModal = false;
  let selectedPayment = null;
  // $:selectedSale = null;

  // Function untuk menampilkan detail
  function showPaymentDetail(payment) {
    const row = findSaleRowByPaymentId(payment?.id);
    const sale = mapSaleRowToSale(row);
    selectedPayment = {
      id: sale.id,
      customer: payment?.customer ?? 'Walk-in',
      cashier: sale.cashier || user?.fullName || user?.userName || 'Staff',
      salesDate: sale.salesDate,
      paidAt: sale.paidAt,
      menuIds: sale.menuIds,
      paymentMethod: sale.paymentMethod,
      paymentStatus: sale.paymentStatus,
      transferIDR: sale.transferIDR,
      qrisIDR: sale.qrisIDR,
      totalPriceIDR: sale.totalPriceIDR,
      cashIDR: sale.cashIDR,
      debtIDR: sale.debtIDR,
      topupIDR: sale.topupIDR,
      donation: sale.donation,
      changeIDR: sale.changeIDR,
      note: sale.note,
    };
    showDetailModal = true;

    console.log("Selected Payment", selectedPayment);
  }

  // function showSaleDetail(sale) {

  //   console.log("Payment detail id", sale.id);
  //   console.log("Sales detail", sales);

  //    // Cari data sale yang sesuai berdasarkan payment
  //   selectedSale = sales.find(sales => sales[0] === sale.id);
  //   selectedPayment = null;
  //   showDetailModal = true;
  // }

  function closeDetailModal() {
    showDetailModal = false;
    selectedPayment = null;
    // selectedSale = null;
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
          <div class="card-content-sale">
            <div class="items-scroll">
              <div class="items-list">
                {#each todaySales as sale (sale.id)}
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
              </div>
            </div>
            <div class="total-section total-sticky">
              <div class="total-row">
                <span>Total Hari Ini:</span>
                <span class="total-amount total-blue">Rp {formatCurrency(totalSales)}</span>
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
          <div class="card-content-payment">
            <div class="items-scroll">
              <div class="items-list">
                {#each todayPayments as payment (payment.id)}
                  <div class="item-row item-green clickable-card" on:click={() => showPaymentDetail(payment)}>
                    <div class="item-info">
                      <p class="item-name">{payment.customer}</p>
                      <p class="item-details">
                        {payment.method} • {payment.time}
                      </p>
                    </div>
                    <p class="item-price price-green">Rp {formatCurrency(payment.amount)}</p>
                  </div>
                {/each}
              </div>
            </div>
            <div class="total-section total-sticky">
              <div class="total-row">
                <span>Total Pembayaran:</span>
                <span class="total-amount total-green">Rp {formatCurrency(totalPayments)}</span>
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
            <div class="card-content-unpaid">
              <div class="items-scroll">
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
            <div class="card-content-overpaid">
              <div class="items-scroll">
                <div class="items-list">
                  {#each overpaidItems as item (item.id)}
                    <div class="item-row item-orange">
                      <div class="item-info">
                        <p class="item-name">{item.customer}</p>
                        <p class="item-details">
                          {item.method} • {item.time}
                        </p>
                        <p class="item-date">{item.date}</p>
                      </div>
                      <p class="item-price price-orange">+Rp {formatCurrency(item.excess)}</p>
                    </div>
                  {/each}
                </div>
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
              <div class="items-scroll">
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
              <div class="items-scroll">
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
        </div>

        <!-- Main Dingdong - Full Width -->
        <!-- <div class="card">
          <div class="card-header">
            <div class="header-title">
              <Icon src={FaSolidGamepad} className="icon" color="#4f46e5;" />
              <h2>Main Dingdong</h2>
              <span class="badge badge-indigo">{gameItems.length}</span>
            </div>
          </div>
          <div class="card-content">
            <div class="items-scroll">
              <div class="items-list">
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
        </div> -->

      </div>
    </div>
  </div>


  <!-- Detail Modal -->
  {#if showDetailModal}
  <div class="modal-overlay">
    <div class="modal detail-modal">
      <div class="modal-header">
        <h2>Detail Transaksi</h2>
        <button class="close-btn" on:click={closeDetailModal}>
          <Icon src={FiX} size={18} />
        </button>
      </div>
      
      <div class="modal-content">
        {#if selectedPayment}
          <!-- Receipt-like layout -->
          <div class="receipt">
            <div class="receipt-header">
              <h3>STRUK PEMBELIAN</h3>
              <p class="receipt-date">{selectedPayment.salesDate}</p>
            </div>
            
            <div class="receipt-customer">
              <p><strong>Pelanggan:</strong> {selectedPayment.customer || selectedPayment?.customer || 'Walk-in'}</p>
              <p><strong>Kasir:</strong> {selectedPayment.cashier || user?.name || 'Staff'}</p>
            </div>
            
            <div class="receipt-items">
              <div class="items-header">
                <span>Item</span>
                <span>Qty</span>
                <span>Harga</span>
                <span>Total</span>
              </div>
              
              <!-- Parse menu items dari selectedPayment -->
              {#if selectedPayment.menuIds && Array.isArray(selectedPayment.menuIds)}
                {#each selectedPayment.menuIds as menuId}
                  {@const menuItem = menusAsObject[menuId]}
                  {#if menuItem}
                    <div class="receipt-item">
                      <span class="item-name">{menuItem.name}</span>
                      <span class="item-qty">1</span>
                      <span class="item-price">Rp {formatCurrency(menuItem.price || 0)}</span>
                      <span class="item-total">Rp {formatCurrency(menuItem.price || 0)}</span>
                    </div>
                  {/if}
                {/each}
              {:else}
                <div class="receipt-item">
                  <span class="item-name">{selectedPayment.name || 'Item'}</span>
                  <span class="item-qty">{selectedPayment.qty || 1}</span>
                  <span class="item-price">Rp {formatCurrency(selectedPayment.total || 0)}</span>
                  <span class="item-total">Rp {formatCurrency(selectedPayment.total || 0)}</span>
                </div>
              {/if}
            </div>
            
            <div class="receipt-summary">
              <div class="summary-line">
                <span>Subtotal:</span>
                <span>Rp {formatCurrency(selectedPayment.totalPriceIDR || selectedPayment.total || 0)}</span>
              </div>
              
              {#if selectedPayment}
                <div class="summary-line">
                  <span>Metode Bayar:</span>
                  <span>{selectedPayment.paymentMethod}</span>
                </div>
                <div class="summary-line">
                  <span>Jumlah Bayar:</span>
                  {#if selectedPayment.paymentMethod === 'Cash'}
                    <span>Rp {formatCurrency(selectedPayment.cashIDR)}</span>
                  {:else if selectedPayment.paymentMethod === 'QRIS'}
                    <span>Rp {formatCurrency(selectedPayment.qrisIDR)}</span>
                  {:else if selectedPayment.paymentMethod === 'Transfer'}
                    <span>Rp {formatCurrency(selectedPayment.transferIDR)}</span>
                  {:else if selectedPayment.paymentMethod === 'Debt'}
                    <span>Rp {formatCurrency(selectedPayment.debtIDR)}</span>
                  {:else if selectedPayment.paymentMethod === 'Topup'}
                    <span>Rp {formatCurrency(selectedPayment.topupIDR)}</span>
                  {:else if selectedPayment.paymentMethod === 'Donation'}
                    <span>Rp {formatCurrency(selectedPayment.donationIDR)}</span>
                  {:else}
                    <span>Rp {formatCurrency(selectedPayment.totalPriceIDR)}</span>
                  {/if}
                </div>
                {#if selectedPayment.paymentMethod === 'Cash'}
                  <div class="summary-line">
                    <span>Kembalian:</span>
                    <span>Rp {formatCurrency(selectedPayment.changeIDR)}</span>
                  </div>
                {/if}
              {/if}
              
              <div class="summary-total">
                <span><strong>TOTAL:</strong></span>
                <span><strong>Rp {formatCurrency(selectedPayment.totalPriceIDR || selectedPayment.total || 0)}</strong></span>
              </div>
            </div>
            
            {#if selectedPayment.note}
              <div class="receipt-note">
                <p><strong>Catatan:</strong> {selectedPayment.note}</p>
              </div>
            {/if}
            
            <div class="receipt-footer">
              <p>Terima kasih atas kunjungan Anda!</p>
              <p>ID Transaksi: {selectedPayment.id}</p>
            </div>
          </div>
        {/if}
      </div>
      
      <div class="modal-actions">
        <button class="btn btn-secondary" on:click={closeDetailModal}>Tutup</button>
        <!-- <button class="btn btn-primary">Print Struk</button> -->
      </div>
    </div>
  </div>
  {/if}


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
        <SaleForm on:saleSubmit={handleSaleSubmit} tenants={tenants} {user} />
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
          on:paymentSubmit={handlePaymentSubmit}
          sales={sales}
          tenants={tenants}
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
    flex: 1;
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

  /* Pastikan card-content jadi column dan punya tinggi maksimum */
  .card-content-sale, .card-content-payment {
    display: flex;
    flex-direction: column;
    /* sesuaikan tinggi max-nya dengan layout kamu */
    max-height: 300px;              /* atau pakai calc/60vh */
    padding: 0;                     /* biar sticky nempel rapi, padding bisa dipindah ke child */
  }

  .card-content-unpaid, .card-content-overpaid {
    display: flex;
    flex-direction: column;
    /* sesuaikan tinggi max-nya dengan layout kamu */
    max-height: 210px;              /* atau pakai calc/60vh */
    padding: 0;                     /* biar sticky nempel rapi, padding bisa dipindah ke child */
  }

  /* Area yang di-scroll khusus list */
  .items-scroll {
    overflow: auto;                 /* scroll hanya di sini */
    padding: 1rem;                  /* padding untuk isi list */
    flex: 1;                        /* ambil ruang tersisa */
    min-height: 0;                  /* penting untuk flex container yang scroll */
  }

  /* List tetap seperti semula */
  .items-list {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  /* Total sticky di bawah kartu */
  .total-sticky {
    position: sticky;
    bottom: 0;
    background: white;              /* tutup konten di belakang saat scroll */
    border-top: 1px solid #e5e7eb;
    padding: 0.75rem 1rem;
    /* opsional: bayangan tipis */
    box-shadow: 0 -2px 6px rgba(0,0,0,0.04);
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

  .clickable-card {
    cursor: pointer;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }

  .clickable-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }

  .detail-modal {
    max-width: 500px;
    max-height: 90vh;
    overflow-y: auto;
  }

  .receipt {
    font-family: 'Courier New', monospace;
    border: 2px dashed #ccc;
    padding: 20px;
    background: white;
  }

  .receipt-header {
    text-align: center;
    border-bottom: 1px dashed #ccc;
    padding-bottom: 10px;
    margin-bottom: 15px;
  }

  .receipt-header h3 {
    margin: 0;
    font-size: 18px;
    font-weight: bold;
  }

  .receipt-date {
    margin: 5px 0 0 0;
    font-size: 12px;
    color: #666;
  }

  .receipt-customer {
    margin-bottom: 15px;
    font-size: 14px;
  }

  .receipt-customer p {
    margin: 5px 0;
  }

  .items-header {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr 1fr;
    gap: 10px;
    padding: 10px 0;
    border-bottom: 1px solid #ccc;
    font-weight: bold;
    font-size: 12px;
  }

  .receipt-item {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr 1fr;
    gap: 10px;
    padding: 8px 0;
    border-bottom: 1px dotted #ddd;
    font-size: 12px;
  }

  .receipt-summary {
    margin-top: 15px;
    padding-top: 10px;
    border-top: 1px dashed #ccc;
  }

  .summary-line {
    display: flex;
    justify-content: space-between;
    margin: 5px 0;
    font-size: 14px;
  }

  .summary-total {
    display: flex;
    justify-content: space-between;
    margin: 10px 0;
    padding-top: 10px;
    border-top: 1px solid #333;
    font-size: 16px;
  }

  .receipt-note {
    margin-top: 15px;
    padding: 10px;
    background: #f9f9f9;
    border-left: 3px solid #007bff;
    font-size: 12px;
  }

  .receipt-footer {
    text-align: center;
    margin-top: 20px;
    padding-top: 15px;
    border-top: 1px dashed #ccc;
    font-size: 11px;
  }

  .modal-actions {
    display: flex;
    gap: 10px;
    justify-content: flex-end;
    padding: 20px;
    border-top: 1px solid #eee;
  }

  .btn-secondary {
    background: #6c757d;
    color: white;
  }

  .btn-primary {
    background: #007bff;
    color: white;
  }

</style>
