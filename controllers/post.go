package controllers

import (
	"github.com/onsd/EquipmentManagement/models"
)

type Post struct{}

func NewPost() Post {
	return Post{}
}

func (c Post) PostNewItem(itemName string ,originUserId int) bool {
	repo := models.NewPostRepository()
	ok := repo.PostNewItem(itemName,originUserId)
	return ok
}

func (c Post) GetItemByUId(uid int) interface{} {
	repo := models.NewPostRepository()
	item := repo.GetItemByUID(uid)

	return item
}
func (c Post) PostNewUser(userName string) interface{} {
	repo := models.NewPostRepository()
	user := repo.PostNewUser(userName)
	return user
}

func (c Post) GetItemAll() interface{} {
	repo := models.NewPostRepository()
	items := repo.GetItemAll()

	return items
}
func (c Post) GetUserAll() interface{} {
	repo := models.NewPostRepository()
	users := repo.GetUserAll()
	return users
}
func (c Post) GetLendingAll() interface{}{
	repo := models.NewPostRepository()
	lendings := repo.GetLendingsAll()
	return lendings
}
func (c Post) GetHistoriesAll() interface{}{
	repo := models.NewPostRepository()
	histories := repo.GetHistoriesAll()
	return histories
}
func (c Post) UpdateItemByUID(borrow bool, uid int, userid int) interface{} {
	repo := models.NewPostRepository()
	result := repo.UpdateItemByUID(borrow, uid, userid)
	if result == nil {
		return nil
	} else {
		return result
	}
}
