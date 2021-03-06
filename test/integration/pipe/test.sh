#!/bin/sh

set -eu

TEMP_DIR=$(mktemp -d /tmp/XXXX)
TEST_DIR="$(cd "$(dirname "${0}")"; echo "$(pwd)")"

ALIASES=$(cd "${TEST_DIR}/../../..//dist"; echo "$(pwd)/aliases -c ${TEST_DIR}/aliases.yaml")
DIFF=$(if which colordiff >/dev/null; then echo "colordiff -Buw --strip-trailing-cr"; else echo "diff -Bw"; fi)
MASK="sed -e s|${HOME}|[HOME]|g -e s|${TEMP_DIR}|[TEMP_DIR]|g"

${ALIASES} gen --export-path ${TEMP_DIR} >/dev/null

echo "foo" | ${TEMP_DIR}/alpine1 /bin/sh -c "alpine2 /bin/sh -c 'cat -'" | ${DIFF} ${TEST_DIR}/stdout -
echo "foo" | ${ALIASES} run /usr/local/bin/alpine1 /bin/sh -c "alpine2 /bin/sh -c 'cat -'" | ${DIFF} ${TEST_DIR}/stdout -