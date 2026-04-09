<script>
  /** @typedef {import('../_types/masters.js').Access} Access */
  /** @typedef {import('../_types/users.js').User} User */
  /**
   * @typedef {Object} BackupSummary
   * @property {string} code
   * @property {string} date
   * @property {string[]} tables
   * @property {string} size
   */

  import LayoutMain from '../_layouts/main.svelte';
  import SettingsSubMenu from '../_partials/SettingsSubMenu.svelte';
  import { AdminSettings } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier';

  let user = /** @type {User} */ ({/* user */});
  let segments = /** @type {Access} */ ({/* segments */});
  let backupSummaries = /** @type {BackupSummary[]} */ ([/* backupSummaries */]);
  let isBackingUp = false;

  const roomsJsDownloadUrl = '/admin/settings/download/rooms.js';

  async function backupDatabase() {
    if (isBackingUp) return;

    isBackingUp = true;
    await AdminSettings(
      {
        cmd: 'backupDatabase',
      },
      /** @returns {Promise<void>} */
      async function (o) {
        isBackingUp = false;
        if (o.error) {
          notifier.showError(o.error);
          return;
        }

        notifier.showSuccess(o.message || 'Backup completed');
        window.location.reload();
      }
    );
  }
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

    <section class="panel">
      <div class="panel-copy">
        <h2>Database Backup</h2>
        <p>Create a fresh backup in the server backup directory using the current admin connection.</p>
      </div>

      <button class="download-button" on:click={backupDatabase} disabled={isBackingUp}>
        {#if isBackingUp}
          Backing up...
        {:else}
          Backup database
        {/if}
      </button>
    </section>

    <section class="panel panel-column">
      <div class="panel-copy">
        <h2>Existing Backups</h2>
        <p>Rendered directly from the local `backup` directory when the page is loaded.</p>
      </div>

      <div class="table-wrap">
        <table class="backup-table">
          <thead>
            <tr>
              <th>Code</th>
              <th>Date</th>
              <th>Tables</th>
              <th>Size</th>
            </tr>
          </thead>
          <tbody>
            {#if (backupSummaries || []).length > 0}
              {#each backupSummaries as backup}
                <tr>
                  <td>{backup.code}</td>
                  <td>{backup.date}</td>
                  <td>{(backup.tables || []).join(', ')}</td>
                  <td>{backup.size}</td>
                </tr>
              {/each}
            {:else}
              <tr>
                <td colspan="4" class="empty-state">No backups found.</td>
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
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

  .panel-column {
    flex-direction: column;
    align-items: stretch;
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
    border: 0;
    border-radius: 10px;
    text-decoration: none;
    color: white;
    background: var(--blue-005);
    font-weight: 600;
    cursor: pointer;
  }

  .download-button:hover {
    background: var(--blue-006);
  }

  .download-button:disabled {
    background: var(--gray-005);
    cursor: wait;
  }

  .table-wrap {
    width: 100%;
    overflow-x: auto;
  }

  .backup-table {
    width: 100%;
    border-collapse: collapse;
    background: white;
    border-radius: 10px;
    overflow: hidden;
  }

  .backup-table th,
  .backup-table td {
    padding: 12px 14px;
    text-align: left;
    border-bottom: 1px solid var(--gray-003);
    vertical-align: top;
  }

  .backup-table th {
    background: var(--gray-002);
  }

  .backup-table tbody tr:last-child td {
    border-bottom: 0;
  }

  .empty-state {
    color: var(--gray-007);
    text-align: center;
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
