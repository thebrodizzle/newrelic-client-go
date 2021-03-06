// Code generated by tutone: DO NOT EDIT
package nerdstorage

// NerdStorageScope - The access level of the NerdStorage data.
type NerdStorageScope string

var NerdStorageScopeTypes = struct {
	// Account-level storage.
	ACCOUNT NerdStorageScope
	// Actor-level storage.
	ACTOR NerdStorageScope
	// Entity-level storage.
	ENTITY NerdStorageScope
}{
	// Account-level storage.
	ACCOUNT: "ACCOUNT",
	// Actor-level storage.
	ACTOR: "ACTOR",
	// Entity-level storage.
	ENTITY: "ENTITY",
}

// NerdStorageCollectionMember -
type NerdStorageCollectionMember struct {
	// The NerdStorage document.
	Document NerdStorageDocument `json:"document,omitempty"`
	// The documentId.
	ID string `json:"id,omitempty"`
}

// NerdStorageEntityScope -
type NerdStorageEntityScope struct {
	//
	Collection []NerdStorageCollectionMember `json:"collection,omitempty"`
	//
	Document NerdStorageDocument `json:"document,omitempty"`
}

// NerdStorageScopeInput - The data access level and ID for the selected scope.
type NerdStorageScopeInput struct {
	// The ID for the selected scope.
	ID string `json:"id"`
	// The NerdStorage data access level.
	Name NerdStorageScope `json:"name"`
}

// NerdStorageDocument - This scalar represents a NerdStorage document.
type NerdStorageDocument string
