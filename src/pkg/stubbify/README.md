# Stubify
An initial package that does url shortening using Base62 encoding/decoding

### Endpoints
    GET /stubify/:stub -> Redirects to a the Base62 decoded string
        Example: curl -i http://localhost:8080/stubify/G7SHuwHpwDIzG3FHZ9HNjDFhHaF
            HTTP/1.1 302 Found
            Date: Fri, 28 May 2021 22:59:01 GMT
            Content-Length: 0
            X-Request-Id: d6be4ef2-166e-4d0d-8651-e65fa0107511
            Vary: Origin
            Access-Control-Allow-Origin:
            Location: https://google.com



    POST /api/v1/stubify -> Returns a shortened url linking to the supplied body.destination
        curl -X POST \
             -d '{ "destination":"https://google.com" }' \
             -H 'Content-Type: application/json' \
              http://localhost:8080/api/v1/stubify 
            {
                "shortened_url":"http://localhost:8080/stubify/G7SHuwHpwDIzG3FHZ9HNjDFhHaF00v"
            }    