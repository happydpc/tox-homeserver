source ~/triline/shell/android-ndk-env.sh
source ~/triline/shell/android-go-env.sh

# android API<19 no mmap64

env | grep CGO_
set -x

# test libav link
export PKG_CONFIG_PATH=/androidsys/lib/pkgconfig  # this help some #cgo pkg-config link
# test compile suvcap, must name as libxxx.so. or it will not appear in app/lib/arm/ dir.
# $CC -xc -g -fPIE -pie suvcap/suvcap.c -o libs/libsuvcapd.so \
#    $CGO_CFLAGS $CGO_LDFLAGS `pkg-config --libs libavformat libavdevice libavfilter`

go install -p 1 -v -i -x -pkgdir ~/oss/pkg/android_arm tox-homeserver/avhlp

go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtqt
go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/kitech/qt.go/qtrt
go install -v -i -pkgdir ~/oss/pkg/android_arm github.com/mattn/go-sqlite3

# go build -p 1 -v  -pkgdir ~/oss/pkg/android_arm .
rm -vf libmain.so

# note: all linked so's soname should be xx.so, can not with version suffix, xx.so.1
time go build -p 1 -v -i -pkgdir ~/oss/pkg/android_arm -buildmode=c-shared -o libmain.so .
chmod +x libmain.so

$CC -xc andwrapmain.c.nogo -shared   -o libgolem.so -lmain -L. -Wl,-soname,libgolem.so
ccret=$?

if [ x"$1" == x"2" ] && [ x"$ccret" == x"0" ]; then
    echo "Ready to build apk..."
    sleep 2;
    sh make2.sh
fi


