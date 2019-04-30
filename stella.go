package stella

import (
	"encoding/json"
	"fmt"
	sc "github.com/elinamaas/stella/cmd"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port  string = "80"
	token string
)

func init() {
	token = os.Getenv("STELLA_TOKEN")
	if "" == token {
		panic("STELLA_TOKEN is not set!")
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func stellaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if token != r.FormValue("token") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)
	balloonWithCow, err := sc.Stellasay(text)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("```%s```", balloonWithCow),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func main() {
	http.HandleFunc("/", stellaHandler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
