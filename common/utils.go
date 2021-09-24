package common

type ErrorResponse struct {
	Message string `json:"message"`
}

type School struct {
	Id           int     `json:"id,omitempty"`
	URL          string  `json:"url,omitempty"`
	Name         string  `json:"name,omitempty"`
	SchoolId     string  `json:"school_id,omitempty"`
	Type         string  `json:"type,omitempty"`
	IsCatholic   bool    `json:"is_catholic,omitempty"`
	Language     string  `json:"language,omitempty"`
	Level        string  `json:"level,omitempty"`
	City         string  `json:"city,omitempty"`
	CitySlug     string  `json:"city_slug,omitempty"`
	Board        string  `json:"board,omitempty"`
	FraserRating float64 `json:"fraser_rating,omitempty"`
	EQAORating   float64 `json:"eqao_rating,omitempty"`
	Address      string  `json:"address,omitempty"`
	Grades       string  `json:"grades,omitempty"`
	Website      string  `json:"website,omitempty"`
	PhoneNumber  string  `json:"phone_number,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
}
