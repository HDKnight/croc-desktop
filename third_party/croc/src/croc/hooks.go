package croc

import (
	"os"
	"path"
	"path/filepath"

	"github.com/schollz/croc/v11/src/models"
	"github.com/schollz/croc/v11/src/utils"
	"github.com/schollz/progressbar/v3"
)

// Files at least this large get a GUI hashing progress report. Smaller files
// finish quickly enough that a separate phase would only flicker.
const guiHashProgressMin int64 = 32 << 20

// Hooks bridges an embedded client to a GUI. Install before Send or Receive;
// callbacks may run on transfer goroutines. Nil callbacks retain CLI behavior.
// Receive metadata is validated before any acceptance callback is invoked.
type Hooks struct {
	OnProgress          func(ProgressEvent)
	OnStateChange       func(string)
	OnAcceptRequest     func(AcceptRequest) bool
	OnOverwriteRequest  func(path string, resumePct float64) bool
	OnConfirmSendToPeer func(machineID string) bool
	// OnText receives the validated text instead of printing it to stdout.
	// Croc still owns and removes the temporary receive file.
	OnText func(string)
}

type ProgressEvent struct {
	Filename   string `json:"filename"`
	BytesDone  int64  `json:"bytesDone"`
	BytesTotal int64  `json:"bytesTotal"`
	FilesDone  int    `json:"filesDone"`
	FilesTotal int    `json:"filesTotal"`
	// PeerToPeer is true when file bytes are flowing directly between the two
	// machines. False means the configured relay is carrying the file.
	PeerToPeer bool `json:"peerToPeer"`
}

type AcceptRequest struct {
	Files        []FileInfo
	TotalSize    int64
	TotalFolders int
	SenderID     string
	IsText       bool
}

func (c *Client) SetHooks(h *Hooks) { c.hooks = h }

func (c *Client) emitState(state string) {
	if c.hooks != nil && c.hooks.OnStateChange != nil {
		c.hooks.OnStateChange(state)
	}
}

// Snapshot metadata when the control loop installs a bar. Data workers must
// not read the control loop's mutable file index or completion map.
func (c *Client) progressMetadata() ProgressEvent {
	event := ProgressEvent{FilesDone: c.FilesToTransferCurrentNum, FilesTotal: len(c.FilesToTransfer)}
	if i := c.FilesToTransferCurrentNum; i >= 0 && i < len(c.FilesToTransfer) {
		event.Filename = c.FilesToTransfer[i].Name
		event.BytesTotal = c.FilesToTransfer[i].Size
	}
	return event
}

func (c *Client) emitProgress(bar *progressbar.ProgressBar) {
	if c.hooks == nil || c.hooks.OnProgress == nil || bar == nil {
		return
	}
	c.barMu.RLock()
	event, current := c.guiProgress, c.bar == bar
	c.barMu.RUnlock()
	if !current {
		return
	}
	event.BytesDone = int64(bar.State().CurrentBytes)
	event.PeerToPeer = c.peerToPeerDataPath()
	c.hooks.OnProgress(event)
}

// hashForTransfer hashes filename for a send or receive. Embedded clients
// report progress instead of drawing a console bar.
func (c *Client) hashForTransfer(filename, algorithm string, cliProgress bool) ([]byte, error) {
	ctx := c.clientContext()
	reported := false
	if c.hooks != nil && c.hooks.OnProgress != nil {
		cliProgress = false
		if info, statErr := os.Stat(filename); statErr == nil && info.Size() >= guiHashProgressMin {
			reported = true
			c.emitState("hashing")
			base := filepath.Base(filename)
			ctx = utils.WithHashProgress(ctx, func(done, total int64) {
				if c.hooks == nil || c.hooks.OnProgress == nil {
					return
				}
				c.hooks.OnProgress(ProgressEvent{
					Filename:   base,
					BytesDone:  done,
					BytesTotal: total,
				})
			})
		}
	}
	sum, err := utils.HashFileCtx(ctx, filename, algorithm, cliProgress)
	// Hashing runs before the peer is connected. Return to the waiting or
	// connecting status unless file bytes have already started moving.
	if reported && !c.activeTransferStarted() {
		if c.Options.IsSender {
			c.emitState("waiting")
		} else {
			c.emitState("connecting")
		}
	}
	return sum, err
}

func (c *Client) askReceiveOverwrite(fileInfo FileInfo, resumable bool) (bool, error) {
	if c.hooks == nil || c.hooks.OnOverwriteRequest == nil {
		return askReceiveOverwrite(c.clientContext(), fileInfo, resumable)
	}
	percent := float64(100)
	destination := path.Join(fileInfo.FolderRemote, fileInfo.Name)
	if resumable {
		ranges := utils.MissingChunks(destination, fileInfo.Size, models.TCP_BUFFER_SIZE/2)
		missing := utils.ChunkRangesBytes(ranges, fileInfo.Size, models.TCP_BUFFER_SIZE/2)
		percent = 100 - float64(missing)/float64(fileInfo.Size)*100
	}
	accepted := c.hooks.OnOverwriteRequest(destination, percent)
	return accepted, c.clientContext().Err()
}
