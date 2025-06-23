<script>
  /** @typedef {import('../_types/masters.js').Access} Access */

  import { onMount } from 'svelte';
  import { UserLogout } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiBuildingsHome2Line, RiMapMapPin2Line, RiUserFacesUser3Line,
    RiUserFacesUserFollowLine, RiBuildingsHotelLine, RiBusinessCalendarScheduleLine,
    RiSystemLogoutBoxRLine, RiFinanceWallet3Line,
    RiBusinessInboxUnarchiveLine, RiOthersDoorOpenLine, RiUserFacesGroupLine,
    RiDocumentFileChartLine, RiSystemMenu2Line, RiFinanceCashLine, RiBusinessBarChartBoxLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { LuHandPlatter } from '../node_modules/svelte-icons-pack/dist/lu';
  import { isOpenSideMenu } from '../_components/xState';

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

<aside class="side-menu-desktop">
  <div class="container">
    <nav class="nav-menu">
      <a href="/" class:active={pathLv1 === ''}>
        <Icon src={RiBuildingsHome2Line} size="20" />
        <span>Home</span>
      </a>
      <a href="/staff/occupancyReport" class:active={pathAll === '/staff/occupancyReport'}>
        <Icon src={RiDocumentFileChartLine} size="20" />
        <span>Occupancy Report</span>
      </a>
      <a href="/staff/missingDataReport" class:active={pathAll === '/staff/missingDataReport'}>
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
        <a href="/admin/menu" class:active={pathLv2 === 'menu'}>
          <Icon src={RiSystemMenu2Line} size="20" />
          <span>Menu</span>
        </a>
        <a href="/admin/sale" class:active={pathLv2 === 'sale'}>
          <Icon src={RiFinanceCashLine} size="20" />
          <span>Sales</span>
        </a>
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
        <a href="/admin/bookingLogs" class:active={String(pathLv2).includes('Logs')||false}>
          <Icon src={RiBusinessBarChartBoxLine} size="20" />
          <span>Logs</span>
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

<button
  aria-label="backdrop toggle menu"
  on:click={() => $isOpenSideMenu = !$isOpenSideMenu}
  class="backdrop {$isOpenSideMenu ? 'open' : ''}"
  >
</button>

<aside class="side-menu-mobile {$isOpenSideMenu ? 'open' : ''}">
  <div class="container">
    <nav class="nav-menu">
      <a href="/" class:active={pathLv1 === ''}>
        <Icon src={RiBuildingsHome2Line} size="20" />
        <span>Home</span>
      </a>
      <a href="/staff/occupancyReport" class:active={pathAll === '/staff/occupancyReport'}>
        <Icon src={RiDocumentFileChartLine} size="20" />
        <span>Occupancy Report</span>
      </a>
      <a href="/staff/missingDataReport" class:active={pathAll === '/staff/missingDataReport'}>
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
        <a href="/admin/menu" class:active={pathLv2 === 'menu'}>
          <Icon src={RiSystemMenu2Line} size="20" />
          <span>Menu</span>
        </a>
        <a href="/admin/sale" class:active={pathLv2 === 'sale'}>
          <Icon src={RiFinanceCashLine} size="20" />
          <span>Sales</span>
        </a>
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
        <a href="/admin/bookingLogs" class:active={String(pathLv2).includes('Logs')||false}>
          <Icon src={RiBusinessBarChartBoxLine} size="20" />
          <span>Logs</span>
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
  aside.side-menu-desktop {
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

  .backdrop {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
    display: none;
    height: 100dvh;
  }

  .backdrop.open {
    display: block;
  }

  aside.side-menu-mobile {
    left: -100%;
    width: 70%;
    position: fixed;
    height: 100dvh;
    top: 0;
    bottom: 0;
    z-index: 1000;
    background-color: #FFF;
    border-top-right-radius: 10px;
    border-bottom-right-radius: 10px;
    overflow-y: auto;
  }

  aside.side-menu-mobile.open {
    transition: left 0.3s ease-in-out;
    left: 0;
  }

  @media only screen and (max-width : 768px) {
    aside.side-menu-desktop {
      display: none;
    }

    aside .container {
      padding: 20px 0;
    }

    aside .container .nav-menu {
      padding: 0 10px;
    }

    aside .container .nav-menu-title {
      padding: 10px 25px;
    }
  }
</style>