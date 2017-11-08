package session

import(
	"log"
	"net/http"
	"time"
	"model/cryption"
	"model/mongodb"
	"gopkg.in/mgo.v2/bson"
	"crypto/rand"
	"encoding/base64"
)

/*
login->make session->(clean old session)->pass cookie(SessionID, )
view site->check cookie(SessionID)->auto login(?)

//*/
type Session struct{
	SessionID string
	Username string
	IPaddress string
	Useragent string
	Time time.Time
}

const(
	database = "test"
	collection = "session"
)
func insertSession(s Session){
	dbsession := GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	//log.Println("now insert", ?)
	err := c.Insert(&s)
	if err != nil{
		log.Println("insert session: ",err)
	}
}
func deleteSession(id string) error { // not yet check
	dbsession := mongodb.GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	err := c.Remove(bson.M{"sessionid": id})
	return err
}
func getSession(id string) (Session, error) {
	c_string := base64.URLEncoding.EncodeToString(cryption.DecryptionToByte(id, "sess10n"))
	
	dbsession := mongodb.GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	result := Session{}
	err := c.Find(bson.M{"sessionid":c_string}).Select(bson.M{"_id":0}).One(&result)
	log.Println(result)
	return result, err
}
func updateSession(s Session) err{ 
	dbsession := mongodb.GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	target := bson.M{"sessionid": s.SessionID}
	change := bson.M{"$set": &s}
	err := c.Update(target, change)
	return err
}
func MakeSession(r *http.Request, user string) string {
	session := Session{}
	session.Username = user
	session.IPaddress = r.RemoteAddr /// fix: I want a deeper IP
	session.Useragent = r.UserAgent()
	session.Time = time.Now()
	for {
		b := cryption.RandomBytes()
		s_id := base64.URLEncoding.EncodeToString(b)
		s_id_encrypt := base64.URLEncoding.EncodeToString(cryption.Encryption(b, "sess10n"))
		_, duplicate := getSession(s_id)
		if duplicate != nil {
			session.SessionID = s_id
			break
		}
	}
	insertSession(session)
	return s_id_encrypt
}

func CookieCheck(r *http.Request, id string) (string, bool) {
	//check timestamp outdated
	s, err := getSession(id)
	if r.UserAgent() != s.Useragent {
		//someone log in from other place
		err := deleteSession(s.SessionID)
		if err != nil { /// need check something like its no result or other
			log.Println(err)
		}
		return false
	}
	if time.Since(s.Time).Hours > 24 {
		s.Time = time.Now()
		updateSession(s)
	}
	return s.Username, true
}

func PutCookie(w http.ResponseWriter, name, value string) { // Secure, HttpOnly need to set
	/*expiration := time.Now()
	expiration = expiration.AddDate(7, 0, 0)//*/
	//cookie := http.Cookie{Name: name, Value: value}
	http.SetCookie(w, &http.Cookie{Name: name, Value: value})
}