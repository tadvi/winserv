# winserv

A minimal Windows only service stub for Go.
It only exposes no-op OnExit function if built on Linux.

Enable running programs as services without modifying them.

import "github.com/tadvi/winserv"

If you need more control over the exit behavior, set

```
winserv.OnExit(func() {
	// Do something.
	// Within 10 seconds call:
	os.Exit(0)
})
```

## Credits

This is reworked version of
[kardianos/minwinsvc](https://github.com/kardianos)
