<script>
  /** @typedef {import('../_types/users').User} User */
  /** @typedef {import('../_types/masters').Access} Access */

  import SideMenu from '../_partials/SideMenu.svelte';
  import Navbar from '../_partials/Navbar.svelte';
  import Footer from '../_partials/Footer.svelte';
  import { isShrinkMenu } from '../_states/page';

  export let user   = /** @type {User} */ ({});
  export let access = /** @type {Access} */ ({});
</script>

<div class="root-layout">
  <div class="root-container">
    <Navbar username={user.userName} role={user.role}/>
    <div class="root-content { $isShrinkMenu ? 'shrink' : 'expand' }">
      <SideMenu {access} />
      <main class="content">
        <slot />
        <Footer />
      </main>
    </div>
  </div>
</div>

<style>
  .root-layout {
    display: block;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		height: 100vh;
		width: 100vw;
  }

  .root-layout .root-container {
    height: 100%;
		width: 100%;
		display: flex;
  }

  .root-layout .root-container .root-content {
		display: flex;
		flex-direction: column;
		-webkit-box-orient: vertical;
		-webkit-box-direction: normal;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
    max-width: 100%;
    margin-top: var(--navbar-height);
    margin-left: var(--sidemenu-width);
    margin-bottom: 0;
    overflow: hidden;
  }

  .root-layout .root-container .root-content .content {
    display: flex;
		flex-direction: column;
		justify-content: space-between;
		min-height: 100%;
    overflow-y: auto;
    padding: 0;
    margin: 0;
    font-size: var(--font-base);
  }
</style>