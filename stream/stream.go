package stream

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/satoshun/twitter/types"
)

type TwitterStreamAPI struct {
	*types.TwitterConfig
	Timestamp string
	Nonce     string
	track     string
}

const (
	FILTER_URL             = "https://stream.twitter.com/1.1/statuses/filter.json"
	SAMPLE_URL             = "https://stream.twitter.com/1.1/statuses/sample.json"
	OAUTH_VERSION          = "1.0"
	OAUTH_SIGNATURE_METHOD = "HMAC-SHA1"
)

func getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func generateNonce() string {
	rand.Seed(time.Now().Unix())
	return strconv.FormatInt(rand.Int63(), 16)
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(10*time.Second))
}

func NewTwitterStream(token, secret, consumerKey, consumerSecret string) *TwitterStreamAPI {
	return &TwitterStreamAPI{
		TwitterConfig: &types.TwitterConfig{
			Token:          token,
			TokenSecret:    secret,
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
		},
		Timestamp: getTimestamp(),
		Nonce:     generateNonce(),
	}
}

func (t *TwitterStreamAPI) Filter(track string) <-chan types.Tweet {
	t.track = track
	tweets := make(chan types.Tweet, 1000)

	go func() {
		defer close(tweets)
		waitingMill := 250
		retry := 0

		f := func() {
			req, _ := http.NewRequest("GET", t.targetPath(), nil)
			resp, err := t.do(req)
			if resp.StatusCode == 401 {
				panic("Authorization Error")
			}

			if err != nil {
				return
			}

			defer resp.Body.Close()
			var beforeB []byte
			reader := bufio.NewReaderSize(resp.Body, 4096*4)

			for {
				b, _, err := reader.ReadLine()

				if err != nil {
					break
				}

				if beforeB != nil {
					b = append(beforeB, b...)
				}

				tweet := new(types.Tweet)
				err = json.Unmarshal(b, tweet)
				if err != nil {
					beforeB = make([]byte, len(b))
					copy(beforeB, b)
					if len(b) >= 100000 {
						break
					}

					continue
				}

				retry = 0
				beforeB = nil
				tweets <- *tweet
			}
		}

		for {
			f()
			retry++
			time.Sleep(time.Duration(waitingMill*retry*2) * time.Millisecond)
			if retry >= 10 {
				break
			}
		}
	}()

	return tweets
}

func (t *TwitterStreamAPI) Sample() <-chan types.Tweet {
	return t.Filter("")
}

func (t *TwitterStreamAPI) do(req *http.Request) (*http.Response, error) {
	t.setHeader(req)

	transport := &http.Transport{
		Dial: dialTimeout,
	}
	client := &http.Client{Transport: transport}

	return client.Do(req)
}

func (t *TwitterStreamAPI) setHeader(req *http.Request) {
	a := "OAuth " +
		fmt.Sprintf("oauth_consumer_key=%s, ", t.ConsumerKey) +
		fmt.Sprintf("oauth_nonce=%s, ", t.Nonce) +
		fmt.Sprintf("oauth_signature=%s, ", t.signature()) +
		fmt.Sprintf("oauth_signature_method=%s, ", OAUTH_SIGNATURE_METHOD) +
		fmt.Sprintf("oauth_timestamp=%s, ", t.Timestamp) +
		fmt.Sprintf("oauth_token=%s, ", t.Token) +
		fmt.Sprintf("oauth_version=%s", OAUTH_VERSION)

	req.Header.Set("Authorization", a)
}

func (t *TwitterStreamAPI) signature() string {
	values := url.Values{}
	values.Add("oauth_consumer_key", t.ConsumerKey)
	values.Add("oauth_nonce", t.Nonce)
	values.Add("oauth_signature_method", OAUTH_SIGNATURE_METHOD)
	values.Add("oauth_timestamp", t.Timestamp)
	values.Add("oauth_token", t.Token)
	values.Add("oauth_version", OAUTH_VERSION)

	if t.track != "" {
		values.Add("track", t.track)
	}

	baseString := "GET&" + url.QueryEscape(t.targetURL()) + "&" + url.QueryEscape(values.Encode())
	signKey := []byte(t.ConsumerSecret + "&" + t.TokenSecret)
	mac := hmac.New(sha1.New, signKey)
	mac.Write([]byte(baseString))
	enc := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	enc = url.QueryEscape(enc)

	return enc
}

func (t *TwitterStreamAPI) targetURL() string {
	if t.track == "" {
		return SAMPLE_URL
	}

	return FILTER_URL
}

func (t *TwitterStreamAPI) targetPath() string {
	if t.track == "" {
		return t.targetURL()
	}

	return t.targetURL() + "?track=" + t.track
}
