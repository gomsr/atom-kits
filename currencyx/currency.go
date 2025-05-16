package currencyx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const keyUrl = "https://v6.exchangerate-api.com/v6/%s/latest/%s"

// Cache holds the exchange rates with a TTL
type Cache struct {
	Rates     map[string]float64
	Timestamp time.Time
}

var (
	CacheTTL     = time.Minute * 30    // Cache time-to-live
	DefaultRates = map[string]float64{ // 202409002 rates
		"USD": 1,
		"AED": 3.6725,
		"AFN": 70.5967,
		"ALL": 90.0267,
		"AMD": 387.6498,
		"ANG": 1.7900,
		"AOA": 919.1545,
		"ARS": 952.8300,
		"AUD": 1.4769,
		"AWG": 1.7900,
		"AZN": 1.6994,
		"BAM": 1.7698,
		"BBD": 2.0000,
		"BDT": 119.4636,
		"BGN": 1.7698,
		"BHD": 0.3760,
		"BIF": 2878.9290,
		"BMD": 1.0000,
		"BND": 1.3057,
		"BOB": 6.9190,
		"BRL": 5.6129,
		"BSD": 1.0000,
		"BTN": 83.9074,
		"BWP": 13.2731,
		"BYN": 3.2564,
		"BZD": 2.0000,
		"CAD": 1.3486,
		"CDF": 2846.4764,
		"CHF": 0.8504,
		"CLP": 913.9279,
		"CNY": 7.1011,
		"COP": 4145.0138,
		"CRC": 519.0559,
		"CUP": 24.0000,
		"CVE": 99.7792,
		"CZK": 22.6377,
		"DJF": 177.7210,
		"DKK": 6.7515,
		"DOP": 59.6098,
		"DZD": 133.9012,
		"EGP": 48.5959,
		"ERN": 15.0000,
		"ETB": 111.1551,
		"EUR": 0.9050,
		"FJD": 2.1999,
		"FKP": 0.7619,
		"FOK": 6.7512,
		"GBP": 0.7619,
		"GEL": 2.6899,
		"GGP": 0.7619,
		"GHS": 15.6729,
		"GIP": 0.7619,
		"GMD": 70.1383,
		"GNF": 8699.0228,
		"GTQ": 7.7292,
		"GYD": 209.2370,
		"HKD": 7.7982,
		"HNL": 24.7883,
		"HRK": 6.8180,
		"HTG": 131.7185,
		"HUF": 354.9326,
		"IDR": 15503.9359,
		"ILS": 3.6469,
		"IMP": 0.7619,
		"INR": 83.9045,
		"IQD": 1309.3857,
		"IRR": 42012.9876,
		"ISK": 138.1787,
		"JEP": 0.7619,
		"JMD": 157.0576,
		"JOD": 0.7090,
		"JPY": 146.2394,
		"KES": 128.7571,
		"KGS": 85.1136,
		"KHR": 4067.4058,
		"KID": 1.4765,
		"KMF": 445.1834,
		"KRW": 1335.9617,
		"KWD": 0.3053,
		"KYD": 0.8333,
		"KZT": 481.9196,
		"LAK": 21913.9659,
		"LBP": 89500.0000,
		"LKR": 298.9345,
		"LRD": 195.0560,
		"LSL": 17.8092,
		"LYD": 4.7601,
		"MAD": 9.7440,
		"MDL": 17.4009,
		"MGA": 4541.9700,
		"MKD": 55.3955,
		"MMK": 2101.1640,
		"MNT": 3352.2975,
		"MOP": 8.0321,
		"MRU": 39.7709,
		"MUR": 46.3954,
		"MVR": 15.4389,
		"MWK": 1735.9596,
		"MXN": 19.7225,
		"MYR": 4.3202,
		"MZN": 63.6764,
		"NAD": 17.8092,
		"NGN": 1594.4098,
		"NIO": 36.8990,
		"NOK": 10.6019,
		"NPR": 134.2518,
		"NZD": 1.6006,
		"OMR": 0.3845,
		"PAB": 1.0000,
		"PEN": 3.7477,
		"PGK": 3.8976,
		"PHP": 56.2018,
		"PKR": 278.7689,
		"PLN": 3.8718,
		"PYG": 7600.7611,
		"QAR": 3.6400,
		"RON": 4.4919,
		"RSD": 105.7406,
		"RUB": 90.7650,
		"RWF": 1336.6490,
		"SAR": 3.7500,
		"SBD": 8.5130,
		"SCR": 13.8141,
		"SDG": 458.8843,
		"SEK": 10.2653,
		"SGD": 1.3041,
		"SHP": 0.7619,
		"SLE": 22.6041,
		"SLL": 22604.0599,
		"SOS": 571.3137,
		"SRD": 28.9949,
		"SSP": 3130.9772,
		"STN": 22.1701,
		"SYP": 12852.7576,
		"SZL": 17.8092,
		"THB": 34.0695,
		"TJS": 10.6061,
		"TMT": 3.4995,
		"TND": 3.0512,
		"TOP": 2.2984,
		"TRY": 34.0982,
		"TTD": 6.7706,
		"TVD": 1.4765,
		"TWD": 31.9248,
		"TZS": 2719.8335,
		"UAH": 41.0456,
		"UGX": 3720.1853,
		"UYU": 40.3458,
		"UZS": 12697.0095,
		"VES": 36.6454,
		"VND": 24877.7719,
		"VUV": 117.8039,
		"WST": 2.7036,
		"XAF": 593.5778,
		"XCD": 2.7000,
		"XDR": 0.7425,
		"XOF": 593.5778,
		"XPF": 107.9839,
		"YER": 250.2550,
		"ZAR": 17.8116,
		"ZMW": 26.2067,
		"ZWL": 13.8546,
	}
)

type ExchangeRateResp struct {
	Result          string           `json:"result"`
	BaseCode        string           `json:"base_code"`
	ConversionRates *ConversionRates `json:"conversion_rates"`
}

type ConversionRates struct {
	USD float64 `json:"USD"`
	AED float64 `json:"AED"`
	AFN float64 `json:"AFN"`
	ALL float64 `json:"ALL"`
	AMD float64 `json:"AMD"`
	ANG float64 `json:"ANG"`
	AOA float64 `json:"AOA"`
	ARS float64 `json:"ARS"`
	AUD float64 `json:"AUD"`
	AWG float64 `json:"AWG"`
	AZN float64 `json:"AZN"`
	BAM float64 `json:"BAM"`
	BBD float64 `json:"BBD"`
	BDT float64 `json:"BDT"`
	BGN float64 `json:"BGN"`
	BHD float64 `json:"BHD"`
	BIF float64 `json:"BIF"`
	BMD float64 `json:"BMD"`
	BND float64 `json:"BND"`
	BOB float64 `json:"BOB"`
	BRL float64 `json:"BRL"`
	BSD float64 `json:"BSD"`
	BTN float64 `json:"BTN"`
	BWP float64 `json:"BWP"`
	BYN float64 `json:"BYN"`
	BZD float64 `json:"BZD"`
	CAD float64 `json:"CAD"`
	CDF float64 `json:"CDF"`
	CHF float64 `json:"CHF"`
	CLP float64 `json:"CLP"`
	CNY float64 `json:"CNY"`
	COP float64 `json:"COP"`
	CRC float64 `json:"CRC"`
	CUP float64 `json:"CUP"`
	CVE float64 `json:"CVE"`
	CZK float64 `json:"CZK"`
	DJF float64 `json:"DJF"`
	DKK float64 `json:"DKK"`
	DOP float64 `json:"DOP"`
	DZD float64 `json:"DZD"`
	EGP float64 `json:"EGP"`
	ERN float64 `json:"ERN"`
	ETB float64 `json:"ETB"`
	EUR float64 `json:"EUR"`
	FJD float64 `json:"FJD"`
	FKP float64 `json:"FKP"`
	FOK float64 `json:"FOK"`
	GBP float64 `json:"GBP"`
	GEL float64 `json:"GEL"`
	GGP float64 `json:"GGP"`
	GHS float64 `json:"GHS"`
	GIP float64 `json:"GIP"`
	GMD float64 `json:"GMD"`
	GNF float64 `json:"GNF"`
	GTQ float64 `json:"GTQ"`
	GYD float64 `json:"GYD"`
	HKD float64 `json:"HKD"`
	HNL float64 `json:"HNL"`
	HRK float64 `json:"HRK"`
	HTG float64 `json:"HTG"`
	HUF float64 `json:"HUF"`
	IDR float64 `json:"IDR"`
	ILS float64 `json:"ILS"`
	IMP float64 `json:"IMP"`
	INR float64 `json:"INR"`
	IQD float64 `json:"IQD"`
	IRR float64 `json:"IRR"`
	ISK float64 `json:"ISK"`
	JEP float64 `json:"JEP"`
	JMD float64 `json:"JMD"`
	JOD float64 `json:"JOD"`
	JPY float64 `json:"JPY"`
	KES float64 `json:"KES"`
	KGS float64 `json:"KGS"`
	KHR float64 `json:"KHR"`
	KID float64 `json:"KID"`
	KMF float64 `json:"KMF"`
	KRW float64 `json:"KRW"`
	KWD float64 `json:"KWD"`
	KYD float64 `json:"KYD"`
	KZT float64 `json:"KZT"`
	LAK float64 `json:"LAK"`
	LBP float64 `json:"LBP"`
	LKR float64 `json:"LKR"`
	LRD float64 `json:"LRD"`
	LSL float64 `json:"LSL"`
	LYD float64 `json:"LYD"`
	MAD float64 `json:"MAD"`
	MDL float64 `json:"MDL"`
	MGA float64 `json:"MGA"`
	MKD float64 `json:"MKD"`
	MMK float64 `json:"MMK"`
	MNT float64 `json:"MNT"`
	MOP float64 `json:"MOP"`
	MRU float64 `json:"MRU"`
	MUR float64 `json:"MUR"`
	MVR float64 `json:"MVR"`
	MWK float64 `json:"MWK"`
	MXN float64 `json:"MXN"`
	MYR float64 `json:"MYR"`
	MZN float64 `json:"MZN"`
	NAD float64 `json:"NAD"`
	NGN float64 `json:"NGN"`
	NIO float64 `json:"NIO"`
	NOK float64 `json:"NOK"`
	NPR float64 `json:"NPR"`
	NZD float64 `json:"NZD"`
	OMR float64 `json:"OMR"`
	PAB float64 `json:"PAB"`
	PEN float64 `json:"PEN"`
	PGK float64 `json:"PGK"`
	PHP float64 `json:"PHP"`
	PKR float64 `json:"PKR"`
	PLN float64 `json:"PLN"`
	PYG float64 `json:"PYG"`
	QAR float64 `json:"QAR"`
	RON float64 `json:"RON"`
	RSD float64 `json:"RSD"`
	RUB float64 `json:"RUB"`
	RWF float64 `json:"RWF"`
	SAR float64 `json:"SAR"`
	SBD float64 `json:"SBD"`
	SCR float64 `json:"SCR"`
	SDG float64 `json:"SDG"`
	SEK float64 `json:"SEK"`
	SGD float64 `json:"SGD"`
	SHP float64 `json:"SHP"`
	SLE float64 `json:"SLE"`
	SLL float64 `json:"SLL"`
	SOS float64 `json:"SOS"`
	SRD float64 `json:"SRD"`
	SSP float64 `json:"SSP"`
	STN float64 `json:"STN"`
	SYP float64 `json:"SYP"`
	SZL float64 `json:"SZL"`
	THB float64 `json:"THB"`
	TJS float64 `json:"TJS"`
	TMT float64 `json:"TMT"`
	TND float64 `json:"TND"`
	TOP float64 `json:"TOP"`
	TRY float64 `json:"TRY"`
	TTD float64 `json:"TTD"`
	TVD float64 `json:"TVD"`
	TWD float64 `json:"TWD"`
	TZS float64 `json:"TZS"`
	UAH float64 `json:"UAH"`
	UGX float64 `json:"UGX"`
	UYU float64 `json:"UYU"`
	UZS float64 `json:"UZS"`
	VES float64 `json:"VES"`
	VND float64 `json:"VND"`
	VUV float64 `json:"VUV"`
	WST float64 `json:"WST"`
	XAF float64 `json:"XAF"`
	XCD float64 `json:"XCD"`
	XDR float64 `json:"XDR"`
	XOF float64 `json:"XOF"`
	XPF float64 `json:"XPF"`
	YER float64 `json:"YER"`
	ZAR float64 `json:"ZAR"`
	ZMW float64 `json:"ZMW"`
	ZWL float64 `json:"ZWL"`
}

func DoExchangeRate4Usd(key string) (*ExchangeRateResp, error) {
	return DoExchangeRate("USD", key)
}

func DoExchangeRate(base string, apiKey string) (*ExchangeRateResp, error) {

	url := fmt.Sprintf(keyUrl, apiKey, base)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return nil, err
	}

	rateResp := &ExchangeRateResp{}
	err = json.Unmarshal(body, rateResp)
	if err != nil {
		return nil, err
	}

	if rateResp.Result != "success" {
		return nil, errors.New("failed to get exchange rate")
	}

	return rateResp, nil
}

func Convert2Map(rates *ConversionRates) (map[string]float64, error) {

	if rates == nil {
		return nil, errors.New("rates is nil")
	}

	return map[string]float64{
		"USD": rates.USD,
		"AED": rates.AED,
		"AFN": rates.AFN,
		"ALL": rates.ALL,
		"AMD": rates.AMD,
		"ANG": rates.ANG,
		"AOA": rates.AOA,
		"ARS": rates.ARS,
		"AUD": rates.AUD,
		"AWG": rates.AWG,
		"AZN": rates.AZN,
		"BAM": rates.BAM,
		"BBD": rates.BBD,
		"BDT": rates.BDT,
		"BGN": rates.BGN,
		"BHD": rates.BHD,
		"BIF": rates.BIF,
		"BMD": rates.BMD,
		"BND": rates.BND,
		"BOB": rates.BOB,
		"BRL": rates.BRL,
		"BSD": rates.BSD,
		"BTN": rates.BTN,
		"BWP": rates.BWP,
		"BYN": rates.BYN,
		"BZD": rates.BZD,
		"CAD": rates.CAD,
		"CDF": rates.CDF,
		"CHF": rates.CHF,
		"CLP": rates.CLP,
		"CNY": rates.CNY,
		"COP": rates.COP,
		"CRC": rates.CRC,
		"CUP": rates.CUP,
		"CVE": rates.CVE,
		"CZK": rates.CZK,
		"DJF": rates.DJF,
		"DKK": rates.DKK,
		"DOP": rates.DOP,
		"DZD": rates.DZD,
		"EGP": rates.EGP,
		"ERN": rates.ERN,
		"ETB": rates.ETB,
		"EUR": rates.EUR,
		"FJD": rates.FJD,
		"FKP": rates.FKP,
		"FOK": rates.FOK,
		"GBP": rates.GBP,
		"GEL": rates.GEL,
		"GGP": rates.GGP,
		"GHS": rates.GHS,
		"GIP": rates.GIP,
		"GMD": rates.GMD,
		"GNF": rates.GNF,
		"GTQ": rates.GTQ,
		"GYD": rates.GYD,
		"HKD": rates.HKD,
		"HNL": rates.HNL,
		"HRK": rates.HRK,
		"HTG": rates.HTG,
		"HUF": rates.HUF,
		"IDR": rates.IDR,
		"ILS": rates.ILS,
		"IMP": rates.IMP,
		"INR": rates.INR,
		"IQD": rates.IQD,
		"IRR": rates.IRR,
		"ISK": rates.ISK,
		"JEP": rates.JEP,
		"JMD": rates.JMD,
		"JOD": rates.JOD,
		"JPY": rates.JPY,
		"KES": rates.KES,
		"KGS": rates.KGS,
		"KHR": rates.KHR,
		"KID": rates.KID,
		"KMF": rates.KMF,
		"KRW": rates.KRW,
		"KWD": rates.KWD,
		"KYD": rates.KYD,
		"KZT": rates.KZT,
		"LAK": rates.LAK,
		"LBP": rates.LBP,
		"LKR": rates.LKR,
		"LRD": rates.LRD,
		"LSL": rates.LSL,
		"LYD": rates.LYD,
		"MAD": rates.MAD,
		"MDL": rates.MDL,
		"MGA": rates.MGA,
		"MKD": rates.MKD,
		"MMK": rates.MMK,
		"MNT": rates.MNT,
		"MOP": rates.MOP,
		"MRU": rates.MRU,
		"MUR": rates.MUR,
		"MVR": rates.MVR,
		"MWK": rates.MWK,
		"MXN": rates.MXN,
		"MYR": rates.MYR,
		"MZN": rates.MZN,
		"NAD": rates.NAD,
		"NGN": rates.NGN,
		"NIO": rates.NIO,
		"NOK": rates.NOK,
		"NPR": rates.NPR,
		"NZD": rates.NZD,
		"OMR": rates.OMR,
		"PAB": rates.PAB,
		"PEN": rates.PEN,
		"PGK": rates.PGK,
		"PHP": rates.PHP,
		"PKR": rates.PKR,
		"PLN": rates.PLN,
		"PYG": rates.PYG,
		"QAR": rates.QAR,
		"RON": rates.RON,
		"RSD": rates.RSD,
		"RUB": rates.RUB,
		"RWF": rates.RWF,
		"SAR": rates.SAR,
		"SBD": rates.SBD,
		"SCR": rates.SCR,
		"SDG": rates.SDG,
		"SEK": rates.SEK,
		"SGD": rates.SGD,
		"SHP": rates.SHP,
		"SLE": rates.SLE,
		"SLL": rates.SLL,
		"SOS": rates.SOS,
		"SRD": rates.SRD,
		"SSP": rates.SSP,
		"STN": rates.STN,
		"SYP": rates.SYP,
		"SZL": rates.SZL,
		"THB": rates.THB,
		"TJS": rates.TJS,
		"TMT": rates.TMT,
		"TND": rates.TND,
		"TOP": rates.TOP,
		"TRY": rates.TRY,
		"TTD": rates.TTD,
		"TVD": rates.TVD,
		"TWD": rates.TWD,
		"TZS": rates.TZS,
		"UAH": rates.UAH,
		"UGX": rates.UGX,
		"UYU": rates.UYU,
		"UZS": rates.UZS,
		"VES": rates.VES,
		"VND": rates.VND,
		"VUV": rates.VUV,
		"WST": rates.WST,
		"XAF": rates.XAF,
		"XCD": rates.XCD,
		"XDR": rates.XDR,
		"XOF": rates.XOF,
		"XPF": rates.XPF,
		"YER": rates.YER,
		"ZAR": rates.ZAR,
		"ZMW": rates.ZMW,
		"ZWL": rates.ZWL,
	}, nil
}
