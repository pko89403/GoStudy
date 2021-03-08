# go-chi 테스트


## running
~~~sh
make dev // localhost:9000
~~~

## go test
~~~sh
go test ./platform/newsfeed // 테스트 코드
go test -cover ./... // 전체 테스트 코드
~~~

## curl post request 
~~~sh
curl -X POST localhost:9000/newsfeed -d '{"title" : "Hello", "post" : "World!"}' 
~~~

### 참고 동영상
[유튜브링크](https://www.youtube.com/watch?v=zeme_TmXyBk)