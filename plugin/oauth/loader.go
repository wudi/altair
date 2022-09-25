package oauth

import (
	"fmt"

	"github.com/kodefluence/altair/core"
	"github.com/kodefluence/altair/module"
)

// Provide create new oauth plugin provider
func Load(appBearer core.AppBearer, dbBearer core.DatabaseBearer, pluginBearer core.PluginBearer, apiError module.ApiError) error {
	if appBearer.Config().PluginExists("oauth") == false {
		return nil
	}

	version, err := pluginBearer.PluginVersion("oauth")
	if err != nil {
		return err
	}

	switch version {
	case "1.0":
		return version_1_0(appBearer, dbBearer, pluginBearer, apiError)
	default:
		return fmt.Errorf("undefined template version: %s for metric plugin", version)
	}
}
