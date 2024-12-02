# URL Shortener

- **Built using:** Go
- **Description:** Simple & light-weight URL shortener application.

## How to use

### Shorten
- Make your URLs shorter by sending a request to
```bash
POST /api/shorten
```
- Send a `post` request to `localhost:8080/api/shorten` :
```bash
curl -X POST http://localhost:8080/api/shorten?url=https://www.lego.com/en-us/product/polaroid-onestep-sx-70-camera-21345

# will return
Shortened URL: XXshort1Q
```
-  will return the shortened version of the URL: `https://www.lego.com/en-us/product/polaroid-onestep-sx-70-camera-21345`
- hence: `localhost:8080/XXshort1Q`
- `XXshort1Q` is the short code for this URL example

### Redirecting
- When you browse a shortened URL the go application will redirect you to the original URL.
- Browsing `localhost:8080/XXshort1Q` will redirect to `https://www.lego.com/en-us/product/polaroid-onestep-sx-70-camera-21345`.
- Run the application, follow the steps and try it for yourself!

### Get URLs
- You can fetch all the URLs that has a short code.
```bash
curl -X GET http://localhost:8080/urls
```
```bash
# return
"https://www.lego.com/en-us/product/polaroid-onestep-sx-70-camera-21345": "http://localhost:8080/XXshort1Q",
"https://google.com": "http://localhost:8080/XXshort2Q",
```
