package types

type PropertyFilters struct {
	Type     string
	Category string
	Status   string
	State    string
	City     string
	PriceMin int
	PriceMax int
	Search   string
	Limit    int
	Page     int
	Verified string
}
type Property struct {
	ID               int     `json:"property_id"`
	Owner            int     `json:"owner_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PropertyType     string  `json:"property_type"`
	PropertyCategory string  `json:"property_category"`
	State            string  `json:"state"`
	City             string  `json:"city"`
	Status           string  `json:"status"`
	Verified         bool    `json:"is_verified"`
	Image            string  `json:"photo_url"`
}

type AllProperty struct {
	Properties []*Property `json:"properties"`
	Count      int         `json:"count"`
}
type PropertyHouse struct {
	ID               int     `json:"property_id"`
	Owner            int     `json:"owner_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PropertyType     string  `json:"property_type"`
	PropertyCategory string  `json:"property_category"`
	State            string  `json:"state"`
	City             string  `json:"city"`
	Postal           string  `json:"postal_code"`
	StreetNo         int     `json:"street_no"`
	StreetName       string  `json:"street_name"`
	HouseNo          int     `json:"house_no"`
	RoomCount        int     `json:"room_count"`
	BathroomCount    int     `json:"bathroom_count"`
	Size             int     `json:"size"`
	BalconyCount     int     `json:"balcony_count"`
	ParkingFacility  bool    `json:"parking"`
	FloorCount       int     `json:"floor_count"`
	Status           string  `json:"status"`
	Verified         bool    `json:"is_verified"`
}
type PropertyCommercial struct {
	ID               int     `json:"property_id"`
	Owner            int     `json:"owner_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PropertyType     string  `json:"property_type"`
	PropertyCategory string  `json:"property_category"`
	State            string  `json:"state"`
	City             string  `json:"city"`
	Postal           string  `json:"postal_code"`
	StreetNo         int     `json:"street_no"`
	StreetName       string  `json:"street_name"`
	HouseNo          int     `json:"house_no"`
	Size             int     `json:"size"`
	ParkingFacility  bool    `json:"parking"`
	FloorNo          int     `json:"floor_no"`
	Status           string  `json:"status"`
	Verified         bool    `json:"is_verified"`
}
type PropertyApartment struct {
	ID               int     `json:"property_id"`
	Owner            int     `json:"owner_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PropertyType     string  `json:"property_type"`
	PropertyCategory string  `json:"property_category"`
	State            string  `json:"state"`
	City             string  `json:"city"`
	Postal           string  `json:"postal_code"`
	StreetNo         int     `json:"street_no"`
	StreetName       string  `json:"street_name"`
	HouseNo          int     `json:"house_no"`
	RoomCount        int     `json:"room_count"`
	BathroomCount    int     `json:"bathroom_count"`
	Size             int     `json:"size"`
	BalconyCount     int     `json:"balcony_count"`
	ParkingFacility  bool    `json:"parking"`
	FloorNo          int     `json:"floor_no"`
	Status           string  `json:"status"`
	Verified         bool    `json:"is_verified"`
}
type PropertyHousePayload struct {
	Owner            int      `json:"owner_id"`
	Title            string   `json:"title" validate:"required"`
	Description      string   `json:"description" validate:"required"`
	Price            float64  `json:"price" validate:"required"`
	PropertyType     string   `json:"property_type" validate:"required"`
	PropertyCategory string   `json:"property_category" validate:"required"`
	State            string   `json:"state" validate:"required"`
	City             string   `json:"city" validate:"required"`
	Postal           string   `json:"postal_code" validate:"required"`
	StreetNo         int      `json:"street_no" validate:"required"`
	StreetName       string   `json:"street_name" validate:"required"`
	HouseNo          int      `json:"house_no" validate:"required"`
	RoomCount        int      `json:"room_count" validate:"required"`
	BathroomCount    int      `json:"bathroom_count" validate:"required"`
	Size             int      `json:"size" validate:"required"`
	BalconyCount     int      `json:"balcony_count" validate:"required"`
	ParkingFacility  bool     `json:"parking" validate:"required"`
	FloorCount       int      `json:"floor_count" validate:"required"`
	Image            []string `json:"image" validate:"required"`
}
type PropertyApartmentPayload struct {
	Owner            int      `json:"owner_id"`
	Title            string   `json:"title" validate:"required"`
	Description      string   `json:"description" validate:"required"`
	Price            float64  `json:"price" validate:"required"`
	PropertyType     string   `json:"property_type" validate:"required"`
	PropertyCategory string   `json:"property_category" validate:"required"`
	State            string   `json:"state" validate:"required"`
	City             string   `json:"city" validate:"required"`
	Postal           string   `json:"postal_code" validate:"required"`
	StreetNo         int      `json:"street_no" validate:"required"`
	StreetName       string   `json:"street_name" validate:"required"`
	HouseNo          int      `json:"house_no" validate:"required"`
	RoomCount        int      `json:"room_count" validate:"required"`
	BathroomCount    int      `json:"bathroom_count" validate:"required"`
	Size             int      `json:"size" validate:"required"`
	BalconyCount     int      `json:"balcony_count" validate:"required"`
	ParkingFacility  bool     `json:"parking" validate:"required"`
	FloorNo          int      `json:"floor_no" validate:"required"`
	Image            []string `json:"image" validate:"required"`
}
type PropertyCommercialPayload struct {
	Owner            int      `json:"owner_id"`
	Title            string   `json:"title" validate:"required"`
	Description      string   `json:"description" validate:"required"`
	Price            float64  `json:"price" validate:"required"`
	PropertyType     string   `json:"property_type" validate:"required"`
	PropertyCategory string   `json:"property_category" validate:"required"`
	State            string   `json:"state" validate:"required"`
	City             string   `json:"city" validate:"required"`
	Postal           string   `json:"postal_code" validate:"required"`
	StreetNo         int      `json:"street_no" validate:"required"`
	StreetName       string   `json:"street_name" validate:"required"`
	HouseNo          int      `json:"house_no" validate:"required"`
	Size             int      `json:"size" validate:"required"`
	ParkingFacility  bool     `json:"parking" validate:"required"`
	FloorNo          int      `json:"floor_no" validate:"required"`
	Image            []string `json:"image" validate:"required"`
}
