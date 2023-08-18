# translate-backend
backend for translate app

# Tech stack
- Golang: see https://go.dev/
- Gin framework: see https://gin-gonic.com/ 
- postgres DB: see https://www.postgresql.org/

# Setting up database
1. run ```createdb -h <DB_HOSTNAME> -p <DB_PORT> -U <DB_USER> diary_app --password```
2. create a .env.local file and copy contents of .env over. Replace fields accordingly

# API spec
```/languages```

```GET```
Response:
```
{
  "id": uint,
  "name": string,
  "code": string
}
```

```POST```
Request:
```
{
  "name": string,
  "code": string
}
```
Response:
Same as ```GET```


```/clients```

```GET```
Response:
```
{
  "id": uint,
  "name": string,
}
```

```POST```
Request:
```
{
  "name": string
}
```
Response:
Same as ```GET```


```/translations```

```GET```
query params: 
```.../translations?from={sourceLanguageID}&to={targetLanguageID}&client={clientID}&q={something+to+translate}```
Response:
```
{
  "id": uint,
  "content": string,
  "languageID": uint,
}
```

```POST```
Request:
```
{
  "sourceLanguageID": uint,
  "text": string,
  "targetLanguageID": uint,
  "translation": string,
  "clientID": uint
}
```
Response:
```
{
  
}
```
