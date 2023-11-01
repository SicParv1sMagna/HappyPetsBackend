package model

// UploadImageRequest represents the request structure for the UploadImage operation.
type UploadImageRequest struct {
	UserID int    `form:"userID" binding:"required"`
	PetID  int    `form:"petID" binding:"required"`
	Image  []byte `form:"image" binding:"required"`
}

// UploadImageResponse represents the response structure for the UploadImage operation.
type UploadImageResponse struct {
	ImageURL string `json:"imageURL"`
	Message  string `json:"message"`
}

// RemoveImageRequest represents the request structure for the RemoveImage operation.
type RemoveImageRequest struct {
	UserID int `uri:"userID" binding:"required"`
	PetID  int `uri:"petID" binding:"required"`
}

// RemoveImageResponse represents the response structure for the RemoveImage operation.
type RemoveImageResponse struct{}
