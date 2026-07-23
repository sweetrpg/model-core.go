package util

import (
	"github.com/sweetrpg/common.go/util"
	"github.com/sweetrpg/model-core.go/models"
	"github.com/sweetrpg/model-core.go/vo"
)

func FromPropertyModels(models []models.Property) []vo.PropertyVO {
	return util.Map(models, FromPropertyModel)
}

func FromPropertyModel(model models.Property) *vo.PropertyVO {
	return &vo.PropertyVO{
		Name:  model.Name,
		Kind:  model.Kind,
		Value: model.Value,
	}
}

func ToPropertyModels(vos []vo.PropertyVO) []models.Property {
	return util.Map(vos, ToPropertyModel)
}

func ToPropertyModel(vo vo.PropertyVO) *models.Property {
	return &models.Property{
		Name:  vo.Name,
		Kind:  vo.Kind,
		Value: vo.Value,
	}
}
