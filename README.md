# Go Little Alchemy
Based on recipes from [littlealchemy.com](https://littlealchemy.com/).

``` go
import littlealchemy "github.com/alancnet/go-little-alchemy"

littlealchemy.GetBaseElements() //=> ["water" "fire" "earth" "air"]
littlealchemy.Combine("water", "fire") //=> ["steam"]
```