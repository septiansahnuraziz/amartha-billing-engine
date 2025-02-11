package httpresponse

import (
	"amartha-billing-engine/config"
	"amartha-billing-engine/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WrapperDto struct {
	Message   string
	Data      any
	ErrorCode string
}

type HttpResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewHttpResponse() *HttpResponse {
	return new(HttpResponse)
}

func (h *HttpResponse) WithFileDownload(ctx *gin.Context, file []byte, fileName string) {
	ctx.Header("Content-Disposition", utils.WriteStringTemplate("attachment; filename=%s", fileName))
	ctx.Header("Content-Type", "application/octet-stream")
	if _, err := ctx.Writer.Write(file); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to copy object content: %v", err)
		return
	}
}

func (h *HttpResponse) WithMessage(message string) *HttpResponse {
	h.Message = message
	return h
}

func (h *HttpResponse) WithData(data any) *HttpResponse {
	h.Data = data
	return h
}

type WrapperResponseDTO struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
	Build   string `json:"build"`
	HttpResponse
}

func (h *HttpResponse) ToWrapperResponseDTO(ctx *gin.Context, httpStatus int) {
	var wrapperResponseDTO WrapperResponseDTO
	wrapperResponseDTO.AppName = config.AppName()
	wrapperResponseDTO.Build = config.AppBuild()
	wrapperResponseDTO.Version = config.AppVersion()
	wrapperResponseDTO.Data = h.Data
	wrapperResponseDTO.Message = h.Message

	ctx.JSON(httpStatus, wrapperResponseDTO)
}
