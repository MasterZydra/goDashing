# Source code annotations

**Contentlist**  
- [Using disk as data source](#using-disk-as-data-source)
- [Splitting application.js into multiple files](#splitting-applicationjs-into-multiple-files)

## Using disk as data source
The package `packr` is used to contain the assets files into the binary itself. The files requested by the application could be loaded from that box. As visible in commit [f5009bd](https://github.com/Hein-Software-Solutions/goDashing/commit/f5009bd) the decision was not to do it this way.  
The reasons are the following:
- The directory on the disc allows modifications. New widgets can easily be added and changes in the widgets can be done without recompiling the binary.
- The file server.go gets simplified. An implementation with both box and disk as data source would be possible. But the logic to only send the content once would be more complex.
- Theoretically the reading speed should not be that different. If a file is read often, the OS will hold it in the RAM or cache to reduce read actions. The included files in the binary also get loaded in the RAM only when in use and moved onto the disk if the page is not used very often.

## Splitting application.js into multiple files
The [application.js](../../assets/public/js/application.js) contains all required dependencies like jQuery, Batman, Gridster or D3. This makes it a huge file - with over 30000 lines. It was to big to be loaded by GitHub ([application.js in full length](../../../149d2fd7ae585ddad473e7ef6ae1d7315db38d1c/assets/public/js/application.js)). The separation into mutiple files, like the dependency [Chart.min.js](../../assets/public/js/Chart.min.js), has a some advantages:
- The maintenance gets much easier. The library file can simply be replaced with e.g. a newer version or fixes in the source code can be done in a file with a manageable number of lines.
- The actual application logic of the dashboard is readable because it is the only content in the file [application.js](../../assets/public/js/application.js).
- The separation provides a better overview of all used libraries, instead of searching in that original huge file. Now a view in the [public/js/](../../assets/public/js) shows all libraries.
- The removal of a library is as simple as deleting the file and removing the entry in the [layout.gerd](../../assets/dashboards/layout.gerb)
