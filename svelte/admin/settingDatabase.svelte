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
  $: latestBackupCode = backupSummaries?.[0]?.code || '20250508002713';

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
        <p>Download the current `rooms.js` source used by `benalu.dev` for review or manual updates.</p>
        <p><strong>CLI:</strong> <code>cd /home/kyz/go/src/kostjc && go run main.go gen-rooms.js ../benalu.dev/src/</code></p>
      </div>

      <a class="download-button" href={roomsJsDownloadUrl}>
        Download rooms.js for benalu.dev
      </a>
    </section>

    <section class="panel">
      <div class="panel-copy">
        <p>Create a fresh backup in the server backup directory using the current admin connection.</p>
        <p><strong>CLI backup:</strong> <code>cd /home/kyz/go/src/kostjc && go run main.go backup</code></p>
        <p><strong>CLI download:</strong> <code>make download-backup</code></p>
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
        <p><strong>CLI restore:</strong> <code>cd /home/kyz/go/src/kostjc && go run main.go restore {latestBackupCode}</code></p>
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
    color: #111111 !important;
  }

  .panel-copy h2,
  .panel-copy p {
    margin: 0;
    color: #111111 !important;
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

  .panel-copy code {
    padding: 2px 6px;
    border-radius: 6px;
    background: var(--gray-002);
    color: #111111;
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
