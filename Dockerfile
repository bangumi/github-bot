FROM gcr.io/distroless/base-debian11

ENTRYPOINT ["/app/github-bot"]

COPY /dist/github-bot /app/github-bot
