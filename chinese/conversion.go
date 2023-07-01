package chinese

import (
	"log"

	"github.com/liuzl/gocc"
)

func S2HK(in string) (string, error) { return ConvertAny(in, `s2hk`) }

/*
s2t Simplified Chinese to Traditional Chinese
t2s Traditional Chinese to Simplified Chinese
s2tw Simplified Chinese to Traditional Chinese (Taiwan Standard)
tw2s Traditional Chinese (Taiwan Standard) to Simplified Chinese
s2hk Simplified Chinese to Traditional Chinese (Hong Kong Standard)
hk2s Traditional Chinese (Hong Kong Standard) to Simplified Chinese
s2twp Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom
tw2sp Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom
t2tw Traditional Chinese (OpenCC Standard) to Taiwan Standard
t2hk Traditional Chinese (OpenCC Standard) to Hong Kong Standard
*/
func ConvertAny(in, method string) (string, error) {
	convert, err := gocc.New(method)
	if err != nil {
		log.Printf("[CONVERT] %v \n", err)
		return ``, err
	}
	return convert.Convert(in)
}
