package service

import "sycret-test-task/internal/model"

type DocsGenerator interface {
	Generate(input model.Input) (model.Output, error)
}

type Services struct {
	DocsGenerator
}

func NewServices() *Services {
	return &Services{
		DocsGenerator: NewDocsGeneratorService(),
	}
}
