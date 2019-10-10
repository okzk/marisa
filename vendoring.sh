#!/bin/sh -ex


cd "$(dirname $0)"
rm -rf marisa-trie tmp
git clone https://github.com/s-yata/marisa-trie.git tmp

cd tmp
REV=$(git rev-parse HEAD)
git archive --prefix=marisa-trie/ HEAD | tar x -C ../
cd ..
rm -rf tmp
echo "marisa-trie@${REV} is vendored." > vendoring_info.txt




#REV=$(git ls-remote https://github.com/s-yata/marisa-trie.git HEAD | cut -f1)
#echo $REV
#
#svn export https://github.com/s-yata/marisa-trie.git/branches/master marisa-trie/
## git archive --prefix=marisa-trie/ --remote=https://github.com/s-yata/marisa-trie.git mas  > hogehoge.tar
#
#echo "vendaring https://github.com/s-yata/marisa-trie/commit/${REV}"
