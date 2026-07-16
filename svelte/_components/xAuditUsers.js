export function auditUsersFromResponse(o) {
  return o?.auditUsers || {};
}

export function auditUserRefs(auditUsers) {
  const refs = auditUsers || {};
  return {
    createdBy: refs,
    updatedBy: refs,
    deletedBy: refs,
  };
}
