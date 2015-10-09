package treasures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
)

type Client struct {
	id, auth       string
	sessionId      int
	Episode, Level int
	vm             *otto.Otto
}

func NewClient() *Client {
	vm := otto.New()

	vm.Run(md5js)
	vm.Run(hexMD5)
	vm.Run(calcSig)

	vm.Run(`network = "vkontakte"`)
	vm.Run(`sessionId = 0`)

	client := &Client{
		vm: vm,
	}

	return client
}

type AuthInfo struct {
	SessionId int `json:"sessionId"`
	Episode   int `json:"episode"`
	Level     int `json:"level"`
}

func (c *Client) Authorize(id, auth string) error {
	signature := c.getSignature(map[string]string{
		"adsRef":      "",
		"userInviter": "",
		"os":          "Windows 64-bit",
		"browser":     "Firefox",
		"version":     "44",
	})

	params := url.Values{
		"adsRef":                   {""},
		"userInviter":              {""},
		"os":                       {"Windows 64-bit"},
		"browser":                  {"Firefox"},
		"version":                  {"44"},
		"mysig":                    {signature},
		"mysigparams":              {"adsRef|browser|os|userInviter|version"},
		"script":                   {"../../../levelbase/src/services/authorization.php"},
		"network":                  {"vkontakte"},
		"sessionid":                {"0"},
		"doublePointsBeforeLetter": {"1"},
	}

	url := fmt.Sprintf(
		"https://www111.orangeapps.ru/gems/gems/site/service/service.php?api_url=https://api.vk.com/api.php&sid=&access_token=&api_id=3882511&api_settings=2097159&viewer_id=%s&viewer_type=2&user_id=%s&group_id=0&is_app_user=1&auth_key=%s&language=0&parent_language=0&is_secure=1",
		id,
		id,
		auth,
	)

	r, err := http.PostForm(url, params)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var authInfo AuthInfo
	err = json.Unmarshal([]byte(body), &authInfo)
	if err != nil {
		return err
	}

	c.Level = authInfo.Level
	c.Episode = authInfo.Episode
	c.sessionId = authInfo.SessionId
	c.vm.Run(fmt.Sprintf("sessionId = %d", c.sessionId))

	return nil
}

func (c *Client) getSignature(params map[string]string) (signature string) {
	paramsList := make([]string, len(params))
	i := 0
	for key, value := range params {
		paramsList[i] = fmt.Sprintf(`%s:"%s"`, key, value)
		i++
	}
	args := strings.Join(paramsList, ",")

	command := fmt.Sprintf("CalcSig({%s})", args)

	result, _ := c.vm.Run(command)

	signature, _ = result.ToString()

	return
}

func (c *Client) SetCoins(coins int) error {
	serverTime := fmt.Sprintf("%d", time.Now().Unix())
	signature := c.getSignature(map[string]string{
		"server_time": serverTime,
		"coins":       strconv.Itoa(coins),
	})

	params := url.Values{
		"coins":       {strconv.Itoa(coins)},
		"server_time": {serverTime},
		"mysig":       {signature},
		"mysigparams": {"coins|server_time"},
		"script":      {"../../../levelbase/src/services/updateuser.php"},
	}

	resp, err := c.post(params)
	if err != nil {
		println(resp)
		panic(err)
	}

	return nil
}

func (c *Client) FinishLevel(episode, level, userLevel, score int) error {
	signature := c.getSignature(map[string]string{
		"episode":          strconv.Itoa(episode),
		"level":            strconv.Itoa(level),
		"userLevel":        strconv.Itoa(userLevel),
		"score":            strconv.Itoa(score),
		"withoutLoseSeria": "0",
	})

	params := url.Values{
		"episode":          {strconv.Itoa(episode)},
		"level":            {strconv.Itoa(level)},
		"userLevel":        {strconv.Itoa(userLevel)},
		"score":            {strconv.Itoa(score)},
		"withoutLoseSeria": {"0"},
		"mysig":            {signature},
		"mysigparams":      {"episode|level|score|userLevel|withoutLoseSeria"},
		"script":           {"../../../levelbase/src/services/addscore.php"},
	}

	_, err := c.post(params)

	return err
}

func (c *Client) BuyKeys(episode, coins int) error {
	signature := c.getSignature(map[string]string{
		"coins":   strconv.Itoa(coins),
		"episode": strconv.Itoa(episode),
		"level":   "21",
	})

	params := url.Values{
		"coins":       {strconv.Itoa(coins)},
		"episode":     {strconv.Itoa(episode)},
		"level":       {"21"},
		"mysig":       {signature},
		"mysigparams": {"coins|episode|level"},
		"script":      {"../../../levelbase/src/services/buykeys.php"},
	}

	_, err := c.post(params)

	return err
}

func (c *Client) post(params url.Values) (resp string, err error) {
	params.Set("sessionid", strconv.Itoa(c.sessionId))
	params.Set("network", "vkontakte")
	params.Set("doublePointsBeforeLetter", "1")

	url := fmt.Sprintf(
		"https://www111.orangeapps.ru/gems/gems/site/service/service.php?api_url=https://api.vk.com/api.php&sid=&access_token=&api_id=3882511&api_settings=2097159&viewer_id=%s&viewer_type=2&secret=0ce7b1e40f&user_id=%s&group_id=0&is_app_user=1&auth_key=%s&language=0&parent_language=0&ad_info=ElsdCQdcQlxlBQdbAwJSXHt6C0Q8HTJXUVBBJRVBNwoIFjI2HA8E&is_secure=1&ads_app_id=3882511_ea21030310f4b19893&referrer=user_apps&lc_name=cc6596bb&hash=",
		c.id,
		c.id,
		c.auth,
	)

	r, err := http.PostForm(url, params)
	if err != nil {
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	resp = string(body)

	return
}
