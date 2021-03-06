# winserv

A minimal Windows only service stub for Go.
It only exposes no-op OnExit function if built on other OS.

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

Make sure to exit with os.Exit(0) or similar at the end of OnExit.

If winserv.Interactive == true we are not running as Service.

## Credits

This is reworked version of
[kardianos/minwinsvc](https://github.com/kardianos)
