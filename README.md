# cmlabs-backend-crawler-freelance-test

This project to scrap a website where all the content  into file html as output



### Prerequisite
- Install all library `go mod vendor`

### Start
```shell
make start
```
The app will run on port 5111 as default


### Do Scrapping
```curl
curl --location 'http://localhost:5111/scraping' \
--header 'Content-Type: application/json' \
--data '{
    "urls": [
        "https://cmlabs.co",
        "https://sequence.day",
        "https://id.wikipedia.org/wiki/Wikipedia"
    ]
}'
```

