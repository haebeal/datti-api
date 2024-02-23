// Code generated by ent, DO NOT EDIT.

package groupuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/datti-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUID, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldGroupID, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v bool) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldOwner, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldDeletedAt, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldUID, v))
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldContains(FieldUID, v))
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldHasPrefix(FieldUID, v))
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldHasSuffix(FieldUID, v))
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEqualFold(FieldUID, v))
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldContainsFold(FieldUID, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldGroupID, v))
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldContains(FieldGroupID, v))
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldHasPrefix(FieldGroupID, v))
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldHasSuffix(FieldGroupID, v))
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEqualFold(FieldGroupID, v))
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldContainsFold(FieldGroupID, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v bool) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v bool) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldOwner, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupUser) predicate.GroupUser {
	return predicate.GroupUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupUser) predicate.GroupUser {
	return predicate.GroupUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupUser) predicate.GroupUser {
	return predicate.GroupUser(sql.NotPredicates(p))
}
