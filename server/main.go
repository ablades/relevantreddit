package main

import (
	"fmt"
)

/*
//TokenRequest stores authentication request
type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

//Store user credentials for easier access
type credentials struct {
	SecretKey string
	Client    string
	Username  string
	Password  string
}

//Sends an http request returns response in bytes
func sendRequest(request *http.Request) []byte {

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

//Load enviornment variables
func loadEnvironment() credentials {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	var cred = credentials{}

	cred.SecretKey = os.Getenv("SCRIPT_REDDIT_SECRET")
	cred.Client = os.Getenv("SCRIPT_CLIENT")
	cred.Username = os.Getenv("USERNAME")
	cred.Password = os.Getenv("PASSWORD")

	return cred

}

//Request a token from reddit server
func requestToken(creds credentials) token {
	//Load Environment Variables

	body := strings.NewReader(fmt.Sprintf("grant_type=password&username=%s&password=%s", creds.Username, creds.Password))

	//Create new http post request
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	if err != nil {
		log.Fatal(err)
	}

	//Set authorization request
	req.SetBasicAuth(creds.Client, creds.SecretKey)

	//header entries
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	content := sendRequest(req)
	//Create empty token request variable
	var tokenRequest = token{}

	//parse json into token struct
	json.Unmarshal(content, &tokenRequest)

	return tokenRequest
}

//variadic argument - closest thing to an optional argument accepts a variable number of arguments usd as a list
func useToken(t token, creds credentials, after ...string) subreddits {
	//variables init
	var req *http.Request
	var err error

	//Pull a users subreddits
	if len(after) == 1 {
		req, err = http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=%s", after[0]), nil)
	} else {
		req, err = http.NewRequest("GET", "https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100", nil)
	}
	if err != nil {
		log.Fatal(err)
	}
	//Set required headers
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", t.TokenType, t.AccessToken))

	//send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//convert response
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rc = subreddits{}

	json.Unmarshal(content, &rc)
	subcribedReddits(rc)

	return rc
}

*/
//displays subreddit
func subcribedReddits(rc subreddits, u *UserProfile) {
	//Loop through all of a requests subreddits
	for _, item := range rc.Data.Children {
		fmt.Printf("Added: %s to user profile. \n", item.Data.DisplayNamePrefixed)

		u.Subreddits[item.Data.DisplayNamePrefixed] = nil

		//parseSubreddit(item)
	}

	fmt.Println(rc.Data.After)
}

/*
func parseSubreddit(reddit struct{}) {

}
*/
func main() {
	loadEnvironment()
	//credentials := loadEnvironment()
	redditMiddleware()
	//authRequest(credentials)
	/*
		credentials := loadEnvironment()

		token := requestToken(credentials)

		//Use token once
		rc := useToken(token, credentials)

		//send multiple requests till all are pulled
		for rc.Data.After != "" {
			rc = useToken(token, credentials, rc.Data.After)
		}*/
}

// You can get up to 100 by passing limit, like:

// http://www.reddit.com/r/pics/.json?limit=100
// If you want more than that, look at the after parameter in the JSON that comes back, and call it again with that, like this:

// http://www.reddit.com/r/pics/.json?limit=100&after=t3_abcde
