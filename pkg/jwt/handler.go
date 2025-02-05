package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)


type JWTData struct{
	Secret string
}

func NewJWTInit(secret string)*JWTData{
	return &JWTData{
		Secret: secret,
	}
}

func(j *JWTData)CreateJWT(email string)(string,error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
	})
	res,err := t.SignedString([]byte(j.Secret))
	if err != nil{
		return err.Error(),nil
	}
	fmt.Println(res)
	return res,nil
}

func(j *JWTData)ParseJWT(token string)(bool,string){
	t ,err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
	    return j.Secret,nil		
	})
	if err != nil{
		return false,err.Error()
	}
	res := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, res.(string)
}