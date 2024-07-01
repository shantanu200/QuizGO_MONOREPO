package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Text        string               `bson:"text" json:"text"`
	Description string               `bson:"description" json:"description"`
	Topic       []primitive.ObjectID `bson:"topics" json:"topics"`
	MediaFile   []string             `bson:"mediaFiles" json:"mediaFiles"`
	Option      []Option             `bson:"options" json:"options"`
	Answer      int32                `bson:"answer" json:"answer"`
	Solution    Solution             `bson:"solution" json:"solution"`
}

type Option struct {
	ID        int32  `bson:"id" json:"id"`
	Text      string `bson:"text" json:"text"`
	MediaFile string `bson:"mediaFile" json:"mediaFile"`
}

type Solution struct {
	Text      string `bson:"text" json:"text"`
	MediaFile string `bson:"mediaFile" json:"mediaFile"`
	Note      string `bson:"note" json:"note"`
}
