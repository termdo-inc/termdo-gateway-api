package helpers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"termdo.com/gateway-api/source/app/constants"
	"termdo.com/gateway-api/source/app/utils"
)

func NormalizeTokenResponse(
	ctx *gin.Context,
	rescp *utils.ResponseCapture,
	remoteToken *string,
) {
	var body map[string]any

	if err := json.Unmarshal(rescp.Buffer.Bytes(), &body); err != nil {
		body = make(map[string]any)
	}

	var token *string
	if remoteToken != nil {
		token = remoteToken
	} else if value, ok := body[constants.BodyTokenKey]; ok {
		switch v := value.(type) {
		case string:
			token = &v
		case *string:
			if v != nil {
				token = v
			}
		}
	}

	if IsBrowserClient(ctx) {
		if token != nil {
			SetAuthCookie(ctx, *token)
		}
		delete(body, constants.BodyTokenKey)
	} else {
		if remoteToken != nil {
			body[constants.BodyTokenKey] = remoteToken
		} else if token != nil {
			body[constants.BodyTokenKey] = token
		}
	}

	newBody, _ := json.Marshal(body)

	rescp.Buffer.Reset()
	rescp.Buffer.Write(newBody)
}
