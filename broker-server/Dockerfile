# Go 빌드를 위한 기본 이미지 선택
FROM golang:1.20 AS build

# 작업 디렉토리 설정
WORKDIR /app
# Go 모듈 파일 복사 및 의존성 다운로드
COPY go.mod .env ./
RUN go mod tidy

# 애플리케이션 소스 코드 복사
COPY . .

# Go 애플리케이션 빌드
RUN go build -o myapp ./cmd/main.go

# 포트 노출 (애플리케이션이 8080 포트를 사용한다고 가정)
EXPOSE 8080

# 애플리케이션 실행 명령
CMD ["/app/myapp"]