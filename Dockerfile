# Telling to use Docker's golang ready image
FROM golang

# Create app folder 
RUN mkdir /tracker

# Copy our file in the host contianer to our contianer
COPY . /tracker

# Set /app to the go folder as workdir
WORKDIR /tracker

# Generate binary file from our /app
RUN go build cmd/tracker/main.go

# Expose the port 3000
EXPOSE 9000:9000

# Run the app binarry file 
CMD ["./main"]