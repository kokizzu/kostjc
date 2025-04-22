/**
 * @typedef {Object} Location
 * @property {string|number} id
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
 * @property {string|number} id
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
 * @property {string|number} id
 * @property {string} buildingName
 * @property {number} locationId
 * @property {number[]} facilities
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
 * @typedef {Object} Booking
 * @property {string|number} id
 * @property {string} dateStart
 * @property {string} dateEnd
 * @property {number} basePriceIDR
 * @property {string} facilitiesObj
 * @property {number} totalPriceIDR
 * @property {string} paidAt
 * @property {string} tenantId
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
 * @typedef {Object} Payment
 * @property {string|number} id 
 * @property {string} bookingId
 * @property {string} paymentAt
 * @property {number} paidIDR
 * @property {string} paymentMethod
 * @property {string} paymentStatus
 * @property {string} note
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {};
