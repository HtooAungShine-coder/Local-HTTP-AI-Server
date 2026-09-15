package main

import (
	"fmt"
	"luna-ai/API"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index2.html")
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		message := r.FormValue("message")
		// Log the incoming frontend query string
		fmt.Printf("[FRONTEND] User input received: %s\n", message)

		answer, err := API.SendMessage(message)

		if err != nil {
			fmt.Printf("[BACKEND EXECUTOR FAILURE]: %v\n", err)
			http.Error(w, fmt.Sprintf("Error calling LLM: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Println("[BACKEND SUCCESS] Response successfully fetched from OpenRouter.")

		// Write the text answer directly back to the HTTP response context
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(answer))
	})

	fmt.Println("Luna AI is running...")
	fmt.Println("Luna Ai is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
