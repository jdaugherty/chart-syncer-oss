package indexer

import (
	"net/url"

	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/retry"
)

// newRemoteRepository creates a remote.Repository for an OCI registry with basic auth
func newRemoteRepository(u *url.URL, username, password string, insecure bool) (*remote.Repository, error) {
	// remote.NewRepository expects a reference in format "host/path", not a full URL
	ref := u.Host + u.Path
	repo, err := remote.NewRepository(ref)
	if err != nil {
		return nil, err
	}

	if insecure {
		repo.PlainHTTP = true
	}

	repo.Client = &auth.Client{
		Client: retry.DefaultClient,
		Cache:  auth.NewCache(),
		Credential: auth.StaticCredential(repo.Reference.Registry, auth.Credential{
			Username: username,
			Password: password,
		}),
	}

	return repo, nil
}
