package ipx

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	if info, err := Query("", "127.0.0.1"); err != nil {
		panic(err)
	} else {
		fmt.Println(info)
	}

	if info, err := Query("", "138.197.101.126"); err != nil {
		panic(err)
	} else {
		fmt.Println(info)
	}
}

func TestQueryWithKey(t *testing.T) {

	if info, err := QueryWithKey("138.197.101.126", "xxxx"); err != nil {
		panic(err)
	} else {
		fmt.Println(info)
	}

	if info, err := QueryWithKey("127.0.0.1", "xx"); err != nil {
		panic(err)
	} else {
		fmt.Println(info)
	}
}
