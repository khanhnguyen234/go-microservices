package common
 
import (
	"context"
	"time"
)
func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	return ctx
}