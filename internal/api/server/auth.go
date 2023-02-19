package server

import (
	"encoding/base64"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

const basicAuthCredentials = "admin:admin"

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
)

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Basic")
	return token == base64.StdEncoding.EncodeToString([]byte(basicAuthCredentials))
}
