package airbox

import (
	// "bytes"
	// "crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var alphasslCA = `-----BEGIN CERTIFICATE-----
MIIETTCCAzWgAwIBAgILBAAAAAABRE7wNjEwDQYJKoZIhvcNAQELBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw0xNDAyMjAxMDAw
MDBaFw0yNDAyMjAxMDAwMDBaMEwxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMSIwIAYDVQQDExlBbHBoYVNTTCBDQSAtIFNIQTI1NiAtIEcy
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2gHs5OxzYPt+j2q3xhfj
kmQy1KwA2aIPue3ua4qGypJn2XTXXUcCPI9A1p5tFM3D2ik5pw8FCmiiZhoexLKL
dljlq10dj0CzOYvvHoN9ItDjqQAu7FPPYhmFRChMwCfLew7sEGQAEKQFzKByvkFs
MVtI5LHsuSPrVU3QfWJKpbSlpFmFxSWRpv6mCZ8GEG2PgQxkQF5zAJrgLmWYVBAA
cJjI4e00X9icxw3A1iNZRfz+VXqG7pRgIvGu0eZVRvaZxRsIdF+ssGSEj4k4HKGn
kCFPAm694GFn1PhChw8K98kEbSqpL+9Cpd/do1PbmB6B+Zpye1reTz5/olig4het
ZwIDAQABo4IBIzCCAR8wDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8C
AQAwHQYDVR0OBBYEFPXN1TwIUPlqTzq3l9pWg+Zp0mj3MEUGA1UdIAQ+MDwwOgYE
VR0gADAyMDAGCCsGAQUFBwIBFiRodHRwczovL3d3dy5hbHBoYXNzbC5jb20vcmVw
b3NpdG9yeS8wMwYDVR0fBCwwKjAooCagJIYiaHR0cDovL2NybC5nbG9iYWxzaWdu
Lm5ldC9yb290LmNybDA9BggrBgEFBQcBAQQxMC8wLQYIKwYBBQUHMAGGIWh0dHA6
Ly9vY3NwLmdsb2JhbHNpZ24uY29tL3Jvb3RyMTAfBgNVHSMEGDAWgBRge2YaRQ2X
yolQL30EzTSo//z9SzANBgkqhkiG9w0BAQsFAAOCAQEAYEBoFkfnFo3bXKFWKsv0
XJuwHqJL9csCP/gLofKnQtS3TOvjZoDzJUN4LhsXVgdSGMvRqOzm+3M+pGKMgLTS
xRJzo9P6Aji+Yz2EuJnB8br3n8NA0VgYU8Fi3a8YQn80TsVD1XGwMADH45CuP1eG
l87qDBKOInDjZqdUfy4oy9RU0LMeYmcI+Sfhy+NmuCQbiWqJRGXy2UzSWByMTsCV
odTvZy84IOgu/5ZR8LrYPZJwR2UcnnNytGAMXOLRc3bgr07i5TelRS+KIz6HxzDm
MTh89N1SyvNTBCVXVmaU6Avu5gMUTu79bZRknl7OedSyps9AsUSoPocZXun4IRZZ
Uw==
-----END CERTIFICATE-----`

type Dialer func(network, addr string) (net.Conn, error)

func makeDialer() Dialer {
	return func(network, addr string) (net.Conn, error) {
		CA_Pool := x509.NewCertPool()
		CA_Pool.AppendCertsFromPEM([]byte(alphasslCA))

		config := &tls.Config{RootCAs: CA_Pool}
		c, err := tls.Dial(network, addr, config)
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

// func main() {
// 	fingerprint := []byte{0x53, 0x8d, 0xe6, 0x6e, 0x1d, 0xaf, 0xf6, 0x25, 0xd6, 0x78, 0xb0, 0xb3, 0x71, 0x4, 0xe5, 0x41, 0xd8, 0xc9, 0x68, 0x1f, 0xa6, 0x6, 0x24, 0x6a, 0xf, 0xf9, 0xea, 0xa0, 0x36, 0x55, 0xdc, 0xc1}
// 	client := &http.Client{}
// 	client.Transport = &http.Transport{
// 		DialTLS: makeDialer(fingerprint, false),
// 	}
// 	req, err := http.NewRequest("GET", "https://www.google.com", nil)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(resp)
// }

func GetDevice(token, mac string) (*DeviceResponse, error) {
	client := &http.Client{}
	// Edimax cert has some problem
	// https://www.ssllabs.com/ssltest/analyze.html?d=airbox.edimaxcloud.com
	client.Transport = &http.Transport{
		DialTLS: makeDialer(),
	}
	url := fmt.Sprintf("https://airbox.edimaxcloud.com/devices?token=%s&id=%s", token, mac)
	log.Println("url: ", url)

	request, _ := http.NewRequest(http.MethodGet, url, nil)

	response, err := client.Do(request)
	if err != nil {
		// fmt.Println(err)
		log.Println(err)
		return nil, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	responsePayload := &DeviceResponse{}
	log.Println("response:", string(data))
	json.Unmarshal(data, responsePayload)
	return responsePayload, nil

}
