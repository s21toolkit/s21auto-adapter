.RECIPEPREFIX := >

MODULE := github.com/s21toolkit/s21adapter

CMD_ADAPTER := $(MODULE)/cmd/adapter
CMD_SPEC := $(MODULE)/cmd/spec

S21ADAPTER := s21adapter
S21ADAPTER_SPEC := s21adapter_spec

.PHONY: default all
default: $(S21ADAPTER)
all: $(S21ADAPTER) $(S21ADAPTER_SPEC)

.PHONY: $(S21ADAPTER)
$(S21ADAPTER): tidy
>	go build -o $@ $(CMD_ADAPTER)

.PHONY: $(S21ADAPTER_SPEC)
$(S21ADAPTER_SPEC): tidy
>	go build -o $@ $(CMD_SPEC)

.PHONY: tidy
tidy:
>	go mod tidy

.PHONY: clean
clean:
>	rm -f $(S21ADAPTER) $(S21ADAPTER_SPEC)
