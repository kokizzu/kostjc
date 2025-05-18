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
 * @property {number} descriptionEN
 * @property {string} facilityType
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
 * @property {number[]} extraTenants
 * @property {number|string} roomId
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

/**
 * @typedef {Object} Stock
 * @property {string|number} id
 * @property {string} stockName
 * @property {string} stockAddedAt
 * @property {number} quantity 
 * @property {number} priceIDR 
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
 * @typedef {Object} Room
 * @property {string|number} id
 * @property {string} roomName
 * @property {number} basePriceIDR
 * @property {string} currentTenantId
 * @property {string} firstUseAt
 * @property {string} roomSize
 * @property {string} imageUrl
 * @property {number} createdAt 
 * @property {string} createdBy 
 * @property {number} updatedAt 
 * @property {string} updatedBy 
 * @property {number} deletedAt 
 * @property {string} deletedBy 
 * @property {string} restoredBy
 * @property {string} buildingId
 * @property {string} lastUseAt
 */
module.exports = {};

/**
 * @typedef {Object} BookingDetail
 * @property {number} roomId
 * @property {string} roomName
 * @property {number} tenantId
 * @property {string} tenantName
 * @property {string} dateStart
 * @property {string} dateEnd
 * @property {number} amountPaid
 * @property {number} totalPrice
 * @property {number} deletedAt 
 */
module.exports = {};