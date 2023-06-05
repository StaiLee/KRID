package module

import "os"

var post Publication   //posts data to print
var templ Templ        //struct to send to print data
var Notif Notification //struct to send to print data
var user User          //store a user data
var templatesDir = os.Getenv("TEMPLATES_DIR")
