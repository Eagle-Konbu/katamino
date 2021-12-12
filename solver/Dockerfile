FROM ekidd/rust-musl-builder:stable AS builder

COPY . .

RUN cargo build --release && \
    strip /home/rust/target/x86_64-unknown-linux-musl/release/solver

FROM scratch

COPY --from=builder \
    /home/rust/target/x86_64-unknown-linux-musl/release/solver .

ENV RUST_BACKTRACE=full

ENTRYPOINT ["./solver"]