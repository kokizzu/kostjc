/**
 * @typedef {Object} User
 * @property {string | number} id
 * @property {string} email
 * @property {string} fullName
 * @property {string} userName
 * @property {string} password
 * @property {string} role
 * @property {number} createdAt
 * @property {number} createdBy
 * @property {number} updatedAt
 * @property {number} updatedBy
 * @property {number} deletedAt
 * @property {number} passwordSetAt
 * @property {number} lastLoginAt
 * @property {number} verifiedAt
 */
module.exports = {}

/**
 * @typedef {Object} Session
 * @property {string} sessionToken
 * @property {number | string} userId
 * @property {number} expiredAt
 * @property {string} device
 * @property {number} loginAt
 * @property {string} loginIPs
 */
module.exports = {}

/**
 * @typedef {Object} Tenant
 * @property {string | number} id
 * @property {string} tenantName
 * @property {string} ktpRegion
 * @property {string} ktpNumber
 * @property {string} ktpName
 * @property {string} ktpPlaceBirth
 * @property {string} ktpDateBirth
 * @property {string} ktpGender
 * @property {string} ktpAddress
 * @property {string} ktpRtRw
 * @property {string} ktpKelurahanDesa
 * @property {string} ktpKecamatan
 * @property {string} ktpReligion
 * @property {string} ktpMaritalStatus
 * @property {string} ktpCitizenship
 * @property {string} ktpOccupation
 * @property {string} telegramUsername
 * @property {string} whatsappNumber
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {}

/**
 * @typedef {Object} TenantNearbyBirthday
 * @property {string} ktpNumber
 * @property {string} tenantName
 * @property {string} birthDay
 * @property {number} age
 * @property {string} roomName
 */
module.exports = {}