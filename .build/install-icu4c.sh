cwd=$(pwd)
cd /tmp
wget http://download.icu-project.org/files/icu4c/58.2/icu4c-58_2-src.tgz
tar -xf icu4c-58_2-src.tgz
cd icu/source
./configure --prefix=/usr && make
sudo make install
cd $cwd
