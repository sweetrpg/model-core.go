package util

import (
	"testing"

	"github.com/sweetrpg/model-core.go/models"
	"github.com/sweetrpg/model-core.go/vo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PropertyTestSuite struct {
	suite.Suite
}

func (suite *PropertyTestSuite) TestFromPropertyModel() {
	model := models.Property{Name: "Strength", Kind: "Integer", Value: "10"}
	result := FromPropertyModel(model)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "Strength", result.Name)
	assert.Equal(suite.T(), "Integer", result.Kind)
	assert.Equal(suite.T(), "10", result.Value)
}

func (suite *PropertyTestSuite) TestFromPropertyModels() {
	models := []models.Property{
		{Name: "Strength", Kind: "Integer", Value: "10"},
		{Name: "Dexterity", Kind: "Float", Value: "12.5"},
	}
	result := FromPropertyModels(models)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "Strength", result[0].Name)
	assert.Equal(suite.T(), "Dexterity", result[1].Name)
}

func (suite *PropertyTestSuite) TestFromPropertyModelsEmpty() {
	result := FromPropertyModels([]models.Property{})
	assert.Empty(suite.T(), result)
}

func (suite *PropertyTestSuite) TestToPropertyModel() {
	propVO := vo.PropertyVO{Name: "Strength", Kind: "Integer", Value: "10"}
	result := ToPropertyModel(propVO)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "Strength", result.Name)
	assert.Equal(suite.T(), "Integer", result.Kind)
	assert.Equal(suite.T(), "10", result.Value)
}

func (suite *PropertyTestSuite) TestToPropertyModels() {
	vos := []vo.PropertyVO{
		{Name: "Strength", Kind: "Integer", Value: "10"},
		{Name: "Dexterity", Kind: "Float", Value: "12.5"},
	}
	result := ToPropertyModels(vos)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "Strength", result[0].Name)
	assert.Equal(suite.T(), "Dexterity", result[1].Name)
}

func (suite *PropertyTestSuite) TestToPropertyModelsEmpty() {
	result := ToPropertyModels([]vo.PropertyVO{})
	assert.Empty(suite.T(), result)
}

func TestPropertyTestSuite(t *testing.T) {
	suite.Run(t, new(PropertyTestSuite))
}
