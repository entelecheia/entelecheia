package profile

import "testing"

func TestDefaultSiteIsHome(t *testing.T) {
	site := DefaultSite()
	if site.URL != "https://youngjoon-lee.com" {
		t.Fatalf("expected home URL, got %q", site.URL)
	}
}

func TestFindSiteMatchesAliasesCaseInsensitively(t *testing.T) {
	tests := map[string]string{
		"home":         "https://youngjoon-lee.com",
		"Homepage":     "https://youngjoon-lee.com",
		"CV":           "https://youngjoon-lee.com/cv/",
		"resume":       "https://youngjoon-lee.com/cv/",
		"CHU":          "https://halla.ai",
		"koica-tiu":    "https://koica-tiu.halla.ai",
		"repositories": "https://youngjoon-lee.com/repositories",
		"gh":           "https://github.com/entelecheia",
		"LI":           "https://linkedin.com/in/entelecheia",
	}

	for alias, want := range tests {
		site, ok := FindSite(alias)
		if !ok {
			t.Fatalf("expected alias %q to resolve", alias)
		}
		if site.URL != want {
			t.Fatalf("alias %q: expected %q, got %q", alias, want, site.URL)
		}
	}
}

func TestFindSiteRejectsUnknownAlias(t *testing.T) {
	if site, ok := FindSite("missing"); ok {
		t.Fatalf("expected unknown alias to fail, got %#v", site)
	}
}

func TestRequiredSitesExist(t *testing.T) {
	for _, alias := range []string{"home", "cv", "halla", "koica", "rise", "staix", "research", "repos", "github", "linkedin"} {
		if _, ok := FindSite(alias); !ok {
			t.Fatalf("expected required alias %q to exist", alias)
		}
	}
}
