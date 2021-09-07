package gateway

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hungtran150/api-app/third_party"
	"go.uber.org/zap"
)

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func FormWrapper(gwmux *runtime.ServeMux, log *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Sugar().Debug("Got request: %#v\n", r)
		if strings.HasPrefix(r.URL.Path, "/api") {
			if strings.ToLower(strings.Split(r.Header.Get("Content-Type"), ";")[0]) == "application/x-www-form-urlencoded" {
				convertFormToJson(w, r, log)
			}
			gwmux.ServeHTTP(w, r)
			return
		}
		getOpenAPIHandler().ServeHTTP(w, r)
	})
}

func convertFormToJson(w http.ResponseWriter, r *http.Request, log *zap.Logger) {
	log.Info("Rewriting form data as json")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Sugar().Error("Bad form request", err.Error())
		return
	}
	bodyMap := make(map[string]interface{}, len(r.Form))
	for k, v := range r.Form {
		if len(v) > 0 {
			bodyMap[k] = v[0]
		}
	}
	payloadString := bodyMap["payload"].(string)
	payloadMap := make(map[string]interface{}, len(payloadString))
	// Convert JSON string to Map
	err := json.Unmarshal([]byte(payloadString), &payloadMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Convert map to JSON byte
	jsonBody, err := json.Marshal(payloadMap)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Construct new body
	r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
	r.ContentLength = int64(len(jsonBody))
	r.Header.Set("Content-Type", "application/json")
}