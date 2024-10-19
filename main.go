package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Output struct {
	Result string `json:"result"`
}

// Обработчик HTTP-запроса
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	if !r.URL.Query().Has("string") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Query params don't have property \"string\"!"))
		return
	}
	str := r.URL.Query().Get("string")

	if str == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("string is empty!"))
		return
	}

	var output Output
	for _, symbol := range str {
		number, err := strconv.Atoi(string(symbol))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong format of string!"))
			return
		}
		output.Result += strconv.Itoa(number * number * number)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respBytes, _ := json.Marshal(output)
	w.Write(respBytes)
}

func main() {
	// Регистрируем обработчик для пути "/calculate"
	http.HandleFunc("/kub", CalculateHandler)

	// Запускаем веб-сервер на порту 8081
	fmt.Println("starting server...")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
