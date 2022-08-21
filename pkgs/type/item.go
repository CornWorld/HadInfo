package _type

type ItemId string
type ItemType string

const (
	ItemTypeText ItemType = "text"
)

type Item struct {
	Id       ItemId
	ItemType ItemType
	Value    string
}
