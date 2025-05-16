package currencyx

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	"github.com/gookit/goutil/jsonutil"
	"testing"
)

func TestDoExchangeRate(t *testing.T) {

	key, _ := jasypt.New().Decrypt("QON+AGZKt9HAqLxc5i7Ye+PyCaWP9nfoZ2av+CkqOdhl+9HgegOYsxZC40oGWKVrjUq3NAj+sg2xGWMf7wFh9Q==")
	rate, err := DoExchangeRate4Usd(key)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonutil.MustString(rate))

	convert2Map, err := Convert2Map(rate.ConversionRates)
	if err != nil {
		return
	}
	fmt.Println(jsonutil.MustString(convert2Map))
}
