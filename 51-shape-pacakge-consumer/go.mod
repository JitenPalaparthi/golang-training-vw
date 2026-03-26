module shapes

go 1.26.0

require github.com/JitenPalaparthi/shapes-pacakge-demo v1.0.0 // indirect

replace github.com/JitenPalaparthi/shapes-pacakge-demo v1.0.0 => ../50-shapes-package
// replace is only for development testing, possibly after fixing bugs, or change something, add new feature,
// just to test from the dev machine bypassing all the ci pipelines and versioning