<script>
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/users.js').User} User */

  import LayoutMain from '../_layouts/main.svelte';
  import SettingsSubMenu from '../_partials/SettingsSubMenu.svelte';

  let user = /** @type {User} */ ({/* user */});
  let segments = /** @type {Access} */ ({/* segments */});

  const roomsJsDownloadUrl = '/admin/settings/download/rooms.js';
</script>

<LayoutMain access={segments} user={user}>
  <div class="setting-container">
    <SettingsSubMenu />

    <section class="panel">
      <div class="panel-copy">
        <h2>Database Downloads</h2>
        <p>Download the current frontend [rooms.js] source for review or manual updates.</p>
      </div>

      <a class="download-button" href={roomsJsDownloadUrl}>
        Download rooms.js
      </a>
    </section>
  </div>
</LayoutMain>

<style>
  .setting-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .panel {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 16px;
    padding: 20px;
    border: 1px solid var(--gray-003);
    border-radius: 12px;
    background: var(--gray-001);
  }

  .panel-copy {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .panel-copy h2,
  .panel-copy p {
    margin: 0;
  }

  .download-button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    min-width: 180px;
    padding: 12px 16px;
    border-radius: 10px;
    text-decoration: none;
    color: white;
    background: var(--blue-005);
    font-weight: 600;
  }

  .download-button:hover {
    background: var(--blue-006);
  }

  @media only screen and (max-width: 768px) {
    .panel {
      flex-direction: column;
      align-items: flex-start;
    }

    .download-button {
      width: 100%;
    }
  }
</style>
