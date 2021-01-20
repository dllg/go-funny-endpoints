package funny

import (
	"testing"

	"github.com/dllg/go-funny-endpoints/httpclient"
	"github.com/golang/mock/gomock"
)

func TestAPIMethods(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := httpclient.NewMockHTTPClient(mockCtrl)

	tests := []struct {
		name     string
		method   msgfunc
		mockCall func()
		want     string
	}{
		{
			name:   "GetAdviceFromAPI",
			method: GetAdviceFromAPI,
			mockCall: func() {
				reply := "{\"slip\": { \"id\": 183, \"advice\": \"Always get two ciders.\"}}"
				mock.EXPECT().SendGetRequest("https://api.adviceslip.com/advice", nil).Return([]byte(reply), nil).Times(1)
			},
			want: "Always get two ciders.",
		},
		{
			name:   "GetChuckNorrisJokeFromAPI",
			method: GetChuckNorrisJokeFromAPI,
			mockCall: func() {
				reply := "{\"categories\":[],\"created_at\":\"2020-01-05 13:42:24.696555\",\"icon_url\":\"https://assets.chucknorris.host/img/avatar/chuck-norris.png\",\"id\":\"3Odt13-SQ06Pq-RUPwRQ4w\",\"updated_at\":\"2020-01-05 13:42:24.696555\",\"url\":\"https://api.chucknorris.io/jokes/3Odt13-SQ06Pq-RUPwRQ4w\",\"value\":\"Peace disturbs Chuck Norris.\"}"
				mock.EXPECT().SendGetRequest("https://api.chucknorris.io/jokes/random", nil).Return([]byte(reply), nil).Times(1)
			},
			want: "Peace disturbs Chuck Norris.",
		},
		{
			name:   "GetDadJokeFromAPI",
			method: GetDadJokeFromAPI,
			mockCall: func() {
				reply := "{\"id\":\"0DQKB51oGlb\",\"joke\":\"What did one nut say as he chased another nut?  I'm a cashew!\",\"status\":200}"
				mock.EXPECT().SendGetRequest("https://icanhazdadjoke.com/", map[string]string{"Accept": "application/json"}).Return([]byte(reply), nil).Times(1)
			},
			want: "What did one nut say as he chased another nut?  I'm a cashew!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCall()
			if got := tt.method(mock); got != tt.want {
				t.Errorf("%v: got \"%v\", want \"%v\"", tt.name, got, tt.want)
			}
		})
	}
}

func TestGetMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := httpclient.NewMockHTTPClient(mockCtrl)

	tests := []struct {
		name     string
		index    int
		mockCall func()
		want     string
	}{
		{
			name:  "GetAdviceFromAPI",
			index: 0,
			mockCall: func() {
				reply := "{\"slip\": { \"id\": 183, \"advice\": \"Always get two ciders.\"}}"
				mock.EXPECT().SendGetRequest("https://api.adviceslip.com/advice", nil).Return([]byte(reply), nil).Times(1)
			},
			want: "Advice: Always get two ciders.",
		},
		{
			name:  "GetChuckNorrisJokeFromAPI",
			index: 1,
			mockCall: func() {
				reply := "{\"categories\":[],\"created_at\":\"2020-01-05 13:42:24.696555\",\"icon_url\":\"https://assets.chucknorris.host/img/avatar/chuck-norris.png\",\"id\":\"3Odt13-SQ06Pq-RUPwRQ4w\",\"updated_at\":\"2020-01-05 13:42:24.696555\",\"url\":\"https://api.chucknorris.io/jokes/3Odt13-SQ06Pq-RUPwRQ4w\",\"value\":\"Peace disturbs Chuck Norris.\"}"
				mock.EXPECT().SendGetRequest("https://api.chucknorris.io/jokes/random", nil).Return([]byte(reply), nil).Times(1)
			},
			want: "Chuck Norris Joke: Peace disturbs Chuck Norris.",
		},
		{
			name:  "GetDadJokeFromAPI",
			index: 2,
			mockCall: func() {
				reply := "{\"id\":\"0DQKB51oGlb\",\"joke\":\"What did one nut say as he chased another nut?  I'm a cashew!\",\"status\":200}"
				mock.EXPECT().SendGetRequest("https://icanhazdadjoke.com/", map[string]string{"Accept": "application/json"}).Return([]byte(reply), nil).Times(1)
			},
			want: "Dad Joke: What did one nut say as he chased another nut?  I'm a cashew!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCall()
			if got := getMessage(mock, tt.index); got != tt.want {
				t.Errorf("%v: got \"%v\", want \"%v\"", tt.name, got, tt.want)
			}
		})
	}
}
