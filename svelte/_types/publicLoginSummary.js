/**
 * @typedef {Object} PublicUnpaidTenant
 * @property {string} TenantName
 * @property {string} RoomName
 * @property {string} LastDate
 * @property {number} DaysAgo
 *
 * @typedef {Object} PublicRoomNotice
 * @property {string} RoomName
 * @property {string} LastTenant
 * @property {string} LastDate
 * @property {number} DaysAgo
 *
 * @typedef {Object} PublicOccupant
 * @property {string} TenantName
 * @property {string[]} RoomNames
 *
 * @typedef {Object} PublicLoginSummary
 * @property {PublicUnpaidTenant[]} UnpaidTenants
 * @property {PublicRoomNotice[]} Rooms
 * @property {PublicOccupant[]} Occupants
 */
