// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datti-api/ent/friend"
)

// Friend is the model entity for the Friend schema.
type Friend struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// FriendUID holds the value of the "friend_uid" field.
	FriendUID string `json:"friend_uid,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt     *time.Time `json:"delete_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Friend) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friend.FieldID:
			values[i] = new(sql.NullInt64)
		case friend.FieldUID, friend.FieldFriendUID:
			values[i] = new(sql.NullString)
		case friend.FieldCreateAt, friend.FieldUpdateAt, friend.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Friend fields.
func (f *Friend) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friend.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case friend.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				f.UID = value.String
			}
		case friend.FieldFriendUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field friend_uid", values[i])
			} else if value.Valid {
				f.FriendUID = value.String
			}
		case friend.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				f.CreateAt = value.Time
			}
		case friend.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				f.UpdateAt = value.Time
			}
		case friend.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				f.DeleteAt = new(time.Time)
				*f.DeleteAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Friend.
// This includes values selected through modifiers, order, etc.
func (f *Friend) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Friend.
// Note that you need to call Friend.Unwrap() before calling this method if this Friend
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Friend) Update() *FriendUpdateOne {
	return NewFriendClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Friend entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Friend) Unwrap() *Friend {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Friend is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Friend) String() string {
	var builder strings.Builder
	builder.WriteString("Friend(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("uid=")
	builder.WriteString(f.UID)
	builder.WriteString(", ")
	builder.WriteString("friend_uid=")
	builder.WriteString(f.FriendUID)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(f.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(f.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := f.DeleteAt; v != nil {
		builder.WriteString("delete_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Friends is a parsable slice of Friend.
type Friends []*Friend
