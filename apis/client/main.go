package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, (2 * time.Second))
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	if err != nil {
		log.Fatal(err)

		return
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)

		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)

		return
	}

	io.Copy(os.Stdout, res.Body)
}
