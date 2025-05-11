<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Booking} Booking */
  /** @typedef {import('./_types/property.js').Facility} Facility */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTableBooking.svelte';
  import { onMount } from 'svelte';
  import { AdminBooking } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import PopUpAddBooking from './_components/PopUpAddBooking.svelte';
    import { CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let booking     = /** @type {Booking} */ ({/* booking */});
  let bookings    = /** @type {any[][]} */([/* bookings */]);
  let tenants     = /** @type {Record<Number, string>} */({/* tenants */});
  let rooms       = /** @type {Record<Number, string>} */({/* rooms */});
  let facilities  = /** @type {Facility[]} */ ([/* facilities */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement| PopUpAddBooking |any
  } */ (null);
  let isSubmitAddBooking = /** @type boolean */ (false);

  const hashSearchByTenant  = 'by-tenant';
  const hashSearchByRoom    = 'by-room';

  let hash = /** @type {string} */ (location.hash || '');

  onMount(() => {
    if ( hash[ 0 ] === '#' ) hash = hash.substring( 1 );

    if (hash.includes(hashSearchByTenant)) {
      const tenantId = hash.replace(hashSearchByTenant+':', '');
      pager.filters = {
        tenantId: [tenantId]
      }
      OnRefresh(pager);
    }

    if (hash.includes(hashSearchByRoom)) {
      const roomId = hash.replace(hashSearchByRoom+':', '');
      pager.filters = {
        roomId: [roomId]
      }
      OnRefresh(pager);
    }
    
    isPopUpFormReady = true
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await AdminBooking( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBooking = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        bookings = o.bookings;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      booking: {
        id: row[0]
      },
      cmd: CmdRestore
    });
    await AdminBooking(i,
      /** @type {import('./jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        bookings = o.bookings;
        notifier.showSuccess(`Booking '${row[1]}' restored !!`);

    
        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      booking: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminBooking(i,
      /** @type {import('./jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        bookings = o.bookings;
        notifier.showSuccess(`Booking '${row[1]}' deleted !!`);

    
        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const booking = {
      id: payloads[0],
      roomId: payloads[1],
      dateStart: payloads[3],
      dateEnd: payloads[4],
      tenantId: payloads[5],
      basePriceIDR: payloads[6],
      facilitiesObj: String(payloads[7]),
      totalPriceIDR: payloads[8],
      paidAt: payloads[9],
      extraTenants: payloads[10],
    }
    const i = /** @type {any}*/ ({
      pager,
      booking,
      cmd: CmdUpsert
    });
    await AdminBooking(i,
      /** @type {import('./jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        bookings = o.bookings;
        notifier.showSuccess(`Booking #${booking.id} updated !!`);
        OnRefresh(pager);
      }
    );
  }

  async function OnAddBooking(/** @type {Booking} */ booking, /** @type {number[]} */ facilities) {
    isSubmitAddBooking = true;
    const i = /** @type {any} */ ({
      pager,
      facilities,
      booking,
      cmd: CmdUpsert
    });

    await AdminBooking(i,
      /** @type {import('../jsApi.GEN').AdminBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBooking = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        bookings = o.bookings;
        notifier.showSuccess(`Booking created !!`);

        popUpForms.Reset();
        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddBooking
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddBooking}
    OnSubmit={OnAddBooking}
    facilities={facilities}
    tenants={tenants}
    rooms={rooms}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-booking">
    <h2>Booking Management</h2>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'tenantId': tenants,
        'roomId': rooms
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={bookings}

      tenants={tenants}
      
      CAN_SEARCH_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
    <button
      class="btn"
      on:click={() => popUpForms.Show()}
      title="add booking"
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
  .master-booking {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-booking h2 {
    margin: 0;
  }
</style>