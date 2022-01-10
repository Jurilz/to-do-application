# build container
FROM rust:1.56 as builder
WORKDIR /usr/src/rust-backend
COPY . .
# installs depencies
RUN cargo install --path .

# run container
FROM debian:buster-slim
# update container and runtime
RUN apt update
RUN apt install -y libpq-dev
RUN rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/local/cargo/bin/rust-backend /usr/local/bin/rust-backend
WORKDIR /usr/local/bin

CMD ["rust-backend"]