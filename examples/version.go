//
//  Report 0MQ version.
//

package main

import (
	zmq "zmq4"

	"fmt"
)

func main() {
	major, minor, patch := zmq.Version()
	fmt.Printf("Current 0MQ version is %d.%d.%d\n", major, minor, patch)
}
