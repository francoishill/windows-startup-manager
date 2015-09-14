package HttpHelperService

import (
	"net/http"
)

type HttpHelperService interface {
	RenderJson(w http.ResponseWriter, data interface{})
	RenderHtml(w http.ResponseWriter, templateName string, data interface{})

	GetRequiredUrlQueryValue_String(r *http.Request, keyName string) string
	GetRequiredUrlQueryValue_Int64(r *http.Request, keyName string) int64

	GetRequiredUrlParamValue_String(r *http.Request, paramName string) string
	GetRequiredUrlParamValue_Int64(r *http.Request, paramName string) int64
}
