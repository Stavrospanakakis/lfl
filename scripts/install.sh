#!/bin/sh

command_exists() {
	command -v "$@" >/dev/null 2>&1
}

error() {
	echo ${RED}"Error: $@"${RESET} >&2
}

command_exists git || {
		error "Git is not installed. Please install git first."
		exit 1
}

# Clone repository
git clone https://github.com/Stavrospanakakis/lfl.git

command_exists go || {
    error "Go is not installed. Please install go first."
    rm -rf $(pwd)/lfl
    exit 1
}

# Install
echo "Installing..."
export GOPATH="/usr/local"
export PATH=$PATH:$GOPATH/bin
cd $(pwd)/lfl
go install github.com/Stavrospanakakis/lfl

# Copy configuration file to home
echo "Creating configuration file..."
cp $(pwd)/config/.lfl.json $HOME
cd ..

# Remove repository
rm -rf $(pwd)/lfl

echo "lfl installed successfully! Try lfl --help."