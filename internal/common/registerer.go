// Package common
package common

import "context"

type Registerer interface {
	Set(int, int, int, func(context.Context) error)
}
