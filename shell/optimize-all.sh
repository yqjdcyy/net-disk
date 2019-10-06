#! /bin/bash

IMG_PATH=${1:-/data/tmp/img/run-all/v2}
IMG_SIZE_LIMIT=${2:-1M}

echo -e "Path=\t$IMG_PATH"
echo -e "Size=\t$IMG_SIZE_LIMIT"

optimizeJpg() {
    find ${IMG_PATH} -size +${IMG_SIZE_LIMIT} | grep -E 'jpg|jpeg' > jpg.log
    cat jpg.log | xargs -I {} convert -quality 75 {} {}
    cat jpg.log | xargs -I {} echo -e "\t{}\tconverted"
}

optimizePng() {
    find ${IMG_PATH} -size +${IMG_SIZE_LIMIT} | grep -E 'png' > png.log
    cat png.log | xargs -I {} convert -quality 75 {} {}.jpg
    cat png.log | xargs -I {} echo -e "\t{}\tconverted"
    cat png.log | xargs -I {} rm {}
    cat png.log | xargs -I {} echo -e "\t{}\tdeleted"
}

optimizeJpg
optimizePng
