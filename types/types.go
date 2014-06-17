// Copyright (c) 2013 Aditya Mukerjee, Quotidian Ventures
// https://github.com/ChimeraCoder/anaconda

package types

type TwitterConfig struct {
	Token          string `json:"access_token"`
	TokenSecret    string `json:"access_token_secret"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret_key"`
}

type Tweet struct {
	Contributors         []int64     `bson:"contributors" json:"contributors"`
	Coordinates          interface{} `bson:"coordinates" json:"coordinates"`
	CreatedAt            string      `bson:"created_at" json:"created_at"`
	Entities             Entities    `bson:"entities" json:"entities"`
	FavoriteCount        int         `bson:"favorite_count" json:"favorite_count"`
	Favorited            bool        `bson:"favorited" json:"favorited"`
	Geo                  interface{} `bson:"geo" json:"geo"`
	Id                   int64       `bson:"id" json:"id"`
	IdStr                string      `bson:"id_str" json:"id_str"`
	InReplyToScreenName  string      `bson:"in_reply_to_screen_name" json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64       `bson:"in_reply_to_status_id" json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string      `bson:"in_reply_to_status_id_str" json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64       `bson:"in_reply_to_user_id" json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string      `bson:"in_reply_to_user_id_str" json:"in_reply_to_user_id_str"`
	PossiblySensitive    bool        `bson:"possibly_sensitive" json:"possibly_sensitive"`
	RetweetCount         int         `bson:"retweet_count" json:"retweet_count"`
	Retweeted            bool        `bson:"retweeted" json:"retweeted"`
	RetweetedStatus      *Tweet      `bson:"retweeted_status" json:"retweeted_status"`
	Source               string      `bson:"source" json:"source"`
	Text                 string      `bson:"text" json:"text"`
	Truncated            bool        `bson:"truncated" json:"truncated"`
	User                 User        `json:"user"`
}

type Entities struct {
	Hashtags []struct {
		Indices []int
		Text    string
	}
	Urls []struct {
		Indices      []int
		Url          string
		Display_url  string
		Expanded_url string
	}
	User_mentions []struct {
		Name        string
		Indices     []int
		Screen_name string
		Id          int64
		Id_str      string
	}
	Media []struct {
		Id              int64
		Id_str          string
		Media_url       string
		Media_url_https string
		Url             string
		Display_url     string
		Expanded_url    string
		Sizes           MediaSizes
		Type            string
		Indices         []int
	}
}

type MediaSizes struct {
	Medium MediaSize
	Thumb  MediaSize
	Small  MediaSize
	Large  MediaSize
}

type MediaSize struct {
	W      int
	H      int
	Resize string
}

type User struct {
	ContributorsEnabled            bool   `json:"contributors_enabled"`
	CreatedAt                      string `json:"created_at"`
	DefaultProfile                 bool   `json:"default_profile"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	Description                    string `json:"description"`
	FavouritesCount                int    `json:"favourites_count"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	FollowersCount                 int    `json:"followers_count"`
	Following                      bool   `json:"following"`
	FriendsCount                   int    `json:"friends_count"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	Id                             int64  `json:"id"`
	IdStr                          string `json:"id_str"`
	IsTranslator                   bool   `json:"is_translator"`
	Lang                           string `json:"lang"`
	ListedCount                    int64  `json:"listed_count"`
	Location                       string `json:"location"`
	Name                           string `json:"name"`
	Notifications                  bool   `json:"notifications"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileImageURL                string `json:"profile_image_url"`
	ProfileImageUrlHttps           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	Protected                      bool   `json:"protected"`
	ScreenName                     string `json:"screen_name"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
	Status                         *Tweet `json:"status"` // Only included if the user is a friend
	StatusesCount                  int64  `json:"statuses_count"`
	TimeZone                       string `json:"time_zone"`
	URL                            string `json:"url"`
	UtcOffset                      int    `json:"utc_offset"`
	Verified                       bool   `json:"verified"`
}
