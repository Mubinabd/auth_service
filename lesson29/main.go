package main

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	user     = "sardorbek"
	dbname   = "postgres"
	password = "1111"
	port     = 5432
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Book struct {
	Id        int
	Name      string
	Page      int
	StudentId int
}

// func (Book) TableName() string {
// 	return "book"
// }

func main() {

	dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		host, user, dbname, password, port)

	db, err := gorm.Open(postgres.Open(dbCon), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// err = db.AutoMigrate(&Product{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	// select * from products where id = 1 order by id limit 1
	db.First(&product, "code = ? and price = ?", "D42", 100) // find product with integer primary key
	fmt.Println(product)

	var b Book

	// select * from products where id = 1 limit 1
	// db.Take(&b)
	// fmt.Println(b)

	// db = db.Last(&b, "fds= ?", 43)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// fmt.Println(b, res.Error)
	// var res []Book

	// db = db.Where(&Book{Name: "efhuef", Page: 103}).Find(&res)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// fmt.Println(res)

	db.Where(&Book{ Page: 293}).Model(&b).Update("student_id", 2)
  // Update user
//   updatedUser := models.User{
//     ID:       "*****", // Don't forget to replace it to yours
//     Name:     "George",
//     Username: "gorge",
//     Password: "newpass",
//     Phone:    "00094223006",
//   }
//   um.Update(&updatedUser)
//   cf.CheckErr(db.Error)
//   fmt.Println("User updated")
	// db.Delete(&Book{}, 3)
}
// func (um *UserManager) Update(user *models.User) error {
// 	return um.db.Save(user).Error
//   }