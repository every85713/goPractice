package mongodb

import (
    //"fmt"
	"log"
	"time"
	//"regexp"
	"model/cryption"
	"model/session"
	"gopkg.in/mgo.v2/bson"
)
type authenticate struct{
	Done bool
	Code string
	Lifetime time.Time
}
type UserInfo struct{
	Username string
	Password string
	//Nickname string
	//Intro string
	Check authenticate
	//Cookie string
}

const(
	database = "test"
	collection = "user"
) 

func InsertUser(username string, password string) { 
	Escape(&username)
	Escape(&password)
	log.Println("now prepare insert")
	session := GetmgoSession()
	defer session.Close()
	cryPassword := cryption.EncryptionByString(password, username)
	auth := authenticate{false, cryption.RandomBytes(), time.Now()}
	user := UserInfo{username, cryPassword, auth}
	c := session.DB(database).C(collection)
	err := c.Insert(user)
	if err != nil{
		log.Println(err)
	}
}

func UserData(username string) (UserInfo, error){ // need filter useless info
	Escape(&username)
	session := GetmgoSession()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username}).Select(bson.M{"password":0}).One(&result)
	if err != nil{
		return false, result
	}
	log.Println(result)
	return result, err
	//find
}

func UserLogin(w http.ResponseWriter, username, password string) (bool, bool, UserInfo){
	Escape(&username)
	Escape(&password)
	session := GetmgoSession()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username,"password":password}).Select(bson.M{"_id":0}).One(&result)
	if err != nil || password != cryption.Decryption(result.Password) {
		log.Println("Wrong Password or No this Account")
		return false, false, result
	}if result.Check.Done == false { //neeeeeed morerererere parameters
		log.Println("Authentication Undone")
		return true, false, result
	} else{
		session.PutCookie(w, "s_id", session.MakeSession(r, username))
		return true, true, result
	}
}

func ConfirmUser(code string) bool{
	Escape(&code)
	dbsession := mongodb.GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	result = UserInfo{}
	err := c.Find(bson.M{"Code":code}).Select(bson.M{"_id":0}).One(&result)
	if err != nil || time.Since(result.Check.Lifetime).Hours > 24{
		//need a new link
		return false // outdated Confirm Link
	}else{
		result.Check.Done = true
		target := bson.M{"Code": code}
		change := bson.M{"$set": &result}
		err := c.Update(target, change)
		return true
	}
}

func AutoLogin(cookie string) (string, bool){
	Escape(&cookie)
	username, boo := session.CookieCheck(r, cookie)
	if boo {
		return username, boo
	}else{
		return "", boo
	}
	//find Cookie
}//*/