<script>
  /** @typedef {import('../_types/users').User} User */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  export let userId      = /** @type {number} */ (0);
  let newPassword = /** @type {string} */ ('');

  export let OnSubmit = async function(/** @type {User} */ user) {
    console.log('OnSubmit :::', user);
  }

  async function submitAdd() {
    const user = /** @type {User} */ ({
      id: userId,
      password: newPassword
    });

    await OnSubmit(user);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
    newPassword = '';
  }
const cancel = () => isShow = false;
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
<div class="popup">
  <header class="header">
    <h2>Update Password for User {`#${userId}`}</h2>
    <button on:click={Hide}>
      <Icon size="22" color="var(--red-005)" src={IoClose}/>
    </button>
  </header>
  <div class="forms">
    <InputBox
      id="newPassword"
      label="New Password"
      type="password"
      bind:value={newPassword}
      placeholder="secret#0x3Hjj"
    />
  </div>
  <div class="foot">
    <div class="left">
    </div>
    <div class="right">
      <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
      <button class="ok" on:click|preventDefault={submitAdd} disabled={isSubmitted}>
        {#if !isSubmitted}
          <span>Submit</span>
        {/if}
        {#if isSubmitted}
          <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
        {/if}
      </button>
    </div>
  </div>
</div>
</div>

<style>
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

:global(.spin) {
  animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
}

.popup-container {
  display: none;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  z-index: 2000;
  background-color: rgba(0 0 0 / 40%);
  backdrop-filter: blur(1px);
  justify-content: center;
  padding: 50px;
  overflow: auto;
}

.popup-container.show {
  display: flex;
}

.popup-container .popup {
  border-radius: 8px;
  background-color: #FFF;
  height: fit-content;
  width: 500px;
  display: flex;
  flex-direction: column;
}

.popup-container .popup header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  border-bottom: 1px solid var(--gray-004);
}

.popup-container .popup header h2 {
  margin: 0;
}

.popup-container .popup header button {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 5px;
  border-radius: 50%;
  border: none;
  background-color: transparent;
  cursor: pointer;
}

.popup-container .popup header button:hover {
  background-color: #ef444420;
}

.popup-container .popup header button:active {
  background-color: #ef444430;
}

.popup-container .popup .forms {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.popup-container .popup .foot {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 10px;
  align-items: center;
  padding: 15px 20px;
  border-top: 1px solid var(--gray-004);
}

.popup-container .popup .foot .right {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
}

.popup-container .popup .foot button {
  padding: 8px 18px;
  border-radius: 9999px;
  border: none;
  color: #FFF;
  cursor: pointer;
  font-weight: 600;
}

.popup-container .popup .foot button.ok {
  background-color: var(--green-006);
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup-container .popup .foot button.ok:hover {
  background-color: var(--green-005);
}

.popup-container .popup .foot button.ok:disabled {
  cursor: not-allowed;
  background-color: var(--gray-003);
  color: var(--gray-007);
}

.popup-container .popup .foot button.cancel {
  background-color: transparent;
  color: var(--gray-007);
}

.popup-container .popup .foot button.cancel:hover {
  background-color: var(--gray-001);
}

.facilities-form {
  display: flex;
  flex-direction: column;
  gap: 3px;
}

@media only screen and (max-width : 768px) {
  .popup-container {
    padding: 10px;
  }

  .popup-container .popup {
      width: 100%;
    }
}
</style>