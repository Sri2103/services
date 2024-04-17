package categories_service

type Category struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Url  string `json:"url,omitempty"`
	Img  string `json:"img,omitempty"`
}


