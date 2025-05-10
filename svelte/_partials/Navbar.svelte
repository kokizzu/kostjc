<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiUserFacesUserLine } from '../node_modules/svelte-icons-pack/dist/ri';

  export let username = /** @type {string} */ ('unknown');
  export let role = /** @type {'Staff' | 'Admin' | string} */ ('Staff');
</script>

<header>
  <div class="label">
    <a href="/" class="logo">
      <img src="/assets/house.svg" alt="House" />
      <h2>KostJC</h2>
    </a>
  </div>
  <div class="info">
    <span class="role {role}" title="You are {role}">{role}</span>
    <a href="/user" class="user">
      <span>{username}</span>
      <Icon src={RiUserFacesUserLine} size="15" />
    </a>
  </div>
</header>

<style>
  header {
    top: 0;
    left: 0;
    right: 0;
    position: fixed;
    width: 100%;
    height: var(--navbar-height);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 50px;
    border-bottom: 1px solid var(--gray-002);
    background-color: #FFF;
    z-index: 100;
  }

  header .label {
    width: fit-content;
    display: flex;
    flex-direction: row;
    gap: 0;
    align-items: center;
    padding-left: 20px;
  }

  header .label .logo {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 8px;
    font-size: 17px;
    color: var(--gray-008);
    text-decoration: none;
    width: calc(var(--sidemenu-width) - 20px);
  }

  header .label .logo img {
    width: 30px;
    height: auto;
  }

  header .info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    font-size: 13px;
  }

  header .info .role {
    padding: 8px 15px;
    border-radius: 9999px;
    font-weight: 500;
    cursor: pointer;
  }

  header .info .role.Staff {
    color: var(--orange-007);
    background-color: var(--orange-transparent);
  }

  header .info .role.Admin {
    color: var(--green-007);
    background-color: var(--green-transparent);
  }

  header .info a.user {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: center;
    gap: 6px;
    text-decoration: none;
    color: var(--gray-008);
    padding: 8px 20px;
    border-radius: 9999px;
    border: 1px solid var(--gray-002);
  }

  header .info a.user:hover {
    background-color: var(--gray-001);
  }
</style>