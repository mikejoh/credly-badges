FROM golang:1.22.4

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o credly-badges .
RUN chmod +x ./credly-badges
RUN mv ./credly-badges /usr/local/bin/

ENTRYPOINT ["credly-badges"]