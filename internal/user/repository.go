package user

import (
	"fmt"

	"gorm.io/gorm"
)




type UserRepoSitory struct{
	database *gorm.DB
}
func NewRepoUser(dataBase *gorm.DB)*UserRepoSitory{
		return &UserRepoSitory{
			database: dataBase,
		}
}


func(urepo *UserRepoSitory)CreateUser(user *User)(*User,error){
	result := urepo.database.Create(user)
	if result.Error != nil{
		return nil,result.Error
	}
	return user,nil
}

func(urepo *UserRepoSitory)DeleteUserById(id int)(error){
	var user User
	result := urepo.database.Delete(&user,id)
	if result.Error != nil{
		return result.Error
	}
	return result.Error
}

func(urepo *UserRepoSitory)FindByEmail(email string)(*User){
	var payload User
	result := urepo.database.First(&payload,"email = ?",email)
	if result.Error != nil{
		fmt.Println(result.Error.Error())
		return nil
	}
	return &payload

}