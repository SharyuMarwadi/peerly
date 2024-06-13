package dto

type UpdateQueryRequest struct {
	Text         string `db:"text" json:"text"`
	Description  string `db:"description" json:"description"`
	ThumbnailUrl string `db:"thumbnailurl" json:"thumbnail_url"`
}
