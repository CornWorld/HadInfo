package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hadInfo/db"
	_type "hadInfo/pkgs/type"
	"net/http"
)

func queryItem(itemId _type.ItemId) _type.Item {
	r := db.SqlQuery(db.QueryItem, itemId)
	defer r.Close()

	ret := _type.Item{}

	r.Next()
	if err := r.Scan(&ret.Id, &ret.ItemType, &ret.Value); err != nil || ret.Id != itemId {
		panic(err)
	}
	if v, err := base64.StdEncoding.DecodeString(ret.Value); err != nil {
		ret.Value = string(v)
	}

	return ret
}

func queryPaste(pasteId _type.PasteId) _type.Paste {
	r := db.SqlQuery(db.QueryPaste, pasteId)
	defer r.Close()

	ret := _type.Paste{}
	meta := ""

	r.Next()
	if err := r.Scan(&ret.Id, &ret.Title, &ret.Author, &meta, &ret.ItemId); err != nil || ret.Id != pasteId {
		panic(err)
	}

	err := json.Unmarshal([]byte(meta), &ret.Meta)
	if err != nil {
		panic(err)
	}

	return ret
}

func handleViewPastePage(c *gin.Context) {
	paste := queryPaste(_type.PasteId(c.Param("pasteId")))
	item := queryItem(paste.ItemId)
	if item.ItemType == "text" {
		v, _ := base64.StdEncoding.DecodeString(item.Value)

		c.HTML(http.StatusOK, "view.html", gin.H{
			"title":  paste.Title,
			"author": paste.Author,
			"meta":   paste.Meta,
			"text":   string(v),
		})
	}
}

func handleViewPasteForm(c *gin.Context) {
	paste := queryPaste(_type.PasteId(c.Param("pasteId")))
	item := queryItem(paste.ItemId)
	if item.ItemType == "text" {
		v := base64.StdEncoding.EncodeToString([]byte(item.Value))

		c.JSON(http.StatusOK, gin.H{
			"title":  paste.Title,
			"author": paste.Author,
			"meta":   paste.Meta,
			"text":   v,
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "fallback",
	})
}
