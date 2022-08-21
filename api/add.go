package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hadInfo/db"
	"hadInfo/pkgs"
	"hadInfo/pkgs/type"
	"net/http"
)

func addItem(item *_type.Item) bool {
	if item.Id == "" {
		for {
			item.Id = _type.ItemId(pkgs.RandomString(32))
			if ok := db.SqlQueryExist(db.QueryItem, item.Id); !ok {
				break
			}
		}
	}

	v := base64.StdEncoding.EncodeToString([]byte(item.Value))
	db.SqlExec(db.InsItem, item.Id, item.ItemType, v)

	return true
}

func addPaste(paste *_type.Paste) bool {
	if paste.Id == "" {
		for {
			paste.Id = _type.PasteId(pkgs.RandomString(16))
			if !db.SqlQueryExist(db.QueryItem, paste.Id) {
				break
			}
		}
	}

	meta, err := json.Marshal(paste.Meta)
	if err != nil {
		panic("json marshal error")
	}
	num, err := db.SqlExec(db.InsPaste, paste.Id, paste.Title, paste.Author, string(meta),
		paste.ItemId).RowsAffected()
	if err == nil && num == 1 {

		return true
	}

	return false
}

func handleAddForm(c *gin.Context) {
	item := _type.Item{
		ItemType: _type.ItemTypeText,
		Value:    c.PostForm("content"),
	}
	// TODO: verify the input
	v, err := base64.StdEncoding.DecodeString(item.Value)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "cannot decode the item content",
		})
		return
	}
	item.Value = string(v)

	if ok := addItem(&item); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot save item",
		})
		return
	}

	p := _type.Paste{
		Title:  c.DefaultPostForm("title", "Untitled paste"),
		Author: c.DefaultPostForm("author", "unknown"),
		Meta: _type.PasteMate{
			CodeLang: c.DefaultPostForm("lang", "plain"),
		},
		ItemId: item.Id,
	}
	if ok := addPaste(&p); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot save paste",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "done",
		"pasteId": p.Id,
	})
}
