.PHONY: all clean run V1_V1 V1_V2 V2_V1 V2_V2

all: V1_V1 V1_V2 V2_V1 V2_V2

# V1 plugin + V1 host
V1_V1:
	go build -tags v1 -o V1_V1/srelib ./plugin/impl/
	go build -tags v1 -o V1_V1/host .

# V1 plugin + V2 host (host wants V2, plugin only has V1)
V1_V2:
	go build -tags v1 -o V1_V2/srelib ./plugin/impl/
	go build -tags v2 -o V1_V2/host .

# V2 plugin + V1 host
V2_V1:
	go build -tags v2 -o V2_V1/srelib ./plugin/impl/
	go build -tags v1 -o V2_V1/host .

# V2 plugin + V2 host (both negotiate V2)
V2_V2:
	go build -tags v2 -o V2_V2/srelib ./plugin/impl/
	go build -tags v2 -o V2_V2/host .

run: all
	@echo "=== V1 plugin + V1 host ==="
	@cd V1_V1 && ./host
	@echo ""
	@echo "=== V1 plugin + V2 host ==="
	@cd V1_V2 && ./host
	@echo ""
	@echo "=== V2 plugin + V1 host ==="
	@cd V2_V1 && ./host
	@echo ""
	@echo "=== V2 plugin + V2 host ==="
	@cd V2_V2 && ./host

clean:
	rm -f V1_V1/srelib V1_V1/host
	rm -f V1_V2/srelib V1_V2/host
	rm -f V2_V1/srelib V2_V1/host
	rm -f V2_V2/srelib V2_V2/host
