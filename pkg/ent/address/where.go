// Code generated by ent, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Sri2103/services/pkg/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// StreetAddress applies equality check predicate on the "street_address" field. It's identical to StreetAddressEQ.
func StreetAddress(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetAddress, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// StreetAddressEQ applies the EQ predicate on the "street_address" field.
func StreetAddressEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreetAddress, v))
}

// StreetAddressNEQ applies the NEQ predicate on the "street_address" field.
func StreetAddressNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStreetAddress, v))
}

// StreetAddressIn applies the In predicate on the "street_address" field.
func StreetAddressIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStreetAddress, vs...))
}

// StreetAddressNotIn applies the NotIn predicate on the "street_address" field.
func StreetAddressNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStreetAddress, vs...))
}

// StreetAddressGT applies the GT predicate on the "street_address" field.
func StreetAddressGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStreetAddress, v))
}

// StreetAddressGTE applies the GTE predicate on the "street_address" field.
func StreetAddressGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStreetAddress, v))
}

// StreetAddressLT applies the LT predicate on the "street_address" field.
func StreetAddressLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStreetAddress, v))
}

// StreetAddressLTE applies the LTE predicate on the "street_address" field.
func StreetAddressLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStreetAddress, v))
}

// StreetAddressContains applies the Contains predicate on the "street_address" field.
func StreetAddressContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStreetAddress, v))
}

// StreetAddressHasPrefix applies the HasPrefix predicate on the "street_address" field.
func StreetAddressHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStreetAddress, v))
}

// StreetAddressHasSuffix applies the HasSuffix predicate on the "street_address" field.
func StreetAddressHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStreetAddress, v))
}

// StreetAddressEqualFold applies the EqualFold predicate on the "street_address" field.
func StreetAddressEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStreetAddress, v))
}

// StreetAddressContainsFold applies the ContainsFold predicate on the "street_address" field.
func StreetAddressContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStreetAddress, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldState, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(sql.NotPredicates(p))
}