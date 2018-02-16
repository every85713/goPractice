package mongodb

import (
    //"fmt"
	"log"
	"time"
	"regexp"
	"model/cryption"
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

//need check what else elements can add into this
var vaildRegister = regexp.MustCompile("^([a-zA-Z0-9_]+)$")
//want to add !@#$$%^&*() things

func InsertUser(username string, password string) { 
	Escape(&username)
	//Escape(&password)
	log.Println("now prepare insert")
	session := GetmgoSession()
	defer session.Close()
	cryPassword := cryption.EncryptionByString(password, username)
	auth := authenticate{false, string(cryption.RandomBytes()), time.Now()}
	user := UserInfo{username, cryPassword, auth}
	c := session.DB(database).C(collection)
	err := c.Insert(user)
	if err != nil{
		log.Println(err)
	}
}

func UserRegister(username, password, password_two string) (string){
	//check availability of these tring
	result := ""
	userValid := vaildRegister.FindStringSubmatch(username)
	passValid := vaildRegister.FindStringSubmatch(password)
	
	_, err := UserData(username)
	if  err != nil{
		result += "This username has been used. "
	}
	if password != password_two {
		result += "Two passwords are not the same. "
	}
	if len(username) > 25 || len(username) < 5 || userValid == nil {
		result += "Invalid Username. Please follow the rule. "
	}
	if len(password) > 30 || len(password) < 8 || passValid == nil {
		result += "Invalid Password. Please follow the rule. "
	}
	if result == "" {
		InsertUser(username, password)
	}
	return result
}

func UserData(username string) (UserInfo, error){ // need filter useless info
	Escape(&username)
	session := GetmgoSession()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username}).Select(bson.M{"password":0}).One(&result)
	if err != nil{
		return result, err
	}
	//Escape(&result.Username) 
	log.Println(result)
	return result, err
	//find
}

func UserLogin(username, password string) (bool, bool, UserInfo){
	Escape(&username)
	//Escape(&password)
	session := GetmgoSession()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username,"password":password}).Select(bson.M{"_id":0}).One(&result)
	if err != nil || password != cryption.Decryption(result.Password, username) {
		log.Println("Wrong Password or No this Account")
		return false, false, result
	}
	if result.Check.Done == false { //neeeeeed morerererere parameters
		log.Println("Authentication Undone")
		return true, false, result
	} else{
		return true, true, result
	}
}

func ConfirmUser(code string) bool{
	Escape(&code)
	dbsession := GetmgoSession()
	defer dbsession.Close()
	c := dbsession.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"Code":code}).Select(bson.M{"_id":0}).One(&result)
	if err != nil || time.Since(result.Check.Lifetime).Hours() > 24 {
		//need a new link
		return false // outdated Confirm Link
	}else{
		result.Check.Done = true
		target := bson.M{"Code": code}
		change := bson.M{"$set": &result}
		err := c.Update(target, change)
		if err != nil{
			log.Println(err)//something wrong
			return false
		}
		return true
	}
}
