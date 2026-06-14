package topic

import (
	"fmt"
	"net/http"
	"strconv"
)

func DeleteTopic(w http.ResponseWriter, r *http.Request) {

	num, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Error to get"))
	} else {
		msg := fmt.Sprintf("Delete Topic By Id :%d", num)
		w.Write([]byte(msg))
	}
}
