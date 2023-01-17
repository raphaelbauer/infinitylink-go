package httpmockable

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "infinitylink-go/hamcrest"
)

func TestHttpReal(t *testing.T) {
	t.Run("Get works)", func(t *testing.T) {
		//given
		httpImpl := HttpImpl{}

		// when
		// real call... maybe not so cool... but for academic reasons we keep it...
		response, _ := httpImpl.Get("https://example.com")
		body, _ := ioutil.ReadAll(response.Body)
		bodyAsString := string(body)

		//then
		Expect(response.StatusCode, t).ToEqual(200)
		fmt.Println("here!!")
		fmt.Println(bodyAsString)
		Expect(bodyAsString, t).ToContain("Example Domain")

	})

}
