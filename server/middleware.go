package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//Handles request/response intermidary actions
func redditMiddleware() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/RedditLogin", handleRedditLogin)
	http.HandleFunc("/RedditCallback", handleRedditCallback)
	fmt.Println(http.ListenAndServe(":3000", nil))

}

//temp till connect react frontend
func handleMain(w http.ResponseWriter, r *http.Request) {
	const htmlIndex = `
	<html>
		<body>
			<a href="/RedditLogin">Authenticate with Reddit</a>
		</body>
	</html>
	`
	fmt.Fprintf(w, htmlIndex)
}

//Redirect user to authorization page
func handleRedditLogin(w http.ResponseWriter, r *http.Request) {

	url, err := url.Parse("reddit.com/api/v1/authorize")
	if err != nil {
		log.Fatal(err)
	}
	url.Scheme = "https"
	q := url.Query()
	q.Add("client_id", creds.Client)
	q.Add("response_type", "code")
	q.Add("state", "foobar")              //verify user is user CSRF
	q.Add("redirect_uri", creds.Redirect) //temp redirect url
	q.Add("duration", "temporary")        //temp for now may switch to perm later
	q.Add("scope", "mysubreddits identity history")

	url.RawQuery = q.Encode()
	fmt.Printf("Redirecting to: %s \n", url)

	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}

//Handle unser response from reddit
func handleRedditCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	//Get first parameter with query name
	code := r.FormValue("code")
	state := r.FormValue("state") //TODO : Verify states are same
	fmt.Println(state)
	fmt.Println(code)
	//
	token := requestToken(code)

	var appUser UserProfile
	appUser.Subreddits = make(map[string][]string)

	url := "https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100"
	rc := useToken(token, url)

	subcribedReddits(rc, &appUser)

	//send multiple requests till all are pulled
	for rc.Data.After != "" {
		url = fmt.Sprintf("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=%s", rc.Data.After)
		rc = useToken(token, url)

		subcribedReddits(rc, &appUser)
	}

	endpoint := "https://oauth.reddit.com/api/v1/me"

	appUser.RedditName = getUserInfo(token, endpoint).Name

	fmt.Print(appUser)

}
