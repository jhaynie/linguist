# make file will basically patch the built general tso and merge it into our own tree to make it easier to go get

PKG_ROOT=$(GOPATH)/src/github.com/generaltso/linguist

all:
	@mkdir -p generaltso/linguist generaltso/linguist/data generaltso/linguist/tokenizer
	@cp $(PKG_ROOT)/linguist.go \
		$(PKG_ROOT)/static.go \
		$(PKG_ROOT)/analyse.go \
		$(PKG_ROOT)/exclude.go \
		$(PKG_ROOT)/LICENSE \
		generaltso/linguist
	@cp -R $(PKG_ROOT)/tokenizer/*.go generaltso/linguist/tokenizer
	@cp $(PKG_ROOT)/data/data.go \
		$(PKG_ROOT)/data/classifier \
		generaltso/linguist/data
	@sed -i '' -E 's/github.com\/generaltso\/linguist/github.com\/jhaynie\/linguist\/generaltso\/linguist/g' generaltso/linguist/*.go


test:
	go test -v .