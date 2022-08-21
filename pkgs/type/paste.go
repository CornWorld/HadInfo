package _type

type PasteId string

type Paste struct {
	Id     PasteId
	Title  string
	Author string
	Meta   PasteMate
	ItemId ItemId
}

type PasteMate struct {
	CodeLang string
}
