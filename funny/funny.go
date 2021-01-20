package funny

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/dllg/go-funny-endpoints/httpclient"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type slip struct {
	Advice string `json:"advice" binding:"required"`
}
type advice struct {
	Slip slip `json:"slip" binding:"required"`
}

// GetAdviceFromAPI will get an advice from api.adviceslip.com
func GetAdviceFromAPI(hc httpclient.HTTPClient) string {
	body, err := hc.SendGetRequest("https://api.adviceslip.com/advice", nil)
	if err != nil {
		return err.Error()
	}
	var a advice
	if err := json.Unmarshal(body, &a); err != nil {
		return err.Error()
	}
	return a.Slip.Advice
}

type chucknorris struct {
	Value string `json:"value" binding:"required"`
}

// GetChuckNorrisJokeFromAPI will get a joke from api.chucknorris.io
func GetChuckNorrisJokeFromAPI(hc httpclient.HTTPClient) string {
	body, err := hc.SendGetRequest("https://api.chucknorris.io/jokes/random", nil)
	if err != nil {
		return err.Error()
	}
	var a chucknorris
	if err := json.Unmarshal(body, &a); err != nil {
		return err.Error()
	}
	return a.Value
}

type dadjoke struct {
	Joke string `json:"joke" binding:"required"`
}

// GetDadJokeFromAPI will get a dad joke from icanhazdadjoke.com
func GetDadJokeFromAPI(hc httpclient.HTTPClient) string {
	body, err := hc.SendGetRequest("https://icanhazdadjoke.com/", map[string]string{"Accept": "application/json"})
	if err != nil {
		return err.Error()
	}
	var a dadjoke
	if err := json.Unmarshal(body, &a); err != nil {
		return err.Error()
	}
	return a.Joke
}

type msgfunc func(hc httpclient.HTTPClient) string
type message struct {
	msgtype string
	msgfunc msgfunc
}

func getMessage(hc httpclient.HTTPClient, index int) string {
	f := []message{
		{"Advice", GetAdviceFromAPI},
		{"Chuck Norris Joke", GetChuckNorrisJokeFromAPI},
		{"Dad Joke", GetDadJokeFromAPI},
	}
	if index >= 0 && index < len(f) {
		log.WithFields(logrus.Fields{
			"messageType": f[index].msgtype,
		}).Info("Generating message")

		return f[index].msgtype + ": " + f[index].msgfunc(hc)
	}
	return ""
}

// GetRandomMessage will get a random message from one of the apis
func GetRandomMessage(hc httpclient.HTTPClient) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	return getMessage(hc, randomIndex)
}
