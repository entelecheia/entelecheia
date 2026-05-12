package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunDefaultPrintsIntro(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run(nil, &stdout, &stderr, func(string) error { return nil })

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, stderr.String())
	}
	out := stdout.String()
	for _, want := range []string{"Young Joon Lee", "AI", "https://youngjoon-lee.com", "https://youngjoon-lee.com/cv/"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected intro to contain %q, got:\n%s", want, out)
		}
	}
}

func TestRunVersionCommands(t *testing.T) {
	for _, args := range [][]string{{"version"}, {"-v"}, {"--version"}} {
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		code := run(args, &stdout, &stderr, func(string) error { return nil })

		if code != 0 {
			t.Fatalf("args %v: expected exit 0, got %d; stderr=%s", args, code, stderr.String())
		}
		if got := strings.TrimSpace(stdout.String()); got != Version {
			t.Fatalf("args %v: expected version %q, got %q", args, Version, got)
		}
	}
}

func TestRunLinksPrintsSites(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"links"}, &stdout, &stderr, func(string) error { return nil })

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, stderr.String())
	}
	out := stdout.String()
	for _, want := range []string{"home, homepage", "cv, curriculum-vitae, resume", "https://youngjoon-lee.com", "halla.ai"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected links to contain %q, got:\n%s", want, out)
		}
	}
}

func TestRunOpenUsesResolvedURL(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var opened string

	code := run([]string{"open", "halla"}, &stdout, &stderr, func(url string) error {
		opened = url
		return nil
	})

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, stderr.String())
	}
	if opened != "https://halla.ai" {
		t.Fatalf("expected halla URL, got %q", opened)
	}
	if !strings.Contains(stdout.String(), "Opening") {
		t.Fatalf("expected open confirmation, got %q", stdout.String())
	}
}

func TestRunOpenUnknownAlias(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"open", "missing"}, &stdout, &stderr, func(string) error {
		t.Fatal("open should not be called for unknown aliases")
		return nil
	})

	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
	if !strings.Contains(stderr.String(), "unknown site alias") {
		t.Fatalf("expected unknown alias error, got %q", stderr.String())
	}
}

func TestRunOpenReportsOpenError(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"open"}, &stdout, &stderr, func(string) error {
		return errors.New("boom")
	})

	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
	if !strings.Contains(stderr.String(), "failed to open") {
		t.Fatalf("expected open failure, got %q", stderr.String())
	}
}
