FROM golang:latest as builder
WORKDIR /build
COPY /services/renderer/. .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main


FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
COPY /services/renderer/template.json ./template.json
COPY /services/renderer/Xolonium-Bold.ttf ./Xolonium-Bold.ttf
COPY /services/renderer/images ./images
EXPOSE 80
ENTRYPOINT ["./main"]