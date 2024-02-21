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

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUpdateAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldDeleteAt, v))
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

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldUpdateAt, v))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.GroupUser {
	return predicate.GroupUser(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.GroupUser {
	return predicate.GroupUser(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.GroupUser {
	return predicate.GroupUser(sql.FieldNotNull(FieldDeleteAt))
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
