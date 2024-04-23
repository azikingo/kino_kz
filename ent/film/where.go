// Code generated by ent, DO NOT EDIT.

package film

import (
	"kinokz/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldDescription, v))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldAuthorID, v))
}

// ManagerID applies equality check predicate on the "manager_id" field. It's identical to ManagerIDEQ.
func ManagerID(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldManagerID, v))
}

// StadiumID applies equality check predicate on the "stadium_id" field. It's identical to StadiumIDEQ.
func StadiumID(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldStadiumID, v))
}

// Star applies equality check predicate on the "star" field. It's identical to StarEQ.
func Star(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldStar, v))
}

// AttackRating applies equality check predicate on the "attack_rating" field. It's identical to AttackRatingEQ.
func AttackRating(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldAttackRating, v))
}

// MidfieldRating applies equality check predicate on the "midfield_rating" field. It's identical to MidfieldRatingEQ.
func MidfieldRating(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldMidfieldRating, v))
}

// DefenceRating applies equality check predicate on the "defence_rating" field. It's identical to DefenceRatingEQ.
func DefenceRating(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldDefenceRating, v))
}

// GoalkeeperRating applies equality check predicate on the "goalkeeper_rating" field. It's identical to GoalkeeperRatingEQ.
func GoalkeeperRating(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldGoalkeeperRating, v))
}

// OverallRating applies equality check predicate on the "overall_rating" field. It's identical to OverallRatingEQ.
func OverallRating(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldOverallRating, v))
}

// LogoURL applies equality check predicate on the "logo_url" field. It's identical to LogoURLEQ.
func LogoURL(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldLogoURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Film {
	return predicate.Film(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Film {
	return predicate.Film(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Film {
	return predicate.Film(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Film {
	return predicate.Film(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Film {
	return predicate.Film(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Film {
	return predicate.Film(sql.FieldContainsFold(FieldDescription, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldAuthorID, vs...))
}

// AuthorIDIsNil applies the IsNil predicate on the "author_id" field.
func AuthorIDIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldAuthorID))
}

// AuthorIDNotNil applies the NotNil predicate on the "author_id" field.
func AuthorIDNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldAuthorID))
}

// ManagerIDEQ applies the EQ predicate on the "manager_id" field.
func ManagerIDEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldManagerID, v))
}

// ManagerIDNEQ applies the NEQ predicate on the "manager_id" field.
func ManagerIDNEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldManagerID, v))
}

// ManagerIDIn applies the In predicate on the "manager_id" field.
func ManagerIDIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldManagerID, vs...))
}

// ManagerIDNotIn applies the NotIn predicate on the "manager_id" field.
func ManagerIDNotIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldManagerID, vs...))
}

// ManagerIDIsNil applies the IsNil predicate on the "manager_id" field.
func ManagerIDIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldManagerID))
}

// ManagerIDNotNil applies the NotNil predicate on the "manager_id" field.
func ManagerIDNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldManagerID))
}

// StadiumIDEQ applies the EQ predicate on the "stadium_id" field.
func StadiumIDEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldStadiumID, v))
}

// StadiumIDNEQ applies the NEQ predicate on the "stadium_id" field.
func StadiumIDNEQ(v int64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldStadiumID, v))
}

// StadiumIDIn applies the In predicate on the "stadium_id" field.
func StadiumIDIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldStadiumID, vs...))
}

// StadiumIDNotIn applies the NotIn predicate on the "stadium_id" field.
func StadiumIDNotIn(vs ...int64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldStadiumID, vs...))
}

// StadiumIDGT applies the GT predicate on the "stadium_id" field.
func StadiumIDGT(v int64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldStadiumID, v))
}

// StadiumIDGTE applies the GTE predicate on the "stadium_id" field.
func StadiumIDGTE(v int64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldStadiumID, v))
}

// StadiumIDLT applies the LT predicate on the "stadium_id" field.
func StadiumIDLT(v int64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldStadiumID, v))
}

// StadiumIDLTE applies the LTE predicate on the "stadium_id" field.
func StadiumIDLTE(v int64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldStadiumID, v))
}

// StadiumIDIsNil applies the IsNil predicate on the "stadium_id" field.
func StadiumIDIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldStadiumID))
}

// StadiumIDNotNil applies the NotNil predicate on the "stadium_id" field.
func StadiumIDNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldStadiumID))
}

// StarEQ applies the EQ predicate on the "star" field.
func StarEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldStar, v))
}

// StarNEQ applies the NEQ predicate on the "star" field.
func StarNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldStar, v))
}

// StarIn applies the In predicate on the "star" field.
func StarIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldStar, vs...))
}

// StarNotIn applies the NotIn predicate on the "star" field.
func StarNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldStar, vs...))
}

// StarGT applies the GT predicate on the "star" field.
func StarGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldStar, v))
}

// StarGTE applies the GTE predicate on the "star" field.
func StarGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldStar, v))
}

// StarLT applies the LT predicate on the "star" field.
func StarLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldStar, v))
}

// StarLTE applies the LTE predicate on the "star" field.
func StarLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldStar, v))
}

// StarIsNil applies the IsNil predicate on the "star" field.
func StarIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldStar))
}

// StarNotNil applies the NotNil predicate on the "star" field.
func StarNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldStar))
}

// AttackRatingEQ applies the EQ predicate on the "attack_rating" field.
func AttackRatingEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldAttackRating, v))
}

// AttackRatingNEQ applies the NEQ predicate on the "attack_rating" field.
func AttackRatingNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldAttackRating, v))
}

// AttackRatingIn applies the In predicate on the "attack_rating" field.
func AttackRatingIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldAttackRating, vs...))
}

// AttackRatingNotIn applies the NotIn predicate on the "attack_rating" field.
func AttackRatingNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldAttackRating, vs...))
}

// AttackRatingGT applies the GT predicate on the "attack_rating" field.
func AttackRatingGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldAttackRating, v))
}

// AttackRatingGTE applies the GTE predicate on the "attack_rating" field.
func AttackRatingGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldAttackRating, v))
}

// AttackRatingLT applies the LT predicate on the "attack_rating" field.
func AttackRatingLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldAttackRating, v))
}

// AttackRatingLTE applies the LTE predicate on the "attack_rating" field.
func AttackRatingLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldAttackRating, v))
}

// AttackRatingIsNil applies the IsNil predicate on the "attack_rating" field.
func AttackRatingIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldAttackRating))
}

// AttackRatingNotNil applies the NotNil predicate on the "attack_rating" field.
func AttackRatingNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldAttackRating))
}

// MidfieldRatingEQ applies the EQ predicate on the "midfield_rating" field.
func MidfieldRatingEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldMidfieldRating, v))
}

// MidfieldRatingNEQ applies the NEQ predicate on the "midfield_rating" field.
func MidfieldRatingNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldMidfieldRating, v))
}

// MidfieldRatingIn applies the In predicate on the "midfield_rating" field.
func MidfieldRatingIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldMidfieldRating, vs...))
}

// MidfieldRatingNotIn applies the NotIn predicate on the "midfield_rating" field.
func MidfieldRatingNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldMidfieldRating, vs...))
}

// MidfieldRatingGT applies the GT predicate on the "midfield_rating" field.
func MidfieldRatingGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldMidfieldRating, v))
}

// MidfieldRatingGTE applies the GTE predicate on the "midfield_rating" field.
func MidfieldRatingGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldMidfieldRating, v))
}

// MidfieldRatingLT applies the LT predicate on the "midfield_rating" field.
func MidfieldRatingLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldMidfieldRating, v))
}

// MidfieldRatingLTE applies the LTE predicate on the "midfield_rating" field.
func MidfieldRatingLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldMidfieldRating, v))
}

// MidfieldRatingIsNil applies the IsNil predicate on the "midfield_rating" field.
func MidfieldRatingIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldMidfieldRating))
}

// MidfieldRatingNotNil applies the NotNil predicate on the "midfield_rating" field.
func MidfieldRatingNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldMidfieldRating))
}

// DefenceRatingEQ applies the EQ predicate on the "defence_rating" field.
func DefenceRatingEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldDefenceRating, v))
}

// DefenceRatingNEQ applies the NEQ predicate on the "defence_rating" field.
func DefenceRatingNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldDefenceRating, v))
}

// DefenceRatingIn applies the In predicate on the "defence_rating" field.
func DefenceRatingIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldDefenceRating, vs...))
}

// DefenceRatingNotIn applies the NotIn predicate on the "defence_rating" field.
func DefenceRatingNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldDefenceRating, vs...))
}

// DefenceRatingGT applies the GT predicate on the "defence_rating" field.
func DefenceRatingGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldDefenceRating, v))
}

// DefenceRatingGTE applies the GTE predicate on the "defence_rating" field.
func DefenceRatingGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldDefenceRating, v))
}

// DefenceRatingLT applies the LT predicate on the "defence_rating" field.
func DefenceRatingLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldDefenceRating, v))
}

// DefenceRatingLTE applies the LTE predicate on the "defence_rating" field.
func DefenceRatingLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldDefenceRating, v))
}

// DefenceRatingIsNil applies the IsNil predicate on the "defence_rating" field.
func DefenceRatingIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldDefenceRating))
}

// DefenceRatingNotNil applies the NotNil predicate on the "defence_rating" field.
func DefenceRatingNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldDefenceRating))
}

// GoalkeeperRatingEQ applies the EQ predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingNEQ applies the NEQ predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingIn applies the In predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldGoalkeeperRating, vs...))
}

// GoalkeeperRatingNotIn applies the NotIn predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldGoalkeeperRating, vs...))
}

// GoalkeeperRatingGT applies the GT predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingGTE applies the GTE predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingLT applies the LT predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingLTE applies the LTE predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldGoalkeeperRating, v))
}

// GoalkeeperRatingIsNil applies the IsNil predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldGoalkeeperRating))
}

// GoalkeeperRatingNotNil applies the NotNil predicate on the "goalkeeper_rating" field.
func GoalkeeperRatingNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldGoalkeeperRating))
}

// OverallRatingEQ applies the EQ predicate on the "overall_rating" field.
func OverallRatingEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldOverallRating, v))
}

// OverallRatingNEQ applies the NEQ predicate on the "overall_rating" field.
func OverallRatingNEQ(v float64) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldOverallRating, v))
}

// OverallRatingIn applies the In predicate on the "overall_rating" field.
func OverallRatingIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldOverallRating, vs...))
}

// OverallRatingNotIn applies the NotIn predicate on the "overall_rating" field.
func OverallRatingNotIn(vs ...float64) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldOverallRating, vs...))
}

// OverallRatingGT applies the GT predicate on the "overall_rating" field.
func OverallRatingGT(v float64) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldOverallRating, v))
}

// OverallRatingGTE applies the GTE predicate on the "overall_rating" field.
func OverallRatingGTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldOverallRating, v))
}

// OverallRatingLT applies the LT predicate on the "overall_rating" field.
func OverallRatingLT(v float64) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldOverallRating, v))
}

// OverallRatingLTE applies the LTE predicate on the "overall_rating" field.
func OverallRatingLTE(v float64) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldOverallRating, v))
}

// OverallRatingIsNil applies the IsNil predicate on the "overall_rating" field.
func OverallRatingIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldOverallRating))
}

// OverallRatingNotNil applies the NotNil predicate on the "overall_rating" field.
func OverallRatingNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldOverallRating))
}

// LogoURLEQ applies the EQ predicate on the "logo_url" field.
func LogoURLEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldLogoURL, v))
}

// LogoURLNEQ applies the NEQ predicate on the "logo_url" field.
func LogoURLNEQ(v string) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldLogoURL, v))
}

// LogoURLIn applies the In predicate on the "logo_url" field.
func LogoURLIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldLogoURL, vs...))
}

// LogoURLNotIn applies the NotIn predicate on the "logo_url" field.
func LogoURLNotIn(vs ...string) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldLogoURL, vs...))
}

// LogoURLGT applies the GT predicate on the "logo_url" field.
func LogoURLGT(v string) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldLogoURL, v))
}

// LogoURLGTE applies the GTE predicate on the "logo_url" field.
func LogoURLGTE(v string) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldLogoURL, v))
}

// LogoURLLT applies the LT predicate on the "logo_url" field.
func LogoURLLT(v string) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldLogoURL, v))
}

// LogoURLLTE applies the LTE predicate on the "logo_url" field.
func LogoURLLTE(v string) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldLogoURL, v))
}

// LogoURLContains applies the Contains predicate on the "logo_url" field.
func LogoURLContains(v string) predicate.Film {
	return predicate.Film(sql.FieldContains(FieldLogoURL, v))
}

// LogoURLHasPrefix applies the HasPrefix predicate on the "logo_url" field.
func LogoURLHasPrefix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasPrefix(FieldLogoURL, v))
}

// LogoURLHasSuffix applies the HasSuffix predicate on the "logo_url" field.
func LogoURLHasSuffix(v string) predicate.Film {
	return predicate.Film(sql.FieldHasSuffix(FieldLogoURL, v))
}

// LogoURLIsNil applies the IsNil predicate on the "logo_url" field.
func LogoURLIsNil() predicate.Film {
	return predicate.Film(sql.FieldIsNull(FieldLogoURL))
}

// LogoURLNotNil applies the NotNil predicate on the "logo_url" field.
func LogoURLNotNil() predicate.Film {
	return predicate.Film(sql.FieldNotNull(FieldLogoURL))
}

// LogoURLEqualFold applies the EqualFold predicate on the "logo_url" field.
func LogoURLEqualFold(v string) predicate.Film {
	return predicate.Film(sql.FieldEqualFold(FieldLogoURL, v))
}

// LogoURLContainsFold applies the ContainsFold predicate on the "logo_url" field.
func LogoURLContainsFold(v string) predicate.Film {
	return predicate.Film(sql.FieldContainsFold(FieldLogoURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Film {
	return predicate.Film(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Film {
	return predicate.Film(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Film {
	return predicate.Film(sql.FieldLTE(FieldCreatedAt, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Film {
	return predicate.Film(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Film {
	return predicate.Film(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasManager applies the HasEdge predicate on the "manager" edge.
func HasManager() predicate.Film {
	return predicate.Film(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ManagerTable, ManagerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasManagerWith applies the HasEdge predicate on the "manager" edge with a given conditions (other predicates).
func HasManagerWith(preds ...predicate.User) predicate.Film {
	return predicate.Film(func(s *sql.Selector) {
		step := newManagerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Film) predicate.Film {
	return predicate.Film(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Film) predicate.Film {
	return predicate.Film(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Film) predicate.Film {
	return predicate.Film(sql.NotPredicates(p))
}
