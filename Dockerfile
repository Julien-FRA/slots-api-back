FROM golang:1.18 as BUILDER

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Je vais faire une build en 2 étapes
# https://dave.cheney.net/2016/01/18/cgo-is-not-go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY ./app .
RUN go mod download \
    && go mod verify \
    && go build -o /build/buildedApp main.go

FROM scratch as FINAL

COPY --from=BUILDER /build/buildedApp .

ENTRYPOINT ["./buildedApp"]