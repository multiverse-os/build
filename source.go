package build

import (
	"time"

	log "github.com/multiverse-os/log"
	watch "github.com/multiverse-os/os/file/watch"
	terminal "github.com/multiverse-os/os/terminal"
)

func (self *Project) Watch() {
	watcher := watch.New()
	watcher.RateLimit(1)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Info(("[Event] File operation: '" + event.FileOperation.String() + "'"))
				terminal.Bash("go build")
			case err := <-watcher.Errors:
				log.FatalError(err)
			case <-watcher.Shutdown:
				return
			}
		}
	}()
	//log.Info("[BUILD] Watching project path: '" + self.Path + '"')
	if err := watcher.AddRecursive(self.Path); err != nil {
		log.FatalError(err)
	}
	for path, file := range watcher.WatchedFiles() {
		// TODO: Parse .gitignore and use it to determine removals
		self.Files = append(self.Files, file)
		log.Info(("Source file added to project: '" + path + "'"))
		log.Info("Source file.Name(): " + file.Name())
	}
	if err := watcher.Start(time.Millisecond * 100); err != nil {
		log.FatalError(err)
	}
}
