// @ts-nocheck

function datetime(unixSec, humanize) {
  if (!unixSec) return '';
  if (typeof unixSec === 'string') return unixSec; // might not be unix time

  const dt = new Date(unixSec * 1000);

  if (!humanize) {
    const year = dt.getFullYear();
    const month = String(dt.getMonth() + 1).padStart(2, '0');
    const day = String(dt.getDate()).padStart(2, '0');
    const hours = String(dt.getHours()).padStart(2, '0');
    const minutes = String(dt.getMinutes()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}`;
  }

  const options = { day: '2-digit', month: 'long', year: 'numeric' };
  return dt.toLocaleDateString(undefined, options);
}

function localeDatetime(unixSec) {
  if (!unixSec) return '';
  const dt = new Date(unixSec * 1000);
  const day = dt.toLocaleDateString('default', { weekday: 'long' });
  const date = dt.getDate();
  const month = dt.toLocaleDateString('default', { month: 'short' });
  const year = dt.getFullYear();
  let hh = dt.getHours();
  if (hh < 10) hh = '0' + hh;
  let mm = dt.getMinutes();
  if (mm < 10) mm = '0' + mm;
  const formattedDate = `${date} ${month} ${year} - ${hh}:${mm}`;
  return formattedDate;
}

function utcDatetime(unixSec) {
  if (!unixSec) return '';
  let dt = new Date(unixSec * 1000);
  if (typeof unixSec !== 'number') dt = new Date(unixSec);

  const year = dt.toLocaleDateString('en-US', { year: '2-digit', timeZone: 'UTC' });
  const month = dt.toLocaleDateString('en-US', { month: '2-digit', timeZone: 'UTC' });
  const date = dt.getUTCDate();
  let hh = dt.getUTCHours();
  if (hh < 10) hh = '0' + hh;
  let mm = dt.getUTCMinutes();
  if (mm < 10) mm = '0' + mm;

  const formattedDate = `${date}/${month}/${year} - ${hh}:${mm}`;
  return formattedDate;
}

function isoTime(unixSec) {
  if (!unixSec) return '';
  const dt = new Date(unixSec * 1000);
  const day = dt.toLocaleDateString('default', { weekday: 'long' });
  const date = dt.getDate();
  const month = dt.toLocaleDateString('default', { month: '2-digit' });
  const year = dt.getFullYear();
  let hh = dt.getHours();
  if (hh < 10) hh = '0' + hh;
  let mm = dt.getMinutes();
  if (mm < 10) mm = '0' + mm;
  const formattedDate = `${year}-${month}-${date} ${hh}:${mm}`;
  return formattedDate;
}

function isoDate(unixSec) {
  if (!unixSec) return '';
  const dt = new Date(unixSec * 1000);
  const day = dt.toLocaleDateString('default', { weekday: 'long' });
  const date = dt.getDate();
  const month = dt.toLocaleDateString('default', { month: '2-digit' });
  const year = dt.getFullYear();
  const formattedDate = `${year}-${month}-${date}`;
  return formattedDate;
}

function dateISOFormat(/** @type number */ dayTo = 0) {
  const dt = new Date();
  dt.setDate(dt.getDate() + dayTo);

  const date = String(dt.getDate()).padStart(2, '0');
  const month = String(dt.getMonth() + 1).padStart(2, '0');
  const year = dt.getFullYear();

  return `${year}-${month}-${date}`;
}

function dateISOFormatFromYYYYMMDD(/** @type {string} */ dateStr, /** @type {number} */ dayTo = 0) {
  const dt = new Date(dateStr);
  dt.setDate(dt.getDate() + dayTo);

  const date = String(dt.getDate()).padStart(2, '0');
  const month = String(dt.getMonth() + 1).padStart(2, '0');
  const year = dt.getFullYear();

  return `${year}-${month}-${date}`;
}

function loadScript(/** @type {string} */ url, /** @type {Function} */ callback) {
  let script = /** @type {HTMLScriptElement} */ (document.createElement('script'));
  script.type = 'text/javascript';
  script.src = url;
  script.onload = () => callback();
  document.getElementsByTagName('head')[0].appendChild(script);
}

/**
 * @description Format price
 * @param {number | bigint | Intl.StringNumericLiteral} price 
 * @param {string} currency 
 * @returns {string}
 */
function formatPrice(price, currency) {
  try {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(price);
  } catch(err) {
    console.log('formatPrice failed', err, price, currency);
    return `${currency} ${price}`
  }
}

/**
 * @description Convert array of any to array of numbers
 * @param {any[]} arr 
 * @returns {number[]}
 */
function arrToArrNum(arr) {
  return arr.map(Number);
}

module.exports = {
  datetime: datetime,
  localeDatetime: localeDatetime,
  utcDatetime: utcDatetime,
  isoTime: isoTime,
  isoDate: isoDate,
  dateISOFormat: dateISOFormat,
  dateISOFormatFromYYYYMMDD: dateISOFormatFromYYYYMMDD,
  loadScript: loadScript,
  formatPrice: formatPrice,
  arrToArrNum: arrToArrNum
};
