package midapp

import (
	"context"
	"fmt"
	"net/http"
	"pelatihan_golang/pkg/fireapp"
	"pelatihan_golang/pkg/jwtapp"
)

// Middleware untuk memeriksa apakah pengguna sudah terautentikasi
func IsAuthenticatedWeb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetSession().Get(r, "session")
		if len(session.Values) == 0 {
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		tokenString, ok := session.Values["token"].(string)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, err := jwtapp.VerifyJWTToken(tokenString, "")
		if err != nil {
			http.Error(w, "Error token verifications", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.User)
		next(w, r.WithContext(ctx))
	}
}

func IsAuthenticatedAPI(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")
		fmt.Println("Authorization:", idToken)
		if idToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := fireapp.GetAuth().VerifyIDToken(context.Background(), idToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to VerifyIDToken ID token: %v", err), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "user", token)
		next(w, r.WithContext(ctx))
	}
}
