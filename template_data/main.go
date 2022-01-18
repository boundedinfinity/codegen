package template_data

import "embed"

//go:embed go/* ts/*
var Raw embed.FS
