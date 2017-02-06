package vkapiclient

type VkApiErrorRequestParam struct {
	Key   string `json:\"key\"`
	Value string `json:\"value\"`
}

type VkApiError struct {
	Code          int    `json:\"error_code\"`
	Message       string `json:\"error_msg\"`
	Text          string `json:\"error_text\"`
	RequestParams []VkApiErrorRequestParam
}

type PhotosPhotoSizes struct {
	// URL of the image.
	Src string `json:\"id\"`
	// Width in px.
	Width string `json:\"width\"`
	// Height in px.
	Height string `json:\"height\"`
	// Size type.
	Type string `json:\"type\"`
}

type PhotosPhoto struct {
	// Photo ID.
	Id int `json:\"id\"`
	// Album ID
	AlbumId int `json:\"album_id\"`
	// Photo owner's ID.
	OwnerId int `json:\"owner_id\"`
	// ID of the user who have uploaded the photo.
	UserId int                `json:\"user_id\"`
	Sizes  []PhotosPhotoSizes `json:\"sizes\"`
	// URL of image with 75 px width.
	Photo75 string `json:\"photo_75\"`
	// URL of image with 130 px width.
	Photo130 string `json:\"photo_130\"`
	// URL of image with 604 px width.
	Photo604 string `json:\"photo_604\"`
	// URL of image with 807 px width.
	Photo807 string `json:\"photo_807\"`
	// URL of image with 1280 px width.
	Photo1280 string `json:\"photo_1280\"`
	// URL of image with 2560 px width.
	Photo2560 string `json:\"photo_2560\"`
	// Post ID.
	PostId int `json:\"post_id\"`
	// Original photo width.
	Width int `json:\"width\"`
	// Original photo height.
	Height int `json:\"height"`
	// Photo caption.
	Text string `json:\"text\"`
	// Date when uploaded.
	Date int `json:\"date\"`
	// Latitude.
	Lat float64 `json:\"lat\"`
	// Longitude.
	Long float64 `json:\"long\"`
	// Access key for the photo.
	AccessKey string `json:\"access_key\"`
}

type PhotosGetByIdResponse struct {
	Error    VkApiError    `json:\"error\"`
	Response []PhotosPhoto `json:\"response\"`
}

type BaseBoolInt struct {
	Type int
}

type BaseObjectCount struct {
	// Items count
	Count int
}

type BaseLikes struct {
	// Information whether current user likes the photo.
	UserLikes BaseBoolInt `json:\"UserLikes\"`
	// Likes number.
	Count int
}

type PhotosPhotoFull struct {
	// Photo ID.
	Id int `json:\"id\"`
	// Album ID.
	AlbumId int `json:\"album_id\"`
	// Photo owner's ID.
	OwnerId int `json:\"owner_id\"`
	// ID of the user who have uploaded the photo.
	UserId  int                `json:\"user_id\"`
	Sizes   []PhotosPhotoSizes `json:\"sizes\"`
	Photo75 string             `json:\"photo_75\"`
	// URL of image with 130 px width.
	Photo130 string `json:\"photo_130\"`
	// URL of image with 604 px width.
	Photo604 string `json:\"photo_604\"`
	// URL of image with 807 px width.
	Photo807 string `json:\"photo_807\"`
	// URL of image with 1280 px width.
	Photo1280 string `json:\"photo_1280\"`
	// URL of image with 2560 px width.
	Photo2560 string `json:\"photo_2560\"`
	// Post ID.
	PostId int `json:\"post_id\"`
	// Original photo width.
	Width int `json:\"width\"`
	// Original photo height.
	Height int `json:\"height"`
	// Photo caption.
	Text string `json:\"text\"`
	// Date when uploaded.
	Date int `json:\"date\"`
	// Latitude.
	Lat float64 `json:\"lat\"`
	// Longitude.
	Long float64 `json:\"long\"`
	// Access key for the photo.
	AccessKey string          `json:\"access_key\"`
	Likes     BaseLikes       `json:\"likes"`
	Reposts   BaseObjectCount `json:\"reposts\"`
	Comments  BaseObjectCount `json:\"comments\"`
	// Information whether current user can comment the photo.
	CanComment BaseBoolInt     `json:\"can_comment\"`
	Tags       BaseObjectCount `json:\"tags\"`
}

type PhotosGetByIdExtendedResponse struct {
	Error    VkApiError      `json:\"error\"`
	Response PhotosPhotoFull `json:\"response\"`
}
