package streaming

import (
	"fmt"
)

type NoConnectionError struct{}

func (err *NoConnectionError) Error() string {
	return fmt.Sprintf("No stream connection exists.")
}
