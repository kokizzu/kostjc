/**
 * @typedef { 'button' | 'checkbox' | 'color' | 'date' | 'datetime-local'
* | 'email' | 'file' | 'hidden' | 'image' | 'month' | 'number'
* | 'password' | 'radio' | 'range' | 'reset' | 'search' | 'submit'
* | 'tel' | 'text' | 'time' | 'url' | 'week' | 'combobox-obj' | 'combobox-arr' | 'combobox'
* | 'select-obj' | 'select-arr' | 'select' | 'textarea' | 'datetime' | 'percentage'
* | 'currency' | 'multiselect'
* } InputType
*/
module.exports = {};

/**
 * @typedef { 'string' | 'int' | 'float' | 'intArr' | 'mapping' | 'currency' } DataType
*/
module.exports = {};

/**
* @typedef {Object} Access
* @property {boolean} admin
* @property {boolean} staff
* @property {boolean} user
*/
module.exports = {};

/**
* @typedef {Object} PagerOut
* @property {number} page
* @property {number} perPage
* @property {number} pages
* @property {number} countResult
* @property {Record<string, any[]>} filters
* @property {string[]} order
*/
module.exports = {};

/**
* @typedef {Object} PagerIn
* @property {number} page
* @property {number} perPage
* @property {Record<string, string[]>} filters
* @property {string[]} order
*/
module.exports = {}

/**
 * @typedef {Object} Field
 * @property {string} name
 * @property {string} label
 * @property {string} description
 * @property {DataType} type
 * @property {InputType} inputType
 * @property {boolean} readOnly
 * @property {string} mapping
 * @property {string[]} validations
 * @property {string} refEndpoint
 * @property {string} ref
 */
module.exports = {};

/**
* @typedef {Object} Coordinate
* @property {string} lat
* @property {string} lng
*/
module.exports = {};

/**
* @typedef {Object} Currency
* @property {string} name
* @property {string} code
*/
module.exports = {};

/**
* @typedef {Object} CountryData
* @property {string} country
* @property {string} iso_2
* @property {string} iso_3
* @property {string} country_code
* @property {string} region
* @property {string} region_code
* @property {string} unit_measurement
* @property {Coordinate} coordinate
* @property {Currency} currency
*/
module.exports = {};

/**
 * @typedef {Object} ExtendedAction 
 * @property {import("svelte-icons-pack").IconType} icon
 * @property {boolean} isTargetBlank - if true, open link in new window
 * @property {(row: any) => string} link
 * @property {string} tooltip
 */
module.exports = {}

/**
 * @typedef {Object} ExtendedActionButton
 * @property {import("svelte-icons-pack").IconType} icon
 * @property {(row: any) => void} action
 * @property {string} tooltip
 */
module.exports = {}