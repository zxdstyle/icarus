package base

import (
	"errors"
	"github.com/golang-module/carbon/v2"
)

var ErrorModelTypeNotCorrect = errors.New("model is not correct,check your code")

type RepositoryModel interface {
	SetID(ID uint)
	GetID() uint
	SetCreatedAt(createdAt carbon.Carbon)
	GetCreatedAt() carbon.Carbon
	SetUpdatedAt(updatedAt carbon.Carbon)
	GetUpdatedAt() carbon.Carbon
}

type RepositoryModels interface {
	GetRepositoryModel(i int) RepositoryModel
	SetRepositoryModel(i int, model RepositoryModel) error
	Length() int
}

type Model struct {
	ID        uint             `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt *carbon.DateTime `gorm:"index;not null" json:"created_at,omitempty"`
	UpdatedAt *carbon.DateTime `json:"updated_at,omitempty"`
}

func (m *Model) SetID(ID uint) {
	m.ID = ID
}

func (m Model) GetID() uint {
	return m.ID
}

func (m *Model) SetCreatedAt(createdAt carbon.Carbon) {
	m.CreatedAt = &carbon.DateTime{Carbon: createdAt}
}

func (m *Model) GetCreatedAt() carbon.Carbon {
	return m.CreatedAt.Carbon
}

func (m *Model) SetUpdatedAt(updatedAt carbon.Carbon) {
	m.UpdatedAt = &carbon.DateTime{Carbon: updatedAt}
}

func (m *Model) GetUpdatedAt() carbon.Carbon {
	return m.UpdatedAt.Carbon
}

type CreateOnlyModel struct {
	ID        uint             `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt *carbon.DateTime `gorm:"index;not null" json:"created_at,omitempty"`
}

func (m *CreateOnlyModel) SetID(ID uint) {
	m.ID = ID
}

func (m CreateOnlyModel) GetID() uint {
	return m.ID
}

func (m *CreateOnlyModel) SetCreatedAt(createdAt carbon.Carbon) {
	m.CreatedAt = &carbon.DateTime{Carbon: createdAt}
}

func (m *CreateOnlyModel) GetCreatedAt() carbon.Carbon {
	return m.CreatedAt.Carbon
}

func (m *CreateOnlyModel) SetUpdatedAt(updatedAt *carbon.Carbon) {
}

func (m *CreateOnlyModel) GetUpdatedAt() *carbon.Carbon {
	return &carbon.Carbon{}
}
