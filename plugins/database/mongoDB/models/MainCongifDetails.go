package models

// MainCongifDetails is used to export data to template
type MainCongifDetails struct {
	CollectionNames []string `json:"collectionNames,omitempty" bson:"collectionNames,omitempty"`
	ContentDetails  []Thread `json:"contentDetails,omitempty" bson:"contentDetails,omitempty"`
	User            *User    `json:"user,omitempty" bson:"user,omitempty"`
}
