package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// generateV4GetObjectSignedURL generates object signed URL with GET method.
func generateV4GetObjectSignedURL(w io.Writer, bucket, object string) (string, error) {
	// bucket := "bucket-name"
	// object := "object-name"

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	// Signing a URL requires credentials authorized to sign a URL. You can pass
	// these in through SignedURLOptions with one of the following options:
	//    a. a Google service account private key, obtainable from the Google Developers Console
	//    b. a Google Access ID with iam.serviceAccounts.signBlob permissions
	//    c. a SignBytes function implementing custom signing.
	// In this example, none of these options are used, which means the SignedURL
	// function attempts to use the same authentication that was used to instantiate
	// the Storage client. This authentication must include a private key or have
	// iam.serviceAccounts.signBlob permissions.
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: "",
		Expires:        time.Now().Add(15 * time.Minute),
	}

	u, err := client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}

	fmt.Fprintln(w, "Generated GET signed URL:")
	fmt.Fprintf(w, "%q\n", u)
	fmt.Fprintln(w, "You can use this URL with any user agent, for example:")
	fmt.Fprintf(w, "curl %q\n", u)
	return u, nil
}

// generateSignedURL to generate signed url to give access for bucket objects in GCP.
func generateSignedURL() {
	// Define the bucket name and object name.
	bucket := "bucket-default-object-access-control-random"
	object := "your-object-name"

	// Open a file to store the signed URL.
	file, err := os.Create("signed_url.txt")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Generate the signed URL and write it to the file.
	signedURL, err := generateV4GetObjectSignedURL(file, bucket, object)
	if err != nil {
		fmt.Printf("Failed to generate signed URL: %v\n", err)
		return
	}

	fmt.Printf("Signed URL successfully written to file: %s\n", signedURL)
}
