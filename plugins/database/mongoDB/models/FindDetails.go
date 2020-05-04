package models

// FindDetails is used to export the data to the template
type FindDetails struct {
	CollectionNames []string  `json:"collectionNames,omitempty" bson:"collectionNames,omitempty"`
	ContentDetails  *Thread   `json:"contentDetails,omitempty" bson:"contentDetails,omitempty"`
	Comments        []Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	User            *User     `json:"user,omitempty" bson:"user,omitempty"`
	LogInUser       *User     `json:"loginuser,omitempty" bson:"loginuser,omitempty"`
}
