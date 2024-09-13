package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
)

const envFile = ".tmux-spotify-env"

// const version="<<development version>>"

var (
	auth  spotify.Authenticator
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(homeDir + "/" + envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	auth = spotify.NewAuthenticator("http://localhost:8080/callback",
		spotify.ScopeUserReadCurrentlyPlaying)
	auth.SetAuthInfo(clientID, clientSecret)

	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	for {
		updateTmuxStatus(client)
		time.Sleep(5 * time.Second)
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func updateTmuxStatus(client *spotify.Client) {
	currentlyPlaying, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		log.Println("Error getting currently playing track:", err)
		return
	}

	if currentlyPlaying.Item == nil {
		return
	}

	trackInfo := fmt.Sprintf("%s - %s", currentlyPlaying.Item.Name, currentlyPlaying.Item.Artists[0].Name)

	cmd := exec.Command("tmux", "set", "-g", "status-right", fmt.Sprintf("#[fg=blue]♫ %s", trackInfo))
	err = cmd.Run()
	if err != nil {
		log.Println("Error updating tmux status:", err)
	}
}
