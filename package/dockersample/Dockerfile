# syntax=docker/dockerfile:1
FROM golang:1.19

RUN mkdir /app

# Root 디렉터리에 있는 모든 것을 app 경로로 복사
ADD . /app 

WORKDIR /app

COPY go.mod ./

# 종속성을 가져오려면 이 go mod 다운로드 명령을 추가하세요.
RUN go mod download 

# 현재 경로에 있는 go file을 컴파일
RUN go build -o main . 

# 실행
CMD ["/app/main"] 