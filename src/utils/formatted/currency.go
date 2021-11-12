package formatted

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

func IndonesiaCurrrency(val float64) string {
	return fmt.Sprintf("Rp %s", strings.ReplaceAll(humanize.Commaf(val), ",", "."))
}
