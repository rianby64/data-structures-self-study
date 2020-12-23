
test:
	podman build -f ContainerFile -t go-ds-self-study .
	podman run -it --rm go-ds-self-study
