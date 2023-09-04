.RECIPEPREFIX := >

MODULE := github.com/s21toolkit/s21adapter

CMD_ADAPTER := $(MODULE)/cmd/adapter
CMD_SPEC := $(MODULE)/cmd/spec

.PHONY: default all
default: adapter
all: adapter spec

.PHONY: adapter
adapter: tidy
>	go build $(CMD_ADAPTER)

.PHONY: spec
spec: tidy
>	go build $(CMD_SPEC)

.PHONY: tidy
tidy:
>	go mod tidy

.PHONY: clean
clean:
>	rm -f $(notdir $(CMD_ADAPTER) $(CMD_SPEC))
