package gateway

import (
	"bytes"
	"io/fs"
	"mime"
	"net/http"
	"strings"
	// "time"

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

func MainHandler(gwmux *runtime.ServeMux) http.Handler {
	oa := getOpenAPIHandler()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
		oa.ServeHTTP(w, r)
	})
}

func AddLogger(logger *zap.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare fields to log
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
		proto := r.Proto
		method := r.Method
		// statusCode := r.Response.StatusCode
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")
		body := r.Body
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
    	bodyString := buf.String()
		contentType := r.Header.Get("Content-Type")
		auth := r.Header.Get("Authorization")


		// Log HTTP request
		logger.Debug("request started",
			// zap.Int("statusCode", statusCode),
			zap.String("http-scheme", scheme),
			zap.String("http-proto", proto),
			zap.String("http-method", method),
			zap.String("remote-addr", remoteAddr),
			zap.String("user-agent", userAgent),
			zap.String("uri", uri),
			zap.String("content-type", contentType),
			zap.String("Authorization", auth),
			zap.String("body", bodyString),
		)

		// t1 := time.Now()

		h.ServeHTTP(w, r)
		body.Close()

		// Log HTTP response
		// logger.Debug("request completed",
		// 	zap.String("http-scheme", scheme),
		// 	zap.String("http-proto", proto),
		// 	zap.String("http-method", method),
		// 	zap.String("remote-addr", remoteAddr),
		// 	zap.String("user-agent", userAgent),
		// 	zap.String("uri", uri),
		// 	zap.Float64("elapsed-ms", float64(time.Since(t1).Nanoseconds())/1000000.0),
		// )
	})
}