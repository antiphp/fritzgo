FROM gcr.io/distroless/static:nonroot

WORKDIR /app/
COPY fritzgo /app/

ENTRYPOINT ["/app/fritzgo"]
