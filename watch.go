package build

import (
	"time"

	log "github.com/multiverse-os/log"
	watch "github.com/multiverse-os/os/file/watch"
	terminal "github.com/multiverse-os/os/terminal"
)

func (self *Project) Hooks(operation FileOperation) []FileHook {
	return self.FileHooks[operation]
}

func (self *Project) On(operation FileOperation, hook FileHook) {
	self.FileHooks[operation] = append(self.FileHooks[operation], hook)
}

func (self *Project) Watch() {
	watcher := watch.New()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Info(("[Event] File operation: '" + event.FileOperation.String() + "'"))
				for _, hook := range self.Hooks(event.FileOperation) {
					hook()
				}
				terminal.Bash("go build")
			case err := <-watcher.Errors:
				log.FatalError(err)
			case <-watcher.Shutdown:
				return
			}
		}
	}()
	//log.Info("[BUILD] Watching project path: '" + self.Path + '"')
	if err := watcher.Watch(self.Path, true); err != nil {
		log.FatalError(err)
	}
	for path, file := range watcher.Files() {
		// TODO: Parse .gitignore and use it to determine removals
		self.Files = append(self.Files, file)
		log.Info(("Source file added to project: '" + path + "'"))
		log.Info("Source file.Name(): " + file.Name())
	}
	if err := watcher.Start(time.Millisecond * 100); err != nil {
		log.FatalError(err)
	}
}
