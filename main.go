package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"userinfo-api/config"
)

func main() {
	config.Load()
	token := config.Get("token")

	client, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
		return
	}

	client.AddHandler(handleAPI)
	err = client.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("[BOT] Online")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = client.Close()
}

func handleAPI(s *discordgo.Session, e *discordgo.Ready) {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		_ = json.NewEncoder(writer).Encode(Alive{Alive: true})
	}).Methods("GET")

	r.HandleFunc("/api/user/{user}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)

		user, err := s.User(vars["user"])
		if err != nil {
			if err.Error() == "HTTP 404 Not Found, {\"message\": \"Unknown User\", \"code\": 10013}" {
				writer.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(writer).Encode(UserNotFoundResponse{
					Success: false,
					Message: "User not found",
				})
				return
			} else {
				log.Println(err)
			}
		}

		_ = json.NewEncoder(writer).Encode(Response{
			Success: true,
			Message: UserResponse{
				Tag: user.Username + "#" + user.Discriminator,
				Id:  user.ID,
				Bot: user.Bot,
				User: User{
					Username:      user.Username,
					Discriminator: user.Discriminator,
					Avatar:        user.AvatarURL("128"),
				},
			},
		})
		return
	}).Methods("GET")

	err := http.ListenAndServe(":2011", r)
	if err != nil {
		log.Fatal(err)
	}
}
