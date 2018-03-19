package controller

import(
	"log"
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"model/user"
	"model/session"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//i need i locker for continuosely login failed
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username")) ///attention, need escape back
	password := template.HTMLEscapeString(r.Form.Get("password")) ///
	correct, auth, _ := user.UserLogin(username, password)
	if !correct {
		http.Redirect(w, r, "/login/wrong", http.StatusFound)
	}else if !auth {
		http.Redirect(w, r, "/login/notauth", http.StatusFound)
	}else{
		//mongodb.InsertUser(username, password)
		log.Println(username, ": Login success")
		if r.Form.Get("keep") == "true"{ //
		//https://stackoverflow.com/questions/34879373/golang-how-to-get-all-the-value-of-checkbox-r-formvalue-is-not-working
			session.PutCookie(w, "s_id", session.MakeSession(r, username))
		}
		http.Redirect(w, r, "/", http.StatusFound)//want to redirect to previous page
		//using GET like ?continue=url to do it
		//fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path) 
	}
}
func LoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderTemplate(w, "login", nil)
}
func LoginError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderTemplate(w, "login", nil)
}

