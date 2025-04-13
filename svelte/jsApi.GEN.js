
const axios = require("axios");


// rearrange response to be data first instead of axios error
function wrapErr( cb ) {
  return function( err ) {
    let data = ((err.response || {}).data || {})
    if( !data.error ) data.error = err.code
    data._axios = err
    cb( data )
  }
}

// rearrange response to be data first instead of axios error
function wrapOk( cb ) {
  return function( resp ) {
    let data = resp.data || {}
    data._axios = resp
    cb( data )
  }
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
/**
 * @typedef {Object} GuestLoginIn
 * @property {String} email
 * @property {String} password
 */
const GuestLoginIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestLoginOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.userName
 * @property {Object} segments
 */
const GuestLoginOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    userName: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback GuestLoginCallback
 * @param {GuestLoginOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestLoginIn} i
 * @param {GuestLoginCallback} cb
 * @returns {Promise}
 */
exports.GuestLogin = async function GuestLogin( i, cb ) {
  return await axios.post( '/guest/login', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestRegisterIn
 * @property {String} email
 * @property {String} password
 */
const GuestRegisterIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestRegisterOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.userName
 * @property {String} verifyEmailUrl
 */
const GuestRegisterOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    userName: '', // string
  }, // rqAuth.Users
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestRegisterCallback
 * @param {GuestRegisterOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestRegisterIn} i
 * @param {GuestRegisterCallback} cb
 * @returns {Promise}
 */
exports.GuestRegister = async function GuestRegister( i, cb ) {
  return await axios.post( '/guest/register', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResendVerificationEmailIn
 * @property {String} email
 */
const GuestResendVerificationEmailIn = {
  email: '', // string
}
/**
 * @typedef {Object} GuestResendVerificationEmailOut
 * @property {Object} ok
 * @property {String} verifyEmailUrl
 */
const GuestResendVerificationEmailOut = {
  ok: false, // bool
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestResendVerificationEmailCallback
 * @param {GuestResendVerificationEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResendVerificationEmailIn} i
 * @param {GuestResendVerificationEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestResendVerificationEmail = async function GuestResendVerificationEmail( i, cb ) {
  return await axios.post( '/guest/resendVerificationEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResetPasswordIn
 * @property {String} secretCode
 * @property {String} hash
 * @property {String} password
 */
const GuestResetPasswordIn = {
  secretCode: '', // string
  hash: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestResetPasswordOut
 * @property {Object} ok
 */
const GuestResetPasswordOut = {
  ok: false, // bool
}
/**
 * @callback GuestResetPasswordCallback
 * @param {GuestResetPasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResetPasswordIn} i
 * @param {GuestResetPasswordCallback} cb
 * @returns {Promise}
 */
exports.GuestResetPassword = async function GuestResetPassword( i, cb ) {
  return await axios.post( '/guest/resetPassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestVerifyEmailIn
 * @property {String} secretCode
 * @property {String} hash
 */
const GuestVerifyEmailIn = {
  secretCode: '', // string
  hash: '', // string
}
/**
 * @typedef {Object} GuestVerifyEmailOut
 * @property {Object} ok
 * @property {String} email
 */
const GuestVerifyEmailOut = {
  ok: false, // bool
  email: '', // string
}
/**
 * @callback GuestVerifyEmailCallback
 * @param {GuestVerifyEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestVerifyEmailIn} i
 * @param {GuestVerifyEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestVerifyEmail = async function GuestVerifyEmail( i, cb ) {
  return await axios.post( '/guest/verifyEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserBuildingIn
 * @property {String} cmd
 * @property {Object} withMeta
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {number} building.id
 * @property {String} building.buildingName
 * @property {number} building.locationId
 * @property {String} building.facilitiesObj
 * @property {number} building.createdAt
 * @property {number} building.createdBy
 * @property {number} building.updatedAt
 * @property {number} building.updatedBy
 * @property {number} building.deletedAt
 * @property {number} building.deletedBy
 * @property {number} building.restoredBy
 */
const UserBuildingIn = {
  cmd: '', // string
  withMeta: false, // bool
  pager: { // zCrud.PagerIn
    page: 0, // int
    perPage: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerIn
  building: { // rqProperty.Buildings
    id: 0, // uint64
    buildingName: '', // string
    locationId: 0, // uint64
    facilitiesObj: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Buildings
}
/**
 * @typedef {Object} UserBuildingOut
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {number} pager.pages
 * @property {number} pager.total
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {Object} meta.fields
 * @property {Object} meta.mutex
 * @property {String} meta.cachedSelect
 * @property {Object} locations
 * @property {number} building.id
 * @property {String} building.buildingName
 * @property {number} building.locationId
 * @property {String} building.facilitiesObj
 * @property {number} building.createdAt
 * @property {number} building.createdBy
 * @property {number} building.updatedAt
 * @property {number} building.updatedBy
 * @property {number} building.deletedAt
 * @property {number} building.deletedBy
 * @property {number} building.restoredBy
 * @property {Object} buildings
 */
const UserBuildingOut = {
  pager: { // zCrud.PagerOut
    page: 0, // int
    perPage: 0, // int
    pages: 0, // int
    total: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerOut
  meta: { // zCrud.Meta
    fields: { // []Field
    }, // []Field
    mutex: { // sync.Mutex
    }, // sync.Mutex
    cachedSelect: '', // string
  }, // zCrud.Meta
  locations: { // []rqProperty.Locations
  }, // []rqProperty.Locations
  building: { // rqProperty.Buildings
    id: 0, // uint64
    buildingName: '', // string
    locationId: 0, // uint64
    facilitiesObj: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Buildings
  buildings: { // [][]any
  }, // [][]any
}
/**
 * @callback UserBuildingCallback
 * @param {UserBuildingOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserBuildingIn} i
 * @param {UserBuildingCallback} cb
 * @returns {Promise}
 */
exports.UserBuilding = async function UserBuilding( i, cb ) {
  return await axios.post( '/user/building', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserFacilityIn
 * @property {String} cmd
 * @property {Object} withMeta
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {number} facility.id
 * @property {String} facility.facilityName
 * @property {number} facility.extraChargeIDR
 * @property {number} facility.createdAt
 * @property {number} facility.createdBy
 * @property {number} facility.updatedAt
 * @property {number} facility.updatedBy
 * @property {number} facility.deletedAt
 * @property {number} facility.deletedBy
 * @property {number} facility.restoredBy
 */
const UserFacilityIn = {
  cmd: '', // string
  withMeta: false, // bool
  pager: { // zCrud.PagerIn
    page: 0, // int
    perPage: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerIn
  facility: { // rqProperty.Facilities
    id: 0, // uint64
    facilityName: '', // string
    extraChargeIDR: 0, // int64
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Facilities
}
/**
 * @typedef {Object} UserFacilityOut
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {number} pager.pages
 * @property {number} pager.total
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {Object} meta.fields
 * @property {Object} meta.mutex
 * @property {String} meta.cachedSelect
 * @property {number} facility.id
 * @property {String} facility.facilityName
 * @property {number} facility.extraChargeIDR
 * @property {number} facility.createdAt
 * @property {number} facility.createdBy
 * @property {number} facility.updatedAt
 * @property {number} facility.updatedBy
 * @property {number} facility.deletedAt
 * @property {number} facility.deletedBy
 * @property {number} facility.restoredBy
 * @property {Object} facilities
 */
const UserFacilityOut = {
  pager: { // zCrud.PagerOut
    page: 0, // int
    perPage: 0, // int
    pages: 0, // int
    total: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerOut
  meta: { // zCrud.Meta
    fields: { // []Field
    }, // []Field
    mutex: { // sync.Mutex
    }, // sync.Mutex
    cachedSelect: '', // string
  }, // zCrud.Meta
  facility: { // rqProperty.Facilities
    id: 0, // uint64
    facilityName: '', // string
    extraChargeIDR: 0, // int64
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Facilities
  facilities: { // [][]any
  }, // [][]any
}
/**
 * @callback UserFacilityCallback
 * @param {UserFacilityOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserFacilityIn} i
 * @param {UserFacilityCallback} cb
 * @returns {Promise}
 */
exports.UserFacility = async function UserFacility( i, cb ) {
  return await axios.post( '/user/facility', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserLocationIn
 * @property {String} cmd
 * @property {Object} withMeta
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {number} location.id
 * @property {String} location.name
 * @property {String} location.address
 * @property {String} location.gmapLocation
 * @property {number} location.createdAt
 * @property {number} location.createdBy
 * @property {number} location.updatedAt
 * @property {number} location.updatedBy
 * @property {number} location.deletedAt
 * @property {number} location.deletedBy
 * @property {number} location.restoredBy
 */
const UserLocationIn = {
  cmd: '', // string
  withMeta: false, // bool
  pager: { // zCrud.PagerIn
    page: 0, // int
    perPage: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerIn
  location: { // rqProperty.Locations
    id: 0, // uint64
    name: '', // string
    address: '', // string
    gmapLocation: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Locations
}
/**
 * @typedef {Object} UserLocationOut
 * @property {number} pager.page
 * @property {number} pager.perPage
 * @property {number} pager.pages
 * @property {number} pager.total
 * @property {Object} pager.filters
 * @property {Array<String>} pager.order
 * @property {Object} meta.fields
 * @property {Object} meta.mutex
 * @property {String} meta.cachedSelect
 * @property {number} location.id
 * @property {String} location.name
 * @property {String} location.address
 * @property {String} location.gmapLocation
 * @property {number} location.createdAt
 * @property {number} location.createdBy
 * @property {number} location.updatedAt
 * @property {number} location.updatedBy
 * @property {number} location.deletedAt
 * @property {number} location.deletedBy
 * @property {number} location.restoredBy
 * @property {Object} locations
 */
const UserLocationOut = {
  pager: { // zCrud.PagerOut
    page: 0, // int
    perPage: 0, // int
    pages: 0, // int
    total: 0, // int
    filters: { // map[string][]string
    }, // map[string][]string
    order: [], // []string
  }, // zCrud.PagerOut
  meta: { // zCrud.Meta
    fields: { // []Field
    }, // []Field
    mutex: { // sync.Mutex
    }, // sync.Mutex
    cachedSelect: '', // string
  }, // zCrud.Meta
  location: { // rqProperty.Locations
    id: 0, // uint64
    name: '', // string
    address: '', // string
    gmapLocation: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    deletedBy: 0, // uint64
    restoredBy: 0, // uint64
  }, // rqProperty.Locations
  locations: { // [][]any
  }, // [][]any
}
/**
 * @callback UserLocationCallback
 * @param {UserLocationOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserLocationIn} i
 * @param {UserLocationCallback} cb
 * @returns {Promise}
 */
exports.UserLocation = async function UserLocation( i, cb ) {
  return await axios.post( '/user/location', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserLogoutIn
 */
const UserLogoutIn = {
}
/**
 * @typedef {Object} UserLogoutOut
 * @property {number} logoutAt
 */
const UserLogoutOut = {
  logoutAt: 0, // int64
}
/**
 * @callback UserLogoutCallback
 * @param {UserLogoutOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserLogoutIn} i
 * @param {UserLogoutCallback} cb
 * @returns {Promise}
 */
exports.UserLogout = async function UserLogout( i, cb ) {
  return await axios.post( '/user/logout', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserProfileIn
 */
const UserProfileIn = {
}
/**
 * @typedef {Object} UserProfileOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.userName
 * @property {Object} segments
 */
const UserProfileOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    userName: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback UserProfileCallback
 * @param {UserProfileOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserProfileIn} i
 * @param {UserProfileCallback} cb
 * @returns {Promise}
 */
exports.UserProfile = async function UserProfile( i, cb ) {
  return await axios.post( '/user/profile', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserSessionKillIn
 * @property {String} sessionTokenHash
 */
const UserSessionKillIn = {
  sessionTokenHash: '', // string
}
/**
 * @typedef {Object} UserSessionKillOut
 * @property {number} logoutAt
 * @property {number} sessionTerminated
 */
const UserSessionKillOut = {
  logoutAt: 0, // int64
  sessionTerminated: 0, // int64
}
/**
 * @callback UserSessionKillCallback
 * @param {UserSessionKillOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserSessionKillIn} i
 * @param {UserSessionKillCallback} cb
 * @returns {Promise}
 */
exports.UserSessionKill = async function UserSessionKill( i, cb ) {
  return await axios.post( '/user/sessionKill', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}


// Code generated by 1_codegen_test.go DO NOT EDIT.
