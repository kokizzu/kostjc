#!/usr/bin/env bash

cd presentation && go test -bench=BenchmarkGenerateViews && goimports -w web_view.GEN.go
