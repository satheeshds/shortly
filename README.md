# Shortly

Simple API to shorten the URL, tracks the top domains which shortened, and redirect to original if shortened URL is given

## Prerequisite

1. Docker and Docker Compose

## How to run

```bash
docker compose up --build -d
```

## How to stop

```bash
docker compose down
```

## How to test the functionality

1. open [swagger](http://localhost/swagger/index.html) url (<http://localhost/swagger/index.html>)
2. try out the shorten method ![Screen shot](/docs/Tryout.png "Optional title")
3. provide the url which needs to be shortned in *req_url* field.
4. click on execute, which will trigger a post request on server, and returns the shortened url, paste the shortned url in browser, which will redirect to original url.
5. to know the *top 3 domains* which done the most shortening, try and execute the next action in the swagger */shorten/topdomains*  
