const assert = require('node:assert/strict');
const test = require('node:test');

const { compactIDRFromThousands, compactPaymentIDRFromThousands, dateAsUtc, paymentTiming } = require('./billingHistory');

test('billing history formats stored thousands as real IDR amounts', () => {
  assert.equal(compactIDRFromThousands(1500), '1.5M');
  assert.equal(compactIDRFromThousands(500), '500K');
  assert.equal(compactPaymentIDRFromThousands(1500), '+1.5M');
});

test('payment timing reports offsets from the booking start date', () => {
  assert.equal(paymentTiming('2026-07-10', '2026-07-10'), 'ontime');
  assert.equal(paymentTiming('2026-07-11', '2026-07-10'), '+1d late');
  assert.equal(paymentTiming('2026-07-15', '2026-07-10'), '+5d late');
  assert.equal(paymentTiming('2026-07-08', '2026-07-10'), '-2d early');
});

test('payment timing accepts timestamps and rejects invalid dates', () => {
  assert.equal(paymentTiming('2026-08-03T18:25:00Z', '2026-08-01'), '+2d late');
  assert.equal(paymentTiming('not-a-date', '2026-08-01'), 'Payment timing unavailable');
  assert.equal(dateAsUtc('2026-02-30'), null);
});
