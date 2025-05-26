<script>
    import { AdminTenants } from '../jsApi.GEN';
  /** @typedef {import('../_types/users').Tenant} Tenant */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './xFormatter';

  let isShow = /** @type {boolean} */ (false);

  export let isSubmitted  = /** @type {boolean} */ (false);
  
  let tenantId          = /** @type {number} */ (0);
  let tenantName        = /** @type {string} */ ('');
  let ktpRegion         = /** @type {string} */ ('');
  let ktpNumber         = /** @type {string} */ ('');
  let ktpName           = /** @type {string} */ ('');
  let ktpPlaceBirth     = /** @type {string} */ ('');
  let ktpDateBirth      = /** @type {string} */ ('');
  let ktpGender         = /** @type {string} */ ('');
  let ktpAddress        = /** @type {string} */ ('');
  let ktpRtRw           = /** @type {string} */ ('');
  let ktpKelurahanDesa  = /** @type {string} */ ('');
  let ktpKecamatan      = /** @type {string} */ ('');
  let ktpReligion       = /** @type {string} */ ('');
  let ktpMaritalStatus  = /** @type {string} */ ('');
  let ktpCitizenship    = /** @type {string} */ ('');
  let ktpOccupation     = /** @type {string} */ ('');
  let telegramUsername  = /** @type {string} */ ('');
  let whatsappNumber    = /** @type {string} */ ('');

  export let OnSubmit = async function(/** @type {Tenant} */ tenant) {
    console.log('OnSubmit :::', tenant);
  }

  async function submitAdd() {
    const tenant = /** @type {Tenant} */ ({
      id: tenantId,
      tenantName: tenantName,
      ktpRegion: ktpRegion,
      ktpNumber: ktpNumber,
      ktpName: ktpName,
      ktpPlaceBirth: ktpPlaceBirth,
      ktpDateBirth: ktpDateBirth,
      ktpGender: ktpGender,
      ktpAddress: ktpAddress,
      ktpRtRw: ktpRtRw,
      ktpKelurahanDesa: ktpKelurahanDesa,
      ktpKecamatan: ktpKecamatan,
      ktpReligion: ktpReligion,
      ktpMaritalStatus: ktpMaritalStatus,
      ktpCitizenship: ktpCitizenship,
      ktpOccupation: ktpOccupation,
      telegramUsername: telegramUsername,
      whatsappNumber: whatsappNumber
    });

    await OnSubmit(tenant);
  }

  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {
  }
  const cancel = () => isShow = false;

  /**
   * @description 
   * @param {number} id
   * @returns {Promise<void>}
   */
  export async function GetTenantById(id) {
    await AdminTenants({
      cmd: 'form', // @ts-ignore
      tenant: {
        id: id+''
      }
    }, function (/** @type {any} */ o) {
      if (o.error) {
        return;
      }

      const tenant = /** @type {Tenant} */ (o.tenant);

      tenantId = Number(tenant.id);
      tenantName = tenant.tenantName;
      ktpRegion = tenant.ktpRegion;
      ktpNumber = tenant.ktpNumber;
      ktpName = tenant.ktpName;
      ktpPlaceBirth = tenant.ktpPlaceBirth;
      ktpDateBirth = tenant.ktpDateBirth;
      ktpGender = tenant.ktpGender;
      ktpRtRw = tenant.ktpRtRw;
      ktpKelurahanDesa = tenant.ktpKelurahanDesa;
      ktpKecamatan = tenant.ktpKecamatan;
      ktpReligion = tenant.ktpReligion;
      ktpCitizenship = tenant.ktpCitizenship;
      ktpOccupation = tenant.ktpOccupation;
      telegramUsername = tenant.telegramUsername;
      whatsappNumber = tenant.whatsappNumber;
    })
  }

  const KtpGenders = [
    'Laki - Laki',
    'Perempuan'
  ];
  const KtpMaritalStatus = [
    'Belum Kawin',
    'Kawin',
    'Cerai Hidup',
    'Cerai Mati'
  ];
  const KtpReligions = [
    'Islam',
    'Hindu',
    'Kristen',
    'Katholik',
    'Buddha',
    'Konghucu'
  ];
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>Edit Tenant #{tenantId}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        id="tenantName"
        label="Nama Penyewa"
        type="text"
        bind:value={tenantName}
        placeholder="I Komang John Doe"
      />
      <InputBox
        id="ktpRegion"
        label="Regional"
        type="text"
        bind:value={ktpRegion}
        placeholder="I Komang John Doe"
      />
      <InputBox
        id="ktpNumber"
        label="Nomor KTP/NIK"
        type="text"
        bind:value={ktpNumber}
        placeholder="1234567890123456"
      />
      <InputBox
        id="ktpName"
        label="Nama Lengkap (sesuai KTP)"
        type="text"
        bind:value={ktpName}
        placeholder="I Komang John Doe"
      />
      <InputBox
        id="ktpPlaceBirth"
        label="Tempat Lahir"
        type="text"
        bind:value={ktpPlaceBirth}
        placeholder="Denpasar"
      />
      <InputBox
        id="ktpDateBirth"
        label="Tanggal Lahir"
        type="date"
        bind:value={ktpDateBirth}
      />
      <InputBox
        id="ktpGender"
        label="Jenis Kelamin"
        type="combobox-arr"
        bind:value={ktpGender}
        values={KtpGenders}
      />
      <InputBox
        id="ktpAddress"
        label="Alamat"
        type="text"
        placeholder="JL PENDIDIKAN NO 8 BLOK A NO 47"
        bind:value={ktpAddress}
      />
      <InputBox
        id="ktpRtRw"
        label="RT/RW"
        type="text"
        placeholder="001/004"
        bind:value={ktpRtRw}
      />
      <InputBox
        id="ktpKelurahanDesa"
        label="Kelurahan/Desa"
        type="text"
        placeholder="Dalung"
        bind:value={ktpKelurahanDesa}
      />
      <InputBox
        id="ktpKecamatan"
        label="Kecamatan"
        type="text"
        placeholder="Kuta Utara"
        bind:value={ktpKecamatan}
      />
      <InputBox
        id="ktpReligion"
        label="Agama"
        type="combobox-arr"
        bind:value={ktpReligion}
        values={KtpReligions}
      />
      <InputBox
        id="ktpMaritalStatus"
        label="Status Perkawinan"
        type="combobox-arr"
        bind:value={ktpMaritalStatus}
        values={KtpMaritalStatus}
      />
      <InputBox
        id="ktpOccupation"
        label="Pekerjaan"
        type="text"
        placeholder="Pelajar/Mahasiswa"
        bind:value={ktpOccupation}
      />
      <InputBox
        id="ktpCitizenship"
        label="Kewarganegaraan"
        type="text"
        placeholder="WNI"
        bind:value={ktpCitizenship}
      />
      <InputBox
        id="telegramUsername"
        label="Username Telegram"
        type="text"
        placeholder="@johndoe"
        bind:value={telegramUsername}
      />
      <InputBox
        id="whatsappNumber"
        label="Nomor Whatsapp"
        type="text"
        placeholder="+6281234567890"
        bind:value={whatsappNumber}
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
  }
</style>