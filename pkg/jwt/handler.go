package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type (
	JWT struct{
		Secret string
	}
	JWTData struct{
		Email string
	}
)


func NewJWTInit(secret string)*JWT{
	return &JWT{
		Secret: secret,
	}
}

func(j *JWT)CreateJWT(data JWTData)(string,error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":data.Email,
	})
	res,err := t.SignedString([]byte(j.Secret))
	if err != nil{
		return err.Error(),nil
	}
	return res,nil
}

func(j *JWT)ParseJWT(token string)(bool,*JWTData){
	t,err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret),nil
	})
	if err != nil{
		return false,nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JWTData{
		Email: email.(string),
	}
}