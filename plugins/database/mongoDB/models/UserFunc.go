package models

import (
	"context"
		"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "time"
    "net/http"
    "regexp"
//    "strings"
	//	"log"
)

func FindUserByID(res1 string, collection *mongo.Collection) (result *User, err error) {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	err = collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		Logger.Fatalln(err)
	}

	return
}

func FindUser(key string, value interface{}, collection *mongo.Collection) (result []bson.M, err error) {
    filterCursor, err := collection.Find(context.TODO(), bson.M{key: value})
    if err != nil {
        Logger.Fatal(err)
    }
    if err = filterCursor.All(context.TODO(), &result); err != nil {
        Logger.Fatal(err)
    }
	return
}

func FindAllUser(limit int64, collection *mongo.Collection) []*User {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			Logger.Fatalln(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		Logger.Fatalln(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}


func FindAllUserByID(key string, res1 string, collection *mongo.Collection) []*User {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// Pass these options to the Find method
	findOptions := options.Find()
	//	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.M{key: docID}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			Logger.Fatalln(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		Logger.Fatalln(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}


// Create a new session for an existing user
func (user *User) CreateSession() (ssid primitive.ObjectID, uuid string, err error) {
    
     // Create a struct type to handle the session for login
    statement :=  Session {
        Salt: CreateUUID(),
        Email: user.Email,
        UserId: user.ID,
        CreatedAt: time.Now(),
}
 
    ssid, err = AddItem(statement, Sessions)
    if err != nil {
			Logger.Fatalln(err)
		}
    uuid = statement.Salt
    
	return
}



//getting the user by session uuid
func GetUserbyUUID(cookieName string,writer http.ResponseWriter, request *http.Request) (user *User, err error){
    cookie, err := request.Cookie(cookieName)
    if err != http.ErrNoCookie{
        fmt.Println("Cookie found for user by session uuid ")
        session, errr := Findmodel("salt", cookie.Value, Sessions)
        if errr != nil {
            fmt.Println("Cannot Find session in getuserby uuid")
//            Logger.Fatalln(err)
//            http.Redirect(writer, request, "/login", 401)
            return
		}
        fmt.Println("Valid Session Was Found in get user by uuid")
        re := regexp.MustCompile(`"(.*?)"`)
        rStr := fmt.Sprintf(`%v`, session[0]["userid"])
        res1 := re.FindStringSubmatch(rStr)[1]
        user, err = FindUserByID(res1, Users)
        if err != nil {
            fmt.Println("Cannot Find user with uuid")
//            Logger.Fatalln(err)
//            http.Redirect(writer, request, "/login", 401)
            return
        }
//        fmt.Println(user)
        return 
    } else{
            fmt.Println("No Cookie was found with Name in get user by uuid")
//            Logger.Fatalln(err)
//            http.Redirect(writer, request, "/login", 401)
            return
    }
    return 
}
