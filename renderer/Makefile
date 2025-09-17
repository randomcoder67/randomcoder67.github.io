normal:
	go mod tidy
	go build -o gorender main.go setup.go renderer.go parser.go definitions.go

install:
	cp gorender ~/.local/bin/

clean:
	rm gorender

uninstall:
	rm ~/.local/bin/gorender

full: normal install clean
