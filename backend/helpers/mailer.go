package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/golang-jwt/jwt/v5"
)

func createMailerJWT() (string, error) {
	token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "backend",
		"crt": time.Now().Unix(),
		"exp": time.Now().Add(config.ACCESS_TOKEN_TTL).Unix(),
	})

	token, err := token_claim.SignedString([]byte(initializers.CONFIG.MAILER_SECRET))
	if err != nil {
		return "", err
	}

	return token, nil
}

func SendMailReq(email string, emailType int, user *models.User, entityType string, data interface{}) error {
	payload := map[string]interface{}{
        "email": email,
        "type":  emailType,
        "user":  user,
    }

    switch entityType {
    case "coment":
        payload["comment"] = data
    case "post":
        payload["post"] = data
    case "project":
        payload["project"] = data
	case "event":
        payload["event"] = data
	case "opening":
        payload["opening"] = data
	case "announcement":
        payload["announcement"] = data
	case "poll":
        payload["poll"] = data
    default:
        return fmt.Errorf("unsupported additional data type: %s", entityType)
    }

    jsonData, err := json.Marshal(payload)

	if err != nil {
		config.Logger.Errorw("Error calling Mailer", err, "SendMailReq")
		return err
	}

	jwt, err := createMailerJWT()
	if err != nil {
		config.Logger.Errorw("Error calling Mailer", err, "SendMailReq")
		return err
	}

	request, err := http.NewRequest("POST", initializers.CONFIG.MAILER_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		config.Logger.Errorw("Error calling Mailer", err, "SendMailReq")
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+jwt)
	request.Header.Set("api-token", initializers.CONFIG.MAILER_TOKEN)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		config.Logger.Errorw("Error calling Mailer", err, "SendMailReq")
		return err
	}
	defer response.Body.Close()

	var responseBody struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	if response.StatusCode != 200 {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&responseBody); err == nil {
			config.Logger.Errorw("Error calling Mailer", fmt.Errorf(fmt.Sprint("Status Code: ", response.StatusCode, ", Message: ", responseBody.Message)), "SendMailReq")
		} else {
			config.Logger.Errorw("Error calling Mailer", fmt.Errorf(fmt.Sprint("Status Code: ", response.StatusCode)), "SendMailReq")
		}
		return fmt.Errorf("error calling mailer")
	}

	return nil
}
