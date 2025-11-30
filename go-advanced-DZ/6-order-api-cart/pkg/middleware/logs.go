package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		start := time.Now()
		wrapper := &WriteWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		logrus.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(wrapper, r)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
