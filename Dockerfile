# base go image
FROM golang:$(goAlpineVersion) as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apk --no-cache update && apk add --no-cache gcc libc-dev g++ git

RUN go mod tidy

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN CGO_ENABLED=1 go build -tags musl -o application .

RUN chmod +x /app/application
# build a tiny docker image

FROM $(goDockerImage)

RUN apk --no-cache add tzdata

RUN adduser -g 1701 -u 1702 -s /bin/bash -h /app -D app

USER app

WORKDIR /app/

COPY --chown=app:app --from=builder /app/application /app
COPY --chown=app:app --from=builder /app/resources/app_config.env /app
COPY --chown=app:app  --from=builder /app/docs /app/docs

# COPY assets /app/assets

CMD [ "/app/application" ]