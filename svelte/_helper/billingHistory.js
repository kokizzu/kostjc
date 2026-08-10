// @ts-check

const millisecondsPerDay = 1000 * 60 * 60 * 24;

/**
 * Billing values are stored in thousands of IDR. Convert to the real IDR
 * amount before applying compact K/M labels.
 *
 * @param {number} value
 * @returns {string}
 */
function compactIDRFromThousands(value) {
  const number = Number(value || 0) * 1_000;
  const sign = number < 0 ? '-' : '';
  const abs = Math.abs(number);

  if (abs >= 1_000_000) {
    return `${sign}${trimCompactNumber(abs / 1_000_000)}M`;
  }
  if (abs >= 1_000) {
    return `${sign}${trimCompactNumber(abs / 1_000)}K`;
  }

  return `${sign}${abs}`;
}

/**
 * @param {number} value
 * @returns {string}
 */
function compactPaymentIDRFromThousands(value) {
  const number = Number(value || 0);
  return number < 0 ? compactIDRFromThousands(number) : `+${compactIDRFromThousands(number)}`;
}

/**
 * @param {number} value
 * @returns {string}
 */
function trimCompactNumber(value) {
  return value.toFixed(value >= 10 || Number.isInteger(value) ? 0 : 1).replace(/\.0$/, '');
}

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
  compactIDRFromThousands,
  compactPaymentIDRFromThousands,
  dateAsUtc,
  paymentTiming,
};
