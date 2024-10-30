package vo

import (
	"github.com/sweetrpg/common/util"
	"github.com/sweetrpg/model-core/models"
)

type TagVO struct {
	Name  string `json:"name" jsonapi:"attr,name"`
	Value string `json:"value" jsonapi:"attr,value"`
}

func FromTagModels(models []models.Tag) []TagVO {
	return util.Map(models, FromTagModel)
}

func FromTagModel(model models.Tag) *TagVO {
	return &TagVO{
		Name:  model.Name,
		Value: model.Value,
	}
}
