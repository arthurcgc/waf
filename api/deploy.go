package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arthurcgc/waf/internal/pkg/manager"
	echo "github.com/labstack/echo/v4"
)

type DeployOpts struct {
	Name      string
	Replicas  int
	Namespace string
	ProxyPass string `json:"proxy,omitempty"`
}

func (a *Api) deploy(c echo.Context) error {
	var opts DeployOpts
	err := json.NewDecoder(c.Request().Body).Decode(&opts)
	if err != nil {
		return err
	}

	args := manager.DeployArgs{
		Name:      opts.Name,
		Namespace: opts.Namespace,
		Replicas:  opts.Replicas,
		ProxyPass: opts.ProxyPass,
	}

	if err := a.manager.Deploy(c.Request().Context(), args); err != nil {
		return fmt.Errorf("error during deploy: %s", err.Error())
	}

	return c.String(http.StatusCreated, "Created nginx resource!\n")
}
