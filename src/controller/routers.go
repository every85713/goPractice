package controller

import(
	"fmt"
	"log"
	"html/template"
	"io/ioutil"
	"net/http"
    "github.com/julienschmidt/httprouter"
    "model/mongodb"
    "model/user"
    //"views"
)

var templates = template.Must(template.ParseFiles("profile.gtpl", "register.gtpl", "login.gtpl", "index.gtpl", "register_check.gtpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}){ ///may explode A_A
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, tmpl+".gtpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderTemplate(w, "index", nil)
	//fmt.Fprintf(w, "It's Index!! Welcome %s", ps.ByName("name"))
}
func Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := user.UserData(ps.ByName("name"))
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%s is not exist", ps.ByName("name"))//renderTMPL profile not found
	}else{
		renderTemplate(w, "profile", &user)
	}
	//cookie to check whether show edit button or no
}
func Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Number %s Test", ps.ByName("name"))
}
func CSS(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "text/css")
	body, err := ioutil.ReadFile("css/" + ps.ByName("filename"))
	if err != nil {
		fmt.Fprintln(w, err)
	}else{
		fmt.Fprintln(w, string(body))
	}
} // make a check list would be safer
func JS(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "application/javascript")
	body, err := ioutil.ReadFile("js/" + ps.ByName("filename"))
	if err != nil {
		fmt.Fprintln(w, err)
	}else{
		fmt.Fprintln(w, string(body))
	}
}

func Routing() {
	log.Println("Main run")
	mongodb.SetSessionMode()
	defer log.Println("Main end")
	
	/*test
	user.MongoTestAdder()
	session.MongoTestAdder()
	test*/
	
	router := httprouter.New()
	router.GET("/", Index)
	
	router.GET("/login", LoginPage)
	router.GET("/login/:error", LoginError)
	router.POST("/login", Login)
	
	router.GET("/profile/:name", Profile)
	
	router.GET("/register", RegisterPage)
	router.POST("/register", Register)
	router.GET("/register_check/:username", RegisterCheck) // i need mailchimp
	router.GET("/confirm/:code", Confirm)
	
	router.GET("/test/:number", Test)
	router.GET("/css/:filename", CSS)
	router.GET("/js/:filename", JS)
	log.Fatal(http.ListenAndServe(":8090", router))
}
