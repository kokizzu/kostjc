<script>
	/** @typedef {import('./_types/users').User} User */
	/** @typedef {import('./_types/users').Session} Session */
	/** @typedef {import('./_types/masters').Access} Access */

	import LayoutMain from './_layouts/main.svelte';
	import { Icon } from './node_modules/svelte-icons-pack/dist';
	import { datetime } from './_components/xFormatter.js';
	import { onMount } from 'svelte';
	import { UserSessionKill, UserUpdateProfile} from './jsApi.GEN.js';
	import { RiSystemDeleteBinLine } from './node_modules/svelte-icons-pack/dist/ri';
	import { notifier } from './_components/xNotifier.js';
	import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';

	let user 						= /** @type {User} */ ({/* user */});
	let segments  = /** @type {Access} */ ({/* segments */});
	let activeSessions	= /** @type {Session[]} */ ({/* activeSessions */});

	let oldPassword			= /** @type {string} */ ('');
	let newPassword 		= /** @type {string} */ ('');
	let confNewPassword	= /** @type {string} */ ('');
	let oldProfileJson 	= /** @type {string} */ ('');

	onMount(() => {
		oldProfileJson = JSON.stringify(user);
	});

	let username 	= /** @type {string} */ (user.userName || '');
	let fullName	= /** @type {string} */ (user.fullName || '');
	let email 		= /** @type {string} */ (user.email || '');

	let isSubmitProfile		= /** @type {boolean} */ (false);
	async function updateProfile() {
		user.userName = username;
		user.fullName = fullName;
		user.email 		= email;

		if (JSON.stringify(user) === oldProfileJson) {
			notifier.showWarning('No changes');
			return
		}

		isSubmitProfile = true;

		// @ts-ignore
		await UserUpdateProfile(user, async function(/** @type {any} */ res) {
			isSubmitProfile = false;
			if (res.error) {
				notifier.showError(res.error || 'Failed to update profile');
				return
			}

			oldProfileJson = JSON.stringify(res.user);
			user = res.user;

			notifier.showSuccess('Profile updated');
		});
	}

	let isSubmitPassword	= /** @type {boolean} */ (false);
	async function changePassword() {
		if (newPassword !== confNewPassword) {
			notifier.showError('New password and repeat new password must be same');
			return
		}
		isSubmitPassword = true;

		// @ts-ignore
		await UserChangePassword({
			oldPass: oldPassword,
			newPass: newPassword
		}, async function(/** @type {any} */ res) {
			isSubmitPassword = false;
			if (res.error) {
				notifier.showError(res.error || 'Failed to update password');
				return
			}
			oldPassword 		= '';
			newPassword 		= '';
			confNewPassword = '';

			notifier.showSuccess('Password updated');
		});
	}

	async function killSession(/** @type {string} */ sessionToken) {
		await UserSessionKill({
			sessionTokenHash: sessionToken
		}, async function (/** @type {any} */ res) {
			if (res.error) {
				notifier.showError(res.error);
				return
			}

			if (res.sessionTerminated < 1) {
				notifier.showWarning('No session terminated');
				return
			}

			notifier.showInfo(Number(res.sessionTerminated || 0) + ' session terminated');

			// @ts-ignore
			await UserSessionsActive({
				userId: user.id
			}, async function(/** @type {any} */ res) {
				if (res.error) {
					notifier.showError(res.error);
					return
				}
				
				activeSessions = res.sessionsActive;
			});
		});
	}
</script>

<LayoutMain access={segments} user={user}>
	<div class="user-detail-container">
		<div class="user-profile">
			<div class="profile">
				<h2>Profile</h2>
				<div class="profile-form">
					<div class="row">
						<InputBox
							className="input-box-container"
							label="Username"
							id="username"
							type="text"
							bind:value={username}
						/>
						<InputBox
							className="input-box-container"
							label="Full Name"
							id="fullname"
							type="text"
							bind:value={fullName}
						/>
					</div>
					<div class="row">
						<InputBox
							className="input-box-container"
							label="Email"
							id="email"
							type="email"
							bind:value={email}
						/>
						<div class="pill-box">Role: {user.role}</div>
					</div>
					<div class="row">
						<div class="pill-box">Registered At: {
							(user.createdAt && user.createdAt > 0)
							? datetime(user.createdAt) : '0'
						}</div>
						<div class="pill-box">Last Login: {
							(user.lastLoginAt && user.lastLoginAt > 0)
							? datetime(user.lastLoginAt) : '0'
						}</div>
					</div>
					<SubmitButton
						on:click={updateProfile}
						isSubmitted={isSubmitProfile}
						label="Update Profile"
					/>
				</div>
			</div>
			<div class="password">
				<h2>Change Password</h2>
				<InputBox
					label="Old Password"
					id="old-password"
					type="password"
					bind:value={oldPassword}
				/>
				<InputBox
					label="New Password"
					id="new-password"
					type="password"
					bind:value={newPassword}
				/>
				<InputBox
					label="Confirm New Password"
					id="conf-new-password"
					type="password"
					bind:value={confNewPassword}
				/>
				<div class="submit-btn">
					<SubmitButton
						on:click={changePassword}
						isSubmitted={isSubmitPassword}
						label="Update Password"
					/>
				</div>
			</div>
		</div>
		<div class="sessions">
			<h2>Sessions</h2>
			<div class="table-root">
				<div class="table-container">
					<table>
						<thead>
							<tr>
								<th class="no">No</th>
								<th class="a-row">Action</th>
								<th>IP Address</th>
								<th class="datetime">Expired At</th>
								<th style="min-width: 200px;">User Agent</th>
							</tr>
						</thead>
						<tbody>
							{#each (activeSessions || []) as session, i}
								<tr>
									<td class="num-row">{i + 1}</td>
									<td>
										<div class="actions">
											<button class="btn" title="Kill Session"
												on:click={() => killSession(session.sessionToken)}>
												<Icon
													size="17"
													src={RiSystemDeleteBinLine}
													color="var(--gray-008)"
												/>
											</button>
										</div>
									</td>
									<td>{session.loginIPs || '0'}</td>
									<td class="datetime">{
										(session.expiredAt && session.expiredAt > 0)
										? datetime(session.expiredAt) : '0'
									}</td>
									<td>{session.device || '--'}</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</LayoutMain>

<style>
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
	
	.user-detail-container {
		display: flex;
		flex-direction: column;
		gap: 20px;
		padding: 20px;
	}

	.user-detail-container .user-profile {
		width: 100%;
		display: grid;
		grid-template-columns: 65% auto;
		gap: 20px;
		border: 1px solid var(--gray-002);
		border-radius: 8px;
		padding: 20px;
	}

	.user-detail-container .user-profile .profile {
		display: flex;
		flex-direction: column;
		height: fit-content;
		gap: 20px;
	}

	.user-detail-container .user-profile .profile h2 {
		margin: 0;;
	}

	.user-detail-container .user-profile .profile .profile-form {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}

	.user-detail-container .user-profile .profile .profile-form .row {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 10px;
		align-items: end;
	}

	.user-detail-container .user-profile .profile .profile-form .row .pill-box {
		display: flex;
		align-items: center;
		height: fit-content;
		padding: 12px;
		border: 1px solid var(--gray-002);
		border-radius: 8px;
		background-color: var(--gray-001);
	}

	.user-detail-container .user-profile .password {
		display: flex;
		flex-direction: column;
		padding: 20px;
		border-radius: 8px;
		background-color: var(--blue-transparent);
		height: fit-content;
		gap: 10px;
	}

	.user-detail-container .user-profile .password h2 {
		color: var(--blue-006);
		margin: 0 0 10px 0;
	}

	:global(.user-detail-container .user-profile .password .label) {
		color: var(--blue-006);
	}

	:global(.user-detail-container .user-profile .password input) {
		border: 1px solid var(--blue-005) !important;
		background-color: #FFF !important;
	}

	.user-detail-container .user-profile .password .submit-btn {
		display: flex;
		flex-direction: row;
		justify-content: flex-end;
		margin-top: 5px;
	}

	.table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    padding: 0;
    overflow: hidden;
		width: 100%;
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
		font-size: 13px;
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table-container table thead tr th.datetime {
    min-width: 140px !important;
  }

  .table-root .table-container table thead tr th.no {
    width: 30px;
  }

  .table-root .table-container table thead tr th.a-row {
    max-width: fit-content;
    min-width: fit-content;
    width: fit-content;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table-container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr:last-child {
		border-bottom: none !important;
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

	.table-root .table-container table tbody tr td.num-row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
	}

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .table-root .table-container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  .table-root .table-container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .table-root .table-container table tbody tr td .actions .btn:hover {
    background-color: var(--red-transparent);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--red-005);
  }
  
	@media only screen and (max-width : 768px) {
		.user-detail-container .user-profile {
			display: flex;
			flex-direction: column;
			width: 100%;
			gap: 20px;
			border: none;
			padding: 0;
		}

		.user-detail-container .user-profile .profile .profile-form {
			gap: 10px;
		}

		.user-detail-container .user-profile .profile .profile-form .row {
			display: flex;
			flex-direction: column;
			gap: 10px;
			width: 100%;
		}

		:global(.user-detail-container .user-profile .profile .profile-form .row .input-box-container),
		.user-detail-container .user-profile .profile .profile-form .row .pill-box {
			width: 100%;
		}
	}
</style>