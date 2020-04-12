package comments

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	m "kRtrima/plugins/database/mongoDB/models"
	//	"kRtrima/plugins/database/mongoDB"
	//    "html/template"
	"log"
	"net/http"
	"regexp"
)

//func test() {
//	lookupStage := bson.D{{"$lookup", bson.D{{"from", "podcasts"}, {"localField", "podcast"}, {"foreignField", "_id"}, {"as", "podcast"}}}}
//	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$podcast"}, {"preserveNullAndEmptyArrays", false}}}}
//
//	showLoadedCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
//	if err != nil {
//		panic(err)
//	}
//	var showsLoaded []bson.M
//	if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
//		panic(err)
//	}
//	fmt.Println(showsLoaded)
//}

func Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		log.Fatalln(err)
	}

	//    fmt.Println(request.Form["text"][0])
	//    fmt.Println(request.Form["author"][0])
	//
	//    fmt.Println(docID)

	newItem := m.Comment{
		Comment: request.Form["text"][0],
		Author:  request.Form["author"][0],
		Thread:  docID,
	}

	log.Println(m.AddItem(newItem, m.Comments))

	http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)

}
