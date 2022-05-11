set -e
set -u

mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg << 'EOF' # internal
# import config
EOF

echo "WORK=$WORK";

/opt/homebrew/Cellar/go/1.18.1/libexec/pkg/tool/darwin_arm64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p command-line-arguments -complete -buildid cTrX2Dx5ldsBbCTF_mau/cTrX2Dx5ldsBbCTF_mau -goversion go1.18.1 -shared -c=4 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./simple.go

