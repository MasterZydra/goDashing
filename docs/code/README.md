# Source code

**Contentlist**  
- [Using disk as data source](#using-disk-as-data-source)

## Using disk as data source
The package `packr` is used to contain the assets files into the binary itself. The files requested by the application could be loaded from that box. As visible in commit [f5009bd](https://github.com/Hein-Software-Solutions/goDashing/commit/f5009bd) the decision was not to do it this way.  
The reasons are the following:
- The directory on the disc allows modifications. New widgets can easily be added and changes in the widgets can be done without recompiling the binary.
- The file server.go gets simplified. An implementation with both box and disk as data source would be possible. But the logic to only send the content once would be more complex.
- Theoretically the reading speed should not be that different. If a file is read often, the OS will hold it in the RAM or cache to reduce read actions. The included files in the binary also get loaded in the RAM only when in use and moved onto the disk if the page is not used very often.