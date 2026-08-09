// @ts-check

const millisecondsPerDay = 1000 * 60 * 60 * 24;

/**
 * @param {string} value
 * @returns {number | null}
 */
function dateAsUtc(value) {
  const [year, month, day] = (value || '').slice(0, 10).split('-').map(Number);
  if (!year || !month || !day) return null;

  const result = Date.UTC(year, month - 1, day);
  const parsed = new Date(result);
  if (parsed.getUTCFullYear() !== year || parsed.getUTCMonth() !== month - 1 || parsed.getUTCDate() !== day) {
    return null;
  }

  return result;
}

/**
 * @param {string} paymentAt
 * @param {string} bookingStart
 * @returns {string}
 */
function paymentTiming(paymentAt, bookingStart) {
  const paidDate = dateAsUtc(paymentAt);
  const startDate = dateAsUtc(bookingStart);
  if (paidDate === null || startDate === null) return 'Payment timing unavailable';

  const days = Math.round((paidDate - startDate) / millisecondsPerDay);
  if (days === 0) return 'ontime';

  const count = Math.abs(days);
  return days > 0
    ? `+${count}d late`
    : `-${count}d early`;
}

module.exports = {
  dateAsUtc,
  paymentTiming,
};
