package auth

import (
	"fmt"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin/binding"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/util"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	QueryAppKey    = "appkey"
	QueryTimestamp = "timestamp"
	QueryNonce     = "nonce"
	QuerySignature = "signature"
)

type Param struct {
	AppKey    string `form:"appkey" binding:"required"`
	Timestamp int64  `form:"timestamp" binding:"required"`
	Nonce     string `form:"nonce" binding:"required"`
	Signature string `form:"signature" binding:"required"`
}

type AppAuth struct {
	AppKey    string
	AppSecret string
	Cache     cache.CacheStore
}

func NewAppAuth(appKey, appSecret string) *AppAuth {
	return &AppAuth{
		AppKey:    appKey,
		AppSecret: appSecret,
		Cache:     cache.NewInMemoryStore(15 * time.Minute),
	}
}

func (auth *AppAuth) Add(r *http.Request) error {
	timestamp := time.Now().Unix() * 1000
	reqQuery := r.URL.Query()
	reqQuery.Add(QueryAppKey, auth.AppKey)
	reqQuery.Add(QueryTimestamp, strconv.FormatInt(timestamp, 10))
	reqQuery.Add(QueryNonce, auth.generateNonce())
	sign, err := auth.generateSignature(r.Method, r.URL.Path, reqQuery)
	if err != nil {
		return err
	}
	reqQuery.Add(QuerySignature, sign)
	r.URL.RawQuery = reqQuery.Encode()
	return nil
}

func (auth *AppAuth) Verify(r *http.Request) error {
	var param Param
	if err := binding.Query.Bind(r, &param); err != nil {
		return ErrAuthParam
	}
	timeExpired := true
	diff := time.Now().Unix() - param.Timestamp/1000
	// -15*60 <= time <= 15*60 允许客户端系统时间浮动
	if (diff >= 0 && diff <= 15*60) || (diff < 0 && diff >= -15*60) {
		timeExpired = false
	}
	if timeExpired {
		return ErrAuthParam
	}
	var nop struct{}
	if err := auth.Cache.Get(param.Nonce, &nop); err == nil {
		// err is nil if key is exists
		return ErrAuthParam
	}
	signComputed, err := auth.generateSignature(r.Method, r.URL.Path, r.URL.Query())
	if err != nil {
		return ErrAuthParam
	}
	if param.Signature != signComputed {
		return ErrAuthVerify
	}
	return nil
}

func (auth *AppAuth) generateSignature(method, path string, query url.Values) (string, error) {
	var queryList []string
	for key, val := range query {
		if key == QuerySignature {
			continue
		}
		val, err := url.QueryUnescape(val[0])
		if err != nil {
			return "", err
		}
		queryList = append(queryList, fmt.Sprintf("%s=%s", key, val))
	}
	sort.Strings(queryList)

	sortedQuery := strings.Join(queryList, "&")

	signString := fmt.Sprintf("%s\n%s\n%s", strings.ToUpper(method), path, sortedQuery)
	signBytes, err := util.HmacSha256([]byte(signString), []byte(auth.AppSecret))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", signBytes), nil
}

func (auth *AppAuth) generateNonce() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
