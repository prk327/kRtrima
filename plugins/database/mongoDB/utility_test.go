package mongoDB

import (
	"strings"
	"testing"
)

var p2 = Thread{"City At Night", "/resources/images/CityNight.jpg"}

func TestAddItem(t *testing.T) {
	lm := &Msg
	//    connect to db
	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	//	logger.Println(msg)
	//    conect to collection
	*lm, Collection = Cnt_Collection("campground", DB)
	//	logger.Println(msg)
	want := "Inserted a single document"
	if got := AddItem(p2, Collection); strings.Split(got, ":")[0] != want {
		t.Errorf("AddItem() = %q, want %q", got, want)
	}
}

func TestShowCollectionNames(t *testing.T) {
	lm := &Msg
	//    connect to db
	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	//	logger.Println(msg)
	//    conect to collection
	*lm, Collection = Cnt_Collection("Thread", DB)
	//	logger.Println(msg)
	want := "Thread"
	if got := ShowCollectionNames(DB); got[0] != want {
		t.Errorf("ShowCollectionNames() = %q, want %q", got, want)
	}
}

func TestFindAllItem(t *testing.T) {
	lm := &Msg
	//    connect to db
	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	//	logger.Println(msg)
	//    conect to collection
	*lm, Collection = Cnt_Collection("Thread", DB)
	//	logger.Println(msg)
	want := "Circuit"
	if got := FindAllItem(1, Collection); got[0].Name != want {
		t.Errorf("FindAllItem() = %q, want %q", got, want)
	}
}
