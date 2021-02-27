package models

// StyleRequest -
type StyleRequest struct {
	BackgroundID int `json:"background_id"`
	ButtonID     int `json:"button_id"`
	FontID       int `json:"font_id"`
}

// StyleResponse -
type StyleResponse struct {
	ProfileID    int `json:"profile_id"`
	BackgroundID int `json:"background_id"`
	ButtonID     int `json:"button_id"`
	FontID       int `json:"font_id"`
}

// BackgroundStyle -
type BackgroundStyle struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	SubImage string `json:"subImage"`
}

// ButtonStyle -
type ButtonStyle struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// FontStyle -
type FontStyle struct {
	Name       string `json:"name"`
	FontFamily string `json:"fontFamily"`
	Href       string `json:"href"`
}

// StyleDetail -
type StyleDetail struct {
	ProfileID  int             `json:"profileID"`
	Background BackgroundStyle `json:"background"`
	Button     ButtonStyle     `json:"button"`
	Font       FontStyle       `json:"font"`
}
