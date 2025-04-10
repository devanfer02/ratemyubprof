FROM golang:1.22-alpine AS builder 

RUN apk update && apk add --no-cache git 

WORKDIR /app    

COPY go.mod go.sum ./ 

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main 

FROM gcr.io/distroless/base-debian11

WORKDIR /app 

COPY --from=builder /app/main . 
COPY --from=builder /app/env.json .

ENTRYPOINT [ "/app/main" ]