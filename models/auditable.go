// Model package
package models

import "time"

// SystemActor is the canonical created_by/updated_by value for a write not driven by an
// authenticated user - event consumers, cron jobs, backfills. Services stamp this rather than
// leaving *_by empty or inventing a per-service string.
const SystemActor = "system"

// Base auditable model
type Auditable struct {
	CreatedAt time.Time  `bson:"created_at" json:"created_at" jsonapi:"attr,created_at"`
	CreatedBy string     `bson:"created_by" json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at" jsonapi:"attr,updated_at"`
	UpdatedBy string     `bson:"updated_by" json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at" jsonapi:"attr,deleted_at,omitempty"`
	DeletedBy *string    `bson:"deleted_by,omitempty" json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}

// StampCreate sets the create and update audit fields to the same user and time. Call it from a
// data layer's Create* with the acting user's canonical id (or SystemActor for a non-user write).
func StampCreate(a *Auditable, userID string, now time.Time) {
	if a == nil {
		return
	}
	a.CreatedAt = now
	a.CreatedBy = userID
	a.UpdatedAt = now
	a.UpdatedBy = userID
}

// StampUpdate advances the update audit fields, leaving CreatedAt/CreatedBy untouched. Call it
// from every mutating data-layer path with the acting user's canonical id (or SystemActor).
func StampUpdate(a *Auditable, userID string, now time.Time) {
	if a == nil {
		return
	}
	a.UpdatedAt = now
	a.UpdatedBy = userID
}
