FROM golang:1.15 as builder

ARG ARG_GOPROXY
ENV GOPROXY $ARG_GOPROXY

ENV CGO_ENABLED 0

WORKDIR /home/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN make build


FROM orvice/go-runtime

ENV PROJECT_NAME ab-job

COPY --from=builder /home/app/bin/${PROJECT_NAME} /app/bin/${PROJECT_NAME}

ENTRYPOINT exec /app/bin/${PROJECT_NAME}