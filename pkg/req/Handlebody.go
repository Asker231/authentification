package req

import "net/http"


func HandleBody[T any](w http.ResponseWriter,r *http.Request)(body *T,err error){
	//decode
	//validation
	return nil,err
}