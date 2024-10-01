FROM golang:1.22.4

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o credly-badges ./cmd/credly-badges
RUN chmod +x ./credly-badges
RUN mv ./credly-badges /usr/local/bin/

ENTRYPOINT ["credly-badges"]
