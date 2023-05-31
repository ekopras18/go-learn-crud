package baseUtility

import (
	"fmt"
	"net/http"
)

func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func StatusInternalServer(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CatchWithMessage(err error, message string) {
	if err != nil {
		fmt.Println(message)
		// panic(err)
	}
}
