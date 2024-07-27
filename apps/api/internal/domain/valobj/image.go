package valobj

type Image struct {

	// Image URL
	//
	// Example: https://s3.example.com/image.jpg
	Url string `json:"url" validate:"required,url" example:"https://s3.example.com/image.jpg"`

	// Image order
	//
	// Example: 1
	Order int16 `json:"order" validate:"required" example:"1"`
}
