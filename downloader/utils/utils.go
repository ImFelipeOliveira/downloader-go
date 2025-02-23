package utils

import (
	"fmt"
	"net/url"
	"os/user"
	"path"
	"runtime"
	"strings"
)

func VerifyExtension(link, extensions string) bool {
	for _, ext := range strings.Fields(extensions) {
		if strings.HasSuffix(link, strings.TrimSpace(ext)) {
			return true
		}
	}
	return false
}

func FormatUrl(base, link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}
	u, err := url.Parse(base)
	if err != nil {
		fmt.Println("Error: ", err)
		return link
	}

	return u.Scheme + "://" + u.Host + "/" + strings.TrimPrefix(link, "/")
}

func SetDownloadPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("erro ao obter o diretório do usuário: %v", err)
	}

	homeDir := u.HomeDir
	downloadsPath := path.Join(homeDir, "Downloads")
	switch runtime.GOOS {
	case "windows", "linux":
		return downloadsPath, nil
	default:
		return "", fmt.Errorf("sistema operacional não suportado")
	}
}
