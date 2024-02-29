package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pelatihan_golang/pkg/fireapp"
	"pelatihan_golang/pkg/jwtapp"
	"pelatihan_golang/pkg/midapp"
	"pelatihan_golang/pkg/tplapp"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type (
	authHandler struct {
		authConfig *oauth2.Config
		authClient *auth.Client
		stateToken string
	}
)

func NewAuthHandler(r *mux.Router) {
	h := &authHandler{
		authConfig: fireapp.GetOAuthConfig(),
		authClient: fireapp.GetAuth(),
		stateToken: fmt.Sprintf("%v", time.Now().Unix()),
	}
	// home page
	r.HandleFunc("/signin", h.Signin).Methods("GET")
	r.HandleFunc("/welcome", midapp.IsAuthenticatedWeb(h.HomePage)).Methods("GET")
	r.HandleFunc("/public", h.PublicHomePage).Methods("GET")
	r.HandleFunc("/logout", midapp.IsAuthenticatedWeb(h.LogOut)).Methods("GET")
	// Endpoint untuk memulai proses login menggunakan akun Google
	r.HandleFunc("/auth/google/login", h.GoogleLoginHandler).Methods("GET")
	// Endpoint untuk menerima callback setelah login menggunakan akun Google
	r.HandleFunc("/auth/google/callback", h.GoogleCallbackHandler).Methods("GET")
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Handle sign-in logic here (validate credentials, set session, etc.)
		// For simplicity, we'll just redirect to the home page in this example.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tplapp.RenderTemplate(w, "signin.html", nil)
}

// home page
func (h *authHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	tplapp.RenderTemplate(w, "welcome.html", nil)
}

func (h *authHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	// Get the current user from the context
	user, ok := r.Context().Value("user").(*auth.UserRecord)
	if !ok {
		// If the user is not found in the context, consider them not logged in
		http.Redirect(w, r, "/auth/google/login", http.StatusSeeOther)
		return
	}

	// Revoke refresh tokens for the current user
	err := h.authClient.RevokeRefreshTokens(context.Background(), user.UID)
	if err != nil {
		http.Error(w, "Failed to RevokeRefreshTokens", http.StatusInternalServerError)
		return
	}

	// Clear the session token
	session, _ := midapp.GetSession().Get(r, "session")
	session.Values["token"] = nil
	session.Save(r, w)

	// Redirect to the Google login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *authHandler) PublicHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Public Home Page"))
}

// Handler untuk memulai proses login menggunakan akun Google
func (h *authHandler) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := h.authConfig.AuthCodeURL(h.stateToken)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *authHandler) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != h.stateToken {
		http.Error(w, "Invalid state token", http.StatusUnauthorized)
		return
	}

	code := r.FormValue("code")
	token, err := h.authConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}

	client := h.authConfig.Client(context.Background(), token)
	client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	fmt.Println("Auth: ", token.AccessToken)
	fmt.Println("Type: ", token.Type())
	fmt.Println("Valid: ", token.Valid())

	// Menggunakan token untuk mendapatkan informasi pengguna dari Google
	// resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info from Google", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	// Anda dapat menggunakan informasi pengguna dari `userInfo` untuk login atau registrasi pengguna di Firebase
	// Misalnya, Anda dapat menggunakan email untuk mencari atau membuat pengguna di Firebase
	email, _ := userInfo["email"].(string)
	fmt.Println("email: ", email)

	user, err := h.authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
		return
	}
	if user == nil {
		params := (&auth.UserToCreate{}).
			Email(email).
			EmailVerified(true).
			Disabled(false)
		user, err = h.authClient.CreateUser(context.Background(), params)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
			return
		}

	}

	// ID token ti pake oleh client untuk signin.
	// https://firebase.google.com/docs/auth/admin/create-custom-tokens
	idToken, err := jwtapp.GenerateJWTToken(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate ID token: %v", err), http.StatusInternalServerError)
		return
	}

	// expiresIn := time.Hour * 24 * 5
	// cookie, err := h.authClient.SessionCookie(r.Context(), idToken, expiresIn)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
	// 	return
	// }
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "session",
	// 	Value:    cookie,
	// 	MaxAge:   int(expiresIn.Seconds()),
	// 	HttpOnly: true,
	// 	Secure:   true,
	// })

	fmt.Println("idToken:", idToken)
	// Set session expiration to 5 days.
	// expiresIn := time.Hour * 24 * 5

	cToken, err := h.authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate CustomToken custome token: %v", err), http.StatusInternalServerError)
		return
	}

	x, err := h.authClient.VerifyIDToken(context.Background(), cToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to VerifyIDToken custome token: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(x)

	sess, _ := midapp.GetSession().Get(r, "session")
	sess.Values["token"] = idToken
	sess.Save(r, w)

	redirectURL := fmt.Sprintf("/welcome?token=%s", url.QueryEscape(cToken))

	// Setelah berhasil login atau registrasi, Anda dapat mengarahkan pengguna ke halaman beranda atau halaman lainnya
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
