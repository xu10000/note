package help

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Get_claims(req *http.Request) jwt.MapClaims {

	_tokens, _ := req.Header["Authorization"]
	token_string := _tokens[0]
	token_segments := strings.Split(token_string, ".")

	claims, _ := jwt.DecodeSegment(token_segments[1])

	claims_map := jwt.MapClaims{}
	json.Unmarshal(claims, &claims_map)
	return claims_map

}
