FROM gcr.io/distroless/base-debian11@sha256:ac69aa622ea5dcbca0803ca877d47d069f51bd4282d5c96977e0390d7d256455

ENTRYPOINT ["/app/github-bot"]

COPY /dist/github-bot /app/github-bot
