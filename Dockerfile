FROM golang:latest as builder
WORKDIR /build
COPY /src/. .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main


FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
COPY /src/template.json ./template.json
COPY /src/Xolonium-Bold.ttf ./Xolonium-Bold.ttf
COPY /src/images ./images
EXPOSE 80
ENTRYPOINT ["./main"]