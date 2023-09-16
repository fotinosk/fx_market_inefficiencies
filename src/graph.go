package Graph

import(
	"fmt"
	"github.com/fotinosk/fx_market_inefficiencies/APIinterface"
)

func main() {
	nodes := APIinterface.generate_nodes()
	fmt.Println(nodes)
}
