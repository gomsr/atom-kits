package ipx

import (
	"fmt"
	"testing"
)

func TestParseAgent(t *testing.T) {

	agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
	fmt.Println(ParseAgent(agent))

	agent2 := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3803.0 Safari/537.36 Edg/76.0.174.0"
	fmt.Println(ParseAgent(agent2))

	agent3 := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
	fmt.Println(ParseAgent(agent3))
}
