OUTPUT_DIR=./_output

bin=webclient

product: clean
	env GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/${bin}.linux ./

build: clean
	go build -o ${OUTPUT_DIR}/${bin} ./

clean:
	rm -rf ${OUTPUT_DIR}/${bin}


