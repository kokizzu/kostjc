/**
 * @typedef {Object} Menu
 * @property {string|number} id
 * @property {string} name
 * @property {number} hppIDR
 * @property {number} salePriceIDR
 * @property {string} detail
 * @property {string} imageUrl
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
 * @typedef {Object} Sale
 * @property {string|number} id
 * @property {string} cashier
 * @property {string|number} tenantId
 * @property {string} buyerName
 * @property {number[]} menuIds
 * @property {string} paymentMethod
 * @property {string} paymentStatus
 * @property {number} transferIDR
 * @property {number} qrisIDR
 * @property {number} cashIDR
 * @property {number} debtIDR
 * @property {number} topupIDR
 * @property {number} totalPriceIDR
 * @property {number} donation
 * @property {string} salesDate
 * @property {string} paidAt
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
 * @typedef {Object} Payment
 * @property {number} id
 * @property {string} customer
 * @property {number} amount
 * @property {string} method
 * @property {string} status
 * @property {string} time
 */
module.exports = {};


/**
 * @typedef {Object} UnpaidItem
 * @property {number} id
 * @property {string} customer
 * @property {string} item
 * @property {number} amount
 * @property {string} date
 */
module.exports = {};