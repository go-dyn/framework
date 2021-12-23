package config

import (
	"github.com/goravel/framework/support/facades"
	testing2 "github.com/goravel/framework/support/testing"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	err := testing2.CreateEnv()
	assert.Nil(t, err)
	assert.NotPanics(t, func() {
		app := Application{}
		app.Init()
	})
}

func TestEnv(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, "goravel", app.Env("APP_NAME").(string))
	assert.Equal(t, "", app.Env("DB_HOST", "127.0.0.1").(string))
}

func TestAdd(t *testing.T) {
	app := Application{}
	app.Init()
	app.Add("app", map[string]interface{}{
		"env": "local",
	})

	assert.Equal(t, "local", app.GetString("app.env"))
}

func TestGet(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, "goravel", app.Get("APP_NAME").(string))
}

func TestGetString(t *testing.T) {
	configServiceProvider := ServiceProvider{}
	configServiceProvider.Register()
	facadesConfig := facades.Config
	facadesConfig.Add("database", map[string]interface{}{
		"default": facadesConfig.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]interface{}{
			"mysql": map[string]interface{}{
				"host": facadesConfig.Env("DB_HOST", "127.0.0.1"),
			},
		},
	})

	assert.Equal(t, "goravel", facadesConfig.GetString("APP_NAME"))
	assert.Equal(t, "", facadesConfig.GetString("database.connections.mysql.host"))
	assert.Equal(t, "mysql", facadesConfig.GetString("database.default"))
}

func TestGetInt(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, app.GetInt("DB_PORT"), 3306)
}

func TestGetBool(t *testing.T) {
	app := Application{}
	app.Init()

	assert.Equal(t, true, app.GetBool("APP_DEBUG"))

	err := os.Remove(".env")
	assert.Nil(t, err)
}
