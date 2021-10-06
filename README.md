# grpc-hands-on for Golang


### Echo Web Framework
참고 사이트 : https://echo.labstack.com/

### Swagger
참고 사이트 : https://github.com/swaggo/swag
swagger 주소 형식 : http://localhost:3000/docs/index.html
```shell
go get github.com/swaggo/swag/cmd/swag 
go get github.com/swaggo/echo-swagger
swag init    # 주석의 설정을 읽어서 swagger.yaml, swagger.json, docs.go 파일을 생성
```