PROJECT_ROOT := src/
VERSION = 0.0.0

.DEFAULT_GOAL := all

# We want to add compliance for all built binaries
_CHECK_COMPLIANCE = $(shell find src/ -not -path '*/vendor/*' -name '*.go' | xargs -I{} dirname {} |sed 's/src\///g' | uniq | sort)
_TESTS = $(shell find src/ -not -path '*/vendor/*' -name '*_test.go' | xargs -I{} dirname {} | sed 's/src\///g'|uniq | sort)

BINARIES = \
	preston

# Build all binaries
BINS = $(addsuffix .build,$(BINARIES))

# Add relevant tests
GO_TESTS = $(addsuffix .test,$(_TESTS))

include Makefile.gobuild

# Ensure our own code is compliant..
compliant: $(addsuffix .compliant,$(_CHECK_COMPLIANCE))

install: $(BINS)
	test -d $(DESTDIR)/usr/bin || install -D -d -m 00755 $(DESTDIR)/usr/bin; \
	install -m 00755 bin/* $(DESTDIR)/usr/bin/.;

ensure_modules:
	@ ( \
		git submodule init; \
		git submodule update; \
	);

# See: https://github.com/meitar/git-archive-all.sh/blob/master/git-archive-all.sh
release: ensure_modules
	git-archive-all.sh --format tar.gz --prefix preston-$(VERSION)/ --verbose -t HEAD preston-$(VERSION).tar.gz

all: $(BINS)
