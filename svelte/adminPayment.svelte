<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Payment} Payment */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminPayment } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpAddPayment from './_components/PopUpAddPayment.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let payment  = /** @type {Payment} */ ({/* payment */});
  let payments = /** @type {any[][]} */([/* payments */]);
  let bookings = /** @type {Record<Number, string>} */ ({/* bookings */});
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpAddPayment |any
  } */ (null);
  let isSubmitAddPayment = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await AdminPayment( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddPayment = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        payments = o.payments;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      payment: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await AdminPayment(i,
      /** @type {import('./jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        payments = o.payments;
        notifier.showSuccess(`Payment #${row[0]} restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      payment: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminPayment(i,
      /** @type {import('./jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        payments = o.payments;
        notifier.showSuccess(`Payment deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type {any} */ id, /** @type {any[]} */ payloads) {
    const payment = {
      id: id,
      bookingId: payloads[1],
      paymentAt: payloads[2],
      paidIDR: payloads[3],
      paymentMethod: payloads[4],
      paymentStatus: payloads[5],
      note: payloads[6]
    }
    const i = /** @type {any}*/ ({
      pager,
      payment,
      cmd: 'upsert'
    });
    await AdminPayment(i,
      /** @type {import('./jsApi.GEN').AdminPaymentCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        payments = o.payments;
        notifier.showSuccess(`Payment #${payment.id} updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddPayment(/** @type {Payment} */ payment) {
    isSubmitAddPayment = true;
    const i = /** @type {any} */ ({
      pager,
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
        
        pager = o.pager;
        payments = o.payments;
        notifier.showSuccess(`Payment created !!`);

        popUpForms.Reset();

        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddPayment
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddPayment}
    bookings={bookings}
    OnSubmit={OnAddPayment}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-payment">
    <h2>Payment Management</h2>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={payments}
      REFS={{
        'bookingId': bookings,
        'paymentStatus': ['Paid', 'Unpaid', 'Pending', 'Cancelled']
      }}

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
    <button
      class="btn"
      on:click={() => popUpForms.Show()}
      title="add payment"
    >
      <Icon
        color="var(--gray-007)"
        size="18"
        src={RiSystemAddBoxLine}
      />
    </button>
    </MasterTable>
  </div>
</LayoutMain>

<style>
  .master-payment {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-payment h2 {
    margin: 0;
  }
</style>