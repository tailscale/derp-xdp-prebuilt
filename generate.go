// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package prebuilt

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -type config -type counters_key -type counter_key_af -type counter_key_packets_bytes_action -type counter_key_prog_end Bpf xdp.c -- -I headers
