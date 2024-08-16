package goUtility

import (
	"fmt"
)

func HandleError(err error, msg string, usePanic bool) {
	if err != nil {

		if msg != "" {
			fmt.Println(msg)
		}
		if usePanic {
			panic(err)
		}
		fmt.Println(fmt.Errorf("error: %w", err))

	}

}
