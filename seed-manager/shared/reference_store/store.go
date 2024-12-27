package reference_store

type ReferenceStore interface {
	// Store user IDs and other cross-service references
	StoreUserID(email string, userID string) error
	GetUserID(email string) (string, error)

	// Store other cross-service references
	StoreFileID(reference string, fileID string) error
	GetFileID(reference string) (string, error)

	// Clear all stored references
	Clear() error
}

// Implementation could be in-memory for development or Redis/DB for distributed environments
