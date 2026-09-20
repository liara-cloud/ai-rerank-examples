package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("LIARA_API_KEY")
	modelName := os.Getenv("RERANK_MODEL_NAME")

	url := baseURL + "/rerank"

	payload := map[string]any{
		"model": modelName,

		"query": "And who is God?",

		"documents": []string{
			"God means love, purity, intimacy, friendship",
			"God is kind",
			"God loves us and we should love him too",
			"AI means Artificial Intelligence",
		},

		"top_n": 3,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+apiKey,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}