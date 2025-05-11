<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Booking} Booking */
  /** @typedef {import('./_types/property.js').Facility} Facility */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { StaffBooking } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import PopUpAddBooking from './_components/PopUpAddBooking.svelte';
  import { CmdDelete, CmdList, CmdRestore, CmdUpsert } from './_components/xConstant';

  let user        = /** @type {User} */ ({/* user */});
  let segments    = /** @type {Access} */ ({/* segments */});
  let booking     = /** @type {Booking} */ ({/* booking */});
  let bookings    = /** @type {any[][]} */([/* bookings */]);
  let tenants     = /** @type {Record<Number, string>} */({/* tenants */});
  let facilities  = /** @type {Facility[]} */ ([/* facilities */]);
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement| PopUpAddBooking |any
  } */ (null);
  let isSubmitAddBooking = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: CmdList
    };
    await StaffBooking( // @ts-ignore
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
    await StaffBooking(i,
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
      cmd: CmdDelete
    });
    await StaffBooking(i,
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
      buildingName: String(payloads[1]),
      locationId: String(payloads[2]),
      facilitiesObj: String(payloads[3]),
    }
    const i = /** @type {any}*/ ({
      pager,
      booking,
      cmd: CmdUpsert
    });
    await StaffBooking(i,
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
        notifier.showSuccess(`Booking '${booking.buildingName}' updated !!`);

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

    await StaffBooking(i,
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
        popUpForms.Hide();
        OnRefresh(pager);
      }
    );
  }
</script>

{#if isPopUpFormReady}
  <PopUpAddBooking
    bind:this={popUpForms}
    bind:isSubmitted={isSubmitAddBooking}
    OnSubmit={OnAddBooking}
    facilities={facilities}
    tenants={tenants}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-booking">
    <h2>Booking Management</h2>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'tenantId': tenants
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={bookings}

      CAN_EDIT_ROW={false}
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