package models

type MainCongifDetails struct {
	CollectionNames []string  `json:"CollectionNames,omitempty" bson:"CollectionNames,omitempty"`
	ContentDetails  []*Thread `json:"ContentDetails,omitempty" bson:"ContentDetails,omitempty"`
    User            *User      `json:"User,omitempty" bson:"User,omitempty"`
}
