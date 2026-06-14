package topic

import (
	"fmt"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {

	num, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Error to Update"))
	} else {
		msg := fmt.Sprintf("Update Topic By Id :%d", num)
		w.Write([]byte(msg))
	}
}
