package module

import (
	"net/http"
	"time"
)

type Request struct {
	ip          string
	count       int
	time        time.Time
	spentintime int
}

var requests []Request

func Contains(s []Request, ipreq string) int {
	for _, addr := range s {
		if addr.ip == ipreq {
			return addr.count
		}
	}
	return -1
}

func Countdown(nb int) {
	t2 := time.Now()
	t1 := requests[nb].time
	elapsed := t2.Sub(t1)
	elapsedinseconds := int(elapsed)/10000000 + requests[nb].spentintime //conversion des secondes
	if elapsedinseconds > 100 {
		requests[nb].count = 0
		requests[nb].time = time.Now()
		requests[nb].spentintime = 0
	} else { //ajout du temps
		requests[nb].spentintime += elapsedinseconds
	}
}

func Ratelimit(w http.ResponseWriter, r *http.Request) {
	nb := Contains(requests, r.RemoteAddr)
	if nb != -1 { //si l'ip est correct
		Countdown(nb)
		if requests[nb].count < 3 { //nombre de requete bonne
			requests[nb].count++
			return
		} else { //si trop de requete
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
		}
	} else { //si l'ip est inconnu
		requests = append(requests, Request{r.RemoteAddr, 1, time.Now(), 0})
	}
}
