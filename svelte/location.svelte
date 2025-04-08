<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/masters.js').Field} Field */
  /** @typedef {import('./_types/masters.js').PagerIn} PagerIn */
  /** @typedef {import('./_types/masters.js').PagerOut} PagerOut */
  /** @typedef {import('./_types/users.js').User} User */
  /** @typedef {import('./_types/property.js').Location} Location */
  
  import LayoutMain from './_layouts/main.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let location  = /** @type {Location} */ ({/* location */});
  let locations = /** @type {any[][]} */([/* locations */]);
  let fields    = /** @type {Field[]} */ ([/* fields */]);
  let pager     = /** @type {PagerOut} */ ({/* pager */});

  console.log('user', user);
  console.log('segments', segments);
  console.log('location', location);
  console.log('locations', locations);
  console.log('fields', fields);
  console.log('pager', pager);
</script>

<LayoutMain access={segments} user={user}>
  <div>
    <h2>Master Location</h2>
  </div>
</LayoutMain>