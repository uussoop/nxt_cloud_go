package gonextcloud

import "strconv"

// User encapsulate the data needed to create a new Nextcloud's User
type User struct {
	Username    string
	Email       string
	DisplayName string
	Quota       string
	Language    string
	Groups      []string
}

// UserDetails is the raw Nextcloud User response
type UserDetails struct {
	Enabled     bool     `json:"enabled"`
	ID          string   `json:"id"`
	Quota       Quota    `json:"quota"`
	Email       string   `json:"email"`
	Displayname string   `json:"displayname"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	Website     string   `json:"website"`
	Twitter     string   `json:"twitter"`
	Groups      []string `json:"groups"`
	Language    string   `json:"language,omitempty"`

	StorageLocation string        `json:"storageLocation,omitempty"`
	LastLogin       int64         `json:"lastLogin,omitempty"`
	Backend         string        `json:"backend,omitempty"`
	Subadmin        []interface{} `json:"subadmin,omitempty"`
	Locale          string        `json:"locale,omitempty"`
}

// Quota is a use storage Quota
type Quota struct {
	Free     int64   `json:"free"`
	Used     int64   `json:"used"`
	Total    int64   `json:"total"`
	Relative float64 `json:"relative"`
	Quota    int64   `json:"quota"`
}

func (q *Quota) String() string {
	if q.Quota < 0 {
		return "none"
	}
	return strconv.FormatInt(q.Quota, 10)
}
