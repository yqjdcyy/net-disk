#! /bin/bash

export IMG_PATH=${1:-/data/tmp/img/run-all/v2}
export IMG_SIZE_LIMIT=${2:-1M}
export IMG_LAST_HOUR_LIMIT=${3:-1}
export LAST_FILE_DATE=`date '+%Y-%m-%d %H:%M:%S' -d $IMG_LAST_HOUR_LIMIT+' hour ago'`

echo -e "Path=\t$IMG_PATH"
echo -e "Size=\t$IMG_SIZE_LIMIT"
echo -e "Last=\t$IMG_LAST_HOUR_LIMIT"
echo -e "Last.Date=\t$LAST_FILE_DATE"

optimizeJpg() {
    JPG_FILE="jpg.last.log"
    find ${IMG_PATH} -size +${IMG_SIZE_LIMIT} -newermt "${LAST_FILE_DATE}" -type f | grep -E 'jpg|jpeg' > ${JPG_FILE}
    cat ${JPG_FILE} | xargs -I {} convert -quality 75 {} {}
    cat ${JPG_FILE} | xargs -I {} echo -e "\t{}\tconverted"
}

optimizePng() {
    PNG_FILE="png.last.log"
    find ${IMG_PATH} -size +${IMG_SIZE_LIMIT} -newermt "${LAST_FILE_DATE}" -type f | grep -E 'png' > ${PNG_FILE}
    cat ${PNG_FILE} | xargs -I {} convert -quality 75 {} {}.jpg
    cat ${PNG_FILE} | xargs -I {} echo -e "\t{}\tconverted"
    cat ${PNG_FILE} | xargs -I {} rm {}
    cat ${PNG_FILE} | xargs -I {} echo -e "\t{}\tdeleted"
}

optimizeJpg
optimizePng
