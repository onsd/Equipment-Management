package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB
var userList = []string{"taka", "asasigure", "sasamoto", "tt"} // -1:部活所有
var itemList = []string{"ラズベリーパイ１号", "ラズベリーパイ２号", "ラズベリーパイ３号", "C++の本", "Pythonの本"}

type Lending struct {
	ItemId    int
	UserId    int
	CreatedAt time.Time
}

type Item struct {
	ItemId    int    `gorm:"primary_key"`
	ItemName  string `sql:"type:varchar(32);"`
	Status    bool
	OriginId  int
	CreatedAt time.Time
	UpdatedAt time.Time //CheckoutDate
}

type User struct {
	UserId   int    `gorm:"primary_key"`
	UserName string `sql:"type:varchar(32)"`
}

type replyAllList struct {
	ItemId   string
	ItemName string
	Status   bool
	UserName string
}

type replyLendingList struct {
	ItemName  string
	UserName  string
	CreatedAt time.Time
}

type lendingHistory struct {
	ItemId    int
	UserId    int
	CreatedAt time.Time
}

func init() {
	conn, err := gorm.Open("mysql", "root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = conn

	if !db.HasTable("items") {
		db.AutoMigrate(&Item{})
		addItem()
	}
	if !db.HasTable("users") {
		db.AutoMigrate(&User{})
		NONE := User{
			UserId:   -1,
			UserName: "部活所有物品",
		}
		db.Create(NONE)
		addUser()
	}
	if !db.HasTable(&Lending{}) {
		db.AutoMigrate(&Lending{})
	}
	if !db.HasTable(&lendingHistory{}) {
		db.AutoMigrate(&lendingHistory{})
	}

}
func addUser() {
	for i := 0; i < len(userList); i++ {
		newUser := User{
			UserName: userList[i],
		}
		db.Create(&newUser)
	}
}

func addItem() {
	for i := 0; i < len(userList); i++ {
		newUser := Item{
			ItemName: itemList[i],
			Status:   true,
			OriginId: -1,
		}
		db.Create(&newUser)
	}
}

type PostRepository struct{}

func NewPostRepository() PostRepository {
	return PostRepository{}
}

func (m PostRepository) PostNewItem(itemName string, userId int) bool {
	newItem := Item{
		ItemName: itemName,
		Status:   true, //未貸出
		OriginId: userId,
	}
	db.Create(&newItem)
	return true
}
func (m PostRepository) PostNewUser(userName string) *User {
	newUser := User{
		UserName: userName,
	}
	db.Create(&newUser)
	db.Last(&newUser)
	return &newUser
}

func (m PostRepository) GetItemByUID(uid int) *Item {
	item := Item{}
	db.Find(&item, "uid=?", uid)
	return &item
}

func (m PostRepository) GetBorrowedItemAll() *[]replyAllList {
	var replies []replyAllList
	db.Table("").Select("items.item_name,items.status,users.user_name").Joins("join users on items.user_id = users.user_id").Where("items.status = 1").Scan(&replies)

	//db.Find(&items)
	return &replies
}
func (m PostRepository) GetItemAll() *[]replyAllList {
	var replies []replyAllList
	db.Table("items").Select("items.item_id,items.item_name,items.status,users.user_name").Joins("join users on items.origin_id = users.user_id").Scan(&replies)
	return &replies
}
func (m PostRepository) GetUserAll() *[]User {
	var users []User
	db.Find(&users)

	return &users
}
func (m PostRepository) GetLendingsAll() *[]replyLendingList {
	var lendings []replyLendingList
	db.Table("lendings").Select("items.item_name,users.user_name,lendings.created_at").Joins("left outer join items on items.item_id = lendings.item_id").Joins("left join users on users.user_id = lendings.user_id").Scan(&lendings)
	return &lendings
}
func (m PostRepository) GetHistoriesAll() *[]lendingHistory {
	var histories []lendingHistory
	db.Table("lending_histories").Select("items.item_name,users.user_name,lendings.created_at").Joins("left outer join items on items.item_id = lendings.item_id").Joins("left join users on users.user_id = lendings.user_id").Scan(&histories)
	return &histories
}
func (m PostRepository) UpdateItemByUID(borrow bool, uid int, userID int) interface{} { //if status true => borrow ,false => return
	lending := Lending{}
	if borrow {
		history := lendingHistory{}
		lending.UserId = userID
		lending.ItemId = uid
		history.UserId = userID
		history.ItemId = uid
		db.Create(&history)
		db.Create(&lending)
	} else {
		//db.Delete()ではlendingsの中身が全て消えていたため断念
		db.Exec("DELETE FROM lendings WHERE item_id = ?", uid) //lendingsテーブルにuidは一つしか無いはずなのでこれで対応

	}
	db.Find(&lending, "ItemId=?", uid)

	return &lending
}
