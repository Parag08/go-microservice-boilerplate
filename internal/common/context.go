package common

import "context"

type Context struct {
	Context      context.Context
	RequestToken *string
	Version      *string
}
