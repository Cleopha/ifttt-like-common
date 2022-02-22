package common

import (
	"errors"
	"github.com/Cleopha/ifttt-like-common/protos"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

var (
	ErrInvalidTaskFormat = errors.New("")
)

func ParseAction(task *protos.Task) (string, string, *structpb.Struct, error) {
	split := strings.Split(task.Action.String(), "_")
	if len(split) < 2 {
		return "", "", nil, ErrInvalidTaskFormat
	}

	method := ""
	namespace := strings.ToLower(split[0])

	for _, elem := range split[1:] {
		method += string(elem[0]) + strings.ToLower(elem[1:])
	}

	return namespace, method, task.Params, nil
}
