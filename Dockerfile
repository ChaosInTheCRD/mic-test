FROM alpine:latest

COPY bin/mic-test /mic-test
RUN chmod +x /mic-test

# We want each image built to be unique for testing
RUN echo $($RANDOM | md5sum | head -c 20; echo;) >> /home/rand.txt

ENTRYPOINT ["/mic-test"]
