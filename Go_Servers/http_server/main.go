package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/math-form/", http.StripPrefix("/math-form/", fs))

	http.HandleFunc("/add-form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Can't parse body", http.StatusBadRequest)
			return
		}

		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")

		num1, err1 := strconv.Atoi(num1Str)
		num2, err2 := strconv.Atoi(num2Str)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
		}

		result := num1 + num2

		response := map[string]interface{}{
			"result": result,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		a := r.FormValue("a")
		b := r.FormValue("b")

		aInt, err1 := strconv.Atoi(a)
		bInt, err2 := strconv.Atoi(b)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Addition of %d and %d is: %d ", aInt, bInt, aInt+bInt)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page")
	})

	fmt.Println("server is running on port: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is not running")
	}

	// testing
	// testing
}
