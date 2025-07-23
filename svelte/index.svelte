<script>
  /** @typedef {import('./_types/masters.js').Access} Access */
  /** @typedef {import('./_types/users.js').User} User */

  import { onMount, tick } from 'svelte';
  import { GuestLogin, GuestRegister } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { FaSolidCircleNotch } from './node_modules/svelte-icons-pack/dist/fa';
  import InputBox from './_components/InputBox.svelte';
  import LayoutMain from './_layouts/main.svelte';
  import BirthDayAgenda from './_partials/home/BirthDayAgenda.svelte';
  import AvailableRooms from './_partials/home/AvailableRooms.svelte';
    import UnpaidBookingTenants from './_partials/home/UnpaidBookingTenants.svelte';
    import DoubleBookingReports from './_partials/home/DoubleBookingReports.svelte';
    import UpcomingTenants from './_partials/home/UpcomingTenants.svelte';

  let title = '#{title}';
  let user  = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});

  function getCookie(/** @type {string} */ name ) {
    var match = document.cookie.match( new RegExp( '(^| )' + name + '=([^;]+)' ) );
    if( match ) return match[ 2 ];
  }

  let email = '';
  let password = '';
  let confirmPassword = '';

  let passInput = /** @type {HTMLInputElement} */ ({});

  const ModeLogin       = 'LOGIN';
  const ModeRegister    = 'REGISTER';
  const ModeUser        = 'USER';
  const ModeFgPassword  = 'FORGOT_PASSWORD';

  let Mode = ModeLogin;

  let isSubmitting = false;

  async function onHashChange() {
    const auth = getCookie( 'auth' );
    if( auth && user && !auth.startsWith( 'TEMP__' ) ) {
      location.hash = ModeUser;
      Mode = ModeUser;
      return;
    }

    let hash = /** @type {string} */ (location.hash || '');

    if ( hash[ 0 ] === '#' ) hash = hash.substring( 1 );

    switch ( hash ) {
      case ModeLogin:
        Mode = ModeLogin;
        break;
      case ModeRegister:
        Mode = ModeRegister;
        break;
      case ModeFgPassword:
        Mode = ModeFgPassword;
        break;
      default:
        Mode = ModeLogin;
    }
    await tick();
  }

  onMount( async () => {
    await onHashChange();
  });

  async function guestRegister() {
    isSubmitting = true
    if( !email ) {
      isSubmitting = false;
      notifier.showError('Email is required' );
      return
    }
    if( password.length<12 ) {
      isSubmitting = false;
      notifier.showError('Password must be at least 12 characters' );
      return
    }
    if( password!==confirmPassword ) {
      isSubmitting = false;
      notifier.showError('Passwords do not match' );
      return
    }

    const i = {email, password};

    await GuestRegister( i, async function( /** @type any */ o ) {
      if( o.error ) {
        isSubmitting = false;
        notifier.showError(o.error );
		return
      }
      isSubmitting = false;
      notifier.showSuccess('Registered successfully, a registration verification has been sent to your email' );
      Mode = ModeLogin;
      password = '';
      await tick();
      passInput.focus();
    } );
  }

  async function guestLogin() {
    isSubmitting = true;
    if( !email ) {
      isSubmitting = false;
      notifier.showError('Email is required' );
      return
    }
    if( password.length<12 ) {
      isSubmitting = false;
      notifier.showError('Password must be at least 12 characters' );
      return
    }
    const i = {email, password};
    await GuestLogin( i, async function(/** @type any */ o ) {
      if( o.error ) {
        isSubmitting = false;
        notifier.showError( o.error );
        return
      }
      isSubmitting = false;
      notifier.showSuccess('Login successfully' );
      setTimeout( () => {
        user = o.user;
        segments = o.segments;
        onHashChange();

        window.document.location = '/';
      }, 1000 );
    } );
  }
</script>

<svelte:window on:hashchange={onHashChange}/>

{#if Mode === ModeUser}
  <LayoutMain access={segments} user={user}>
    <div class="home-container">
      <BirthDayAgenda />
      <AvailableRooms />
      <UnpaidBookingTenants />
      <DoubleBookingReports />
      <UpcomingTenants />
    </div>
  </LayoutMain>
{:else}
  <section class="auth-section">
    <div class="main-container">
      <div class="title-container">
        <p>KostJC</p>
        <h1>{Mode.split( '_' ).join( ' ' )}</h1>
      </div>
      <div class="form-container">
        <div class="input-container">
          {#if Mode === ModeLogin || Mode === ModeRegister}
            <InputBox
              id="email"
              label="Email"
              type="text"
              bind:value={email}
              autocomplete="on"
            />
          {/if}

          {#if Mode === ModeLogin || Mode === ModeRegister}
            <InputBox
              id="password"
              label="Password"
              type="password"
              bind:value={password}
            />
          {/if}

          {#if Mode === ModeRegister}
            <InputBox
              id="confirmPass"
              label="Confirm Password"
              type="password"
              bind:value={confirmPassword}
            />
          {/if}
        </div>
        <div class="button-container">
          {#if Mode===ModeRegister}
            <button on:click={guestRegister}>
              {#if isSubmitting===true}
                <Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
              {/if}
              {#if isSubmitting===false}
                <span>Register</span>
              {/if}
            </button>
          {/if}
          {#if Mode===ModeLogin}
            <button on:click={guestLogin}>
              {#if isSubmitting===true}
                <Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
              {/if}
              {#if isSubmitting===false}
                <span>Login</span>
              {/if}
            </button>
          {/if}
        </div>
        <!-- Oauth Buttons -->
        {#if Mode===ModeRegister || Mode===ModeLogin}
          <div class="oauth-container">
            <div class="or-separator">
              <span/>
              <p>or</p>
              <span/>
            </div>
          </div>
        {/if}
        <div class="foot-auth">
          {#if Mode!==ModeRegister}
            <p>Have no account? <a href={`#${ModeRegister}`} on:click={() => (Mode = ModeRegister)}>Register</a></p>
          {/if}
          {#if Mode!==ModeLogin}
            <p>Already have account? <a href={`#${ModeLogin}`} on:click={() => (Mode = ModeLogin)}>Login</a></p>
          {/if}
        </div>
      </div>
    </div>
  </section>
{/if}

<style>
  .home-container {
    display: flex;
    flex-direction: column;
    padding: 20px;
    gap: 20px;
  }

  @keyframes spin {
    from {
      transform : rotate(0deg);
    }
    to {
      transform : rotate(360deg);
    }
  }

  :global(.spin) {
    animation : spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .auth-section {
    height: 100%;
    width: 100%;
    background-color: var(--gray-001);
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    display: flex;
    color: var(--gray-008);
  }

  .auth-section .main-container {
    width: 480px;
    height: fit-content;
    padding: 20px;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    background-color: white;
    margin: 50px auto;
    border: 1px solid var(--gray-003);
    gap: 10px;
  }

  .auth-section .main-container .title-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
    text-align: center;
  }

  .auth-section .main-container .title-container p {
    font-size: 16px;
    font-weight: 600;
    color: var(--blue-005);
    margin: 0;
  }

  .auth-section .main-container .title-container h1 {
    margin: 0 0 5px 0;
    font-size: 22px;
    font-weight: 700;
  }

  .auth-section .main-container .form-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
  }

  .auth-section .main-container .form-container .input-container {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .auth-section .main-container .form-container .forgot-password a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .auth-section .main-container .form-container .forgot-password a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }

  .auth-section .main-container .form-container .button-container {
    height: fit-content;
    width: 100%;
    display: flex;
  }

  .auth-section .main-container .form-container .button-container button {
    margin: 0;
    width: 100%;
    padding: 12px;
    font-size: 15px;
    font-weight: 600;
    background-color: var(--blue-006);
    border-radius: 8px;
    color: white;
    border: none;
    cursor: pointer;
  }

  .auth-section .main-container .form-container .button-container button:hover {
    background-color : var(--blue-005);
  }

  .auth-section .main-container .form-container .oauth-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator span {
    flex-grow: 1;
    height: 0;
    border-top: 1px solid var(--gray-002);
    padding: 0;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator p {
    width       : fit-content;
    font-weight : 600;
    padding     : 0 10px;
    margin: 0;
  }

  .auth-section .main-container .form-container .oauth-container .button span {
    margin-left: 8px;
  }

  .auth-section .main-container .form-container .foot-auth {
    display: flex;
    flex-direction: column;
    gap: 5px;
    text-align: center;
  }

  .auth-section .main-container .form-container .foot-auth p {
    font-weight: 500;
    margin: 0;
  }

  .auth-section .main-container .form-container .foot-auth a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .auth-section .main-container .form-container .foot-auth a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }

  @media only screen and (max-width : 768px) {
    .auth-section {
      background-color: #FFF;
    }

    .auth-section .main-container {
      border: none;
    }
  }
</style>