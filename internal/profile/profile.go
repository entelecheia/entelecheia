package profile

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Site struct {
	Aliases []string
	URL     string
	Label   string
}

var Sites = []Site{
	{Aliases: []string{"home", "homepage"}, URL: "https://youngjoon-lee.com", Label: "Personal homepage"},
	{Aliases: []string{"cv", "curriculum-vitae", "resume"}, URL: "https://youngjoon-lee.com/cv/", Label: "Academic CV (EN/KR)"},
	{Aliases: []string{"halla", "chu"}, URL: "https://halla.ai", Label: "Cheju Halla University AI"},
	{Aliases: []string{"koica", "koica-tiu", "tiu"}, URL: "https://koica-tiu.halla.ai", Label: "KOICA-TIU AI Training Center"},
	{Aliases: []string{"rise"}, URL: "https://rise.jeju.ai", Label: "RISE Jeju AI"},
	{Aliases: []string{"staix", "bwlb"}, URL: "https://staixbwlb.com", Label: "STAI x BWLB"},
	{Aliases: []string{"research"}, URL: "https://research.jeju.ai", Label: "Research Jeju AI"},
	{Aliases: []string{"repos", "repositories"}, URL: "https://youngjoon-lee.com/repositories", Label: "Public repos"},
	{Aliases: []string{"github", "gh"}, URL: "https://github.com/entelecheia", Label: "GitHub"},
	{Aliases: []string{"linkedin", "li"}, URL: "https://linkedin.com/in/entelecheia", Label: "LinkedIn"},
}

func DefaultSite() Site {
	return Sites[0]
}

func FindSite(alias string) (Site, bool) {
	needle := strings.ToLower(strings.TrimSpace(alias))
	for _, site := range Sites {
		for _, candidate := range site.Aliases {
			if needle == strings.ToLower(candidate) {
				return site, true
			}
		}
	}
	return Site{}, false
}

func PrintIntro(w io.Writer, color bool) {
	c := colors{enabled: color}
	fmt.Fprintf(w, "%sYoung Joon Lee (이영준)%s\n", c.bold(), c.reset())
	fmt.Fprintf(w, "%sAI Professor & Builder | Associate Professor of Artificial Intelligence, Cheju Halla University%s\n\n", c.dim(), c.reset())

	fmt.Fprintln(w, "Roles")
	fmt.Fprintln(w, "  - Chief AI Officer (CAIO) & Assistant Vice President for AX")
	fmt.Fprintln(w, "  - Director, KOICA-TIU-Cheju Halla AI Training Center")
	fmt.Fprintln(w, "  - Director, Cheju Halla RISE Center")
	fmt.Fprintln(w, "  - Builder of AI education, research, and public-impact systems")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Focus")
	fmt.Fprintln(w, "  - Agentic AI, large language models, and knowledge graphs")
	fmt.Fprintln(w, "  - NLP for economic and financial analysis")
	fmt.Fprintln(w, "  - Ethical, explainable, and human-centered AI")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Flagship Work")
	fmt.Fprintln(w, "  - AI-based autonomous road repair robot")
	fmt.Fprintln(w, "  - Agentic AI-based intelligent education platform")
	fmt.Fprintln(w, "  - KOICA-TIU AI Training Center in Uzbekistan")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "Homepage: %s\n", DefaultSite().URL)
	fmt.Fprintln(w, "CV:       https://youngjoon-lee.com/cv/")
	fmt.Fprintln(w, "GitHub:   https://github.com/entelecheia")
	fmt.Fprintln(w, "Email:    yj.lee@chu.ac.kr")
	fmt.Fprintf(w, "\n%sTip:%s run 'entelecheia open cv' to open my bilingual CV.\n", c.bold(), c.reset())
}

func PrintLinks(w io.Writer) {
	aliasWidth := len("Aliases")
	urlWidth := len("URL")
	labelWidth := len("Label")
	for _, site := range Sites {
		aliasWidth = max(aliasWidth, len(strings.Join(site.Aliases, ", ")))
		urlWidth = max(urlWidth, len(site.URL))
		labelWidth = max(labelWidth, len(site.Label))
	}

	fmt.Fprintf(w, "%-*s  %-*s  %-*s\n", aliasWidth, "Aliases", urlWidth, "URL", labelWidth, "Label")
	fmt.Fprintf(w, "%-*s  %-*s  %-*s\n", aliasWidth, strings.Repeat("-", aliasWidth), urlWidth, strings.Repeat("-", urlWidth), labelWidth, strings.Repeat("-", labelWidth))
	for _, site := range Sites {
		fmt.Fprintf(w, "%-*s  %-*s  %-*s\n", aliasWidth, strings.Join(site.Aliases, ", "), urlWidth, site.URL, labelWidth, site.Label)
	}
}

func SupportsColor(file *os.File) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	info, err := file.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

type colors struct {
	enabled bool
}

func (c colors) bold() string {
	if !c.enabled {
		return ""
	}
	return "\x1b[1m"
}

func (c colors) dim() string {
	if !c.enabled {
		return ""
	}
	return "\x1b[2m"
}

func (c colors) reset() string {
	if !c.enabled {
		return ""
	}
	return "\x1b[0m"
}
