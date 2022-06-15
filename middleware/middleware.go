package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CacheProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.String(), "coin/history") {
			w.Header().Set("Cache-Control", "max-age=1")
			now := time.Now()
			custom := now.Format("2006-01-02 15:04")
			etag := custom
			w.Header().Set("Etag", etag)
			fmt.Println(r.Header)
			if match := r.Header.Get("If-None-Match"); match != "" {
				if strings.Contains(match, etag) {
					w.WriteHeader(http.StatusNotModified)
					return
				}
			}
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))

	})
}

func AuthProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signingString := "1234"
		secret := "SEASEASEASE"
		digest := hmac.New(sha256.New, []byte(secret))
		digest.Write([]byte(signingString))
		signature := base64.StdEncoding.EncodeToString(digest.Sum(nil))
		fmt.Println(signature)
		match := r.Header.Get("Authorization")
		if match != signature {
			//캐쉬사용시 header 브라우저
			//w.WriteHeader(http.StatusUnauthorized)
			//return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))

	})
}
