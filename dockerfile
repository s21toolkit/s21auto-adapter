FROM --platform=$TARGETPLATFORM golang as build
WORKDIR /build
COPY . .
RUN go mod tidy && \
  CGO_ENABLED=0 go build -o ./s21adapter ./cmd/adapter


FROM --platform=$TARGETPLATFORM alpine
COPY --from=build /build/s21adapter /bin/s21adapter
RUN chmod +x /bin/s21adapter

LABEL org.opencontainers.image.source https://github.com/s21toolkit/s21adapter

ENTRYPOINT s21adapter