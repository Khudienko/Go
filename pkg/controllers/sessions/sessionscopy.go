package sessions
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/gorilla/sessions"
//)
//
//var (
//	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
//	key = []byte("super-secret-key")
//	store = sessions.NewCookieStore(key)
//)
//
//func Secret(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Check if user is authenticated
//	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
//		http.Error(w, "Forbidden", http.StatusForbidden)
//		return
//	}
//
//	// Print secret message
//	fmt.Fprintln(w, "The cake is a lie!")
//}
//
//func Login(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Authentication goes here
//	// ...
//
//	// Set user as authenticated
//	session.Values["authenticated"] = true
//	session.Save(r, w)
//}
//
//func Logout(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Revoke users authentication
//	session.Values["authenticated"] = false
//	session.Save(r, w)
//}
