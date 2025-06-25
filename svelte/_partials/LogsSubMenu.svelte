<script>
	let segment2 = /** @type {string} */ (window.location.pathname.split('/').pop() || '');
</script>

<div class="admin-submenu">
	<nav>
		<a href="/admin/bookingLogs" class:active={segment2==='bookingLogs'}>
			<span class="title">Booking</span>
			<span class="block"></span>
		</a>
		<a href="/admin/buildingLogs" class:active={segment2==='buildingLogs'}>
			<span class="title">Building</span>
			<span class="block"></span>
		</a>
		<a href="/admin/facilityLogs" class:active={segment2==='facilityLogs'}>
			<span class="title">Facility</span>
			<span class="block"></span>
		</a>
		<a href="/admin/locationLogs" class:active={segment2==='locationLogs'}>
			<span class="title">Location</span>
			<span class="block"></span>
		</a>
		<a href="/admin/paymentLogs" class:active={segment2==='paymentLogs'}>
			<span class="title">Payment</span>
			<span class="block"></span>
		</a>
		<a href="/admin/roomLogs" class:active={segment2==='roomLogs'}>
			<span class="title">Room</span>
			<span class="block"></span>
		</a>
		<a href="/admin/stockLogs" class:active={segment2==='stockLogs'}>
			<span class="title">Stock</span>
			<span class="block"></span>
		</a>
		<a href="/admin/saleLogs" class:active={segment2==='saleLogs'}>
			<span class="title">Sale</span>
			<span class="block"></span>
		</a>
		<a href="/admin/menuLogs" class:active={segment2==='menuLogs'}>
			<span class="title">Menu</span>
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
