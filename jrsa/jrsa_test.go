package jrsa

import (
	"log"
	"testing"
)

var (
	pub  = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0+t+o4Wb6IEZiR1/7lSSrhz+hKVUl5Xcimw2FszhKkRKhW3I+Mmm7pDbN/aTt67PRTZFmjVU9qWvCXtqTCF6RQJN35wNlCs+HDrf1Pea3EiMd+n8InLcHGuqmIMhy9vRwXZE06TJ9zuSlRxppxv6PEl+4zX2VzW1lz+hyP2lxopml04UvH/QRWk9Q5ZIQfYAm0N7LjHhVucuu5aM87EurAMuH6+m6mwp+0uapNtfSlQfAvNc5+KYsOvn2/7IjmWIA3V/wyxGdbjsKX3Ez5STEE9l6Pj1x5Y9IQTKcAM/J4HZkI8oCXchj9iuJ76z1r+M8dJfb4L95S0B2l8YE+azUwIDAQAB`
	pri  = `MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDT636jhZvogRmJHX/uVJKuHP6EpVSXldyKbDYWzOEqREqFbcj4yabukNs39pO3rs9FNkWaNVT2pa8Je2pMIXpFAk3fnA2UKz4cOt/U95rcSIx36fwictwca6qYgyHL29HBdkTTpMn3O5KVHGmnG/o8SX7jNfZXNbWXP6HI/aXGimaXThS8f9BFaT1DlkhB9gCbQ3suMeFW5y67lozzsS6sAy4fr6bqbCn7S5qk219KVB8C81zn4piw6+fb/siOZYgDdX/DLEZ1uOwpfcTPlJMQT2Xo+PXHlj0hBMpwAz8ngdmQjygJdyGP2K4nvrPWv4zx0l9vgv3lLQHaXxgT5rNTAgMBAAECggEAMcp02Kq9tOUQQN0Df9WzGOGLE1NdCriVNpiyJK1CghHFiJAA1iNa61ZxizaOAmw6GsLjm5XIPqMy+wFaOkTrW2dtMIqTzipSz4Mtr/X1FqidCFebe+SMvsi1LlwxBprBL7k/9pITOdlvWUVxtWenZZ7HB09Yb6kY305+DqE53+mbMdfsj+3QiJ8Q0R3fyDDYiK8stKXhHZAouc40nQpdddt6mhT/H8eeqIlWd1rACbWxVMWK1oEMqAOyZWwkwUR7P0fArcQR3YWsl+bLbVSBXh9PHXNR61I/v904A+ChYWrUAfeB6WwigOHle2ecf2iD6DIZn2d7s5Q6N9qn7tEWgQKBgQDudqUJXKhqQcyDd1OCe8OgEd8Rh3YMckmdtW3EfivjuqPmfJsZXefKoSC4pdlq8L/fUu/1GDST7/vhkAMNcE6UTP+Djcz5HpjwoGZQNYlFRCuQjkPhrlfkoavvxEWPKWIt7FgYB/XfLNLAsnIRuUPTAhAkg+URkzu862ldOOSoswKBgQDjgSLnIhaAV3pjAnLWsqr/+79dUW9bqzx39A4kWCTKMOY9AQVv18B7rfXUpqOqIDm42MO8d7uIT3qkOIXA0DHxfScUwXrXErOukPt4H9B4nUcoorjQJ/AOC+Id2mGpLOeNXqlyvwCDLR3UO3Uiewj0etkLHRN7BBpWdJTW/qLa4QKBgQC8qW2E7KPp/UjasiuFznlFSR5c9fk1e642cfX4QYJj16QYlHj5JsxeCgCjVcOf0f1bqHhkRtRCPtne6Vsg5LumpQx6flOuvMBvj9eimdiSzo5Q/d1CzfaH0kj/lx8ZVrpwbs57pMCn4wSSpBuPXi4E4Vr4KMmwj/XxT6a6tGpj9QKBgFrXwNEcOnHq/FK2spZqZ3+pzaL8loO7H30iddcrXx1hYz4uvzYGp7R7JTNRiv6uuX3HCHmkMbzfR7B1ZWs1dwvflpBiJaPlstvcxp0TzxGqc7SVLUjo+aESO6sB/YcpCBjaGL1Z6WF3zXZme4JWMKA2wZ3/cTzoyX+GM7yZlvvhAoGAOij6ehG9zg2fh5lKRy9e7xxdD6nwaPC0MYnH2ON2ye3Gge52FJ47DMzIUWlLMIqwGnqjCjDlt316UP6+dNtFbzWDqA7G/bQ7H34w1AcqlzUxiY+M8o128iY/d2LSp54VQ4/BxoSP3JL2AwpdSZdFMKyj3jRfEr+P6szhpmug1xw=`
	data = `{"username":"admin","password":"q1w2e3r4t5"}`
)

func TestCreateKeys(t *testing.T) {
	r := JRSA{}
	err := r.CreateKeys(2048)
	if err != nil {
		t.Error(err.Error())
	}
	r.Print()
	resp, err := r.PubEncrypt(data)
	checkErr(err)
	println(`[ENCRYPTED]`, resp)

	resp, err = r.PriDecrypt(resp)
	checkErr(err)
	println(`[DECRYPTED]`, resp)
}
func TestKeyParser(t *testing.T) {
	r := JRSA{}
	checkErr(r.Set(pub, Public))
	checkErr(r.Set(pri, Private))
	r.Print()

	resp, err := r.PubEncrypt(data)
	checkErr(err)
	println(`[ENCRYPTED]`, resp)

	resp, err = r.PriDecrypt(resp)
	checkErr(err)
	println(`[DECRYPTED]`, resp)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v \n", err)
	}
}
