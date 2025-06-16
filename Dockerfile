ARG executableName=go-funny-endpoints

FROM golang:1.24-alpine AS builder

ARG depVersion=0.5.0
ARG executableName
ARG repoName=dllg
ARG projectName=${executableName}
RUN apk add --no-cache alpine-sdk zip

WORKDIR ${GOPATH}/src/github.com/${repoName}/${projectName}

# Get the source files in
COPY . .

# Build
RUN [ "make", "build-linux" ]
RUN mkdir -p /opt/${executableName}/
RUN cp "${GOPATH}"/src/github.com/"${repoName}"/"${projectName}"/build/* /opt/"${projectName}/"

FROM alpine:3.22

ARG executableName
ENV EXECUTABLE_NAME=${executableName}

RUN apk add --no-cache ca-certificates curl

COPY --from=builder /opt/${EXECUTABLE_NAME}/${EXECUTABLE_NAME}-linux /opt/

WORKDIR /opt/

ENV PORT="18080"

EXPOSE 18080

ENTRYPOINT ["sh", "-c", "/opt/${EXECUTABLE_NAME}-linux server"]
