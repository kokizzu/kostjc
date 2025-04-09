<script>
  import { onMount } from 'svelte';
  /** @typedef {import('../_types/masters.js').Access} Access */

  import { UserLogout } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier.js';

  export let access = /** @type {Access} */ ({});

  onMount( () => {
    console.log( access );
  } );

  const pathAll = /** @type {string}*/ (window.location.pathname);
  const pathLv1 = /** @type {string}*/ (window.location.pathname.split( '/' )[ 1 ]);
  const pathLv2 = /** @type {string}*/ (window.location.pathname.split( '/' )[ 2 ]);

  async function logout() {
    await UserLogout( {}, function( /** @type any */ o ) {
      if( o.error ) {
        notifier.showError( o.error );
        console.log( o );
        return;
      }

      notifier.showSuccess( 'Logged out' );

      return new Promise( resolve => {
        setTimeout( () => {
          window.location.href = '/';
          resolve();
        }, 1000 );
      });
    });
  }
</script>

<aside>
  <div class="container">
    <nav class="nav-menu">
      <a href="/" class:active={pathLv1 === ''}>Home</a>
      <a href="/user/location" class:active={pathLv2 === 'location'}>Location</a>
    </nav>
    <span class="separator" />
    <nav class="nav-menu">
      <a href="/user" class:active={pathAll === '/user'}>Profile</a>
      <button class="red" on:click={logout}>Logout</button>
    </nav>
  </div>
</aside>

<style>
  aside {
    position: fixed;
    top: var(--navbar-height);
    left: 0;
    bottom: 0;
    height: 100%;
    width: var(--sidemenu-width);
    border-right: 1px solid var(--gray-002);
    background-color: #FFF;
    font-size: var(--font-md);
  }

  aside .separator {
    width: 80%;
    margin: auto;
    height: 1px;
    background-color: var(--gray-002);
  }

  aside .container {
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu {
    display: flex;
    flex-direction: column;
    gap: 5px;
    height: fit-content;
    flex-wrap: nowrap;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu a,
  aside .container .nav-menu button {
    text-decoration: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: 10px 15px;
    border-radius: 8px;
    cursor: pointer;
    color: var(--gray-008);
    border: none;
    background-color: transparent;
    font-size: var(--font-md);
  }

  aside .container .nav-menu a:hover,
  aside .container .nav-menu button:hover {
    background-color: var(--gray-001);
  }

  aside .container .nav-menu a.active {
    background-color: var(--orange-transparent);
    color: var(--orange-007);
  }

  aside .container .nav-menu button.red:hover {
    background-color: var(--red-transparent);
    color: var(--red-007);
  }
</style>