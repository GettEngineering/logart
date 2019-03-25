# errArt - WIP

*error handling as an art*

## Description:
TODO


## Notes:

1. When `err == nil`, `errart.Wrap(err, "...")` will return `nil`, so `WithField/s(...)` can't be used

2. const errors should be created with `NewConstError(...)`
