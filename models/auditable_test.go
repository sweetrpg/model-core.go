package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuditableTestSuite struct {
	suite.Suite
}

func (suite *AuditableTestSuite) TestSystemActor() {
	assert.Equal(suite.T(), "system", SystemActor)
}

func (suite *AuditableTestSuite) TestStampCreateSetsAllFour() {
	now := time.Date(2026, 9, 2, 12, 0, 0, 0, time.UTC)
	var a Auditable

	StampCreate(&a, "user-1", now)

	assert.Equal(suite.T(), now, a.CreatedAt)
	assert.Equal(suite.T(), "user-1", a.CreatedBy)
	assert.Equal(suite.T(), now, a.UpdatedAt)
	assert.Equal(suite.T(), "user-1", a.UpdatedBy)
	assert.Equal(suite.T(), a.CreatedAt, a.UpdatedAt)
	assert.Nil(suite.T(), a.DeletedAt)
	assert.Nil(suite.T(), a.DeletedBy)
}

func (suite *AuditableTestSuite) TestStampUpdateLeavesCreateUntouched() {
	created := time.Date(2026, 9, 1, 8, 0, 0, 0, time.UTC)
	updated := time.Date(2026, 9, 2, 9, 30, 0, 0, time.UTC)
	a := Auditable{CreatedAt: created, CreatedBy: "user-1", UpdatedAt: created, UpdatedBy: "user-1"}

	StampUpdate(&a, "user-2", updated)

	assert.Equal(suite.T(), created, a.CreatedAt)
	assert.Equal(suite.T(), "user-1", a.CreatedBy)
	assert.Equal(suite.T(), updated, a.UpdatedAt)
	assert.Equal(suite.T(), "user-2", a.UpdatedBy)
}

func (suite *AuditableTestSuite) TestStampCreateWithSystemActor() {
	now := time.Now().UTC()
	var a Auditable

	StampCreate(&a, SystemActor, now)

	assert.Equal(suite.T(), "system", a.CreatedBy)
	assert.Equal(suite.T(), "system", a.UpdatedBy)
}

func (suite *AuditableTestSuite) TestStampNilIsNoOp() {
	assert.NotPanics(suite.T(), func() {
		StampCreate(nil, "user-1", time.Now())
		StampUpdate(nil, "user-1", time.Now())
	})
}

func TestAuditableTestSuite(t *testing.T) {
	suite.Run(t, new(AuditableTestSuite))
}
