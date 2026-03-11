# bookworm, last updated 2026-02-28
FROM debian:13

RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install --no-install-recommends -y \
    sqlite3 \
    golang \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

WORKDIR "/repo/go"

ENV GOCACHE="/repo/ignore/container-cache"

EXPOSE 8080

ENTRYPOINT ["/repo/entrypoint.sh"]
