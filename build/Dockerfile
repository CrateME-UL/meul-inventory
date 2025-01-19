FROM golang:1.23.4-alpine AS build
WORKDIR /go/src/app

COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY scripts/ ./scripts/
COPY web/ ./web/
COPY go.mod .

ARG BUILD_MODE
ARG DB_HOST
ARG DB_NAME
ARG DB_PORT
ARG DB_SSL_MODE

ENV BUILD_DIR="/go/src/app/build" \
    CMD_DIR="/go/src/app/cmd" \
    BUILD_MODE=${BUILD_MODE} \
    DB_HOST=${DB_HOST} \
    DB_NAME=${DB_NAME} \
    DB_PORT=${DB_PORT} \
    DB_SSL_MODE=${DB_SSL_MODE}

RUN chmod +x /go/src/app/scripts/build_inv_release.sh
RUN --mount=type=secret,id=db_password_file \
    --mount=type=secret,id=db_user_file \
    DB_PASSWORD=$(cat /run/secrets/db_password_file) && \
    DB_USER=$(cat /run/secrets/db_user_file) && \
    /go/src/app/scripts/build_inv_release.sh $DB_USER $DB_PASSWORD $BUILD_MODE $DB_HOST $DB_NAME $DB_PORT $DB_SSL_MODE

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/src/app/build/inv-meul-app /
CMD ["/inv-meul-app"]
