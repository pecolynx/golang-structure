//go:generate mockery --output mock --name WorkbookSearchCondition
//go:generate mockery --output mock --name WorkbookSearchResult
//go:generate mockery --output mock --name WorkbookAddParameter
//go:generate mockery --output mock --name WorkbookUpdateParameter
package model

import (
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
)

type DocumentSearchCondition interface {
	GetPageNo() int
	GetPageSize() int
}

type documentSearchCondition struct {
	PageNo   int
	PageSize int
}

func NewDocumentSearchCondition(pageNo, pageSize int) (DocumentSearchCondition, error) {
	m := &documentSearchCondition{
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (p *documentSearchCondition) GetPageNo() int {
	return p.PageNo
}

func (p *documentSearchCondition) GetPageSize() int {
	return p.PageSize
}

type DocumentSearchResult interface {
	GetTotalCount() int
	GetResults() []DocumentModel
}

type documentSearchResult struct {
	TotalCount int
	Results    []DocumentModel
}

func NewDocumentSearchResult(totalCount int, results []DocumentModel) (DocumentSearchResult, error) {
	m := &documentSearchResult{
		TotalCount: totalCount,
		Results:    results,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}
func (m *documentSearchResult) GetTotalCount() int {
	return m.TotalCount
}

func (m *documentSearchResult) GetResults() []DocumentModel {
	return m.Results
}

// type WorkbookAddParameter interface {
// 	GetProblemType() ProblemTypeName
// 	GetName() string
// 	GetLang2() Lang2
// 	GetQuestionText() string
// 	GetProperties() map[string]string
// }

// type workbookAddParameter struct {
// 	ProblemType  ProblemTypeName
// 	Name         string
// 	Lang2        Lang2
// 	QuestionText string
// 	Properties   map[string]string
// }

// func NewWorkbookAddParameter(problemType ProblemTypeName, name string, lang2 Lang2, questionText string, properties map[string]string) (WorkbookAddParameter, error) {
// 	m := &workbookAddParameter{
// 		ProblemType:  problemType,
// 		Name:         name,
// 		Lang2:        lang2,
// 		QuestionText: questionText,
// 		Properties:   properties,
// 	}

// 	if err := libD.Validator.Struct(m); err != nil {
// 		return nil, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
// 	}

// 	return m, nil
// }

// func (p *workbookAddParameter) GetProblemType() ProblemTypeName {
// 	return p.ProblemType
// }

// func (p *workbookAddParameter) GetName() string {
// 	return p.Name
// }

// func (p *workbookAddParameter) GetLang2() Lang2 {
// 	return p.Lang2
// }

// func (p *workbookAddParameter) GetQuestionText() string {
// 	return p.QuestionText
// }

// func (p *workbookAddParameter) GetProperties() map[string]string {
// 	return p.Properties
// }

// type WorkbookUpdateParameter interface {
// 	GetName() string
// 	GetQuestionText() string
// }

// type workbookUpdateParameter struct {
// 	Name         string
// 	QuestionText string
// }

// func NewWorkbookUpdateParameter(name, questionText string) (WorkbookUpdateParameter, error) {
// 	m := &workbookUpdateParameter{
// 		Name:         name,
// 		QuestionText: questionText,
// 	}

// 	if err := libD.Validator.Struct(m); err != nil {
// 		return nil, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
// 	}

// 	return m, nil
// }

// func (p *workbookUpdateParameter) GetName() string {
// 	return p.Name
// }

// func (p *workbookUpdateParameter) GetQuestionText() string {
// 	return p.QuestionText
// }
