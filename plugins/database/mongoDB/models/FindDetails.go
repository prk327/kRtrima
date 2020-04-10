package models

type FindDetails struct {
	CollectionNames []string `json:"CollectionNames,omitempty" bson:"CollectionNames,omitempty"`
	ContentDetails  *Thread  `json:"ContentDetails,omitempty" bson:"ContentDetails,omitempty"`  
}