package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"encoding/base64"

	"google.golang.org/api/iam/v1"
)

// createKey creates a service account key.
func createKey(w io.Writer, serviceAccountEmail string) (*iam.ServiceAccountKey, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %w", err)
	}

	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	request := &iam.CreateServiceAccountKeyRequest{}
	key, err := service.Projects.ServiceAccounts.Keys.Create(resource, request).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.Create: %w", err)
	}

	// Decode the PrivateKeyData which is base64-encoded JSON
	jsonKeyFile, err := base64.StdEncoding.DecodeString(key.PrivateKeyData)
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString: %w", err)
	}

	// Write the key to the writer
	if _, err := w.Write(jsonKeyFile); err != nil {
		return nil, fmt.Errorf("w.Write: %w", err)
	}

	return key, nil
}

// Before running this: run gcloud auth application-default login
func main() {
	f, err := os.Create("gcp-credentials.json")
	if err != nil {
		fmt.Printf("os.Create: %v\n", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = createKey(w, "gcp-sa@project-id.iam.gserviceaccount.com")
	if err != nil {
		fmt.Printf("createKey: %v\n", err)
		return
	}

	// Ensure all buffered data is written to the file
	if err := w.Flush(); err != nil {
		fmt.Printf("w.Flush: %v\n", err)
		return
	}

	fmt.Println("gcp-credentials.json written successfully")
}
