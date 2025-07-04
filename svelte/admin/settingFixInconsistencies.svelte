<script>
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/users.js').User} User */
  
  import LayoutMain from '../_layouts/main.svelte';
    import RoomsInconsistencies from '../_partials/setting/RoomsInconsistencies.svelte';
  import SettingsSubMenu from '../_partials/SettingsSubMenu.svelte';

  let user               = /** @type {User} */ ({/* user */});
  let segments           = /** @type {Access} */ ({/* segments */});
  const tenants          = /** @type {Record<number, string>} */ ({/* tenants */});
  const rooms            = /** @type {Record<number, string>} */ ({/* rooms */});
</script>

<LayoutMain access={segments} user={user}>
  <div class="setting-container">
    <SettingsSubMenu />
    <RoomsInconsistencies {tenants} {rooms} />
  </div>
</LayoutMain>

<style>
  .setting-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
</style>