package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nal/go/tgbot/logger"
)

var (
	tgbotLogger *logger.Logger
	err         error
)

func init() {
	tgbotLogger, err = logger.Init("log")
	if err != nil {
		log.Fatalf("Failed to init logger: %s\n", err)
	}
	fmt.Println("Initalized logger...")
}

// Finally, the main funtion starts our server on port 8089
// TODO: use .env file to store sensitive information
func main() {
	defer tgbotLogger.Defer() // flushes buffer, if any

	tgbotLogger.Info("Running on http://127.0.0.1:8089/")
	log.Fatal(http.ListenAndServe("127.0.0.1:8089", http.HandlerFunc(Handler)))
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	tgbotLogger.Info("Receved new request")
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		// Log error
		tgbotLogger.Info("Could not decode request body: ",
			"body", req.Body,
			"err", err)

		// Return smth useful for testing
		fmt.Fprintf(res, "%q", "Invalid request")
		return
	}

	// Create the request body struct
	response := sendMessageReqBody{
		ChatID: body.Message.Chat.ID,
		Text:   "Hello, " + body.Message.Text,
	}

	if err := sayEcho(response); err != nil {
		tgbotLogger.Info("Failed to send reply in func sayEcho: ",
			"err", err)

		// Return smth useful for testing
		fmt.Fprintf(res, "%q", "Failed to send reply")
		return
	}

	// Output the same response for testing
	fmt.Fprintf(res, "Hello, %q", body.Message.Text)

	// Log a confirmation message if the message was sent successfully
	tgbotLogger.Info("Reply sent...")
}

// sayEcho takes a chatID and sends echo to them
func sayEcho(response sendMessageReqBody) error {
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(&response)
	if err != nil {
		return err
	}

	tgbotToken := "ENV_TGBOT_API_TOKEN"
	tgbotApiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tgbotToken)

	// Send a post request with your token
	res, err := http.Post(tgbotApiUrl, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
