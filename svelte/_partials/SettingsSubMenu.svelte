<script>
	let segment2 = /** @type {string} */ (window.location.pathname.split('/').pop() || '');
</script>

<div class="admin-submenu">
	<nav>
		<a href="/admin/settingFixInconsistencies" class:active={segment2==='settingFixInconsistencies'}>
			<span class="title">Fix Inconsistencies</span>
			<span class="block"></span>
		</a>
    <a href="/admin/settingDatabase" class:active={segment2==='settingDatabase'}>
			<span class="title">--</span>
			<span class="block"></span>
		</a>
	</nav>
</div>

<style>
	.admin-submenu {
		display: flex;
		flex-direction: row;
		width: 100%;
		height: fit-content;
	}

	.admin-submenu nav {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		width: 100%;
		height: fit-content;
		gap: 10px;
	}

	.admin-submenu nav a {
		height: fit-content;
		display: flex;
		flex-direction: column;
		text-decoration: none;
		color: var(--gray-008);
	}

	.admin-submenu nav a:hover {
		background-color: var(--gray-001);
	}

	.admin-submenu nav a.active {
		color: var(--blue-005);
	}

	.admin-submenu nav a .title {
		flex-grow: 1;
		padding: 15px 10px;
		display: flex;
		align-items: center;
	}

	.admin-submenu nav a .block {
		height: 3px;
		background-color: transparent;
		width: 100%;
	}

	.admin-submenu nav a:hover .block {
		background-color: var(--gray-005);
	}
	.admin-submenu nav a.active .block {
		background-color: var(--blue-005);
	}

	@media only screen and (max-width: 768px) {
	}
</style>
