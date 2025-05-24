<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */

  import LayoutMain from './_layouts/main.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
</script>


<LayoutMain access={segments} user={user}>
  <div class="report-container">
    <h1>Missing data report</h1>
  </div>
</LayoutMain>

<style>
</style>