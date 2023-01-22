FROM gcr.io/distroless/base-debian11

ENTRYPOINT ["/app/github-bot"]

WORKDIR "/lib/app"

COPY /dist/github-bot /app/github-bot
