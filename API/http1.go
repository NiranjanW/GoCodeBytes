package API

import (
	"net/http"
	"log"
	"encoding/json"
)
type Msg struct {
	Message string `json:"msg"`

}
//interfaces are not implemented expicitly, they are ducktyped

//Request will contain Body path , header

func main(){
	http.HandleFunc("/hello" , func( w http.ResponseWriter , r *http.Request){
		err := json.NewEncoder(w).Encode(Msg{Message: "Hello World"})

		if err != nil {
			http.Error(w , err.Error(),http.StatusInternalServerError)
		}
	})
	log.Fatalln(http.ListenAndServe("8080",nil))
}