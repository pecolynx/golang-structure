package controller

import (
	"github.com/pecolynx/golang-structure/src/usecase"
)

type DocumentHandler interface {
	// FindDocuments(c *gin.Context)
	// FindDocumentByID(c *gin.Context)
	// AddDocument(c *gin.Context)
	// UpdateDocument(c *gin.Context)
	// RemoveDocument(c *gin.Context)
}

type documentHandler struct {
	// repository             gateway.Repository
	userUsecaseDocument usecase.UserUsecaseDocument
}

func NewDocumentHandler(userUsecaseDocument usecase.UserUsecaseDocument) DocumentHandler {
	return &documentHandler{
		userUsecaseDocument: userUsecaseDocument,
	}
}

// // FindDocuments godoc
// // @Summary Find workbooks
// // @Produce json
// // @Success 200 {object} entity.DocumentSearchResponse
// // @Failure 400
// // @Router /v1/private/workbook/search [post]
// func (h *documentHandler) FindDocuments(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	logger.Info("FindDocuments")

// 	id := c.Param("workbookID")
// 	if id != "search" {
// 		c.Status(http.StatusNotFound)
// 		return
// 	}

// 	controllerhelper.HandleSecuredFunction(c, func(organizationID userD.OrganizationID, operatorID userD.AppUserID) error {
// 		result, err := h.studentUsecaseDocument.FindDocuments(ctx, organizationID, operatorID)
// 		if err != nil {
// 			return liberrors.Errorf("h.studentUsecaseDocument.FindDocuments. err: %w", err)
// 		}

// 		response, err := converter.ToDocumentSearchResponse(result)
// 		if err != nil {
// 			return liberrors.Errorf("converter.ToDocumentSearchResponse. err: %w", err)
// 		}
// 		c.JSON(http.StatusOK, response)
// 		return nil
// 	}, h.errorHandle)
// }

// func (h *privateDocumentHandler) FindDocumentByID(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	logger.Info("FindDocumentByID")

// 	controllerhelper.HandleSecuredFunction(c, func(organizationID userD.OrganizationID, operatorID userD.AppUserID) error {
// 		id := c.Param("workbookID")
// 		workbookID, err := strconv.Atoi(id)
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return nil
// 		}

// 		workbook, err := h.studentUsecaseDocument.FindDocumentByID(ctx, organizationID, operatorID, domain.DocumentID(uint(workbookID)))
// 		if err != nil {
// 			return liberrors.Errorf("failed to FindDocumentByID. err: %w", err)
// 		}

// 		workbookResponse, err := converter.ToDocumentHTTPEntity(workbook)
// 		if err != nil {
// 			return liberrors.Errorf("failed to ToDocumentHTTPEntity. err: %w", err)
// 		}

// 		c.JSON(http.StatusOK, workbookResponse)
// 		return nil
// 	}, h.errorHandle)
// }

// // AddDocument godoc
// // @Summary Create new workbook
// // @Produce json
// // @Param param body entity.DocumentAddParameter true "parameter to create new workbook"
// // @Success 200 {object} controllerhelper.IDResponse
// // @Failure 400
// // @Router /v1/private/workbook [post]
// func (h *privateDocumentHandler) AddDocument(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	logger.Info("AddWokrbook")

// 	controllerhelper.HandleSecuredFunction(c, func(organizationID userD.OrganizationID, operatorID userD.AppUserID) error {
// 		param := entity.DocumentAddParameter{}
// 		if err := c.ShouldBindJSON(&param); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 			logger.Warnf("failed to BindJSON. err: %v", err)
// 			return nil
// 		}

// 		parameter, err := converter.ToDocumentAddParameter(&param)
// 		if err != nil {
// 			return liberrors.Errorf("failed to ToAdd. err: %w", err)
// 		}

// 		workbookID, err := h.studentUsecaseDocument.AddDocument(ctx, organizationID, operatorID, parameter)
// 		if err != nil {
// 			return liberrors.Errorf("failed to addDocument. err: %w", err)
// 		}

// 		c.JSON(http.StatusOK, controllerhelper.IDResponse{ID: uint(workbookID)})
// 		return nil
// 	}, h.errorHandle)
// }

// // UpdateDocument godoc
// // @Summary     Update the workbook
// // @Description update the workbook
// // @Tags        private workbook
// // @Accept      json
// // @Produce     json
// // @Param       workbookID path int true "Document ID"
// // @Param       param body entity.DocumentUpdateParameter true "parameter to update the workbook"
// // @Success     200 {object} controllerhelper.IDResponse
// // @Failure     400
// // @Router      /v1/private/workbook/{workbookID} [put]
// func (h *privateDocumentHandler) UpdateDocument(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	logger.Info("UpdateDocument")

// 	controllerhelper.HandleSecuredFunction(c, func(organizationID userD.OrganizationID, operatorID userD.AppUserID) error {
// 		param := entity.DocumentUpdateParameter{}
// 		if err := c.BindJSON(&param); err != nil {
// 			logger.Warnf("failed to BindJSON. err: %v", err)
// 			return nil
// 		}
// 		workbookID, err := helper.GetUintFromPath(c, "workbookID")
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return nil
// 		}

// 		version, err := helper.GetIntFromQuery(c, "version")
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return nil
// 		}

// 		parameter, err := converter.ToDocumentUpdateParameter(&param)
// 		if err != nil {
// 			return liberrors.Errorf("converter.ToDocumentUpdateParameter. err: %w", err)
// 		}

// 		if err := h.studentUsecaseDocument.UpdateDocument(ctx, organizationID, operatorID, domain.DocumentID(workbookID), version, parameter); err != nil {
// 			return liberrors.Errorf("h.studentUsecaseDocument.UpdateDocument. err: %w", err)
// 		}

// 		c.JSON(http.StatusOK, controllerhelper.IDResponse{ID: workbookID})
// 		return nil
// 	}, h.errorHandle)
// }

// func (h *privateDocumentHandler) RemoveDocument(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	logger.Info("RemoveDocument")

// 	controllerhelper.HandleSecuredFunction(c, func(organizationID userD.OrganizationID, operatorID userD.AppUserID) error {
// 		workbookID, err := helper.GetUintFromPath(c, "workbookID")
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return nil
// 		}

// 		version, err := helper.GetIntFromQuery(c, "version")
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return nil
// 		}

// 		if err := h.studentUsecaseDocument.RemoveDocument(ctx, organizationID, operatorID, domain.DocumentID(workbookID), version); err != nil {
// 			return liberrors.Errorf("h.studentUsecaseDocument.RemoveDocument. err: %w", err)
// 		}

// 		c.Status(http.StatusOK)
// 		return nil
// 	}, h.errorHandle)
// }

// func (h *privateDocumentHandler) errorHandle(c *gin.Context, err error) bool {
// 	ctx := c.Request.Context()
// 	logger := log.FromContext(ctx)
// 	if errors.Is(err, service.ErrDocumentAlreadyExists) {
// 		logger.Warnf("workbookHandler err: %+v", err)
// 		c.JSON(http.StatusConflict, gin.H{"message": "Document already exists"})
// 		return true
// 	} else if errors.Is(err, service.ErrDocumentNotFound) {
// 		logger.Warnf("workbookHandler err: %+v", err)
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
// 		return true
// 	}
// 	logger.WithError(err).Errorf("workbookHandler err: %+v", err)
// 	return false
// }
