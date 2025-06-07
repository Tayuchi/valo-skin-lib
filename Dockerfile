# デプロイ用コンテナに含めるバイナリを作成するコンテナ

FROM golang.1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# デプロイ用のコンテナ

FROM dibian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ローカル開発環境で利用するホットリロード環境

FROM golang:1.24 as dev
WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.7
CMD ["air"]