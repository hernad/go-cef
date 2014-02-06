cp -av /home/bringout/dev/cef/cef_bin/out/Release/*.so .
cp -av /home/bringout/dev/cef/cef_bin/out/Release/*.pak .
cp -av /home/bringout/dev/cef/cef_bin/out/Release/chrome-sandbox .
cp -av /home/bringout/dev/cef/cef_bin/out/Release/locales .

sudo chown root:root chrome-sandbox
sudo chmod 4755 chrome-sandbox


