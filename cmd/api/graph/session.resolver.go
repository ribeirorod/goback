package graph

import (
	"fmt"
	"go-server/cmd/api/middlewares"

	"github.com/graphql-go/graphql"
)

// Resolves SessionInit by checking input for cookie and return a session object
func SessionResolver(params graphql.ResolveParams) (interface{}, error) {
	sid, err := params.Args["sid"].(string)
	if err {
		// get session from context
		k := middlewares.CtxKey("sid")
		sid = params.Context.Value(k).(string)
	}

	sessionManager := middlewares.UserSessions
	session, _ := sessionManager.SessionRead(sid)
	if session == nil {
		return nil, fmt.Errorf("session not found")
	}

	return sid, nil

}
