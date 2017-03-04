#!/bin/sh
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  chmod a+w "${distdir}"
  glide install
  go build -v -o ${distdir}/bmt
}

go_build