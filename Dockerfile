FROM ubuntu

COPY bin/mic-test /mic-test
RUN chmod +x /mic-test

# We want each image built unique for testing
RUN echo $(date) >> /home/rand.txt
RUN echo "hello world" >> /home/hello.txt
