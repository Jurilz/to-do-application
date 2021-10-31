# build container
FROM rust:1.56 as builder

RUN apt update
RUN apt install -y libpq-dev
RUN cargo install diesel_cli --no-default-features --features postgres

WORKDIR /usr/src/rust-backend
COPY . .

# installs depencies
RUN cargo install --path .

# run container
FROM debian:buster-slim
# update container and runtime
RUN apt update
RUN apt install -y libpq-dev

COPY --from=builder /usr/local/cargo/bin/rust-backend /usr/local/bin/rust-backend
EXPOSE 8080
CMD ["rust-backend"]
