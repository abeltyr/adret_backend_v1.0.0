package middleware

import (
	"adr/backend/src/model"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
)

var body []byte

func Auth(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// fetch the token and the needed details
	godotenv.Load(".env")
	region := os.Getenv("AWS_REGION")
	userPool := os.Getenv("AWS_COGNITO_USER_POOL")
	url := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPool)
	token := c.Request().Header.Peek("Authorization")

	if token == nil {
		log.Println("no token")
		return errors.New("please provide authentication token")
	}

	if body == nil {

		// fetch the a pwt need to get the jwk key
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("api fetching issue: %s", err)
			return err
		}

		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to parse body from aws api request: %s", err)
			return err
		}
	}

	// parse the jwk from the api request body
	set, err := jwk.Parse(body)
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return err

	}

	// get the public key from the jwk
	var rsaData *rsa.PublicKey
	var ok bool
	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		var rawkey interface{} // This is the raw key, like *rsa.PrivateKey or *ecdsa.PrivateKey
		if err := key.Raw(&rawkey); err != nil {
			log.Printf("failed to create public key: %s", err)
		}
		// We know this is an RSA Key so...
		rsaData, ok = rawkey.(*rsa.PublicKey)
		if !ok {
			panic(fmt.Sprintf("expected ras key, got %T", rawkey))
		}
	}

	// verify the token against the public key
	payload, err := jws.Verify(token, jwa.RS256, rsaData)
	if err != nil {
		log.Printf("failed to verify message: %s", err)
	}

	//formate the jwt data
	var currentUser model.CognitoUser
	err = json.Unmarshal(payload, &currentUser)
	if err != nil {
		return err
	}

	// // check if the token has expired
	// currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	// if currentUser.Exp*1000 < currentTime {
	// 	return errors.New("token has expired")
	// }

	out, err := json.Marshal(currentUser)
	if err != nil {
		return err
	}

	// save the current user in the request
	c.Locals("user", string(out))

	c.Locals("accessToken", string(token))

	return c.Next()
}
