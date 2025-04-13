/**
 * @typedef {Object} Location
 * @property {string} id
 * @property {string} name
 * @property {string} address
 * @property {string} gmapLocation
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {};

/**
 * @typedef {Object} Facility
 * @property {string} id
 * @property {string} facilityName
 * @property {number} extraChargeIDR
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {};

/**
 * @typedef {Object} Building
 * @property {string} id
 * @property {string} buildingName
 * @property {number} locationId
 * @property {string} facilitiesObj
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {};