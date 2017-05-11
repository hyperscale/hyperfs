
all: build-hyperfs build-hyperfs-api build-hyperfs-index build-hyperfs-storage


build-hyperfs:
	@cd bin/hyperfs; cargo build

build-hyperfs-api:
	@cd bin/hyperfs-api; cargo build

build-hyperfs-index:
	@cd bin/hyperfs-index; cargo build

build-hyperfs-storage:
	@cd bin/hyperfs-storage; cargo build


run-hyperfs: build-hyperfs
	@./target/debug/hyperfs -h
