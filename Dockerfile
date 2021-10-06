FROM golang:1.17 as builder

ENV APPLICATION_NAME=golang-hands-on \
    BUILD_WORK_PATH=app \
    CGO_ENABLED=0 \
    GOOS=linux

# Create and change to the app directory.
WORKDIR /${BUILD_WORK_PATH}

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

RUN go get github.com/swaggo/swag/cmd/swag &&\
    swag init


# Build the binary.
RUN go build -v -o ${APPLICATION_NAME}

FROM alpine:3

ENV	APPLICATION_NAME=golang-hands-on \
    BUILD_WORK_PATH=app \
    # CONFIG_PATH=configs \
    SERVICE_USER=scarif \
   	SERVICE_UID=10001 \
 	  SERVICE_GROUP=scarif \
 	  SERVICE_GID=10001

RUN	addgroup -g ${SERVICE_GID} ${SERVICE_GROUP} && \
 	  adduser -g "${SERVICE_NAME} user" -D -H -G ${SERVICE_GROUP} -s /sbin/nologin -u ${SERVICE_UID} ${SERVICE_USER}
RUN apk add --no-cache ca-certificates

COPY --from=builder /${BUILD_WORK_PATH}/${APPLICATION_NAME} /${APPLICATION_NAME}
# COPY --from=builder /${BUILD_WORK_PATH}/${CONFIG_PATH} /${CONFIG_PATH}

# Ports : 3000 - Echo Framework, 3001 - gRPC
EXPOSE 3000/tcp
EXPOSE 3001/tcp

# Run the web service on container startup.
ENTRYPOINT ./${APPLICATION_NAME}