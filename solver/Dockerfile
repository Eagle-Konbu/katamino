FROM ekidd/rust-musl-builder:stable AS builder

COPY ./Cargo.toml /home/rust

RUN mkdir src && \
    echo "fn main(){}" > /home/rust/src/main.rs && \
    cargo build --release

COPY ./src /home/rust

RUN rm -f /home/rust/target/x86_64-unknown-linux-musl/release/deps/solver*

RUN cargo build --release && \
    strip /home/rust/target/x86_64-unknown-linux-musl/release/solver

FROM scratch

COPY --from=builder \
    /home/rust/src/target/x86_64-unknown-linux-musl/release/solver .

ENTRYPOINT ["./solver"]