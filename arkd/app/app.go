/*
 * Ark - Copyleft of Simone 'evilsocket' Margaritelli.
 * evilsocket at protonmail dot com
 * https://www.evilsocket.net/
 *
 * See LICENSE.
 */
package app

import (
	"encoding/json"
	"fmt"
	"github.com/evilsocket/ark/arkd/models"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	kManifestFileName = "manifest.json"
	kSeedsFileName    = "seeds.json"
)

type Author struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	License string `json:"license"`
}

type Manifest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Author      Author `json:"author"`
}

type App struct {
	Path     string
	Manifest Manifest
	Seeds    []models.Store
}

func Open(path string) (err error, app *App) {
	if path, err = filepath.Abs(path); err != nil {
		return
	}

	stat, err := os.Stat(path)
	if err != nil {
		return
	}

	if stat.IsDir() == false {
		err = fmt.Errorf("Path %s is not a folder.", path)
		return
	}

	seeds_fn := path + "/" + kSeedsFileName
	manifest_fn := path + "/" + kManifestFileName

	seeds := make([]models.Store, 0)
	manifest := Manifest{
		Name:        "?",
		Version:     "0.0.0",
		Description: "",
	}

	if _, err = os.Stat(manifest_fn); err == nil {
		raw, ferr := ioutil.ReadFile(manifest_fn)
		if ferr != nil {
			err = ferr
			return
		}

		if err = json.Unmarshal(raw, &manifest); err != nil {
			return
		}
	}

	if _, err = os.Stat(seeds_fn); err == nil {
		raw, ferr := ioutil.ReadFile(seeds_fn)
		if ferr != nil {
			err = ferr
			return
		}

		if err = json.Unmarshal(raw, &seeds); err != nil {
			return
		}
	}

	app = &App{
		Path:     path,
		Manifest: manifest,
		Seeds:    seeds,
	}

	return nil, app
}

func (app *App) String() string {
	return fmt.Sprintf("%s v%s", app.Manifest.Name, app.Manifest.Version)
}
