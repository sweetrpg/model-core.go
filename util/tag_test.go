package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/sweetrpg/model-core.go/models"
	"github.com/sweetrpg/model-core.go/vo"
)

type TagTestSuite struct {
	suite.Suite
}

func (suite *TagTestSuite) SetupTest() {
	// Setup code if needed
}

func (suite *TagTestSuite) TestFromTagModel() {
	model := models.Tag{Name: "genre", Value: "fantasy"}
	result := FromTagModel(model)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "genre", result.Name)
	assert.Equal(suite.T(), "fantasy", result.Value)
}

func (suite *TagTestSuite) TestFromTagModels() {
	models := []models.Tag{
		{Name: "genre", Value: "fantasy"},
		{Name: "system", Value: "d20"},
	}
	result := FromTagModels(models)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "genre", result[0].Name)
	assert.Equal(suite.T(), "fantasy", result[0].Value)
	assert.Equal(suite.T(), "system", result[1].Name)
	assert.Equal(suite.T(), "d20", result[1].Value)
}

func (suite *TagTestSuite) TestToTagModel() {
	vo := vo.TagVO{Name: "setting", Value: "space"}
	result := ToTagModel(vo)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "setting", result.Name)
	assert.Equal(suite.T(), "space", result.Value)
}

func (suite *TagTestSuite) TestToTagModels() {
	vos := []vo.TagVO{
		{Name: "setting", Value: "space"},
		{Name: "theme", Value: "horror"},
	}
	result := ToTagModels(vos)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "setting", result[0].Name)
	assert.Equal(suite.T(), "space", result[0].Value)
	assert.Equal(suite.T(), "theme", result[1].Name)
	assert.Equal(suite.T(), "horror", result[1].Value)
}

func (suite *TagTestSuite) TestFromTagModelsEmpty() {
	result := FromTagModels([]models.Tag{})
	assert.Empty(suite.T(), result)
}

func (suite *TagTestSuite) TestToTagModelsEmpty() {
	result := ToTagModels([]vo.TagVO{})
	assert.Empty(suite.T(), result)
}

func TestTagTestSuite(t *testing.T) {
	suite.Run(t, new(TagTestSuite))
}
