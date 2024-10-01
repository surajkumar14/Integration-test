FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy

# Copy the run_tests.sh script and make it executable
COPY run_tests.sh /app/run_tests.sh
RUN chmod +x /app/run_tests.sh

# Run the test script before starting the service
CMD ["/bin/sh", "-c", "./run_tests.sh && go run ./main.go"]
