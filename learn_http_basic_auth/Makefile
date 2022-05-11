USERNAME=myusername
PASSWORD=mypassword

http: http.go
	go build -o http http.go

# Authentication is provided as an HTTP request header of the form:
# "Authorization: <scheme> <credentials>"
# 	<scheme> is "Basic"
# 	<credentials> are a base64 encoding of username:password
CREDENTIALS=$(shell /bin/echo -n "$(USERNAME):$(PASSWORD)" | base64)
CURL_VERBOSE_OPTS=-v -i
test: http
	./http :8080 & sleep 1 && curl $(CURL_VERBOSE_OPTS) -H "Authorization: Basic $(CREDENTIALS)" localhost:8080/api ; wait
	./http :8080 & sleep 1 && curl $(CURL_VERBOSE_OPTS) -u "$(USERNAME):$(PASSWORD)" localhost:8080/api2 ; wait
