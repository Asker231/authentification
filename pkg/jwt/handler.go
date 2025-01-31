package jwt

import "github.com/golang-jwt/jwt/v5"


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
	res,err := t.SignedString(j.Secret)
	if err != nil{
		return err.Error(),nil
	}
	return res,nil

}