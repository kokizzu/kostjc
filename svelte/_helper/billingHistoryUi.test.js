const assert = require('node:assert/strict');
const { readFileSync } = require('node:fs');
const { join } = require('node:path');
const test = require('node:test');

const svelteRoot = join(__dirname, '..');
const unpaidRecap = readFileSync(join(svelteRoot, '_partials/home/UnpaidBookingRecap.svelte'), 'utf8');
const unpaidBookings = readFileSync(join(svelteRoot, '_partials/home/UnpaidBookingTenants.svelte'), 'utf8');
const availableRooms = readFileSync(join(svelteRoot, '_partials/home/AvailableRooms.svelte'), 'utf8');
const occupancyReport = readFileSync(join(svelteRoot, 'staffOccupancyReport.svelte'), 'utf8');
const billingPopup = readFileSync(join(svelteRoot, '_components/PopUpShowBookingMissingTenant.svelte'), 'utf8');

test('unpaid recap puts the billing-history action in the leftmost column', () => {
  assert.match(unpaidRecap, /<tr>\s*<th class="action-heading">/);
  assert.match(unpaidRecap, /<tr>\s*<td class="action-cell">/);
  assert.match(unpaidRecap, /showBillingHistory\(item\.tenantId, item\.tenantName\)/);
});

test('occupancy report puts billing history before the tenant name', () => {
  assert.match(
    occupancyReport,
    /class="billing-history-button"[\s\S]*showBillingHistory\(booking\.tenantId, booking\.tenantName\)[\s\S]*<span>\{booking\.tenantName\}/,
  );
});

test('occupancy hover actions remain four buttons in a top-right 2x2 grid', () => {
  const actionBlocks = [...occupancyReport.matchAll(/<div class="actions">([\s\S]*?)<\/div>/g)];
  const bookingActions = actionBlocks.find(match => match[1].includes('Extend Booking'));
  assert.ok(bookingActions, 'booking actions block not found');
  assert.equal((bookingActions[1].match(/<button class="btn"/g) || []).length, 4);
  assert.match(occupancyReport, /grid-template-columns: repeat\(2, 27px\)/);
  assert.match(occupancyReport, /position: absolute;\s*right: 3px;\s*top: 3px;/);
  assert.match(occupancyReport, /\.cell:hover \.actions,[\s\S]*\.cell:focus-within \.actions/);
});

test('available room outstanding warning matches by tenant, not by room', () => {
  assert.match(availableRooms, /function hasUnpaidBooking\(room\)/);
  assert.match(availableRooms, /Number\(room\.currentTenantId \|\| 0\)/);
  assert.match(availableRooms, /Number\(item\.tenantId \|\| 0\) === tenantId/);
  assert.match(availableRooms, /\{#if hasUnpaidBooking\(r\)\}/);
  assert.doesNotMatch(availableRooms, /item\.roomName === roomName/);
});

test('available rooms can open billing history for current tenant', () => {
  assert.match(availableRooms, /<PopUpShowBookingMissingTenant/);
  assert.match(availableRooms, /function canShowBillingHistory\(room\)/);
  assert.match(availableRooms, /function showBillingHistory\(room\)/);
  assert.match(availableRooms, /\{ tenantId, cmd: CmdForm, yearMonth: '' \}/);
  assert.match(availableRooms, /on:click=\{\(\) => showBillingHistory\(r\)\}/);
});

test('available rooms show the compact parenthesized availability age', () => {
  assert.match(availableRooms, /return `\(-<b>\$\{diffDays\}<\/b>d ago\)`/);
  assert.match(availableRooms, /`Available Now \$\{getAvailableDays\(r\.availableAt\)\}`/);
});

test('payment promises are tenant-linked and shown at the bottom of both card types', () => {
  for (const source of [unpaidBookings, availableRooms]) {
    assert.match(source, /paymentPromisePayload/);
    assert.match(source, /<PopUpAddPaymentPromise/);
    assert.match(source, /StaffPaymentPromise/);
    assert.match(source, /class="promise-area"/);
    assert.match(source, /class="promise-add"/);
    assert.match(source, /promiseDayLabel\(promise\.promisePayAt\)/);
    assert.match(source, /promiseStatus\(promise\.promisePayAt\)/);
    assert.match(source, /\{#if promise\.note\}/);
    assert.match(source, /\.promise-chip\.left/);
    assert.match(source, /\.promise-chip\.over/);
  }
  assert.match(unpaidBookings, /const payload = paymentPromisePayload\(promise, paymentPromiseTenantId, paymentPromiseBookingId\)/);
  assert.match(unpaidBookings, /promise: payload/);
  assert.match(availableRooms, /paymentPromiseTenantId = tenantId/);
  assert.match(availableRooms, /const payload = paymentPromisePayload\(promise, paymentPromiseTenantId, paymentPromiseBookingId\)/);
  assert.match(availableRooms, /promise: payload/);
});

test('unpaid booking cards use a top-right 2x2 action grid with billing history', () => {
  assert.match(unpaidBookings, /<PopUpShowBookingMissingTenant/);
  assert.match(unpaidBookings, /on:click=\{\(\) => showBillingHistory\(ub\.tenantId, ub\.tenantName\)\}/);
  assert.match(unpaidBookings, /grid-template-columns: repeat\(2, 27px\)/);
  assert.match(unpaidBookings, /grid-template-rows: repeat\(2, 27px\)/);
  assert.match(unpaidBookings, /position: absolute;\s*top: 8px;\s*right: 8px;/);
  const actionBlocks = [...unpaidBookings.matchAll(/<div class="actions">([\s\S]*?)<\/div>/g)];
  const cardActions = actionBlocks.find(match => match[1].includes('Input Payment'));
  assert.ok(cardActions, 'unpaid booking card actions block not found');
  assert.equal((cardActions[1].match(/<button class="btn"/g) || []).length, 4);
});

test('billing popup shows compact two-column booking and payment detail', () => {
  assert.match(billingPopup, /Billing &amp; Payment History/);
  assert.match(billingPopup, /class="header-right"[\s\S]*<p>\{tenantName\}<\/p>/);
  assert.match(billingPopup, /class="booking-summary"/);
  assert.match(billingPopup, /class="payments"/);
  assert.match(billingPopup, /grid-template-columns: minmax\(260px, max-content\) minmax\(360px, max-content\)/);
  assert.match(billingPopup, /Booking #\{booking\.bookingId\} <span>Room <strong>\{booking\.roomName \|\| '—'\}<\/strong><\/span>/);
  assert.match(billingPopup, /<strong>Bill<\/strong> \{compactIDRFromThousands\(booking\.totalPriceIDR\)\}/);
  assert.match(billingPopup, /<strong>Paid<\/strong> \{compactIDRFromThousands\(booking\.totalPaidIDR\)\}/);
  assert.doesNotMatch(billingPopup, /Outstanding/);
  assert.match(billingPopup, /<span class="debt"><strong>Debt<\/strong> \{compactIDRFromThousands\(debtAmount\(booking\)\)\}<\/span>/);
  assert.match(billingPopup, /\{payment\.paymentMethod\}/);
  assert.match(billingPopup, /class="payment-amount"/);
  assert.match(billingPopup, /compactPaymentIDRFromThousands\(payment\.paidIDR\)/);
  assert.match(billingPopup, /\{payment\.note\}/);
  assert.match(billingPopup, /paymentTiming\(payment\.paymentAt, booking\.dateStart\)/);
  assert.match(billingPopup, /\.timing\.late/);
  assert.match(billingPopup, /\.timing\.early/);
  assert.match(billingPopup, /\.timing\.ontime/);
});
