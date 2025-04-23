package util

import (
	"testing"

	"github.com/sweetrpg/model-core.go/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PropertyTestSuite struct {
	suite.Suite
}

func (suite *PropertyTestSuite) TestPropertyFields() {
	prop := models.Property{
		Name:  "Strength",
		Kind:  "Integer",
		Value: "10",
	}

	assert.Equal(suite.T(), "Strength", prop.Name)
	assert.Equal(suite.T(), "Integer", prop.Kind)
	assert.Equal(suite.T(), "10", prop.Value)
}

func (suite *PropertyTestSuite) TestPropertyEmpty() {
	prop := models.Property{}
	assert.Empty(suite.T(), prop.Name)
	assert.Empty(suite.T(), prop.Kind)
	assert.Empty(suite.T(), prop.Value)
}

func (suite *PropertyTestSuite) TestPropertySetters() {
	prop := models.Property{}
	prop.Name = "Dexterity"
	prop.Kind = "Float"
	prop.Value = "12.5"

	assert.Equal(suite.T(), "Dexterity", prop.Name)
	assert.Equal(suite.T(), "Float", prop.Kind)
	assert.Equal(suite.T(), "12.5", prop.Value)
}

func TestPropertyTestSuite(t *testing.T) {
	suite.Run(t, new(PropertyTestSuite))
}
