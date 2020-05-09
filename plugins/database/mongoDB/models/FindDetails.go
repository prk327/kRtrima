package models

// FindDetails is used to export the data to the template
type FindDetails struct {
	CollectionNames []string  `json:"collectionnames,omitempty" bson:"collectionnames,omitempty"`
	ContentDetails  *Thread   `json:"contentdetails,omitempty" bson:"contentdetails,omitempty"`
	Comments        []Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	SingleComment   *Comment  `json:"singlecomment,omitempty" bson:"singlecomment,omitempty"`
	User            *User     `json:"user,omitempty" bson:"user,omitempty"`
	LogInUser       *User     `json:"loginuser,omitempty" bson:"loginuser,omitempty"`
}
