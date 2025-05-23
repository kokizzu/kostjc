<script>
  import { onMount } from 'svelte';
  /** @typedef {import('../_types/masters.js').Access} Access */

  import { UserLogout } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiBuildingsHome2Line, RiMapMapPin2Line, RiUserFacesUser3Line,
    RiUserFacesUserFollowLine, RiBuildingsHotelLine, RiBusinessCalendarScheduleLine,
    RiSystemLogoutBoxRLine, RiFinanceWallet3Line,
    RiBusinessInboxUnarchiveLine, RiOthersDoorOpenLine, RiUserFacesGroupLine,
    RiDocumentFileChartLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { LuHandPlatter } from '../node_modules/svelte-icons-pack/dist/lu';

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
        <Icon src={RiBuildingsHome2Line} size="20" />
        <span>Home</span>
      </a>
      <a href="/user/occupancyReport" class:active={pathAll === '/user/occupancyReport'}>
        <Icon src={RiDocumentFileChartLine} size="20" />
        <span>Occupancy Report</span>
      </a>
      <a href="/user/missingReport" class:active={pathAll === '/user/missingReport'}>
        <Icon src={RiDocumentFileChartLine} size="20" />
        <span>Missing Data Report</span>
      </a>
      <a href="/user" class:active={pathAll === '/user'}>
        <Icon src={RiUserFacesUser3Line} size="20" />
        <span>Profile</span>
      </a>
    </nav>
    {#if !access.admin}
      <h3 class="nav-menu-title">Staff</h3>
      <nav class="nav-menu">
        <a href="/staff/booking" class:active={pathLv2 === 'booking'}>
          <Icon src={RiBusinessCalendarScheduleLine} size="20" />
          <span>Bookings</span>
        </a>
      </nav>
    {/if}
    {#if access.admin}
      <h3 class="nav-menu-title">Admin</h3>
      <nav class="nav-menu">
        <a href="/admin/location" class:active={pathLv2 === 'location'}>
          <Icon src={RiMapMapPin2Line} size="20" />
          <span>Locations</span>
        </a>
        <a href="/admin/facility" class:active={pathLv2 === 'facility'}>
          <Icon src={LuHandPlatter} size="20" />
          <span>Facilities</span>
        </a>
        <a href="/admin/building" class:active={pathLv2 === 'building'}>
          <Icon src={RiBuildingsHotelLine} size="20" />
          <span>Buildings</span>
        </a>
        <a href="/admin/room" class:active={pathLv2 === 'room'}>
          <Icon src={RiOthersDoorOpenLine} size="20" />
          <span>Rooms</span>
        </a>
        <a href="/admin/booking" class:active={(pathLv1 === 'admin') && (pathLv2 === 'booking')}>
          <Icon src={RiBusinessCalendarScheduleLine} size="20" />
          <span>Bookings</span>
        </a>
        <a href="/admin/payment" class:active={pathLv2 === 'payment'}>
          <Icon src={RiFinanceWallet3Line} size="20" />
          <span>Payments</span>
        </a>
        <a href="/admin/stock" class:active={pathLv2 === 'stock'}>
          <Icon src={RiBusinessInboxUnarchiveLine} size="20" />
          <span>Stocks</span>
        </a>
        <a href="/admin/tenants" class:active={pathLv2 === 'tenants'}>
          <Icon src={RiUserFacesUserFollowLine} size="20" />
          <span>Tenants</span>
        </a>
        <a href="/admin/usersManagement" class:active={pathLv2 === 'usersManagement'}>
          <Icon src={RiUserFacesGroupLine} size="20" />
          <span>Users</span>
        </a>
      </nav>
    {/if}
    <span class="separator" />
    <nav class="nav-menu">
      <button class="red" on:click={logout}>
        <Icon src={RiSystemLogoutBoxRLine} size="20" />
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
    width: var(--sidemenu-width);
    border-right: 1px solid var(--gray-002);
    background-color: #FFF;
    font-size: var(--font-md);
    overflow-y: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: calc(100% - 20px);
    padding: 10px 0 20px 0;
  }

  aside .separator {
    width: 100%;
    margin: auto;
    height: 1px;
    background-color: var(--gray-002);
  }

  aside .container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu-title {
    margin: 0;
    padding: 10px 25px 10px 70px;
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
    padding: 0 10px 0 50px;
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
    font-size: var(--font-lg);
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