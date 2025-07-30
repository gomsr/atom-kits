package ipx

import (
	"fmt"
	"github.com/gookit/goutil/x/fmtutil"
	"testing"
)

func TestQueryFreeByIPInfo(t *testing.T) {

	got, err := QueryByIPInfo("2601:0243:21:1448:da65:6d64:23d7:f37c", "xxx")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmtutil.PrettyJSON(got))
}
