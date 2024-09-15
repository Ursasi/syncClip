package peer

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Foundation

#import <Cocoa/Cocoa.h>

extern void clipboardChanged(char* content);

static void StartClipboardListener() {
    @autoreleasepool {
        NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
        __block NSInteger changeCount = [pasteboard changeCount];

        NSTimer *timer = [NSTimer timerWithTimeInterval:0.5 repeats:YES block:^(NSTimer * _Nonnull timer) {
            NSInteger newChangeCount = [pasteboard changeCount];
            if (newChangeCount != changeCount) {
                changeCount = newChangeCount;
                NSString *content = [pasteboard stringForType:NSPasteboardTypeString];
                if (content != nil) {
                    const char *cContent = [content UTF8String];
                    clipboardChanged((char *)cContent);
                }
            }
        }];

        [[NSRunLoop currentRunLoop] addTimer:timer forMode:NSDefaultRunLoopMode];
        [[NSRunLoop currentRunLoop] run];
    }
}
*/
import "C"
import (
	"fmt"
	"sync"
)

var (
	latestClipboardContent string
	mutex                  sync.Mutex
)

//export clipboardChanged
func clipboardChanged(cContent *C.char) {
	content := C.GoString(cContent)
	handleClipboardChange(content)
}

func handleClipboardChange(content string) {
	mutex.Lock()
	defer mutex.Unlock()
	if content != latestClipboardContent {
		latestClipboardContent = content
		fmt.Println("clipboard changed：", content)
		clipChan <- content
	}
}

func WatchClipboard() {
	C.StartClipboardListener()
}
