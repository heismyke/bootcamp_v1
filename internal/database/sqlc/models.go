// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type MinimumSkill string

const (
	MinimumSkillBeginner     MinimumSkill = "beginner"
	MinimumSkillIntermediate MinimumSkill = "intermediate"
	MinimumSkillAdvanced     MinimumSkill = "advanced"
)

func (e *MinimumSkill) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MinimumSkill(s)
	case string:
		*e = MinimumSkill(s)
	default:
		return fmt.Errorf("unsupported scan type for MinimumSkill: %T", src)
	}
	return nil
}

type NullMinimumSkill struct {
	MinimumSkill MinimumSkill `json:"minimum_skill"`
	Valid        bool         `json:"valid"` // Valid is true if MinimumSkill is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMinimumSkill) Scan(value interface{}) error {
	if value == nil {
		ns.MinimumSkill, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MinimumSkill.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMinimumSkill) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MinimumSkill), nil
}

type UserRole string

const (
	UserRoleUser      UserRole = "user"
	UserRolePublisher UserRole = "publisher"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Bootcamps struct {
	ID          int64          `json:"id"`
	UserID      sql.NullInt64  `json:"user_id"`
	Name        string         `json:"name"`
	Slug        sql.NullString `json:"slug"`
	Description string         `json:"description"`
	// Must start with http:// or https://
	Website string `json:"website"`
	// Use E.164 format (+1234567890)
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	// Latitude value (-90 to 90)
	Latitude string `json:"latitude"`
	// Longitude value (-180 to 180)
	Longitude string `json:"longitude"`
	// Stores city, state, country
	LocationDetails json.RawMessage `json:"location_details"`
	// List of career paths offered
	Careers json.RawMessage `json:"careers"`
	// Min: 1, Max: 10
	AverageRating string       `json:"average_rating"`
	AverageCost   string       `json:"average_cost"`
	Photo         string       `json:"photo"`
	Housing       bool         `json:"housing"`
	JobAssistance bool         `json:"job_assistance"`
	JobGuarantee  bool         `json:"job_guarantee"`
	AcceptGi      bool         `json:"accept_gi"`
	CreatedAt     sql.NullTime `json:"created_at"`
}

type Courses struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Weeks       string `json:"weeks"`
	Tuition     string `json:"tuition"`
	// Allowed values: beginner, intermediate, advanced
	MinimumSkill         MinimumSkill `json:"minimum_skill"`
	ScholarshipAvailable bool         `json:"scholarship_available"`
	BootcampID           int64        `json:"bootcamp_id"`
	CreatedAt            sql.NullTime `json:"created_at"`
}

type Users struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// Must be a valid email
	Email string `json:"email"`
	// Allowed values: user, publisher
	Role                UserRole     `json:"role"`
	Password            string       `json:"password"`
	ResetPasswordToken  string       `json:"reset_password_token"`
	ResetPasswordExpire time.Time    `json:"reset_password_expire"`
	ConfirmEmailToken   string       `json:"confirm_email_token"`
	IsEmailConfirmed    bool         `json:"is_email_confirmed"`
	TwoFactorCode       string       `json:"two_factor_code"`
	TwoFactorCodeExpire time.Time    `json:"two_factor_code_expire"`
	TwoFactorEnable     bool         `json:"two_factor_enable"`
	CreatedAt           sql.NullTime `json:"created_at"`
}
