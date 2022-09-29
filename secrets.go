package gcpgosec

import (
	"context"
	"log"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// GetSecret gets a secret version from Google Cloud Secrets
// projectID     := "javiercm-testproject" //The ProjectID where the secret is
// secretID      := "mysecret" //ID of the secret to retrieve
// secretVersion := "latest" //Specific revision, use latest for the newest one
// It returns the specific secret value
func GetSecret(projectID string, secretID string, secretVersion string) []byte {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	versionPath := strings.Join([]string{
		"projects/",
		projectID,
		"/secrets/",
		secretID,
		"/versions/",
		secretVersion,
	}, "")

	// Get the secret
	getSecretReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: versionPath,
	}

	// Call the API
	result, err := client.AccessSecretVersion(ctx, getSecretReq)
	if err != nil {
		log.Fatalf("Failed to get secret version: %v", err)
	}

	// Return the secret version payload
	secretContent := result.Payload.Data
	return secretContent
}
