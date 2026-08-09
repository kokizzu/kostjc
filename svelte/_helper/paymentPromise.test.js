const assert = require('node:assert/strict');
const test = require('node:test');

const {
  compactDate,
  dateAsUtc,
  paymentPromisePayload,
  promiseDayLabel,
  promiseStatus,
  updatePaymentPromises,
} = require('./paymentPromise');

const today = new Date(2026, 6, 25, 12);

test('payment promise dates use compact labels', () => {
  assert.equal(compactDate('2026-07-10'), '10 Jul');
  assert.equal(compactDate('not-a-date'), '');
  assert.equal(dateAsUtc('2026-02-30'), null);
});

test('payment promise deadline reports days left and overdue', () => {
  assert.equal(promiseDayLabel('2026-07-27', today), '2d left');
  assert.equal(promiseDayLabel('2026-07-25', today), 'today');
  assert.equal(promiseDayLabel('2026-07-22', today), '3d over');
  assert.equal(promiseStatus('2026-07-27', today), 'left');
  assert.equal(promiseStatus('2026-07-25', today), 'left');
  assert.equal(promiseStatus('2026-07-22', today), 'over');
});

test('payment promise payload string-encodes Go json string-tag IDs', () => {
  const payload = paymentPromisePayload(
    { promisedAt: '2026-07-25', promisePayAt: '2026-07-29', note: 'payday' },
    123,
    456,
  );

  assert.equal(payload.tenantId, '123');
  assert.equal(payload.bookingId, '456');
  assert.equal(payload.promisedAt, '2026-07-25');
  assert.equal(payload.promisePayAt, '2026-07-29');
  assert.equal(payload.note, 'payday');
});

test('AJAX response updates only the affected Home payment promise', () => {
  const existing = {
    10: { tenantId: '10', promisePayAt: '2026-07-26' },
  };
  const saved = {
    id: '88',
    tenantId: '20',
    promisedAt: '2026-07-25',
    promisePayAt: '2026-07-29',
    note: 'payday',
  };

  const updated = updatePaymentPromises(existing, saved);

  assert.notEqual(updated, existing);
  assert.deepEqual(updated[10], existing[10]);
  assert.equal(updated[20], saved);
  assert.equal(existing[20], undefined);
});
