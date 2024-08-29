FROM golang:1.22.4

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o credly-badges .
RUN chmod +x ./credly-badges

ENTRYPOINT ["./credly-badges"]