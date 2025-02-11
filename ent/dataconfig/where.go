// Code generated by ent, DO NOT EDIT.

package dataconfig

import (
	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/janus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldLTE(FieldID, id))
}

// DaysToKeep applies equality check predicate on the "days_to_keep" field. It's identical to DaysToKeepEQ.
func DaysToKeep(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldEQ(FieldDaysToKeep, v))
}

// DaysToKeepEQ applies the EQ predicate on the "days_to_keep" field.
func DaysToKeepEQ(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldEQ(FieldDaysToKeep, v))
}

// DaysToKeepNEQ applies the NEQ predicate on the "days_to_keep" field.
func DaysToKeepNEQ(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldNEQ(FieldDaysToKeep, v))
}

// DaysToKeepIn applies the In predicate on the "days_to_keep" field.
func DaysToKeepIn(vs ...int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldIn(FieldDaysToKeep, vs...))
}

// DaysToKeepNotIn applies the NotIn predicate on the "days_to_keep" field.
func DaysToKeepNotIn(vs ...int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldNotIn(FieldDaysToKeep, vs...))
}

// DaysToKeepGT applies the GT predicate on the "days_to_keep" field.
func DaysToKeepGT(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldGT(FieldDaysToKeep, v))
}

// DaysToKeepGTE applies the GTE predicate on the "days_to_keep" field.
func DaysToKeepGTE(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldGTE(FieldDaysToKeep, v))
}

// DaysToKeepLT applies the LT predicate on the "days_to_keep" field.
func DaysToKeepLT(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldLT(FieldDaysToKeep, v))
}

// DaysToKeepLTE applies the LTE predicate on the "days_to_keep" field.
func DaysToKeepLTE(v int) predicate.DataConfig {
	return predicate.DataConfig(sql.FieldLTE(FieldDaysToKeep, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DataConfig) predicate.DataConfig {
	return predicate.DataConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DataConfig) predicate.DataConfig {
	return predicate.DataConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DataConfig) predicate.DataConfig {
	return predicate.DataConfig(sql.NotPredicates(p))
}
