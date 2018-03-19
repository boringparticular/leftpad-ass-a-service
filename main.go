package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		var buffer bytes.Buffer

		query := r.URL.Query()
		str := query.Get("str")
		chq := query.Get("ch")
		ch := string([]rune(chq)[0])
		length := query.Get("len")
		l, err := strconv.Atoi(length)

		if err != nil {
			l = 0
		}

		l = l - len(str)

		for i := 0; i < l; i++ {
			buffer.WriteString(ch)
		}

		buffer.WriteString(str)

		result := map[string]string{"str": buffer.String()}

		json.NewEncoder(w).Encode(result)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8085", nil))
}
