package topic

import (
	"net/http"
)

func CreateTopic(w http.ResponseWriter, r *http.Request) {

	//num, err := strconv.Atoi(r.PathValue("id"))
	// if err != nil {
	// 	w.Write([]byte("Error to Create"))
	// } else {
	// 	msg := fmt.Sprintf("Create Topic By Id :%d", num)
	// 	w.Write([]byte(msg))
	// }
	msg := "Create New Topic By "
	w.Write([]byte(msg))

}
