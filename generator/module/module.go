package module

import (
	"fmt"
	"github.com/lulucas/hasura-pie-cli/errors"
	"github.com/lulucas/hasura-pie-cli/utils"
	"github.com/otiai10/copy"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const (
	repo       = "github.com/lulucas/hasura-pie-modules"
	projectMod = "api/modules"
)

type SyncModuleConfig struct {
	Remote string
	Local  string
}

func GenerateModule(path string) error {
	filename := filepath.Join(path, "module.go")

	if utils.FileExists(filename) {
		return errors.ErrAlreadyExists
	}

	if err := utils.EnsureDir(filename); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("module").Parse(Module)
	if err != nil {
		return err
	}

	if err := t.Execute(f, map[string]string{
		"Module": filepath.Base(path),
	}); err != nil {
		return err
	}

	return nil
}

const (
	syncModuleTempDir = ".modules"
)

func SyncGit() error {
	exists := utils.DirExists(syncModuleTempDir)

	var args []string
	if exists {
		path, err := filepath.Abs(syncModuleTempDir)
		if err != nil {
			return err
		}
		args = append(args, "-C", path, "pull")
	} else {
		args = append(args, "clone", fmt.Sprintf("git@%s.git", repo), syncModuleTempDir)
	}
	cmd := exec.Command("git", args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func SyncModule(modulePath, localPath string) error {
	if localPath == "" {
		localPath = "modules/" + modulePath
	}

	fmt.Printf("sync module %s/%s => %s\n", repo, modulePath, localPath)

	// clear old module
	_ = os.RemoveAll(localPath)

	// extract module
	if err := copy.Copy(filepath.Join(syncModuleTempDir, modulePath), localPath); err != nil {
		return err
	}

	// replace import
	if err := utils.BatchReplaceInFiles(localPath, repo, projectMod); err != nil {
		return err
	}

	return nil
}

func SyncModuleByConfig(cfg []SyncModuleConfig) error {
	for _, c := range cfg {
		if err := SyncModule(c.Remote, c.Local); err != nil {
			return err
		}
	}
	return nil
}
