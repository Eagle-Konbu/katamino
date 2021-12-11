# バイナリのビルド用にrust-musl-builderイメージを使用する
FROM ekidd/rust-musl-builder:latest AS builder

COPY . .

RUN sudo chown -R rust:rust ./

RUN cargo build --release && \
    strip /home/rust/src/target/x86_64-unknown-linux-musl/release/solver

FROM scratch

COPY --from=builder \
    /home/rust/src/target/x86_64-unknown-linux-musl/release/solver /

ENTRYPOINT ["/solver"]