<script>
  /** @typedef {import('../../_types/users').TenantNearbyBirthday} TenantNearbyBirthday */

  const tenantsNearbyBirthdays = /** @type {TenantNearbyBirthday[]} */ ([/* tenantsNearbyBirthdays*/]);

  function formatDateLong(/** @type {string} */ dateStr) {
    const dt = new Date(dateStr);
    return dt.toLocaleDateString('en-GB', {
      weekday: 'long',
      day: '2-digit',
      month: 'long',
      year: 'numeric'
    });
  }

  function isDateNow(/** @type {string} */ dateStr) {
    const inputDate = new Date(dateStr);
    const today = new Date();

    const inputMMDD = `${String(inputDate.getMonth() + 1).padStart(2, '0')}-${String(inputDate.getDate()).padStart(2, '0')}`;
    const todayMMDD = `${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`;

    return inputMMDD === todayMMDD;
  }
</script>

<section class="birthday-agenda">
  <h1>Birthday Agenda</h1>
  {#if tenantsNearbyBirthdays && tenantsNearbyBirthdays.length > 0}
    <div class="cards">
      {#each (tenantsNearbyBirthdays || []) as tnt}
        <div class="card {isDateNow(tnt.birthDay) ? 'today' : ''}">
          <h3>{formatDateLong(tnt.birthDay)}</h3>
          <div class="desc">
            <span class="tenant-name">{tnt.tenantName}</span>
            <span class="room">Room {tnt.roomName}</span>
          </div>
          <img
            src="/assets/icons/birthday-cake.svg"
            alt=""
            class="icon-cake"
          />
        </div>
      {/each}
    </div>
  {:else}
    <div class="no-data">
      <p>No birthdays</p>
    </div>
  {/if}
</section>

<style>
  .birthday-agenda {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .birthday-agenda h1 {
    margin: 0;
    padding: 0;
    font-size: var(--font-xl);
  }

  .birthday-agenda .cards {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
  }

  .birthday-agenda .cards .card {
    background-color: var(--gray-001);
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    position: relative;
    overflow: hidden;
  }

  .birthday-agenda .cards .card h3 {
    margin: 0;
    padding: 0;
    font-size: var(--font-lg);
    z-index: 20;
  }

  .birthday-agenda .cards .card .desc {
    padding-top: 10px;
    border-top: 1px solid var(--gray-004);
    display: flex;
    flex-direction: column;
    gap: 5px;
    z-index: 20;
  }

  .birthday-agenda .cards .card.today {
    color: var(--blue-007);
    background-color: var(--blue-transparent);
  }

  .birthday-agenda .cards .card.today .desc {
    border-top: 1px solid var(--blue-004);
  }

  .birthday-agenda .cards .card .icon-cake {
    display: none;
  }
  .birthday-agenda .cards .card.today .icon-cake {
    display: block;
    position: absolute;
    bottom: 10px;
    right: 15px;
    width: 55px;
    height: auto;
    fill: var(--gray-003);
    z-index: 10;
  }
</style>