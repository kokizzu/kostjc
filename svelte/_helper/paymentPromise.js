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
 * @param {Date=} now
 * @returns {number}
 */
function todayAsUtc(now = new Date()) {
  return Date.UTC(now.getFullYear(), now.getMonth(), now.getDate());
}

/**
 * @param {string} value
 * @returns {string}
 */
function compactDate(value) {
  const parsed = dateAsUtc(value);
  if (parsed === null) return '';

  const dt = new Date(parsed);
  return dt.toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'short',
    timeZone: 'UTC',
  });
}

/**
 * @param {string} promisePayAt
 * @param {Date=} now
 * @returns {string}
 */
function promiseDayLabel(promisePayAt, now = new Date()) {
  const payAt = dateAsUtc(promisePayAt);
  if (payAt === null) return '';

  const days = Math.round((payAt - todayAsUtc(now)) / millisecondsPerDay);
  if (days < 0) return `${Math.abs(days)}d over`;
  if (days === 0) return 'today';

  return `${days}d left`;
}

/**
 * @param {string} promisePayAt
 * @param {Date=} now
 * @returns {'over' | 'left'}
 */
function promiseStatus(promisePayAt, now = new Date()) {
  const payAt = dateAsUtc(promisePayAt);
  if (payAt === null) return 'left';

  return payAt < todayAsUtc(now) ? 'over' : 'left';
}

/**
 * PaymentPromises uses Go json:",string" ID fields, so browser payloads must
 * encode IDs as strings even when local UI state stores them as numbers.
 *
 * @template {Record<string, any>} T
 * @param {T} promise
 * @param {string | number} tenantId
 * @param {string | number} bookingId
 * @returns {T & { tenantId: string, bookingId: string }}
 */
function paymentPromisePayload(promise, tenantId, bookingId) {
  return {
    ...promise,
    tenantId: String(tenantId || 0),
    bookingId: String(bookingId || 0),
  };
}

/**
 * Merge the promise returned by the AJAX mutation into Home's local snapshot.
 * Returning a new object makes the affected Svelte card update immediately.
 *
 * @template {Record<string, any>} T
 * @param {Record<string | number, T>} paymentPromises
 * @param {T} promise
 * @returns {Record<string | number, T>}
 */
function updatePaymentPromises(paymentPromises, promise) {
  const tenantId = Number(promise?.tenantId || 0);
  if (!tenantId) return paymentPromises;

  return {
    ...(paymentPromises || {}),
    [tenantId]: promise,
  };
}

module.exports = {
  compactDate,
  dateAsUtc,
  paymentPromisePayload,
  promiseDayLabel,
  promiseStatus,
  updatePaymentPromises,
};
