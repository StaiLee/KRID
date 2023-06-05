package module

import "html/template"

type Templ struct {
	Post     []Publication // posts
	User     User          // user info
	JsonUser []jsonUser
	Comments []Comment //comments under post
	Report   []Report
	Mods     []Mod         //moderation
	Inter    []Interaction // interaction
	Notif    Notification  // notification
}

type Publication struct {
	Id       int
	Creator  string
	Title    string
	Content  string
	Picture  template.HTML
	Likes    int
	Dislikes int
	Slash    string
	Date     string
}

type Notification struct {
	Post  Publication
	Com   []Comment
	Inter []Interaction
}
type Interaction struct {
	Postid string
	Uuid   string
	Date   string
}

type User struct {
	Username     string
	Creationdate string
	Level        int
}

type Mod struct {
	Uuid         string
	Email        string
	Username     string
	Creationdate string
	Level        int
}

type Report struct {
	ReportContent string
	ReportId      string
	User          string
	Postid        string
}

type Comment struct {
	Creator          string
	Postid           string
	Comment          string
	Likes            int
	Dislikes         int
	File             string
	Creationdateuser string
	Username         string
	Level            int
	Date             string
}

type jsonUser struct {
	Name     string `json:"login"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar_url"`
	Location string `json:"location"`
}
