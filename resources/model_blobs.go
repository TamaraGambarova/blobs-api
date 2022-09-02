/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Blobs struct {
	Key
	Attributes BlobsAttributes `json:"attributes"`
}
type BlobsResponse struct {
	Data     Blobs    `json:"data"`
	Included Included `json:"included"`
}

type BlobsListResponse struct {
	Data     []Blobs  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustBlobs - returns Blobs from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBlobs(key Key) *Blobs {
	var blobs Blobs
	if c.tryFindEntry(key, &blobs) {
		return &blobs
	}
	return nil
}
