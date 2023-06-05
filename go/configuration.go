package module

//FRONT
//page about us, petite présentation + (nos git,linkedin,mails)
//profile is looking rather bad
//ajout report,like et le mettre dans la boite contenant le post
//notif dans le header a gauche de about us j'arrive pas à le placer and place on top of buttons
//bouton ajout de post + joli
//admin panel moderation (report handling, mod list promote or demote, communities management)
//slashs a gerer affichage
//page activité/profil de l'utilisateur (created post, commentaires(show post and comment), post liked or disliked)
//notif dans le header a gauche de about us like dislike post

//BACKEND
//admin functions missing
//encrypt database
//authentification avec google and github NEED IP DOMAIN
//implementer https
//docker
//modele entitées relation
//powerpoint

//communities
//scroll infini 10 par 10 posts
//user profile picture
//ajout d'emojis
//password reset
//email verification
//compteurs de messages utilisateur

import "os"

var post Publication   //posts data to print
var templ Templ        //struct to send to print data
var Notif Notification //struct to send to print data
var user User          //store a user data
var templatesDir = os.Getenv("TEMPLATES_DIR")
