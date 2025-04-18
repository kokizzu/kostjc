<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Booking} Booking */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { UserBooking } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let booking  = /** @type {Booking} */ ({/* booking */});
  let bookings = /** @type {any[][]} */([/* bookings */]);
  let tenants = /** @type {Record<Number, string>} */({/* tenants */});
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement |any
  } */ (null);
  let isSubmitAddBuilding = /** @type boolean */ (false);

  onMount(() => {
    isPopUpFormReady = true
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await UserBooking( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').UserBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBuilding = false;
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
      building: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await UserBooking(i,
      /** @type {import('./jsApi.GEN').UserBookingCallback} */
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
      building: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await UserBooking(i,
      /** @type {import('./jsApi.GEN').UserBookingCallback} */
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
    console.log('Booking ID to Edit: ' + String(id));
    const building = {
      id: payloads[0],
      buildingName: String(payloads[1]),
      locationId: String(payloads[2]),
      facilitiesObj: String(payloads[3]),
    }
    const i = /** @type {any}*/ ({
      pager,
      building,
      cmd: 'upsert'
    });
    await UserBooking(i,
      /** @type {import('./jsApi.GEN').UserBookingCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        bookings = o.bookings;
        notifier.showSuccess(`Booking '${building.buildingName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddBooking(/** @type any[] */ payloads) {
    isSubmitAddBuilding = true;
    const i = /** @type {any} */ ({
      pager,
      booking: {
        id: payloads[0],
      },
      cmd: 'upsert'
    });

    await UserBooking(i,
      /** @type {import('../jsApi.GEN').TenantAdminProductsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBuilding = false;
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
  <PopUpForms
    bind:this={popUpForms}
    heading="Add Facility"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddBuilding}
    OnSubmit={OnAddBooking}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-booking">
    <h2>Master Booking</h2>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'tenantId': tenants
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={bookings}

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
      title="add booking"
    >
      <Icon
        color="var(--gray-007)"
        size="16"
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