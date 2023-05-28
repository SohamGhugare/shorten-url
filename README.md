# URL Shortener

## A lightning fast URL Shortener using Go's Gin Framework and Redis

### Features:

    - Custom url
    - Ratelimiting
    - Visits counter

**Uses docker-compose for service management** <br>
Dockerfiles are included in their respective directories including the docker-compose config.

### API Docs:

1. `POST /api/v1` Shorten a url

   - Required JSON:

     ```json
     {
        "url": actual_url,
        "short": short_url,
        "expiry": expiry(optional)
     }

     ```

2. `GET /:short_url` Redirects to the actual url

3. (WIP) `GET /api/v1/info?url="short_url"` Get info on the registered short url
