FOM ubuntu

COPY bin/mic-test /mic-test
RUN chmod +x /mic-test

# We want each image built to be unique for testing
RUN echo $(date) >> /home/rand.txt
RUN echo "Hello Toby!" >> /home/toby-test.txt
