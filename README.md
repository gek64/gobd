# GO Builder(GOBD)

- A tool can help you easily build golang static files for multiple operating systems and architectures

## Usage

```

```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/gobd https://github.com/gek64/gobd/releases/latest/download/gobd-linux-amd64
chmod +x /usr/local/bin/gobd

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/gobd https://github.com/gek64/gobd/releases/latest/download/gobd-freebsd-amd64
chmod +x /usr/local/bin/gobd
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/gobd.git
cd gobuilder
export CGO_ENABLED=0 
go build -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
