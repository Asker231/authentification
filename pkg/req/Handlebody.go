package req

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
)


func HandleBody[T any](w http.ResponseWriter,r *http.Request)(*T, error){
	//decode
	result,err :=  DecoderBody[T](r)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err

	}
	//validation
	err = Validation(result)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return result,nil
}

func Validation[T any](strct T)error{
	v := validator.New()
	err := v.Struct(strct)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	return err
}

func DecoderBody[T any](r *http.Request)( *T, error){
		var payload T
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil{
			fmt.Println(err.Error())
			return nil,err
		}
		return &payload,nil
}