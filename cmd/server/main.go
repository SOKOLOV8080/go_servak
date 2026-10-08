package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "гав гав мяу мяу мяу")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Сервак запущен на порту 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
