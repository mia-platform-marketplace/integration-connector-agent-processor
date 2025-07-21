# mia_template_service_name_placeholder

## Building and Testing

### Testing

Use the provided Makefile to build your processor:

```bash
make test
```

This will run tests and generate a coverage report.

### Docker Build

Build the Docker image with your custom processor:

```bash
docker build -t my-custom-processor:latest .
```

## Tag a new version

```bash
make version VERSION=x.y.z
```
