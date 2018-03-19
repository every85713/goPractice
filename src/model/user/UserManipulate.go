package user

import (
    //"fmt"
	"log"
	"time"
	"regexp"
	"model/cryption"
	"gopkg.in/mgo.v2/bson"
	"model/mongodb"
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

func insertUser(username string, password string) { 
	mongodb.Escape(&username)
	//mongodb.Escape(&password) no need
	code := cryption.RandomBytesBase64()
	mongodb.Escape(&code)
	
	log.Println("now prepare insert")
	session := mongodb.GetmgoSession()
	defer session.Close()
	cryPassword := cryption.EncryptionByString(password, username)
	auth := authenticate{false, code, time.Now()}
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
	if  err == nil{
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
		insertUser(username, password)
	}
	return result
}

func ConfirmUser(code string) (bool, string){
	mongodb.Escape(&code)
	session := mongodb.GetmgoSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"check.done":false,"check.code":code}).Select(bson.M{"_id":0}).One(&result) // "password": 0????
	if err != nil || time.Since(result.Check.Lifetime).Hours() > 72 {
		newConfirmCode(result.Username)
		return false, "" // outdated Confirm Link
	}else{
		target := bson.M{"check.code": code}
		change := bson.M{"$set": bson.M{"check": authenticate{true, "", time.Now()}}}
		err := c.Update(target, change)
		if err != nil{
			log.Println(err)//something wrong
			return false, ""
		}
		return true, result.Username
	}
}

func newConfirmCode(username string) string {
	//mongodb.Escape(&username) //no need
	auth := authenticate{false, cryption.RandomBytesBase64(), time.Now()}
	
	session := mongodb.GetmgoSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	target := bson.M{"username": username}
	change := bson.M{"$set":bson.M{"check": auth}}
	err := c.Update(target, change)
	if err != nil{
		log.Printf("newConfirmCode ERROR, user: %s", username)
		log.Println(err)
	}
	return auth.Code
}

func GetConfirmCode(username string) (bool, string){
	code := ""
	mongodb.Escape(&username)
	//Escape(&password)
	session := mongodb.GetmgoSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username, "check.done":false}).Select(bson.M{"_id":0,"password":0}).One(&result)
	code = result.Check.Code
	if err != nil { //neeeeeed morerererere parameters
		log.Println("GetConfirmCode Fail")
		log.Println(err)
		return false, "Go to 404?"
	} else{
		if time.Since(result.Check.Lifetime).Hours() > 72{
			code = newConfirmCode(result.Username)
		}
		return true, code
	}
}

func UserLogin(username, password string) (bool, bool, UserInfo){
	mongodb.Escape(&username)
	//mongodb.Escape(&password)
	session := mongodb.GetmgoSession()
	defer session.Close()
	c := session.DB(database).C(collection)
	result := UserInfo{}
	err := c.Find(bson.M{"username":username}).Select(bson.M{"_id":0}).One(&result)
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

func UserData(username string) (UserInfo, error){ // need filter useless info
	mongodb.Escape(&username)
	session := mongodb.GetmgoSession()
	defer session.Close()
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

func MongoTestAdder(){
	mongodb.TAdder()
	log.Println(mongodb.GetT())
}
