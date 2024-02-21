// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datti-api/ent/bankaccount"
)

// BankAccountCreate is the builder for creating a BankAccount entity.
type BankAccountCreate struct {
	config
	mutation *BankAccountMutation
	hooks    []Hook
}

// SetAccountCode sets the "account_code" field.
func (bac *BankAccountCreate) SetAccountCode(s string) *BankAccountCreate {
	bac.mutation.SetAccountCode(s)
	return bac
}

// SetBankCode sets the "bank_code" field.
func (bac *BankAccountCreate) SetBankCode(s string) *BankAccountCreate {
	bac.mutation.SetBankCode(s)
	return bac
}

// SetBranchCode sets the "branch_code" field.
func (bac *BankAccountCreate) SetBranchCode(s string) *BankAccountCreate {
	bac.mutation.SetBranchCode(s)
	return bac
}

// SetCreatedAt sets the "created_at" field.
func (bac *BankAccountCreate) SetCreatedAt(t time.Time) *BankAccountCreate {
	bac.mutation.SetCreatedAt(t)
	return bac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableCreatedAt(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetCreatedAt(*t)
	}
	return bac
}

// SetUpdatedAt sets the "updated_at" field.
func (bac *BankAccountCreate) SetUpdatedAt(t time.Time) *BankAccountCreate {
	bac.mutation.SetUpdatedAt(t)
	return bac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableUpdatedAt(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetUpdatedAt(*t)
	}
	return bac
}

// SetDeletedAt sets the "deleted_at" field.
func (bac *BankAccountCreate) SetDeletedAt(t time.Time) *BankAccountCreate {
	bac.mutation.SetDeletedAt(t)
	return bac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableDeletedAt(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetDeletedAt(*t)
	}
	return bac
}

// SetID sets the "id" field.
func (bac *BankAccountCreate) SetID(s string) *BankAccountCreate {
	bac.mutation.SetID(s)
	return bac
}

// Mutation returns the BankAccountMutation object of the builder.
func (bac *BankAccountCreate) Mutation() *BankAccountMutation {
	return bac.mutation
}

// Save creates the BankAccount in the database.
func (bac *BankAccountCreate) Save(ctx context.Context) (*BankAccount, error) {
	bac.defaults()
	return withHooks(ctx, bac.sqlSave, bac.mutation, bac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bac *BankAccountCreate) SaveX(ctx context.Context) *BankAccount {
	v, err := bac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bac *BankAccountCreate) Exec(ctx context.Context) error {
	_, err := bac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bac *BankAccountCreate) ExecX(ctx context.Context) {
	if err := bac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bac *BankAccountCreate) defaults() {
	if _, ok := bac.mutation.CreatedAt(); !ok {
		v := bankaccount.DefaultCreatedAt()
		bac.mutation.SetCreatedAt(v)
	}
	if _, ok := bac.mutation.UpdatedAt(); !ok {
		v := bankaccount.DefaultUpdatedAt()
		bac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bac *BankAccountCreate) check() error {
	if _, ok := bac.mutation.AccountCode(); !ok {
		return &ValidationError{Name: "account_code", err: errors.New(`ent: missing required field "BankAccount.account_code"`)}
	}
	if v, ok := bac.mutation.AccountCode(); ok {
		if err := bankaccount.AccountCodeValidator(v); err != nil {
			return &ValidationError{Name: "account_code", err: fmt.Errorf(`ent: validator failed for field "BankAccount.account_code": %w`, err)}
		}
	}
	if _, ok := bac.mutation.BankCode(); !ok {
		return &ValidationError{Name: "bank_code", err: errors.New(`ent: missing required field "BankAccount.bank_code"`)}
	}
	if v, ok := bac.mutation.BankCode(); ok {
		if err := bankaccount.BankCodeValidator(v); err != nil {
			return &ValidationError{Name: "bank_code", err: fmt.Errorf(`ent: validator failed for field "BankAccount.bank_code": %w`, err)}
		}
	}
	if _, ok := bac.mutation.BranchCode(); !ok {
		return &ValidationError{Name: "branch_code", err: errors.New(`ent: missing required field "BankAccount.branch_code"`)}
	}
	if v, ok := bac.mutation.BranchCode(); ok {
		if err := bankaccount.BranchCodeValidator(v); err != nil {
			return &ValidationError{Name: "branch_code", err: fmt.Errorf(`ent: validator failed for field "BankAccount.branch_code": %w`, err)}
		}
	}
	if _, ok := bac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BankAccount.created_at"`)}
	}
	if _, ok := bac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BankAccount.updated_at"`)}
	}
	if v, ok := bac.mutation.ID(); ok {
		if err := bankaccount.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "BankAccount.id": %w`, err)}
		}
	}
	return nil
}

func (bac *BankAccountCreate) sqlSave(ctx context.Context) (*BankAccount, error) {
	if err := bac.check(); err != nil {
		return nil, err
	}
	_node, _spec := bac.createSpec()
	if err := sqlgraph.CreateNode(ctx, bac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BankAccount.ID type: %T", _spec.ID.Value)
		}
	}
	bac.mutation.id = &_node.ID
	bac.mutation.done = true
	return _node, nil
}

func (bac *BankAccountCreate) createSpec() (*BankAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &BankAccount{config: bac.config}
		_spec = sqlgraph.NewCreateSpec(bankaccount.Table, sqlgraph.NewFieldSpec(bankaccount.FieldID, field.TypeString))
	)
	if id, ok := bac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bac.mutation.AccountCode(); ok {
		_spec.SetField(bankaccount.FieldAccountCode, field.TypeString, value)
		_node.AccountCode = value
	}
	if value, ok := bac.mutation.BankCode(); ok {
		_spec.SetField(bankaccount.FieldBankCode, field.TypeString, value)
		_node.BankCode = value
	}
	if value, ok := bac.mutation.BranchCode(); ok {
		_spec.SetField(bankaccount.FieldBranchCode, field.TypeString, value)
		_node.BranchCode = value
	}
	if value, ok := bac.mutation.CreatedAt(); ok {
		_spec.SetField(bankaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bac.mutation.UpdatedAt(); ok {
		_spec.SetField(bankaccount.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bac.mutation.DeletedAt(); ok {
		_spec.SetField(bankaccount.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// BankAccountCreateBulk is the builder for creating many BankAccount entities in bulk.
type BankAccountCreateBulk struct {
	config
	err      error
	builders []*BankAccountCreate
}

// Save creates the BankAccount entities in the database.
func (bacb *BankAccountCreateBulk) Save(ctx context.Context) ([]*BankAccount, error) {
	if bacb.err != nil {
		return nil, bacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bacb.builders))
	nodes := make([]*BankAccount, len(bacb.builders))
	mutators := make([]Mutator, len(bacb.builders))
	for i := range bacb.builders {
		func(i int, root context.Context) {
			builder := bacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BankAccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) SaveX(ctx context.Context) []*BankAccount {
	v, err := bacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bacb *BankAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := bacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) ExecX(ctx context.Context) {
	if err := bacb.Exec(ctx); err != nil {
		panic(err)
	}
}
