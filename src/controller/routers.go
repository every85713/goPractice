package controller

import(
	"fmt"
	"log"
	"html/template"
	"net/http"
    "github.com/julienschmidt/httprouter"
    "model/mongodb"
    //"io/ioutil"
    //"views"
)

var templates = template.Must(template.ParseFiles("profile.gtpl", "register.gtpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, user *mongodb.UserInfo){
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, tmpl+".gtpl", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Confirm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if mongodb.ConfirmUser(ps.ByName("code")) {
		//confirm successfully, redirect to login page
		//maybe wait for few seconds would be better
		http.Redirect(w, r, "/login/", http.StatusFound) // should be a special page says it's done
	}
	fmt.Fprintf(w, "%s, please Confirm Your Account", ps.ByName("name"))
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "It's Index!! Welcome %s", ps.ByName("name"))
}
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//i need i locker for continuosely login failed
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username")) ///attention, need escape back
	password := template.HTMLEscapeString(r.Form.Get("password")) ///
	correct, auth, info := mongodb.UserLogin(w, username, password)
	if !correct {
		http.Redirect(w, r, "/login/wrong", http.StatusFound)
	}if !auth {
		http.Redirect(w, r, "/login/notauth", http.StatusFound)
	}else{
		mongodb.InsertUser(username, password)
		log.Println(username, ": Login success")
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
func Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := mongodb.UserData(ps.ByName("name"))
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%s is not exist", ps.ByName("name"))
	}else{
		renderTemplate(w, "profile", &user)
	}
	//cookie to check whether show edit button or no
}
func Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Number %s Test", ps.ByName("name"))
}
func RegisterPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplate(w, "register", nil)/
}
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username")) ///need escape back?
	password := template.HTMLEscapeString(r.Form.Get("password")) ///
	log.Println("Scheme: ", r.URL.Scheme,"username:", username,"password:", password)
	//lack: same username detect
	mongodb.InsertUser(username, password)
	log.Println(username, ": Register success")
	//should redirect to Register Checking Page
	http.Redirect(w, r, "/register_check/" + username, http.StatusFound)
}
func RegisterCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s, please Confirm Your Account", ps.ByName("code"))
}

func Routing() {
	log.Println("Main run")
	mongodb.SetSessionMode()
	defer log.Println("Main end")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/confirm/:code", Confirm)
	router.GET("/login/", LoginPage)
	router.GET("/login/:error", LoginError)
	router.POST("/login/", Login)
	router.GET("/profile/:name", Profile)
	router.GET("/register", RegisterPage)
	router.GET("/register_check/:code", RegisterCheck) // i need mailchimp
	router.POST("/register", Register)
	router.GET("/test/:number", Test)

	log.Fatal(http.ListenAndServe(":8090", router))
}