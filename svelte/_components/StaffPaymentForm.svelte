  <!-- <script>
      /** @typedef {import('../_types/cafe').Sale} Sale */
      import { onMount } from 'svelte';
      import { writable } from 'svelte/store';
      import { createEventDispatcher } from 'svelte';
      import { dateISOFormat } from './xFormatter';
      import InputBox from './InputBox.svelte';
      import { formatCurrency, getBuyerOptionsForUnpaidItems } from '../_helper/sale';

      const dispatch = createEventDispatcher();

      export let sales;
      export let tenants;
      let buyerOptions = getBuyerOptionsForUnpaidItems(sales, tenants);

      const PaymentMethods = [
        'Cash',
        'QRIS',
        'Transfer',
        // 'Debt',
        'TopUp',
        'Split Payment'
      ];

      const PaymentStatuses = {
        Unpaid: 'Unpaid',
        Paid: 'Paid',
        Overpaid: 'Overpaid',
      };


      let id = 0;
      let paymentMethod = PaymentMethods[0];
      let paidStatus = PaymentStatuses.Unpaid;
      let transferIDR = 0;
      let qrisIDR = 0;
      let cashIDR = 0;
      let debtIDR = 0;
      let topupIDR = 0;
      let paidAt = dateISOFormat(0);
      
      let totalPriceIDR = 0;
      let buyerNameDisplay = '-';
      let remainingAmount = 0;

    


      $: {
    const sale = sales.find(s => Number(s[0]) === Number(id));
    if (sale) {
      totalPriceIDR = sale[13];
      const tenantId = sale[2];
      const buyerName = sale[3];
      
      if (tenantId > 0) {
        const tenantInfo = tenants[tenantId];
        if (tenantInfo) {
          const [tenantName] = tenantInfo.split(' / ');
          buyerNameDisplay = tenantName;
        } else {
          buyerNameDisplay = `Tenant #${tenantId}`;
        }
      } else {
        buyerNameDisplay = buyerName || '-';
      }
    } else {
      totalPriceIDR = 0;
      buyerNameDisplay = '-';
    }
  }

      $: totalPayment =
      Number(transferIDR || 0) +
      Number(qrisIDR || 0) +
      Number(cashIDR || 0) +
      Number(debtIDR || 0) +
      Number(topupIDR || 0);

      $: remainingAmount = Number(totalPriceIDR || 0) - totalPayment;

      $: {
        if (remainingAmount > 0)      paidStatus = PaymentStatuses.Unpaid;
        else if (remainingAmount < 0) paidStatus = PaymentStatuses.Overpaid;
        else                          paidStatus = PaymentStatuses.Paid;
      }

      function submitFormPayment() {
          const totalPaymentNow =
            Number(transferIDR || 0) +
            Number(qrisIDR || 0) +
            Number(cashIDR || 0) +
            Number(debtIDR || 0) +
            Number(topupIDR || 0);

      const rem = Number(totalPriceIDR || 0) - totalPaymentNow;
      const finalStatus =
        rem > 0 ? PaymentStatuses.Unpaid :
        rem < 0 ? PaymentStatuses.Overpaid :
                  PaymentStatuses.Paid;

      const saleData = /** @type {Sale|any} */ ({
            id: id+"",
            totalPriceIDR: totalPriceIDR,
            transferIDR: transferIDR,
            qrisIDR: qrisIDR,
            cashIDR: cashIDR,
            debtIDR: debtIDR,
            topupIDR: topupIDR,
            paidAt: paidAt,
            paymentMethod: paymentMethod,
            paymentStatus: finalStatus,
          });

          dispatch('paymentSubmit', saleData);
        }    
      
    </script>
    
    <div class="form">
      <div class="customer-info">
        <div class="info-row">
          <span class="info-label">Buyer:</span>
          <span class="info-value">{ buyerNameDisplay || '-'}</span>
        </div>
        <div class="info-row">
          <span class="info-label">Total Price:</span>
          <span class="info-value price">Rp {formatCurrency(totalPriceIDR)}</span>
        </div>
      </div>

      <InputBox
        id="tenantId"
        label="Buyer"
        isObject={true}
        bind:value={id}
        type="combobox"
        values={buyerOptions}
      />
      
      <div class="payment-methods">
        <h3>Payment Method</h3>

        <div class="form-group">
          <InputBox
        id="paymentMethod"
        label=""
        bind:value={paymentMethod}
        type="combobox-arr"
        values={PaymentMethods}
      />
        </div>

        <div class="form-group">
          <label for="cashIDR">Cash</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="cashIDR" 
              bind:value={cashIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="qrisIDR">QRIS</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="qrisIDR" 
              bind:value={qrisIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div>
        
        <div class="form-group">
          <label for="transferIDR">Transfer</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="transferIDR" 
              bind:value={transferIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="topupIDR">Top Up</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="topupIDR" 
              bind:value={topupIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div>
        
        <!-- <div class="form-group">
          <label for="debtIDR">Debt</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="debtIDR" 
              bind:value={debtIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div> -->
        
        
        <!-- <div class="form-group">
          <InputBox
          id="paidAt"
          label="Paid At"
          bind:value={paidAt}
          type="datetime"
          placeholder="YYYY-MM-DD"
        />
    </div>
      </div>
      
      <div class="payment-summary">
        <div class="summary-row {remainingAmount > 0 ? 'negative' : remainingAmount < 0 ? 'positive' : ''}">
          <span>Remaining Amount:</span>
          <span>
            {#if remainingAmount > 0}
              Less Rp {formatCurrency(remainingAmount)}
            {:else if remainingAmount < 0}
              More Rp {formatCurrency(Math.abs(remainingAmount))}
            {:else}
              Paid
            {/if}
          </span>
        </div>
      </div>
      
      <div class="form-actions">
        <button 
          type="submit" 
          class="btn btn-primary"
          disabled={remainingAmount > 0}
          on:click={submitFormPayment}
        >
          {remainingAmount > 0 ? 'Remaining Payment' : 'Submit Payment'}
        </button>
      </div>
  </div>
    
    <style>
      .form {
        padding: 1rem;
      }
      
      .customer-info {
        background-color: #f3f4f6;
        padding: 0.75rem;
        border-radius: 0.375rem;
        margin-bottom: 1.5rem;
      }
      
      .info-row {
        display: flex;
        justify-content: space-between;
        margin-bottom: 0.5rem;
      }
      
      .info-row:last-child {
        margin-bottom: 0;
      }
      
      .info-label {
        font-weight: 500;
        color: #4b5563;
      }
      
      .info-value {
        font-weight: 600;
      }
      
      .price {
        color: #2563eb;
      }
      
      .payment-methods {
        margin-bottom: 1.5rem;
      }
      
      .payment-methods h3 {
        font-size: 1rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #374151;
      }
      
      .form-group {
        margin-bottom: 1rem;
      }
      
      .form-group label {
        display: block;
        font-size: 13px;
        margin-left: 10px;
        font-weight: 500;
        color: #374151;
        margin-bottom: 0.5rem;
      }
      
      .input-group {
        display: flex;
        align-items: center;
        border: 1px solid #d1d5db;
        border-radius: 0.375rem;
        overflow: hidden;
      }
      
      .input-prefix {
        background-color: #f3f4f6;
        padding: 0.8rem;
        color: #4b5563;
        border-right: 1px solid #d1d5db;
      }
      
      .input-group input {
        flex: 1;
        border: none;
        padding: 0.8rem;
        font-size: 1rem;
      }
      
      .input-group input:focus {
        outline: none;
      }
      
      .payment-summary {
        background-color: #f3f4f6;
        padding: 0.75rem;
        border-radius: 0.375rem;
        margin-bottom: 1.5rem;
      }
      
      .summary-row {
        display: flex;
        justify-content: space-between;
        font-weight: 600;
      }
      
      .negative {
        color: #dc2626;
      }
      
      .positive {
        color: #059669;
      }
      
      .form-actions {
        display: flex;
        justify-content: flex-end;
      }
      
      
      .btn {
        padding: 0.5rem 1rem;
        font-size: 1rem;
        font-weight: 500;
        border-radius: 0.375rem;
        cursor: pointer;
        border: none;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
      }
      
      .btn-primary {
        background-color: #2563eb;
        color: white;
      }
      
      .btn-primary:hover:not(:disabled) {
        background-color: #1d4ed8;
      }
      
      .btn-primary:disabled {
        background-color: #93c5fd;
        cursor: not-allowed;
      }
    </style> -->
    


    <script>
      /** @typedef {import('../_types/cafe').Sale} Sale */
      import { onMount } from 'svelte';
      import { writable } from 'svelte/store';
      import { createEventDispatcher } from 'svelte';
      import { dateISOFormat } from './xFormatter';
      import InputBox from './InputBox.svelte';
      import { formatCurrency, getBuyerOptionsForUnpaidItems } from '../_helper/sale';

      const dispatch = createEventDispatcher();

      export let sales;
      export let tenants;
      let buyerOptions = getBuyerOptionsForUnpaidItems(sales, tenants);

      const PaymentMethods = [
        'Cash',
        'QRIS',
        'Transfer',
        // 'Debt',
        'TopUp',
        'Split Payment'
      ];

      const PaymentStatuses = {
        Unpaid: 'Unpaid',
        Paid: 'Paid',
        Overpaid: 'Overpaid',
      };


      let id = 0;
      let paymentMethod = PaymentMethods[0];
      let paidStatus = PaymentStatuses.Unpaid;
      let transferIDR = 0;
      let qrisIDR = 0;
      let cashIDR = 0;
      let debtIDR = 0;
      let topupIDR = 0;
      let paidAt = dateISOFormat(0);
      
      let totalPriceIDR = 0;
      let buyerNameDisplay = '-';
      let remainingAmount = 0;

      // Reset payment values when payment method changes
      $: {
        if (paymentMethod) {
          // Reset all payment values when method changes
          if (paymentMethod !== 'Cash' && paymentMethod !== 'Split Payment') {
            cashIDR = 0;
          }
          if (paymentMethod !== 'QRIS' && paymentMethod !== 'Split Payment') {
            qrisIDR = 0;
          }
          if (paymentMethod !== 'Transfer' && paymentMethod !== 'Split Payment') {
            transferIDR = 0;
          }
          if (paymentMethod !== 'TopUp' && paymentMethod !== 'Split Payment') {
            topupIDR = 0;
          }
        }
      }



      $: {
    const sale = sales.find(s => Number(s[0]) === Number(id));
    if (sale) {
      totalPriceIDR = sale[14];
      const tenantId = sale[2];
      const buyerName = sale[3];
      
      if (tenantId > 0) {
        const tenantInfo = tenants[tenantId];
        if (tenantInfo) {
          const [tenantName] = tenantInfo.split(' / ');
          buyerNameDisplay = tenantName;
        } else {
          buyerNameDisplay = `Tenant #${tenantId}`;
        }
      } else {
        buyerNameDisplay = buyerName || '-';
      }
    } else {
      totalPriceIDR = 0;
      buyerNameDisplay = '-';
    }
  }

  $: changeAmount = (paymentMethod === 'Cash' && cashIDR > totalPriceIDR) 
  ? cashIDR - totalPriceIDR 
  : 0;

  $: totalPayment = (() => {
  if (paymentMethod === 'Split Payment') {
    return Number(transferIDR || 0) +
           Number(qrisIDR || 0) +
           Number(cashIDR || 0) +
           Number(topupIDR || 0);
  } else if (paymentMethod === 'Cash') {
    // Untuk cash, payment maksimal adalah total price
    return Math.min(Number(cashIDR || 0), Number(totalPriceIDR || 0));
  } else {
    // Untuk method lain
    return Number(transferIDR || 0) +
           Number(qrisIDR || 0) +
           Number(topupIDR || 0);
  }
  })();

      $: remainingAmount = Number(totalPriceIDR || 0) - totalPayment;

      $: {
        if (remainingAmount > 0)      paidStatus = PaymentStatuses.Unpaid;
        else if (remainingAmount < 0) paidStatus = PaymentStatuses.Overpaid;
        else                          paidStatus = PaymentStatuses.Paid;
      }

      function submitFormPayment() {
        const actualCashPayment = paymentMethod === 'Cash'
          ? Math.min(Number(cashIDR || 0), Number(totalPriceIDR || 0))
          : Number(cashIDR || 0);

        const totalPaymentNow =
          Number(transferIDR || 0) +
          Number(qrisIDR || 0) +
          actualCashPayment +
          Number(debtIDR || 0) +
          Number(topupIDR || 0);

        const rem = Number(totalPriceIDR || 0) - totalPaymentNow;
        const finalStatus =
          rem > 0 ? PaymentStatuses.Unpaid :
          rem < 0 ? PaymentStatuses.Overpaid :
                    PaymentStatuses.Paid;

        const saleData = {
          id: id + "",
          totalPriceIDR,
          transferIDR,
          qrisIDR,
          cashIDR,
          changeIDR: changeAmount,
          debtIDR,
          topupIDR,
          paidAt,
          paymentMethod,
          paymentStatus: finalStatus,
        };

        dispatch('paymentSubmit', saleData);
    }
    
      
    </script>
    
    <div class="form">
      <div class="customer-info">
        <div class="info-row">
          <span class="info-label">Buyer:</span>
          <span class="info-value">{ buyerNameDisplay || '-'}</span>
        </div>
        <div class="info-row">
          <span class="info-label">Total Price:</span>
          <span class="info-value price">Rp {formatCurrency(totalPriceIDR)}</span>
        </div>
      </div>

      <InputBox
        id="tenantId"
        label="Buyer"
        isObject={true}
        bind:value={id}
        type="combobox"
        values={buyerOptions}
      />
      
      <div class="payment-methods">
        <h3>Payment Method</h3>
      
        <div class="form-group">
          <InputBox
            id="paymentMethod"
            label=""
            bind:value={paymentMethod}
            type="combobox-arr"
            values={PaymentMethods}
          />
        </div>
      
        <!-- Cash Payment Fields -->
        <!-- Cash Payment Fields -->
        {#if paymentMethod === 'Cash'}
        <div class="form-group">
          <label for="cashIDR">Cash Received</label>
          <div class="input-group">
            <span class="input-prefix">Rp</span>
            <input 
              type="number" 
              id="cashIDR" 
              bind:value={cashIDR} 
              min="0"
              placeholder="0"
            />
          </div>
        </div>

        <!-- Tampilkan change otomatis jika ada -->
        {#if changeAmount > 0}
          <div class="form-group">
            <label>Change to Give</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <span class="change-amount">
                {formatCurrency(changeAmount)}
              </span>
            </div>
          </div>
        {/if}
        {/if}

      
        <!-- QRIS Payment Fields -->
        {#if paymentMethod === 'QRIS'}
          <div class="form-group">
            <label for="qrisIDR">QRIS Amount</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="qrisIDR" 
                bind:value={qrisIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
        {/if}
      
        <!-- Transfer Payment Fields -->
        {#if paymentMethod === 'Transfer'}
          <div class="form-group">
            <label for="transferIDR">Transfer Amount</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="transferIDR" 
                bind:value={transferIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
        {/if}
      
        <!-- TopUp Payment Fields -->
        {#if paymentMethod === 'TopUp'}
          <div class="form-group">
            <label for="topupIDR">Top Up Amount</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="topupIDR" 
                bind:value={topupIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
        {/if}
      
        <!-- Split Payment Fields -->
        {#if paymentMethod === 'Split Payment'}
          <div class="form-group">
            <label for="cashIDR">Cash</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="cashIDR" 
                bind:value={cashIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
      
          <div class="form-group">
            <label for="qrisIDR">QRIS</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="qrisIDR" 
                bind:value={qrisIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
          
          <div class="form-group">
            <label for="transferIDR">Transfer</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="transferIDR" 
                bind:value={transferIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
      
          <div class="form-group">
            <label for="topupIDR">Top Up</label>
            <div class="input-group">
              <span class="input-prefix">Rp</span>
              <input 
                type="number" 
                id="topupIDR" 
                bind:value={topupIDR} 
                min="0"
                placeholder="0"
              />
            </div>
          </div>
        {/if}
      
        <div class="form-group">
          <InputBox
            id="paidAt"
            label="Paid At"
            bind:value={paidAt}
            type="datetime"
            placeholder="YYYY-MM-DD"
          />
        </div>
      </div>
      
      
      <div class="payment-summary">
        <div class="summary-row {remainingAmount > 0 ? 'negative' : remainingAmount < 0 ? 'positive' : ''}">
          <span>Remaining Amount:</span>
          <span>
            {#if remainingAmount > 0}
              Less Rp {formatCurrency(remainingAmount)}
            {:else if remainingAmount < 0}
              More Rp {formatCurrency(Math.abs(remainingAmount))}
            {:else}
              Paid
            {/if}
          </span>
        </div>
      </div>
      
      <div class="form-actions">
        <button 
          type="submit" 
          class="btn btn-primary"
          disabled={remainingAmount > 0}
          on:click={submitFormPayment}
        >
          {remainingAmount > 0 ? 'Remaining Payment' : 'Submit Payment'}
        </button>
      </div>
  </div>
    
    <style>
      .form {
        padding: 1rem;
      }
      
      .customer-info {
        background-color: #f3f4f6;
        padding: 0.75rem;
        border-radius: 0.375rem;
        margin-bottom: 1.5rem;
      }

      .change-amount {
    flex: 1;
    padding: 0.8rem;
    font-size: 1rem;
    font-weight: 600;
    color: #059669;
    background-color: #f0fdf4;
  }

      
      .info-row {
        display: flex;
        justify-content: space-between;
        margin-bottom: 0.5rem;
      }
      
      .info-row:last-child {
        margin-bottom: 0;
      }
      
      .info-label {
        font-weight: 500;
        color: #4b5563;
      }
      
      .info-value {
        font-weight: 600;
      }
      
      .price {
        color: #2563eb;
      }
      
      .payment-methods {
        margin-bottom: 1.5rem;
      }
      
      .payment-methods h3 {
        font-size: 1rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #374151;
      }
      
      .form-group {
        margin-bottom: 1rem;
      }
      
      .form-group label {
        display: block;
        font-size: 13px;
        margin-left: 10px;
        font-weight: 500;
        color: #374151;
        margin-bottom: 0.5rem;
      }
      
      .input-group {
        display: flex;
        align-items: center;
        border: 1px solid #d1d5db;
        border-radius: 0.375rem;
        overflow: hidden;
      }
      
      .input-prefix {
        background-color: #f3f4f6;
        padding: 0.8rem;
        color: #4b5563;
        border-right: 1px solid #d1d5db;
      }
      
      .input-group input {
        flex: 1;
        border: none;
        padding: 0.8rem;
        font-size: 1rem;
      }
      
      .input-group input:focus {
        outline: none;
      }
      
      .payment-summary {
        background-color: #f3f4f6;
        padding: 0.75rem;
        border-radius: 0.375rem;
        margin-bottom: 1.5rem;
      }
      
      .summary-row {
        display: flex;
        justify-content: space-between;
        font-weight: 600;
      }
      
      .negative {
        color: #dc2626;
      }
      
      .positive {
        color: #059669;
      }
      
      .form-actions {
        display: flex;
        justify-content: flex-end;
      }
      
      
      .btn {
        padding: 0.5rem 1rem;
        font-size: 1rem;
        font-weight: 500;
        border-radius: 0.375rem;
        cursor: pointer;
        border: none;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
      }
      
      .btn-primary {
        background-color: #2563eb;
        color: white;
      }
      
      .btn-primary:hover:not(:disabled) {
        background-color: #1d4ed8;
      }
      
      .btn-primary:disabled {
        background-color: #93c5fd;
        cursor: not-allowed;
      }
    </style>
    