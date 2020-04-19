package models

type FindDetails struct {
	CollectionNames []string   `json:"CollectionNames,omitempty" bson:"CollectionNames,omitempty"`
	ContentDetails  *Thread    `json:"ContentDetails,omitempty" bson:"ContentDetails,omitempty"`
	Comments        []*Comment `json:"Comments,omitempty" bson:"Comments,omitempty"`
    User            *User      `json:"user,omitempty" bson:"user,omitempty"`
}

//type FindComment struct {
//	CollectionNames []string `json:"CollectionNames,omitempty" bson:"CollectionNames,omitempty"`
//	ContentDetails  *Comment  `json:"ContentDetails,omitempty" bson:"ContentDetails,omitempty"`
//}
