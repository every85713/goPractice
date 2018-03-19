package controller

import(
	//"fmt"
	"log"
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"model/user"
)

type user_confirm struct{
	Username string
	//Nickname string
	//Intro string
	Code string
	//Cookie string
}

func RegisterPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplate(w, "register", nil)//
}
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username")) ///need escape back?
	password := template.HTMLEscapeString(r.Form.Get("password")) ///
	password_two := template.HTMLEscapeString(r.Form.Get("password_two")) ///
	//log.Println("Scheme: ", r.URL.Scheme,"username:", username,"password:", password)
	//lack: same username detect
	error_msg := user.UserRegister(username, password, password_two)
	if error_msg != ""{
		http.Redirect(w, r, "/register/?error=" + error_msg, http.StatusFound)
	}else{
		log.Println(username, ": Register success")
		//should redirect to Register Checking Page
		http.Redirect(w, r, "/register_check/" + username, http.StatusFound)
	}
}
func RegisterCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	success, code := user.GetConfirmCode(ps.ByName("username"))
	if success == false{
		http.Redirect(w, r, "/index/", http.StatusFound) // should go 404
	}else{
		data := user_confirm{ps.ByName("username"), code}
		renderTemplate(w, "register_check", data)
	}
}
func Confirm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	succ, _ := user.ConfirmUser(ps.ByName("code")) 
	if succ {
		//confirm successfully, redirect to login page
		//maybe wait for few seconds would be better
		http.Redirect(w, r, "/login/", http.StatusFound) // should be a special page says it's done
	}else{
		http.Redirect(w, r, "/index/", http.StatusFound) 
		// should have a page said the code is expired or invalid
	}
}

