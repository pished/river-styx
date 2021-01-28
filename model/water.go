package model

type Water struct {
	ItemId    string `json:"itemId"`
	Type      string `json:"type"`
	Brand     string `json:"brand"`
	Category  string `json:"Category"`
	Flavored  bool   `json:"flavored"`
	Flavoring string `json:"flavoring"`
	Quality   int    `json:"quality"`
}
type WaterForm struct {
	ItemId    string `json:"itemId"`
	Type      string `json:"type" form:"required,max=255"`
	Brand     string `json:"brand" form:"required,max=255"`
	Category  string `json:"Category" form:"required,max=255"`
	Flavored  bool   `json:"flavored"`
	Flavoring string `json:"flavoring" form:"required,max=255"`
	Quality   int    `json:"quality" form:"required,number"`
}
type WaterUpdateForm struct {
	Quality int `json:"quality" form:"required,number"`
}
type WaterModel struct {
	ItemId    string
	Type      string
	Brand     string
	Category  string
	Flavored  bool
	Flavoring string
	Quality   int
}

type Waters []WaterForm

func (w *WaterForm) ToModel() *WaterModel {
	return &WaterModel{
		ItemId:    w.ItemId,
		Type:      w.Type,
		Brand:     w.Brand,
		Category:  w.Category,
		Flavored:  w.Flavored,
		Flavoring: w.Flavoring,
		Quality:   w.Quality,
	}
}

func (w *WaterUpdateForm) ToModel() *WaterModel {
	return &WaterModel{
		Quality: w.Quality,
	}
}
