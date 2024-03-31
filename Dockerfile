FROM ubuntu:latest

# Update and install dependencies
RUN apt-get -y update && \
    apt-get install -y build-essential git cmake libssl-dev wget pkg-config

# Install Go
ENV GO_VERSION=1.22.1
RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -xvf go${GO_VERSION}.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Add Go to PATH
ENV PATH="/usr/local/go/bin:${PATH}"

# Get and install liboqs
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs && \
    cmake -S liboqs -B liboqs/build -DBUILD_SHARED_LIBS=ON && \
    cmake --build liboqs/build --parallel 4 && \
    cmake --build liboqs/build --target install

# Enable a normal user
RUN useradd -m -c "Open Quantum Safe" oqs
USER oqs
WORKDIR /home/oqs

# Get liboqs-go
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs-go.git

# Set ENV variables for liboqs-go configuration
ENV PKG_CONFIG_PATH=/home/oqs/liboqs-go/.config:$PKG_CONFIG_PATH
ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
