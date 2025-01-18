FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main


FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
COPY template.json ./template.json
COPY Xolonium-Bold.ttf ./Xolonium-Bold.ttf
COPY ./images ./images
EXPOSE 80
ENTRYPOINT ["./main"]