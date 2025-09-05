<script>
  /** @typedef {import('../_types/cafe').Sale} Sale */

  import { writable } from 'svelte/store';
  import { dateISOFormat } from './xFormatter';
  import { createEventDispatcher } from 'svelte';
  import InputBox from './InputBox.svelte';
  import { parseMenuMapToOptions } from '../_helper/sale';


  const dispatch = createEventDispatcher();

  export let tenants;

  let menuChoices = /** @type {Record<Number, string>} */({/* menuChoices */});
  $: menuOptions = parseMenuMapToOptions(menuChoices);

  const PaymentStatuses = {
      Unpaid: 'Unpaid',
      Paid: 'Paid',
      Overpaid: 'Overpaid',
    };

  let cashier = '';
  let tenantId = 0;
  let buyerName = '';
  let salesDate = dateISOFormat(0);
  let note = '';
  let menuIds = [];
  let totalPriceIDR = 0;

  let selectedMenus = writable([]);

  // Reactive declaration for selectedMenus
  $: selectedMenusValue = $selectedMenus;

  // Calculate total price when selected menus change
 // @ts-ignore
  // @ts-ignore
   // @ts-ignore
     $: if (selectedMenusValue.length > 0) {
    totalPriceIDR = selectedMenusValue.reduce((total, menuId) => {
      const menu = menuOptions.find(m => m.id === menuId);
      return total + (menu ? menu.price : 0);
    }, 0);

    // Update menuIds
    menuIds = [...selectedMenusValue];
  } else {
    totalPriceIDR = 0;
    menuIds = [];
  }

  //logic event di bawah sini
  function formatCurrency(amount) {
    return new Intl.NumberFormat('id-ID').format(amount);
  }

  // @ts-ignore
  // @ts-ignore
  function toggleMenu(menuId) {
    selectedMenus.update(value => {
      if (value.includes(menuId)) {
        return value.filter(id => id !== menuId);
      } else {
        return [...value, menuId];
      }
    });
  }

  function incrementMenu(menuId) {
  selectedMenus.update(value => {
    const newCount = (value[menuId] || 0) + 1;
    return { ...value, [menuId]: newCount };
  });
}

// @ts-ignore
let selectedMenusTest = writable({});

function decrementMenu(menuId) {
  // @ts-ignore
  selectedMenus.update(value => {
    const current = value[menuId] || 0;

    if (current <= 1) {
      // Hapus key-nya jika count <= 1
      const { [menuId]: _, ...rest } = value;
      return rest;
    }

    return {
      ...value,
      [menuId]: current - 1
    };
  });
}


  // @ts-ignore
    $: {
  let result = [];
  totalPriceIDR = 0;

  for (const id in $selectedMenus) {
    const count = $selectedMenus[id];
    const menu = menuOptions.find(m => m.id == id);
    if (menu) {
      totalPriceIDR += menu.price * count;
      result.push(...Array(count).fill(Number(id)));
    }
  }

  menuIds = result;
}

  function submitFormSale() {
    const saleData = /** @type {Sale|any} */ ({
      cashier: cashier,
      tenantId: tenantId+'',
      buyerName: buyerName,
      menuIds: menuIds,
      paymentStatus: PaymentStatuses.Unpaid,
      salesDate: salesDate,
      note: note,
      totalPriceIDR: totalPriceIDR,
    });

    dispatch('saleSubmit', saleData);
  }

  console.log(menuOptions)
</script>

<div class="form">
  <div class="form-group">
    <label for="cashier">Kasir</label>
    <input type="text" id="cashier" bind:value={cashier} required placeholder="Nama Kasir" />
  </div>

  <div class="form-group">
    <InputBox
          id="tenantId"
          label="Tenant"
          isObject={true}
          bind:value={tenantId}
          type="combobox"
          values={tenants}
        />
  </div>

  <div class="form-group">
    <label for="buyerName">Nama Pembeli</label>
    <input type="text" id="buyerName" bind:value={buyerName} required placeholder="Nama Pembeli" />
  </div>

  <div class="form-group">
    <label for="menuIds">Menus</label>
    <div class="menu-grid">
      {#each menuOptions as menu}
  <div class="menu-card">
    <div class="menu-image-with-photo" style="background-image: url('{menu.imageUrl}')"></div>
    <div class="menu-info">
      <p class="menu-name">{menu.name}</p>
      <p class="menu-price">Rp {formatCurrency(menu.price)}</p>
      <div class="menu-counter">
        <button on:click={() => decrementMenu(menu.id)}>-</button>
        <span>{$selectedMenus[menu.id] || 0}</span>
        <button on:click={() => incrementMenu(menu.id)}>+</button>
      </div>
    </div>
  </div>
{/each}
    </div>
  </div>

  <div class="form-group">
    <label for="salesDate">Tanggal Penjualan</label>
    <input type="date" id="salesDate" bind:value={salesDate} required />
  </div>

  <div class="form-group">
    <label for="note">Catatan</label>
    <textarea id="note" bind:value={note} rows="3" placeholder="Catatan tambahan"></textarea>
  </div>

  <div class="form-group total-price">
    <p>Total Harga:</p>
    <span class="price">Rp {formatCurrency(totalPriceIDR)}</span>
  </div>

  <div class="form-actions">
    <button class="btn btn-primary" on:click={submitFormSale}>Simpan Penjualan</button>
  </div>
</div>

<style>
 .form {
    padding: 1rem;
  }

  .form-group {
    margin-bottom: 1rem;
  }

  .form-group label {
    display: block;
    font-size: 13px;
    margin-left: 10px;
    font-weight: 500;
    margin-bottom: 0.5rem;
    color: #374151;
  }

  .form-group input,
  .form-group textarea {
    width: 100%;
    padding: 1.1rem;
    border: 1px solid #d1d5db;
    border-radius: 0.375rem;
    font-size: 1rem;
  }

  .form-group input:focus,
  .form-group textarea:focus {
    outline: none;
    border-color: #2563eb;
    box-shadow: 0 0 0 1px #2563eb;
  }


  .menu-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr); 
    gap: 1rem;
    max-height: 250px; 
    overflow-y: auto;  
    padding-top: 0.5rem;
    border: 1px solid #e5e7eb;
    border-radius: 0.5rem;
    padding: 1rem;
}


  .menu-card {
  all: unset; /* Menghapus default style button */
  padding: 1rem;
  border-radius: 1rem;
  border-color: #374151;
  box-shadow: 0 0 0 2px transparent;
  transition: box-shadow 0.2s ease;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
}

  .menu-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  }

  .menu-info {
    text-align: center;
  }

  .menu-name {
    font-size: 0.875rem;
    font-weight: 600;
    margin: 0 0 0.25rem 0;
    color: #374151;
  }

  .menu-price {
    font-size: 0.85rem;
    font-weight: 500;
    margin: 0;
    color: #6b7280;
  }



  @media (max-width: 480px) {
    .menu-grid {
      grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
      gap: 0.75rem;
    }

    .menu-card {
      padding: 0.75rem;
    }

  }

  .total-price {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.75rem;
    background-color: #f3f4f6;
    border-radius: 0.375rem;
  }

  .price {
    font-weight: 600;
    font-size: 1.125rem;
    color: #2563eb;
  }

  .form-actions {
    display: flex;
    justify-content: flex-end;
    margin-top: 1.5rem;
  }

  .btn {
    padding: 0.5rem 1rem;
    font-size: 0.875rem;
    font-weight: 500;
    border-radius: 0.375rem;
    cursor: pointer;
    border: none;
  }

  .btn-primary {
    background-color: #2563eb;
    color: white;
  }

  .btn-primary:hover {
    background-color: #1d4ed8;
  }

.menu-counter {
  margin-top: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center; /* optional: kalau ingin tombol+angka tetap di tengah container */
  gap: 0.5rem;
}

.menu-counter button {
  background: #2563eb;
  color: white;
  border: none;
  border-radius: 0.375rem;
  padding: 0.25rem 0.5rem;
  cursor: pointer;
  font-size: 1rem;          /* pastikan sama dengan span */
  line-height: 1;           /* agar tidak melebihi tinggi */
  height: 32px;             /* tambahkan tinggi agar seragam */
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-counter span {
  min-width: 1.5rem;
  text-align: center;
  font-size: 1rem;
  line-height: 1;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-image-with-photo {
  width: 80px;
  height: 80px;
  margin: 0 auto 0.75rem;
  border-radius: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  position: relative;
  background-size: cover;
  background-position: center;
}

.menu-image-with-photo::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, rgba(0,0,0,0.2), rgba(255,255,255,0.1));
}
</style>