<script>
     import { onMount } from 'svelte';
     import { writable } from 'svelte/store';
     import { createEventDispatcher } from 'svelte';

     const dispatch = createEventDispatcher();
   
      export let buyerName;
      export let totalPriceIDR;

      let paidAt = '';
      let transferIDR = 0;
      let qrisIDR = 0;
      let cashIDR = 0;
      let debtIDR = 0;
      let topupIDR = 0;
      let donation = 0;
      
     
     let remainingAmount = writable(totalPriceIDR);

     // Calculate remaining amount when payment methods change
     function calculateRemainingAmount() {
       const totalPayment = 
         Number(transferIDR || 0) + 
         Number(qrisIDR || 0) + 
         Number(cashIDR || 0) + 
         Number(debtIDR || 0) + 
         Number(topupIDR || 0);
       
       remainingAmount.set(totalPriceIDR - totalPayment);
     }
   
     onMount(() => {
       calculateRemainingAmount();
     });
   
     onMount(() => {
       calculateRemainingAmount();
     });
     
     // Format currency
     function formatCurrency(amount) {
       return new Intl.NumberFormat('id-ID').format(amount);
     }


     function submitFormPayment(formData) {
          console.log('Form data submitted:', formData);
          ///tambahkan disini nanti
          dispatch('submit');
     }    
     
   </script>
   
   <div class="form">
     <div class="customer-info">
       <div class="info-row">
         <span class="info-label">Pembeli:</span>
         <span class="info-value">{buyerName || '-'}</span>
       </div>
       <div class="info-row">
         <span class="info-label">Total Harga:</span>
         <span class="info-value price">Rp {formatCurrency(totalPriceIDR)}</span>
       </div>
     </div>
     
     <div class="payment-methods">
       <h3>Metode Pembayaran</h3>
       
       <div class="form-group">
         <label for="transferIDR">Transfer Bank</label>
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
         <label for="cashIDR">Tunai</label>
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
         <label for="debtIDR">Hutang</label>
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
       
       <div class="form-group">
         <label for="donation">Donasi</label>
         <div class="input-group">
           <span class="input-prefix">Rp</span>
           <input 
             type="number" 
             id="donation" 
             bind:value={donation} 
             min="0"
             placeholder="0"
           />
         </div>
       </div>

       <div class="form-group">
        <label for="salesDate">Tanggal Penjualan</label>
        <input 
        type="date" 
        id="salesDate" 
        bind:value={paidAt} 
        required
        />
   </div>
     </div>
     
     <div class="payment-summary">
       <div class="summary-row {$remainingAmount > 0 ? 'negative' : $remainingAmount < 0 ? 'positive' : ''}">
         <span>Sisa Pembayaran:</span>
         <span>
           {#if $remainingAmount > 0}
             Kurang Rp {formatCurrency($remainingAmount)}
           {:else if $remainingAmount < 0}
             Lebih Rp {formatCurrency(Math.abs($remainingAmount))}
           {:else}
             Lunas
           {/if}
         </span>
       </div>
     </div>
     
     <div class="form-actions">
       <button 
         type="submit" 
         class="btn btn-primary"
         disabled={$remainingAmount > 0}
         on:click={submitFormPayment}
       >
         {$remainingAmount > 0 ? 'Pembayaran Kurang' : 'Masukkan Pembayaran'}
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
       font-size: 0.875rem;
       font-weight: 500;
       margin-bottom: 0.5rem;
       color: #374151;
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
       padding: 0.5rem;
       color: #4b5563;
       border-right: 1px solid #d1d5db;
     }
     
     .input-group input {
       flex: 1;
       border: none;
       padding: 0.5rem;
       font-size: 0.875rem;
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
     
     .btn-primary:hover:not(:disabled) {
       background-color: #1d4ed8;
     }
     
     .btn-primary:disabled {
       background-color: #93c5fd;
       cursor: not-allowed;
     }
   </style>
   