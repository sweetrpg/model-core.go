package vo

import "time"

type AuditableVO struct {
	CreatedAt time.Time  `bson:"created_at" json:"created_at" jsonapi:"attr,created_at"`
	CreatedBy string     `bson:"created_by" json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at" jsonapi:"attr,updated_at"`
	UpdatedBy string     `bson:"updated_by" json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at" jsonapi:"attr,deleted_at,omitempty"`
	DeletedBy *string    `bson:"deleted_by,omitempty" json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
