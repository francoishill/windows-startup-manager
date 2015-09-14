package DefaultHttpHelperService

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"

	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/HttpHelper"
)

type defaultHttpHelperService struct {
	renderer *render.Render
}

func (d *defaultHttpHelperService) RenderJson(w http.ResponseWriter, data interface{}) {
	d.renderer.JSON(w, http.StatusOK, data)
}

func (d *defaultHttpHelperService) RenderHtml(w http.ResponseWriter, templateName string, data interface{}) {
	d.renderer.HTML(w, http.StatusOK, templateName, data)
}

func parseInt64(s string) int64 {
	intVal, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return intVal
}

func (d *defaultHttpHelperService) GetRequiredUrlQueryValue_String(r *http.Request, keyName string) string {
	val := r.URL.Query().Get(keyName)
	if val == "" {
		panic(keyName + " cannot be found from URL")
	}
	return val
}

func (d *defaultHttpHelperService) GetRequiredUrlQueryValue_Int64(r *http.Request, keyName string) int64 {
	return parseInt64(d.GetRequiredUrlQueryValue_String(r, keyName))
}

func (d *defaultHttpHelperService) GetRequiredUrlParamValue_String(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	paramValue, varFound := vars[paramName]
	if !varFound {
		panic(paramName + " cannot be found from URL")
	}
	return paramValue
}

func (d *defaultHttpHelperService) GetRequiredUrlParamValue_Int64(r *http.Request, paramName string) int64 {
	return parseInt64(d.GetRequiredUrlParamValue_String(r, paramName))
}

func New(isDevMode bool) HttpHelperService {
	tmpRenderer := render.New(render.Options{
		Directory:     "Views",
		Layout:        "_Layouts/Default",
		Extensions:    []string{".gohtml"},
		IndentJSON:    true,
		IsDevelopment: isDevMode,
	})
	return &defaultHttpHelperService{
		tmpRenderer,
	}
}
