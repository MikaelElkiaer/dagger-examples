# Dockerfile secrets

## Vanilla Docker

```bash
docker build -t dagger-test --secret id=secret-file,src=secret.txt --file ../Dockerfile .
```

## Dagger

```bash
dagger call --debug build --build-context . --dockerfile ../Dockerfile --secret-file file:secret.txt
```
