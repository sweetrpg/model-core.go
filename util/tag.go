package util

import (
	"github.com/sweetrpg/common.go/util"
	"github.com/sweetrpg/model-core.go/models"
	"github.com/sweetrpg/model-core.go/vo"
)

func FromTagModels(models []models.Tag) []vo.TagVO {
	return util.Map(models, FromTagModel)
}

func FromTagModel(model models.Tag) *vo.TagVO {
	return &vo.TagVO{
		Name:  model.Name,
		Value: model.Value,
	}
}

func ToTagModels(vos []vo.TagVO) []models.Tag {
	return util.Map(vos, ToTagModel)
}

func ToTagModel(vo vo.TagVO) *models.Tag {
	return &models.Tag{
		Name:  vo.Name,
		Value: vo.Value,
	}
}
