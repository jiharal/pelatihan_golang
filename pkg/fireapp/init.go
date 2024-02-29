package fireapp

import (
	"context"
	"log"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var (
	firebaseClient *auth.Client
	initOnce       sync.Once
	initComplete   chan struct{}
)

// https://console.cloud.google.com/apis/credentials/oauthclient/391517908412-1rrcr23aouc1j0cnausosgp9kadssdit.apps.googleusercontent.com?authuser=0&hl=en&orgonly=true&project=tlabdev&supportedpurview=organizationId
// Konfigurasi OAuth 2.0 untuk login menggunakan akun Google
var googleOAuthConfig = &oauth2.Config{
	ClientID:     "391517908412-1rrcr23aouc1j0cnausosgp9kadssdit.apps.googleusercontent.com", // Ganti dengan ID klien OAuth Google Anda
	ClientSecret: "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",                                      // Ganti dengan rahasia klien OAuth Google Anda
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}

func initFirebase() {
	opt := option.WithCredentialsFile("./config/tlabdev-firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	firebaseClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase Auth client: %v", err)
	}
	close(initComplete)
}

func GetOAuthConfig() *oauth2.Config {
	return googleOAuthConfig
}

func GetAuth() *auth.Client {
	initOnce.Do(func() {
		initComplete = make(chan struct{})
		go initFirebase()
	})
	<-initComplete
	return firebaseClient
}
