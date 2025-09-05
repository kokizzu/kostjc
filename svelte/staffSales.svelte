<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/cafe.js').Sale} Sale */
  /** @typedef {import('./_types/cafe.js').Menu} Menu */

  import LayoutMain from './_layouts/main.svelte';
  import MasterSales from './_components/MasterSales.svelte';
  import { AdminSale } from './jsApi.GEN';
  import { CmdList, CmdUpsert, CmdUpdatePayment } from './_components/xConstant';
  import { notifier } from './_components/xNotifier';
    
  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let sale  = /** @type {Sale} */ ({/* sale */});
  let sales = /** @type {any[][]} */([/* sales */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});
  let tenants = /** @type {Record<Number, string>} */ ({/* tenants */});

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminSale( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        sales = o.sales;
      }
    );
  }



  async function OnAddSale(/** @type {Sale} */ sale) {
    const i = /** @type {any} */ ({
      pager,
      sale,
      cmd: CmdUpsert
    });

    await AdminSale(i,
      /** @type {import('../jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        sales = o.sales;
        notifier.showSuccess(`Penjualan berhasil disimpan. Lanjut ke pembayaran.`);
        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type {Sale} */ saleData) {
    const sale = {
        id: String(saleData.id),
        transferIDR: Number(saleData.transferIDR || 0),
        qrisIDR: Number(saleData.qrisIDR || 0),
        cashIDR: Number(saleData.cashIDR || 0),
        debtIDR: Number(saleData.debtIDR || 0),
        topupIDR: Number(saleData.topupIDR || 0),
        totalPriceIDR: Number(saleData.totalPriceIDR || 0),
        paidAt: String(saleData.paidAt || ''),
        paymentMethod: String(saleData.paymentMethod || ''),
        paymentStatus: String(saleData.paymentStatus || ''),
    };

    const i = /** @type {any}*/ ({
      pager,
      sale,
      cmd: CmdUpdatePayment
    });
    await AdminSale(i,
      /** @type {import('./jsApi.GEN').AdminSaleCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        sales = o.sales;
    
        notifier.showSuccess(`Pembayaran berhasil disimpan.`);

        OnRefresh(pager);
      }
    );
  }
</script>

<LayoutMain access={segments} {user}>
  <MasterSales
  sales={sales}
  OnSubmit={OnAddSale}
  OnEdit={OnEdit}
  tenants={tenants}
  />
</LayoutMain>