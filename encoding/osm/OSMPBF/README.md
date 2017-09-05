Those two proto files are part of the osmosis project.

They can be found here [https://github.com/openstreetmap/osmosis](https://github.com/openstreetmap/osmosis)

Converted the type of StringTable to repeated sequence of string instead of bytes

```proto
message StringTable {
    repeated string s = 1;
}
```