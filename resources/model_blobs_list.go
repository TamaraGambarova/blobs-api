/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BlobsList struct {
	Key
	Attributes BlobsListAttributes `json:"attributes"`
}
type BlobsListResponse struct {
	Data     BlobsList `json:"data"`
	Included Included  `json:"included"`
}

type BlobsListListResponse struct {
	Data     []BlobsList `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustBlobsList - returns BlobsList from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBlobsList(key Key) *BlobsList {
	var blobsList BlobsList
	if c.tryFindEntry(key, &blobsList) {
		return &blobsList
	}
	return nil
}
