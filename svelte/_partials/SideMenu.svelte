<script>
  import { onMount } from 'svelte';
  /** @typedef {import('../_types/masters.js').Access} Access */

  import { UserLogout } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiBuildingsHome2Line, RiMapMapPin2Line, RiUserFacesUser3Line,
    RiUserFacesUserFollowLine, RiBuildingsHotelLine, RiBusinessCalendarScheduleLine,
    RiFoodCupLine, RiSystemLogoutBoxRLine, RiFinanceWallet3Line
  } from '../node_modules/svelte-icons-pack/dist/ri';

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
      <a href="/" class:active={pathLv1 === ''}>
        <Icon src={RiBuildingsHome2Line} size="17" />
        <span>Home</span>
      </a>
      <a href="/user" class:active={pathAll === '/user'}>
        <Icon src={RiUserFacesUser3Line} size="17" />
        <span>Profile</span>
      </a>
    </nav>
    <h3 class="nav-menu-title">Admin</h3>
    <nav class="nav-menu">
      <a href="/admin/location" class:active={pathLv2 === 'location'}>
        <Icon src={RiMapMapPin2Line} size="17" />
        <span>Locations</span>
      </a>
      <a href="/admin/facility" class:active={pathLv2 === 'facility'}>
        <Icon src={RiFoodCupLine} size="17" />
        <span>Facilities</span>
      </a>
      <a href="/admin/building" class:active={pathLv2 === 'building'}>
        <Icon src={RiBuildingsHotelLine} size="17" />
        <span>Buildings</span>
      </a>
      <a href="/admin/booking" class:active={pathLv2 === 'booking'}>
        <Icon src={RiBusinessCalendarScheduleLine} size="17" />
        <span>Bookings</span>
      </a>
      <a href="/admin/payment" class:active={pathLv2 === 'payment'}>
        <Icon src={RiFinanceWallet3Line} size="17" />
        <span>Payments</span>
      </a>
      <a href="/admin/tenants" class:active={pathLv2 === 'tenants'}>
        <Icon src={RiUserFacesUserFollowLine} size="17" />
        <span>Tenants</span>
      </a>
    </nav>
    <span class="separator" />
    <nav class="nav-menu">
      <button class="red" on:click={logout}>
        <Icon src={RiSystemLogoutBoxRLine} size="17" />
        <span>Logout</span>
      </button>
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
    width: 100%;
    margin: auto;
    height: 1px;
    background-color: var(--gray-002);
  }

  aside .container {
    padding: 10px 0;
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu-title {
    margin: 0;
    padding: 10px 25px;
    border-bottom: 1px solid var(--gray-002);
    font-weight: 600;
    text-transform: uppercase;
    font-size: 13px;
  }

  aside .container .nav-menu {
    display: flex;
    flex-direction: column;
    gap: 5px;
    height: fit-content;
    flex-wrap: nowrap;
    width: 100%;
    margin: 0;
    padding: 0 10px;
  }

  aside .container .nav-menu a,
  aside .container .nav-menu button {
    text-decoration: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: 10px;
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
    background-color: var(--blue-transparent);
    color: var(--blue-007);
  }

  aside .container .nav-menu button.red:hover {
    background-color: var(--red-transparent);
    color: var(--red-007);
  }
</style>