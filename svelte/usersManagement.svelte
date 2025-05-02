<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  
  import LayoutMain from './_layouts/main.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { onMount } from 'svelte';
  import { AdminUsersManagement } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import PopUpForms from './_components/PopUpForms.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let users = /** @type {any[][]} */([/* users */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  console.log('usersManagement: ', users);

  let isPopUpFormReady = /** @type boolean */ (false);
  let popUpForms = /** @type {
    import('svelte').SvelteComponent | HTMLElement | PopUpForms |any
  } */ (null);
  let isSubmitAddUser = /** @type boolean */ (false);

  onMount(() => isPopUpFormReady = true);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await AdminUsersManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').AdminUsersManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddUser = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        users = o.users;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      user: {
        id: row[0]
      },
      cmd: 'restore'
    });
    await AdminUsersManagement(i,
      /** @type {import('./jsApi.GEN').AdminUsersManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        notifier.showSuccess(`User '${row[1]}' restored !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = /** @type {any}*/ ({
      pager,
      user: {
        id: row[0]
      },
      cmd: 'delete'
    });
    await AdminUsersManagement(i,
      /** @type {import('./jsApi.GEN').AdminUsersManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        notifier.showSuccess(`User '${row[1]}' deleted !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    console.log('User ID to Edit: ' + String(payloads));
    const user = {
      id: id,
      email: payloads[1],
      userName: payloads[2],
      fullName: payloads[3],
      role: payloads[4]
    }
    const i = /** @type {any}*/ ({
      pager,
      user,
      cmd: 'upsert'
    });
    await AdminUsersManagement(i,
      /** @type {import('./jsApi.GEN').AdminUsersManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        notifier.showSuccess(`User '${user.fullName}' updated !!`);

        OnRefresh(pager);
      }
    );
  }

  async function OnAddUser(/** @type any[] */ payloads) {
    isSubmitAddUser = true;

    const user = /** @type {any} */ ({
      email: payloads[1],
      userName: payloads[2],
      fullName: payloads[3],
      role: payloads[4]
    });
    const i = /** @type {any} */ ({
      pager,
      user,
      cmd: 'upsert'
    });

    await AdminUsersManagement(i,
      /** @type {import('../jsApi.GEN').AdminUsersManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddUser = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        users = o.users;
        notifier.showSuccess(`User '${user.facilityName}' created !!`);
        
        OnRefresh(pager);
        popUpForms.Reset();
        popUpForms.Hide();
      }
    );
  }
</script>

{#if isPopUpFormReady}
  <PopUpForms
    bind:this={popUpForms}
    heading="Add User"
    FIELDS={fields}
    REFS={{
      'role': ['Admin', 'User']
    }}
    bind:isSubmitted={isSubmitAddUser}
    OnSubmit={OnAddUser}
  />
{/if}

<LayoutMain access={segments} user={user}>
  <div class="master-user">
    <h2>User Management</h2>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={users}
      REFS={{
        'role': ['Admin', 'User']
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
      title="add user"
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
  .master-user {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .master-user h2 {
    margin: 0;
  }
</style>